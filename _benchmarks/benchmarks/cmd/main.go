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

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/montanaflynn/stats"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"

	"github.com/elastic/go-elasticsearch/v8/benchmarks"
	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"

	_ "github.com/elastic/go-elasticsearch/v8/benchmarks/actions"
)

var (
	red           = color.New(color.FgRed).SprintFunc()
	green         = color.New(color.FgGreen).SprintFunc()
	faint         = color.New(color.Faint).SprintFunc()
	bold          = color.New(color.Bold).SprintFunc()
	underline     = color.New(color.Underline).SprintFunc()
	boldUnderline = color.New(color.Bold).Add(color.Underline).SprintFunc()

	defaultRepetitions = 1000

	filterActions string
)

func main() {
	log.SetFlags(0)

	rand.Seed(time.Now().UnixNano())

	start := time.Now().UTC()
	benchmarks.Config = map[string]string{
		"BUILD_ID":                     "",
		"DATA_SOURCE":                  "",
		"CLIENT_BRANCH":                "",
		"CLIENT_COMMIT":                "",
		"CLIENT_BENCHMARK_ENVIRONMENT": "",
		"ELASTICSEARCH_TARGET_URL":     "",
		"ELASTICSEARCH_REPORT_URL":     "",
		"TARGET_SERVICE_TYPE":          "",
		"TARGET_SERVICE_NAME":          "",
		"TARGET_SERVICE_VERSION":       "",
		"TARGET_SERVICE_OS_FAMILY":     "",
	}

	log.Printf(boldUnderline("Running benchmarks for go-elasticsearch@%s; %s/go%s"), elasticsearch.Version, runner.RuntimeOS, runner.RuntimeVersion)

	var missingConfigs []string
	for k := range benchmarks.Config {
		v := os.Getenv(k)
		if v == "" {
			missingConfigs = append(missingConfigs, k)
		} else {
			benchmarks.Config[k] = v
		}
	}

	if len(missingConfigs) > 0 {
		log.Fatalf("ERROR: Required environment variables empty: %s", strings.Join(missingConfigs, ", "))
	}

	if _, err := os.Stat(benchmarks.Config["DATA_SOURCE"]); os.IsNotExist(err) {
		log.Fatalf("ERROR: Data source at [%s] does not exist", benchmarks.Config["DATA_SOURCE"])
	}

	dirs, err := ioutil.ReadDir(benchmarks.Config["DATA_SOURCE"])
	if err != nil {
		log.Fatalf("ERROR: Unable to list files in [%s]", benchmarks.Config["DATA_SOURCE"])
	}
	for _, file := range dirs {
		if !file.IsDir() {
			continue
		}
		c, err := os.Open(filepath.Join(benchmarks.Config["DATA_SOURCE"], file.Name(), "document.json"))
		if err != nil {
			log.Fatalf("ERROR: Unable to open file: %s", err)
		}
		var b bytes.Buffer
		b.ReadFrom(c)
		c.Close()

		benchmarks.DataSources[file.Name()] = &b
	}

	filterActions = os.Getenv("FILTER")

	runnerClientConfig := elasticsearch.Config{
		Addresses:    []string{benchmarks.Config["ELASTICSEARCH_TARGET_URL"]},
		DisableRetry: true,
	}

	reportClientConfig := elasticsearch.Config{
		Addresses:  []string{benchmarks.Config["ELASTICSEARCH_REPORT_URL"]},
		MaxRetries: 10,
	}
	if os.Getenv("DEBUG") != "" {
		runnerClientConfig.Logger = &elastictransport.ColorLogger{Output: os.Stdout}
		reportClientConfig.Logger = &elastictransport.ColorLogger{Output: os.Stdout}
	}

	runnerClient, _ := elasticsearch.NewClient(runnerClientConfig)
	reportClient, _ := elasticsearch.NewClient(reportClientConfig)

	runnerConfig := runner.Config{
		RunnerClient: runnerClient,
		ReportClient: reportClient,

		BuildID:     benchmarks.Config["BUILD_ID"],
		Category:    os.Getenv("CLIENT_BENCHMARK_CATEGORY"),
		Environment: benchmarks.Config["CLIENT_BENCHMARK_ENVIRONMENT"],
		Target: struct {
			OS      runner.ConfigOS
			Service runner.ConfigService
		}{
			OS: runner.ConfigOS{Family: benchmarks.Config["TARGET_SERVICE_OS_FAMILY"]},
			Service: runner.ConfigService{
				Type:    benchmarks.Config["TARGET_SERVICE_TYPE"],
				Name:    benchmarks.Config["TARGET_SERVICE_NAME"],
				Version: benchmarks.Config["TARGET_SERVICE_VERSION"],
				Git: runner.ConfigGit{
					Branch: os.Getenv("TARGET_SERVICE_GIT_BRANCH"),
					Commit: os.Getenv("TARGET_SERVICE_GIT_COMMIT"),
				},
			},
		},
		Runner: struct {
			Service runner.ConfigService
		}{
			Service: runner.ConfigService{
				Git: runner.ConfigGit{Branch: benchmarks.Config["CLIENT_BRANCH"], Commit: benchmarks.Config["CLIENT_COMMIT"]},
			},
		},
	}

	for _, action := range benchmarks.Actions {
		if filterActions != "" {
			if !strings.Contains(action.Name, filterActions) {
				continue
			}
		}

		runnerConfig.Action = action.Name
		runnerConfig.NumRepetitions = action.NumRepetitions
		runnerConfig.SetupFunc = action.SetupFunc
		runnerConfig.RunnerFunc = action.RunnerFunc

		if action.NumOperations > 0 {
			runnerConfig.NumOperations = action.NumOperations
		} else {
			runnerConfig.NumOperations = 1
		}

		if action.Category != "" {
			runnerConfig.Category = action.Category
		} else {
			runnerConfig.Category = os.Getenv("CLIENT_BENCHMARK_CATEGORY")
		}

		if action.Environment != "" {
			runnerConfig.Environment = action.Environment
		} else {
			runnerConfig.Environment = os.Getenv("CLIENT_BENCHMARK_ENVIRONMENT")
		}

		run, err := runner.NewRunner(runnerConfig)
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}

		err = run.Run()

		var (
			w = log.Writer()

			samples []float64
			mean    time.Duration
		)

		for _, s := range run.Stats() {
			samples = append(samples, float64(s.Duration))
		}
		mean = func() time.Duration { v, _ := stats.Mean(samples); return time.Duration(v).Truncate(time.Millisecond) }()

		fmt.Fprintf(w, "  %-15s", fmt.Sprintf("[%s]", action.Name))
		fmt.Fprintf(w, " %-9s", fmt.Sprintf("%d×", action.NumRepetitions))
		fmt.Fprintf(w, " "+faint("mean=")+"%s", mean)
		fmt.Fprintf(w, " "+faint("runner=")+"%s", func() string {
			if len(run.Stats()) < 1 {
				return red("failure")
			}
			for _, s := range run.Stats() {
				if s.Outcome == "failure" {
					return red("failure")
				}
			}
			return green("success")
		}())
		fmt.Fprintf(w, " "+faint("report=")+"%s", func() string {
			if err != nil || len(run.Errs()) > 0 {
				return red("failure")
			}
			return green("success")
		}())
		fmt.Fprintf(w, "\n")

		if err != nil {
			if runnerErr, ok := err.(*runner.Error); ok {
				if os.Getenv("DEBUG") != "" {
					log.Print("Error: ", runnerErr, ": ", runnerErr.Errs())
				} else {
					log.Print("Error: ", runnerErr)
				}
			} else {
				log.Print("Error: ", err)
			}
		}

		if len(run.Errs()) > 0 {
			log.Printf("Errors: %q", run.Errs())
		}
	}

	log.Printf(underline("Finished in %s"), time.Since(start).Truncate(time.Millisecond))
}
