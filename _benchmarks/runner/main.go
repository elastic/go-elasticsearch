// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/montanaflynn/stats"
	"github.com/tidwall/gjson"

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

	rand.Seed(time.Now().UnixNano())

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

	dataSource := os.Getenv("DATA_SOURCE")
	if dataSource == "" {
		log.Fatal("ERROR: Required environment variable [DATA_SOURCE] empty")
	}
	if _, err := os.Stat(dataSource); os.IsNotExist(err) {
		log.Fatalf("ERROR: Data source at [%s] does not exist", dataSource)
	}

	filterOperations = os.Getenv("FILTER")

	runnerClientConfig := elasticsearch.Config{
		Addresses:    []string{targetURL},
		DisableRetry: true,
	}

	reportClientConfig := elasticsearch.Config{
		Addresses:            []string{reportURL},
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
	}

	operations := []struct {
		Action         string
		NumWarmups     int
		NumRepetitions int
		SetupFunc      runner.RunnerFunc
		RunnerFunc     runner.RunnerFunc
	}{
		// ----- Ping() -------------------------------------------------------------------------------
		{
			Action:         "ping",
			NumWarmups:     0,
			NumRepetitions: defaultRepetitions,
			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				res, err := c.RunnerClient.Ping()
				if err == nil && res != nil {
					res.Body.Close()
				}
				return res, err
			},
		},

		// ----- Info() -------------------------------------------------------------------------------
		{
			Action:         "info",
			NumWarmups:     0,
			NumRepetitions: defaultRepetitions,
			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				res, err := c.RunnerClient.Info()
				if err == nil && res != nil {
					res.Body.Close()
				}
				return res, err
			},
		},

		// ----- Get() --------------------------------------------------------------------------------
		{
			Action:         "get",
			NumWarmups:     100,
			NumRepetitions: defaultRepetitions,
			SetupFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName = "test-bench-get"
				)
				res, _ = c.RunnerClient.Indices.Delete([]string{indexName})
				if res != nil && res.Body != nil {
					res.Body.Close()
				}
				res, err = c.RunnerClient.Index(indexName, strings.NewReader(`{"title":"Test"}`), c.RunnerClient.Index.WithDocumentID("1"))
				if err != nil {
					return res, err
				}
				res.Body.Close()
				res, err = c.RunnerClient.Indices.Refresh(c.RunnerClient.Indices.Refresh.WithIndex(indexName))
				if err != nil {
					return res, err
				}
				res.Body.Close()
				return res, err
			},
			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var indexName = "test-bench-get"
				res, err := c.RunnerClient.Get(indexName, "1")
				if err != nil {
					return res, err
				}
				defer res.Body.Close()
				var b bytes.Buffer
				if _, err := b.ReadFrom(res.Body); err != nil {
					return nil, err
				}
				output := gjson.GetBytes(b.Bytes(), "_source.title")
				if output.Str != "Test" {
					return nil, fmt.Errorf("Unexpected output: %s", b.String())
				}
				return res, err
			},
		},

		// ----- Index() --------------------------------------------------------------------------------
		{
			Action:         "index",
			NumWarmups:     100,
			NumRepetitions: defaultRepetitions,
			SetupFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName = "test-bench-index"
				)
				res, _ = c.RunnerClient.Indices.Delete([]string{indexName})
				if res != nil && res.Body != nil {
					res.Body.Close()
				}
				res, err = c.RunnerClient.Indices.Create(indexName)
				if err != nil {
					return res, err
				}
				res.Body.Close()
				return res, err
			},
			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName = "test-bench-index"
				)

				docID := fmt.Sprintf("%04d-%04d", n, rand.Intn(defaultRepetitions))
				docBody, err := os.Open(filepath.Join(dataSource, "small", "document.json"))
				if err != nil {
					return nil, err
				}
				res, err = c.RunnerClient.Index(indexName, docBody, c.RunnerClient.Index.WithDocumentID(docID))
				if err != nil {
					return res, err
				}
				defer res.Body.Close()

				var b bytes.Buffer
				if _, err := b.ReadFrom(res.Body); err != nil {
					return nil, err
				}
				output := gjson.GetBytes(b.Bytes(), "result")
				if output.Str != "created" {
					return nil, fmt.Errorf("Unexpected output: %s", b.String())
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
			if err, ok := err.(*runner.Error); ok {
				if os.Getenv("DEBUG") != "" {
					log.Print("Error: ", err, ": ", err.Errs())
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
