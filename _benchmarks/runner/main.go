// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/estransport"

	"github.com/elastic/go-elasticsearch/v8/_benchmarks/runner"
)

func main() {
	log.SetFlags(0)

	targetURL := os.Getenv("ELASTICSEARCH_TARGET_URL")
	if targetURL == "" {
		log.Fatal("ERROR: Required environment variable [ELASTICSEARCH_TARGET_URL] empty")
	}

	reportURL := os.Getenv("ELASTICSEARCH_REPORT_URL")
	if targetURL == "" {
		log.Fatal("ERROR: Required environment variable [ELASTICSEARCH_REPORT_URL] empty")
	}

	runnerClient, _ := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses:    []string{targetURL},
			DisableRetry: true,
		},
	)

	statsClient, _ := elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{reportURL},
			Logger:    &estransport.ColorLogger{Output: os.Stdout},
		},
	)

	runnerConfig := runner.Config{
		RunnerClient: runnerClient,
		StatsClient:  statsClient,

		NumWarmups:     1,
		NumRepetitions: 10,
		NumIterations:  100,
	}

	runnerConfig.Action = "info"

	runnerConfig.SetupFunc = func(c runner.Config) (*esapi.Response, error) {
		res, err := c.RunnerClient.Info()
		if err == nil && res != nil {
			res.Body.Close()
		}
		return res, err
	}
	runnerConfig.RunnerFunc = func(c runner.Config) (*esapi.Response, error) {
		res, err := c.RunnerClient.Info()
		res.Body.Close()
		return res, err
	}

	runner, err := runner.NewRunner(runnerConfig)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	if err := runner.Run(); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// log.Printf("[%s] %+v\n", runnerConfig.Action, runner.Stats())
}
