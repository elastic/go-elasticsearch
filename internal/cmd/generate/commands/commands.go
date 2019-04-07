package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v8/internal/cmd/generate/utils"
)

var rootCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate allows you to generate APIs and tests",
	// Long:  "TODO",
}

// Execute launches the CLI application.
//
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.PrintErr(err)
		os.Exit(1)
	}
}

// RegisterCmd adds a command to rootCmd.
//
func RegisterCmd(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
