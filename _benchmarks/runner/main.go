// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/elastic/go-elasticsearch/v8/_benchmarks/runner"
)

func main() {
	log.SetFlags(0)

	runnerClient, _ := elasticsearch.NewDefaultClient()
	statsClient, _ := elasticsearch.NewDefaultClient()

	runnerConfig := runner.Config{
		RunnerClient: runnerClient,
		StatsClient:  statsClient,

		NumWarmups:     1,
		NumRepetitions: 5,
		NumIterations:  10,
	}

	runnerConfig.Action = "info"

	runnerConfig.SetupFunc = func(c runner.Config) error {
		res, _ := c.RunnerClient.Info()
		res.Body.Close()
		return nil
	}
	runnerConfig.RunnerFunc = func(c runner.Config) error {
		res, _ := c.RunnerClient.Info()
		res.Body.Close()
		return nil
	}

	runner, err := runner.NewRunner(runnerConfig)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	if err := runner.Run(); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	log.Printf("[%s] %+v\n", runnerConfig.Action, runner.Stats())
}
