// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/montanaflynn/stats"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/estransport"

	"github.com/elastic/go-elasticsearch/v8/_benchmarks/runner"
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

	start := time.Now().UTC()

	log.Printf(boldUnderline("Running benchmarks for go-elasticsearch@%s; %s/go%s"), elasticsearch.Version, runner.OSFamily, runner.RuntimeVersion)

	targetURL := os.Getenv("ELASTICSEARCH_TARGET_URL")
	if targetURL == "" {
		log.Fatal("ERROR: Required environment variable [ELASTICSEARCH_TARGET_URL] empty")
	}

	reportURL := os.Getenv("ELASTICSEARCH_REPORT_URL")
	if targetURL == "" {
		log.Fatal("ERROR: Required environment variable [ELASTICSEARCH_REPORT_URL] empty")
	}

	filterOperations = os.Getenv("FILTER")

	runnerClient, _ := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses:    []string{targetURL},
			DisableRetry: true,
		},
	)

	statsClientConfig := elasticsearch.Config{
		Addresses:            []string{reportURL},
		MaxRetries:           10,
		EnableRetryOnTimeout: true,
	}
	if os.Getenv("DEBUG") != "" {
		statsClientConfig.Logger = &estransport.ColorLogger{Output: os.Stdout}
	}

	statsClient, _ := elasticsearch.NewClient(statsClientConfig)

	runnerConfig := runner.Config{
		RunnerClient: runnerClient,
		StatsClient:  statsClient,
	}

	operations := []struct {
		Action         string
		NumWarmups     int
		NumRepetitions int
		SetupFunc      runner.RunnerFunc
		RunnerFunc     runner.RunnerFunc
	}{
		{
			Action:         "ping",
			NumWarmups:     0,
			NumRepetitions: defaultRepetitions,
			RunnerFunc: func(c runner.Config) (*esapi.Response, error) {
				res, err := c.RunnerClient.Ping()
				if err == nil && res != nil {
					res.Body.Close()
				}
				return res, err
			},
		},

		{
			Action:         "info",
			NumWarmups:     0,
			NumRepetitions: defaultRepetitions,
			RunnerFunc: func(c runner.Config) (*esapi.Response, error) {
				res, err := c.RunnerClient.Info()
				if err == nil && res != nil {
					res.Body.Close()
				}
				return res, err
			},
		},
	}

	for _, operation := range operations {
		if filterOperations != "" {
			if !strings.Contains(filterOperations, operation.Action) {
				continue
			}
		}

		runnerConfig.Action = operation.Action
		runnerConfig.NumRepetitions = operation.NumRepetitions
		runnerConfig.SetupFunc = operation.SetupFunc
		runnerConfig.RunnerFunc = operation.RunnerFunc

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

		fmt.Fprintf(w, "  [%s]", operation.Action)
		fmt.Fprintf(w, " %d×", operation.NumRepetitions)
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
			if err, ok := err.(*runner.Error); ok {
				if os.Getenv("DEBUG") != "" {
					log.Print("Error: ", err, err.Errs())
				} else {
					log.Print("Error: ", err)
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
