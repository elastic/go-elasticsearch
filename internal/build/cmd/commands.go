// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/elastic/go-elasticsearch/v7/internal/build/utils"
)

var rootCmd = &cobra.Command{
	Use:   "build",
	Short: "Build tools for Elasticsearch client, helpers for CI, etc...",
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
