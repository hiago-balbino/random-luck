package cmd

import (
	"github.com/hiago-balbino/random-luck/configuration"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "A main CLI to random luck",
	Long:  "CLI used to create a random luck numbers based on quantity numbers to Mega Sena",
}

// Execute executes the root command.
func Execute() error {
	cobra.OnInitialize(configuration.InitConfigurations)
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(webCmd)

	return rootCmd.Execute()
}
