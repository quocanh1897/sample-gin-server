package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/quocanh1897/sample-gin-server/internal/config"
	"github.com/quocanh1897/sample-gin-server/internal/pkg/logger"
)

var rootCmd = &cobra.Command{
	Short: "Go-based API service for Sample Gin Server",
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	config.Load()
	logger.Init()

	rootCmd.AddCommand(serveCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
