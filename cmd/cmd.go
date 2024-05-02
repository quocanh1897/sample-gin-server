package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"git.taservs.net/platform/edca-api/internal/config"
	"git.taservs.net/platform/edca-api/internal/pkg/logger"
)

var rootCmd = &cobra.Command{
	Short: "Go-based API service for EDCA V2",
	Run: func(cmd *cobra.Command, _ []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

func init() {
	config.Load()
	logger.Init()

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(debugCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
