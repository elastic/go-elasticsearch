// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8/internal/build/cmd"
	"github.com/elastic/go-elasticsearch/v8/internal/build/sdk/artifacts"
	"github.com/elastic/go-elasticsearch/v8/internal/build/sdk/ref"
	"github.com/elastic/go-elasticsearch/v8/internal/version"
	"github.com/spf13/cobra"
)

var (
	output     *string
	commitHash *string
	debug      *bool
	info       *bool
)

func init() {
	output = toolsCmd.Flags().StringP("output", "o", "", "Path to a folder for generated output")
	commitHash = toolsCmd.Flags().StringP("commit_hash", "c", "", "Elasticsearch commit hash (deprecated: no longer supported)")
	debug = toolsCmd.Flags().BoolP("debug", "d", false, "Print the generated source to terminal")
	info = toolsCmd.Flags().Bool("info", false, "Print the API details to terminal")

	cmd.RegisterCmd(toolsCmd)
}

var toolsCmd = &cobra.Command{
	Use:   "download-spec",
	Short: "Download specification artifact for code & tests generation",
	Run: func(cmd *cobra.Command, args []string) {
		command := &Command{
			Output:     *output,
			CommitHash: *commitHash,
			Debug:      *debug,
			Info:       *info,
		}
		err := command.Execute()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	},
}

type Command struct {
	Output     string
	CommitHash string
	Debug      bool
	Info       bool
}

// Execute downloads the rest-resources specification artifact using the artifacts API.
// It retrieves the spec for the version specified by ELASTICSEARCH_BUILD_VERSION env var,
// or falls back to the client version if not set.
func (c Command) Execute() error {
	// Warn if deprecated --commit_hash flag is used
	if c.CommitHash != "" {
		log.Printf("Warning: --commit_hash flag is deprecated and no longer supported. It will be ignored.")
	}

	esBuildVersion := os.Getenv("ELASTICSEARCH_BUILD_VERSION")
	if esBuildVersion == "" {
		esBuildVersion = version.Client
	}

	if c.Debug {
		log.Printf("Using version: %s", esBuildVersion)
	}

	// Parse the version/branch reference, renaming "main" to "master" for branch lookups
	r, err := ref.Parse(esBuildVersion, func(s string) string {
		if s == "main" {
			return "master"
		}
		return s
	})
	if err != nil {
		return fmt.Errorf("invalid version %q: %w", esBuildVersion, err)
	}

	client := artifacts.NewClient()
	if err := client.DownloadRestResources(context.Background(), r, c.Output); err != nil {
		return fmt.Errorf("failed to download rest-resources: %w", err)
	}

	if c.Debug {
		log.Printf("Successfully downloaded rest-resources to %s", c.Output)
	}

	return nil
}
