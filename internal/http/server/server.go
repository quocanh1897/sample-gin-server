package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"git.taservs.net/platform/edca-api/internal/api/controller"
	"git.taservs.net/platform/edca-api/internal/config"
	"git.taservs.net/platform/edca-api/internal/constant"
	"git.taservs.net/platform/edca-api/internal/middleware"
	requestdto "git.taservs.net/platform/edca-api/internal/model/dto/request"
	"git.taservs.net/platform/edca-api/internal/pkg/telemetry"
	"git.taservs.net/platform/edca-api/internal/repository/komrade"
	sessionusecase "git.taservs.net/platform/edca-api/internal/usecase/session"
	appvalidator "git.taservs.net/platform/edca-api/internal/validator"
)

//go:generate mockery --name=Server --case=snake
type Server interface {
	Start() error
}

func New(controller controller.Controller, sessionUseCase sessionusecase.SessionUseCase, komradeRepository komrade.KomradeRepository) Server {
	cfg := config.GetAppConfig()
	server := &serverImpl{
		serverConfig: &cfg,
		controller:   controller,
		middleware:   middleware.New(&cfg, sessionUseCase, komradeRepository),
	}

	server.withRouter()

	return server
}

type serverImpl struct {
	controller   controller.Controller
	router       *gin.Engine
	middleware   middleware.Middleware
	serverConfig *config.AppConfig
}

func (i *serverImpl) Start() error {
	ctx := context.Background()
	// Instantiate the OTLP HTTP exporter
	metricExporter, err := telemetry.InitGPRCMetricExporter(i.serverConfig.Otlp, ctx)
	traceExporter, err := telemetry.InitGPRCTraceExporter(i.serverConfig.Otlp, ctx)
	if err != nil {
		log.Fatalln(err)
	}
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(i.serverConfig.ServiceName),
		semconv.DeploymentEnvironmentKey.String(i.serverConfig.Env),
	)
	traceProvider, err := telemetry.SetupTracer(traceExporter, "edca-api", *i.serverConfig, resource)
	if err != nil {
		log.Fatal(err)
	}
	meterProvider, err := telemetry.SetupMetrics(metricExporter, "edca-api", *i.serverConfig, resource)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
		if err := meterProvider.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	slog.Info(fmt.Sprintf("Listening and serving HTTP on :%d", config.GetHTTPConfig().Port))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetHTTPConfig().Port),
		Handler: i.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
		return err
	}

	log.Println("Server exiting")
	return nil
}

func (i *serverImpl) withRouter() {
	if config.GetAppConfig().Env != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	// router.Use(otelgin.Middleware("edca-api"))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	router.Use(i.middleware.LimiterHandler())
	router.Use(i.middleware.LoggerHandler())
	router.Use(i.middleware.ErrorHandlerHandler())
	router.Use(i.middleware.ContextHandler())
	router.Use(i.middleware.TelemetryHandler())
	router.Use(gin.Recovery())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("is_uuid", appvalidator.IsUUID)
		v.RegisterStructValidation(appvalidator.ValidateResetMFAFactorStruct, requestdto.Factor{})
		v.RegisterValidation("is_device_model", appvalidator.DeviceModelValidator)
		v.RegisterValidation("is_iso8601", appvalidator.ISO8601Validator)
		v.RegisterValidation("is_subscription_type", appvalidator.SubscriptionTypeValidator)
		v.RegisterValidation("is_agency_role", appvalidator.AgencyRoleValidator)
		v.RegisterValidation("is_agency_type", appvalidator.AgencyTypeValidator)
		v.RegisterValidation("is_ax_production_type", appvalidator.AxProductionTypeValidator)
		v.RegisterValidation("is_user_status", appvalidator.UserStatusValidator)
		v.RegisterValidation("is_parser_type", appvalidator.CadRmsParserTypeValidator)
		v.RegisterValidation("is_evidence_type", appvalidator.CadRmsEvidenceTypeValidator)
		v.RegisterValidation("is_evidence_source", appvalidator.CadRmsEvidenceSourceValidator)
		v.RegisterValidation("is_software_mode", appvalidator.SoftwareModeValidator)
		v.RegisterStructValidation(appvalidator.ValidateUserInputMFAFactorStruct, requestdto.VerifyChallenge{})
	}

	sessionValidatorFunc := i.middleware.SessionValidatorHandler()
	scopeValidate := i.middleware.ScopeValidatorFactory()

	router = i.setupRouter(router, sessionValidatorFunc, scopeValidate)

	i.router = router
}

func (i *serverImpl) setupRouter(router *gin.Engine, sessionValidatorFunc gin.HandlerFunc, scopeValidate func([][]constant.Scope) gin.HandlerFunc) *gin.Engine {
	group := router.Group("/api")
	{
		group.GET("/session", sessionValidatorFunc, i.controller.SessionController().GetSession)
		group.GET("v1/deployment", i.controller.UtilityController().GetDeploymentInfo)

		utilityRouter := group.Group("/utilities")
		{
			utilityRouter.GET("health", i.controller.UtilityController().HealthCheck)
			utilityRouter.GET("countries", i.controller.UtilityController().GetCountries)
			utilityRouter.GET("states", i.controller.UtilityController().GetStates)
			utilityRouter.GET("timezones", i.controller.UtilityController().GetTimeZones)
			utilityRouter.GET("features/disabled", i.controller.UtilityController().GetDisabledFeatures)
		}

	}
	return router
}
