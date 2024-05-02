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
	"github.com/quocanh1897/sample-gin-server/internal/api/controller"
	"github.com/quocanh1897/sample-gin-server/internal/config"
	"github.com/quocanh1897/sample-gin-server/internal/constant"
	"github.com/quocanh1897/sample-gin-server/internal/middleware"
	appvalidator "github.com/quocanh1897/sample-gin-server/internal/validator"
)

//go:generate mockery --name=Server --case=snake
type Server interface {
	Start() error
}

func New(controller controller.Controller) Server {
	cfg := config.GetAppConfig()
	server := &serverImpl{
		serverConfig: &cfg,
		controller:   controller,
		middleware:   middleware.New(&cfg),
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
	}))

	router.Use(i.middleware.LoggerHandler())
	router.Use(i.middleware.ErrorHandlerHandler())
	router.Use(i.middleware.ContextHandler())
	router.Use(gin.Recovery())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("is_uuid", appvalidator.IsUUID)
	}

	scopeValidate := i.middleware.ScopeValidatorFactory()

	router = i.setupRouter(router, scopeValidate)

	i.router = router
}

func (i *serverImpl) setupRouter(router *gin.Engine, scopeValidate func([][]constant.Scope) gin.HandlerFunc) *gin.Engine {
	group := router.Group("/api")
	{
		utilityRouter := group.Group("/utilities")
		{
			utilityRouter.GET("health", scopeValidate([][]constant.Scope{{constant.AgencyAnyRead}}), i.controller.UtilityController().HealthCheck)
			utilityRouter.GET("timezones", i.controller.UtilityController().GetTimeZones)
		}
	}
	return router
}
