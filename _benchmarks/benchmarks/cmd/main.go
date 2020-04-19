// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/montanaflynn/stats"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"

	"github.com/elastic/go-elasticsearch/v8/benchmarks"
	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"

	_ "github.com/elastic/go-elasticsearch/v8/benchmarks/operations"
)

var (
	red           = color.New(color.FgRed).SprintFunc()
	green         = color.New(color.FgGreen).SprintFunc()
	faint         = color.New(color.Faint).SprintFunc()
	bold          = color.New(color.Bold).SprintFunc()
	underline     = color.New(color.Underline).SprintFunc()
	boldUnderline = color.New(color.Bold).Add(color.Underline).SprintFunc()

	defaultRepetitions = 1000

	filterOperations string
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
	for k, _ := range benchmarks.Config {
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

	filterOperations = os.Getenv("FILTER")

	runnerClientConfig := elasticsearch.Config{
		Addresses:    []string{benchmarks.Config["ELASTICSEARCH_TARGET_URL"]},
		DisableRetry: true,
	}

	reportClientConfig := elasticsearch.Config{
		Addresses:            []string{benchmarks.Config["ELASTICSEARCH_REPORT_URL"]},
		MaxRetries:           10,
		EnableRetryOnTimeout: true,
	}
	if os.Getenv("DEBUG") != "" {
		runnerClientConfig.Logger = &estransport.ColorLogger{Output: os.Stdout}
		reportClientConfig.Logger = &estransport.ColorLogger{Output: os.Stdout}
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

	for _, operation := range benchmarks.Operations {
		if filterOperations != "" {
			if !strings.Contains(operation.Action, filterOperations) {
				continue
			}
		}

		runnerConfig.Action = operation.Action
		runnerConfig.NumRepetitions = operation.NumRepetitions
		runnerConfig.SetupFunc = operation.SetupFunc
		runnerConfig.RunnerFunc = operation.RunnerFunc

		if operation.NumOperations > 0 {
			runnerConfig.NumOperations = operation.NumOperations
		} else {
			runnerConfig.NumOperations = 1
		}

		if operation.Category != "" {
			runnerConfig.Category = operation.Category
		} else {
			runnerConfig.Category = os.Getenv("CLIENT_BENCHMARK_CATEGORY")
		}

		if operation.Environment != "" {
			runnerConfig.Environment = operation.Environment
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

		fmt.Fprintf(w, "  %-15s", fmt.Sprintf("[%s]", operation.Action))
		fmt.Fprintf(w, " %-9s", fmt.Sprintf("%d×", operation.NumRepetitions))
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
