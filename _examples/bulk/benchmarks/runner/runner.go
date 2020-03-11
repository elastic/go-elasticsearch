// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package runner

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/montanaflynn/stats"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

// NewRunner returns new BulkIndexer benchmarking runner.
//
func NewRunner(cfg Config) (*Runner, error) {
	return &Runner{config: cfg}, nil
}

// Runner represents the BulkIndexer benchmarking runner.
//
type Runner struct {
	config Config

	doc           []byte
	indexSettings strings.Builder

	samples    []float64
	throughput map[string]float64
}

// Config represents configuration for Runner.
//
type Config struct {
	IndexName   string
	DatasetName string
	Client      *elasticsearch.Client
	Decoder     esutil.BulkResponseJSONDecoder

	NumShards     int
	NumReplicas   int
	NumItems      int
	NumRuns       int
	NumWarmupRuns int
	NumWorkers    int
	FlushBytes    int
	Wait          time.Duration
	Mockserver    bool
}

// Report prints statistics from the benchmark runs.
//
func (r *Runner) Report() error {
	r.throughput = map[string]float64{
		"min": func() float64 { v, _ := stats.Min(r.samples); return v }(),
		"max": func() float64 { v, _ := stats.Max(r.samples); return v }(),
		"mdn": func() float64 { v, _ := stats.Median(r.samples); return v }(),
		"p25": func() float64 { v, _ := stats.Percentile(r.samples, 25); return v }(),
		"p50": func() float64 { v, _ := stats.Percentile(r.samples, 50); return v }(),
		"p75": func() float64 { v, _ := stats.Percentile(r.samples, 75); return v }(),
		"p95": func() float64 { v, _ := stats.Percentile(r.samples, 95); return v }(),
	}

	log.Printf(
		"docs/sec: min [%s] max [%s] mean [%s]",
		humanize.Comma(int64(r.throughput["min"])),
		humanize.Comma(int64(r.throughput["max"])),
		humanize.Comma(int64(r.throughput["mdn"])),
	)

	ratio := 50.0 / r.throughput["max"]

	if !math.IsNaN(r.throughput["p25"]) {
		fmt.Println("25%",
			strings.Repeat("▆", int(r.throughput["p25"]*ratio)),
			humanize.Comma(int64(r.throughput["p25"])), "docs/sec")
	}
	if !math.IsNaN(r.throughput["p50"]) {
		fmt.Println("50%",
			strings.Repeat("▆", int(r.throughput["p50"]*ratio)),
			humanize.Comma(int64(r.throughput["p50"])), "docs/sec")
	}
	if !math.IsNaN(r.throughput["p75"]) {
		fmt.Println("75%",
			strings.Repeat("▆", int(r.throughput["p75"]*ratio)),
			humanize.Comma(int64(r.throughput["p75"])), "docs/sec")
	}
	if !math.IsNaN(r.throughput["p95"]) {
		fmt.Println("95%",
			strings.Repeat("▆", int(r.throughput["p95"]*ratio)),
			humanize.Comma(int64(r.throughput["p95"])), "docs/sec")
	}

	return nil
}

// Run executes the benchmark runs.
//
func (r *Runner) Run() error {
	for n := 1; n <= r.config.NumWarmupRuns; n++ {
		if err := r.run(n, false); err != nil {
			log.Fatalf("Runner error: %s", err)
		}
	}

	for n := 1; n <= r.config.NumRuns; n++ {
		if err := r.run(n, true); err != nil {
			log.Fatalf("Runner error: %s", err)
		}
	}

	return nil
}

// setup re-creates the index for a benchmark run.
//
func (r *Runner) setup() error {
	fm, err := os.Open(filepath.Join("data", r.config.DatasetName, "mapping.json"))
	if err != nil {
		return fmt.Errorf("setup: reading mapping: %s", err)
	}

	var mappingEnvelope map[string]interface{}
	json.NewDecoder(fm).Decode(&mappingEnvelope)
	mapping, err := json.Marshal(mappingEnvelope["mappings"])
	if err != nil {
		return fmt.Errorf("setup: encoding mapping: %s", err)
	}

	r.indexSettings.WriteString(`{ "settings": `)
	fmt.Fprintf(&r.indexSettings, `{"number_of_shards": %d, "number_of_replicas": %d, "refresh_interval":"5s"}`, r.config.NumShards, r.config.NumReplicas)
	r.indexSettings.WriteString(`, "mappings":`)
	r.indexSettings.Write(mapping)
	r.indexSettings.WriteString(`}`)

	f, err := os.Open(filepath.Join("data", r.config.DatasetName, "document.json"))
	if err != nil {
		return fmt.Errorf("setup: reading document: %s", err)
	}
	var m map[string]interface{}
	json.NewDecoder(f).Decode(&m)
	doc, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("setup: encoding document: %s", err)
	}
	r.doc = doc

	r.config.Client.Indices.Delete(
		[]string{r.config.IndexName},
		r.config.Client.Indices.Delete.WithIgnoreUnavailable(true))
	res, err := r.config.Client.Indices.Create(
		r.config.IndexName,
		r.config.Client.Indices.Create.WithBody(strings.NewReader(r.indexSettings.String())),
		r.config.Client.Indices.Create.WithWaitForActiveShards("1"))
	if err != nil {
		return fmt.Errorf("setup: encoding document: %s", err)
	}
	res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("setup: encoding document: %s", res.String())
	}

	return nil
}

// run executes a single benchmark run n, recording stats when measure is true.
//
func (r *Runner) run(n int, measure bool) error {
	if err := r.setup(); err != nil {
		return fmt.Errorf("run: %s", err)
	}

	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:         r.config.IndexName,
		Client:        r.config.Client,
		Decoder:       r.config.Decoder,
		NumWorkers:    r.config.NumWorkers,
		FlushBytes:    r.config.FlushBytes,
		FlushInterval: time.Hour, // Disable automatic flushing
	})
	if err != nil {
		return fmt.Errorf("run: %s", err)
	}

	start := time.Now().UTC()
	for i := 1; i <= r.config.NumItems; i++ {
		err := bi.Add(context.Background(), esutil.BulkIndexerItem{
			Action: "index",
			Body:   bytes.NewReader(r.doc),
			OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
				if err != nil {
					log.Printf("ERROR: %s", err)
				} else {
					log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
				}
			},
		})
		if err != nil {
			return fmt.Errorf("run: %s", err)
		}
	}

	if err := bi.Close(context.Background()); err != nil {
		return fmt.Errorf("run: %s", err)
	}

	duration := time.Since(start)

	if measure {
		biStats := bi.Stats()

		var numThroughput uint64
		if r.config.Mockserver {
			numThroughput = biStats.NumAdded
		} else {
			numThroughput = biStats.NumFlushed
		}
		sample := 1000.0 / float64(duration/time.Millisecond) * float64(numThroughput)
		r.samples = append(r.samples, sample)

		log.Printf("%4d) add=%s\tflush=%s\tfail=%s\treqs=%s\tdur=%-6s\t%6s docs/sec\n",
			n,
			formatInt(int(biStats.NumAdded)),
			formatInt(int(biStats.NumFlushed)),
			formatInt(int(biStats.NumFailed)),
			formatInt(int(biStats.NumRequests)),
			duration.Truncate(10*time.Millisecond),
			humanize.Comma(int64(sample)))

		time.Sleep(r.config.Wait)
	}

	return nil
}

// formatInt returns a number like 123456 as a string like 123.45K
//
func formatInt(i int) string {
	return strings.ReplaceAll(strings.ToUpper(humanize.SIWithDigits(float64(i), 2, "")), " ", "")
}
