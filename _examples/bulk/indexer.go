// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build ignore

// This example demonstrates indexing documents using the esutil.BulkIndexer helper.
//
// You can configure the settings with command line flags:
//
//     go run indexer.go --dataset=httplog --runs=10 --count=1000000 --shards=3 --replicas=1 --flush=1000000

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
	"time"

	"github.com/montanaflynn/stats"

	"github.com/dustin/go-humanize"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"
	"github.com/elastic/go-elasticsearch/v8/esutil"

	"github.com/elastic/go-elasticsearch/v8/_examples/bulk/model"
)

var (
	indexName   string
	datasetName string
	numWorkers  int
	flushBytes  int
	numRuns     int
	numItems    int
	numShards   int
	numReplicas int
	debug       bool
)

func init() {
	flag.StringVar(&indexName, "index", "test-bulk-example", "Index name")
	flag.StringVar(&datasetName, "dataset", "small", "Dataset to use for indexing")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of indexer workers")
	flag.IntVar(&flushBytes, "flush", 5e+6, "Flush threshold in bytes")
	flag.IntVar(&numRuns, "runs", 10, "Number of runs")
	flag.IntVar(&numItems, "count", 1000000, "Number of documents to generate")
	flag.IntVar(&numShards, "shards", 3, "Number of index shards")
	flag.IntVar(&numReplicas, "replicas", 0, "Number of index replicas")
	flag.BoolVar(&debug, "debug", false, "Enable logging output")
	flag.Parse()
}

func main() {
	log.SetFlags(0)

	var (
		countSuccessful uint64
		samples         []float64
		throughput      map[string]float64
		indexSettings   strings.Builder
		indexName       = indexName + "-" + datasetName
	)

	cfg := elasticsearch.Config{}
	cfg.Transport = &fasthttpTransport{}
	if debug {
		cfg.Logger = &estransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true}
	}
	es, _ := elasticsearch.NewClient(cfg)

	fm, err := os.Open(filepath.Join("data", datasetName, "mapping.json"))
	if err != nil {
		log.Fatalf("Error reading mapping file: %s", err)
	}
	var mappingEnvelope map[string]interface{}
	json.NewDecoder(fm).Decode(&mappingEnvelope)
	mapping, err := json.Marshal(mappingEnvelope["mappings"])
	if err != nil {
		log.Fatalf("Error encoding mapping: %s", err)
	}

	indexSettings.WriteString(`{ "settings": `)
	fmt.Fprintf(&indexSettings, `{"number_of_shards": %d, "number_of_replicas": %d, "refresh_interval":"5s"}`, numShards, numReplicas)
	indexSettings.WriteString(`, "mappings":`)
	indexSettings.Write(mapping)
	indexSettings.WriteString(`}`)

	log.Printf(
		"%s: [%d] shards [%d] replicas [%d] workers [%s] flush threshold",
		datasetName, numShards, numReplicas, numWorkers, humanize.Bytes(uint64(flushBytes)))
	log.Println(strings.Repeat("▔", 85))

	f, err := os.Open(filepath.Join("data", datasetName, "document.json"))
	if err != nil {
		log.Fatalf("Error reading test document file: %s", err)
	}
	var m map[string]interface{}
	json.NewDecoder(f).Decode(&m)
	doc, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error encoding test document: %s", err)
	}

	es.Indices.Delete([]string{indexName}, es.Indices.Delete.WithIgnoreUnavailable(true))
	res, err := es.Indices.Create(
		indexName,
		es.Indices.Create.WithBody(strings.NewReader(indexSettings.String())),
		es.Indices.Create.WithWaitForActiveShards("1"))
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
	res.Body.Close()
	if res.IsError() {
		log.Fatalf("Error creating index: %s", res.String())
	}

	// TODO(karmi): Run warmup before measuring
	for n := 1; n <= numRuns; n++ {
		bi, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:      indexName,
			Client:     es,
			Decoder:    easyjsonDecoder{},
			NumWorkers: numWorkers,
			FlushBytes: int(flushBytes),
		})

		start := time.Now().UTC()
		for i := 1; i <= numItems; i++ {
			err := bi.Add(context.Background(), esutil.BulkIndexerItem{
				Action: "index",
				Body:   bytes.NewReader(doc),
				OnSuccess: func(item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},
			})
			if err != nil {
				log.Fatalf("Unexpected error: %s", err)
			}
		}

		if err := bi.Close(context.Background()); err != nil {
			log.Fatalf("Unexpected error: %s", err)
		}

		biStats := bi.Stats()

		sample := 1000.0 / float64(time.Since(start)/time.Millisecond) * float64(biStats.NumFlushed)
		samples = append(samples, sample)

		if biStats.NumAdded != uint(numItems) {
			log.Fatalf("Unexpected NumAdded: %d", biStats.NumAdded)
		}

		if biStats.NumFailed != 0 {
			log.Fatalf("Unexpected NumFailed: %d", biStats.NumFailed)
		}

		if countSuccessful != uint64(n*numItems) {
			log.Fatalf("Unexpected countSuccessful: %d", countSuccessful)
		}

		log.Printf("%4d) added=%s flushed=%s failed=%s duration=%-8s throughput=%s docs/sec\n",
			n,
			humanize.Comma(int64(biStats.NumAdded)),
			humanize.Comma(int64(biStats.NumFlushed)),
			humanize.Comma(int64(biStats.NumFailed)),
			time.Since(start).Truncate(10*time.Millisecond),
			humanize.Comma(int64(sample)))
	}

	throughput = map[string]float64{
		"min": func() float64 { v, _ := stats.Min(samples); return v }(),
		"max": func() float64 { v, _ := stats.Max(samples); return v }(),
		"mdn": func() float64 { v, _ := stats.Median(samples); return v }(),
		"p25": func() float64 { v, _ := stats.Percentile(samples, 25); return v }(),
		"p50": func() float64 { v, _ := stats.Percentile(samples, 50); return v }(),
		"p75": func() float64 { v, _ := stats.Percentile(samples, 75); return v }(),
		"p95": func() float64 { v, _ := stats.Percentile(samples, 95); return v }(),
	}

	log.Println(strings.Repeat("▁", 85))
	log.Printf(
		"docs/sec: min [%s] max [%s] mean [%s]",
		humanize.Comma(int64(throughput["min"])),
		humanize.Comma(int64(throughput["max"])),
		humanize.Comma(int64(throughput["mdn"])),
	)

	ratio := 50.0 / throughput["max"]

	fmt.Println("25%",
		strings.Repeat("▆", int(throughput["p25"]*ratio)),
		humanize.Comma(int64(throughput["p25"])), "docs/sec")
	fmt.Println("50%",
		strings.Repeat("▆", int(throughput["p50"]*ratio)),
		humanize.Comma(int64(throughput["p50"])), "docs/sec")
	fmt.Println("75%",
		strings.Repeat("▆", int(throughput["p75"]*ratio)),
		humanize.Comma(int64(throughput["p75"])), "docs/sec")
	fmt.Println("95%",
		strings.Repeat("▆", int(throughput["p95"]*ratio)),
		humanize.Comma(int64(throughput["p95"])), "docs/sec")
}

// easyjsonDecoder implements a JSON decoder for the indexer
// via the "github.com/mailru/easyjson" package.
// See _examples/encoding for a demo.

type easyjsonDecoder struct{}

func (d easyjsonDecoder) UnmarshalFromReader(r io.Reader, blk *esutil.BulkIndexerResponse) error {
	var v model.BulkIndexerResponse
	if err := easyjson.UnmarshalFromReader(r, &v); err != nil {
		return err
	}
	blk.Took = v.Took
	blk.HasErrors = v.HasErrors
	blk.Items = v.Items

	return nil
}

// fasthttpTransport implements HTTP transport for the Elasticsearch client
// via the "github.com/valyala/fasthttp" package.
// See _examples/fasthttp for a demo.

type fasthttpTransport struct{}

func (t *fasthttpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	freq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(freq)

	fres := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(fres)

	t.copyRequest(freq, req)

	err := fasthttp.Do(freq, fres)
	if err != nil {
		return nil, err
	}

	res := &http.Response{Header: make(http.Header)}
	t.copyResponse(res, fres)

	return res, nil
}

func (t *fasthttpTransport) copyRequest(dst *fasthttp.Request, src *http.Request) *fasthttp.Request {
	if src.Method == "GET" && src.Body != nil {
		src.Method = "POST"
	}

	dst.SetHost(src.Host)
	dst.SetRequestURI(src.URL.String())

	dst.Header.SetRequestURI(src.URL.String())
	dst.Header.SetMethod(src.Method)

	for k, vv := range src.Header {
		for _, v := range vv {
			dst.Header.Set(k, v)
		}
	}

	if src.Body != nil {
		dst.SetBodyStream(src.Body, -1)
	}

	return dst
}

func (t *fasthttpTransport) copyResponse(dst *http.Response, src *fasthttp.Response) *http.Response {
	dst.StatusCode = src.StatusCode()

	src.Header.VisitAll(func(k, v []byte) {
		dst.Header.Set(string(k), string(v))
	})

	// Cast to a string to make a copy seeing as src.Body() won't
	// be valid after the response is released back to the pool (fasthttp.ReleaseResponse).
	dst.Body = ioutil.NopCloser(strings.NewReader(string(src.Body())))

	return dst
}
