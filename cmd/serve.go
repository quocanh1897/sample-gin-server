package cmd

import (
	"github.com/quocanh1897/sample-gin-server/internal/api/controller"
	"github.com/quocanh1897/sample-gin-server/internal/http/server"
	sessionservice "github.com/quocanh1897/sample-gin-server/internal/service/session"
	utilityservice "github.com/quocanh1897/sample-gin-server/internal/service/utility"
	"github.com/spf13/cobra"
	"log"
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
	utilityService := utilityservice.NewUtilityService()
	sessionService := sessionservice.NewSessionService()

	utilityController := controller.NewUtilityController(utilityService)
	sessionController := controller.NewSessionController(sessionService)

	return server.New(
		controller.New(
			utilityController,
			sessionController,
		),
		sessionService,
	)
}
