package cmd

import (
	"fmt"
	"github.com/quocanh1897/sample-go-server/internal/http/server"
	"log"
	"log/slog"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start HTTP server on the predefined port",
	Run: func(_ *cobra.Command, _ []string) {
		s := setupServer()
		if err := s.Start(); err != nil {
			log.Fatalf("Failed to start HTTP server with error: %s", err.Error())
		}
	},
}

func setupServer() server.Server {
	appConfig := config.GetAppConfig()
	discovery, err := service.NewDiscovery(
		appConfig.ZooKeeper,
	)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create discovery, reason: %s", err))
		panic(err)
	}
	komradeRepository := komrade.NewKomradeRepository(appConfig.Komrade, discovery)
	userRepository := user.NewUserRepository(&appConfig.User)
	agencyRepository := agency.NewAgencyRepository(&appConfig.Agency)
	applicationRepository := application.NewApplicationRepository(&appConfig.Application)
	auditRepository := audit.NewAuditRepository(appConfig.Audit, discovery)
	solrRepository := solr.NewSolrRepository(&appConfig.Solr)
	mfaVerifierRepository := mfaverifier.NewMfaVerifierRepository(&appConfig.MfaVerifier)
	oAuthRepository := oauth.NewOAuthRepository(&appConfig.OAuth)
	arsenalRepository := arsenal.NewArsenalRepository(appConfig.Arsenal, discovery)
	cacheRepository := cache.NewCacheRepository(&appConfig.Redis, &appConfig.Tls)
	terminatorRepository := terminator.NewTerminatorRepository(appConfig.Terminator, discovery)
	reporterRepository := reporter.NewReporterRepository(appConfig.Reporter, discovery)
	arkhamRepository := arkham.NewArkhamRepository(appConfig.Arkham, discovery)
	sherlockRepository := sherlock.NewSherlockRepository(appConfig.Sherlock, discovery)
	jarvisEndpointRepository := javisendpoint.NewJarvisEndpointClient(appConfig.JarvisEndpoint, discovery)
	ecomsaasRepository := ecomsaas.NewEcomsaasRepository(appConfig.Ecomsaas)
	dreddRepository := dredd.NewDreddRepository(appConfig.Dredd, discovery)
	hrvRepository, err := httprequestvalidator.NewHRVRepository(appConfig.HRVConfig)
	cartographerRepository := cartographer.NewCartographerRepository(&appConfig.Cartographer)
	systemStatusRepository := systemstatus.NewSystemStatusRepository(&appConfig.SystemStatus)

	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create HRV repository, reason: %s", err))
		panic(err)
	}

	utilityUseCase := utilityusecase.NewUtilityUseCase(
		komradeRepository,
		userRepository,
		agencyRepository,
		hrvRepository,
		auditRepository,
		solrRepository,
		cacheRepository,
		oAuthRepository,
		terminatorRepository,
		arkhamRepository,
		sherlockRepository,
		jarvisEndpointRepository,
		ecomsaasRepository,
		dreddRepository,
		cartographerRepository,
		systemStatusRepository,
	)
	agencyUseCase := agencyusecase.NewAgencyUseCase(agencyRepository, solrRepository, komradeRepository, mfaVerifierRepository, arsenalRepository, terminatorRepository, reporterRepository, ecomsaasRepository, jarvisEndpointRepository, dreddRepository, cartographerRepository, sherlockRepository)
	applicationUseCase := applicationusecase.NewApplicationUseCase(applicationRepository)
	userUseCase := userusecase.NewUserUseCase(userRepository, solrRepository, komradeRepository, mfaVerifierRepository, reporterRepository, ecomsaasRepository, agencyRepository)
	oAuthUseCase := oauthusecase.NewOAuthUseCase(oAuthRepository, cacheRepository)
	sessionUseCase := sessionusecase.NewSessionUseCase(komradeRepository, hrvRepository)
	reportUseCase := reportusecase.NewReportUseCase(reporterRepository)
	announcementUseCase := announcementusecase.NewAnnouncementUseCase(agencyRepository)
	cadRmsIntegrationUseCase := cadrmsintegrationusecase.NewCadRmsIntegrationUseCase(jarvisEndpointRepository, dreddRepository)
	deviceUseCase := deviceusecase.NewDeviceUseCase(solrRepository, arsenalRepository, komradeRepository, auditRepository, arkhamRepository, sherlockRepository)
	profileUseCase := profileusecase.NewProfileUseCase(mfaVerifierRepository, ecomsaasRepository)
	softwareUseCase := softwareusecase.NewSoftwareUseCase(ecomsaasRepository)
	systemStatusUseCase := systemstatususecase.NewSystemStatusUseCase(systemStatusRepository)

	utilityController := controller.NewUtilityController(utilityUseCase)
	agencyController := controller.NewAgencyController(agencyUseCase)
	applicationController := controller.NewApplicationController(applicationUseCase)
	oAuthController := controller.NewOAuthController(sessionUseCase, oAuthUseCase)
	sessionController := controller.NewSessionController(sessionUseCase)
	userController := controller.NewUserController(userUseCase)
	reportController := controller.NewReportController(reportUseCase)
	announcementController := controller.NewAnnouncementController(announcementUseCase)
	deviceController := controller.NewDeviceController(deviceUseCase)
	cadRmsIntegrationController := controller.NewCadRmsIntegrationController(cadRmsIntegrationUseCase)
	profileController := controller.NewProfileController(profileUseCase, userUseCase, agencyUseCase)
	softwareController := controller.NewSoftwareController(softwareUseCase)
	systemStatusController := controller.NewSystemStatusController(systemStatusUseCase)

	return server.New(
		controller.New(
			utilityController,
			agencyController,
			applicationController,
			oAuthController,
			sessionController,
			userController,
			reportController,
			deviceController,
			announcementController,
			cadRmsIntegrationController,
			profileController,
			softwareController,
			systemStatusController,
		),
		sessionUseCase,
		komradeRepository,
	)
}
