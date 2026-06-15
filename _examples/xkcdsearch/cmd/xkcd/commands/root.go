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

package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	// IndexName is the Elasticsearch index name.
	IndexName string

	tWidth int
)

var rootCmd = &cobra.Command{
	Use:   "xkcd",
	Short: "xkcd allows you to index and search xkcd.com",
	// Long:  "TODO",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&IndexName, "index", "i", "xkcd", "Index name")
	tWidth, _, _ = terminal.GetSize(int(os.Stdout.Fd()))
}

// Execute launches the CLI application.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
