package cmd

import (
	"github.com/hiago-balbino/random-luck/handler"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A command to run the API",
	Run: func(_ *cobra.Command, _ []string) {
		server := handler.NewServer(handler.API)
		server.Start()
	},
}
