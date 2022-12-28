package cmd

import (
	"github.com/hiago-balbino/random-luck/handler"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "A command to run the Web",
	Run: func(_ *cobra.Command, _ []string) {
		server := handler.NewServer(handler.WEB)
		server.Start()
	},
}
