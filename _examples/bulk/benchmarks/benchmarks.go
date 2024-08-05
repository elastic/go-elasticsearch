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

//go:build ignore
// +build ignore

// This example demonstrates indexing documents using the esutil.BulkIndexer helper.
//
// You can configure the settings with command line flags:
//
//     go run benchmark.go --dataset=httplog --runs=15 --count=1_000_000 --shards=5 --replicas=1 --flush=1MB

package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"

	"github.com/elastic/go-elasticsearch/v8/_examples/bulk/benchmarks/model"
	"github.com/elastic/go-elasticsearch/v8/_examples/bulk/benchmarks/runner"
)

type humanBytes uint64

func (b *humanBytes) String() string { return humanize.Bytes(uint64(*b)) }
func (b *humanBytes) Set(v string) error {
	n, err := humanize.ParseBytes(v)
	*b = humanBytes(n)
	return err
}

var (
	indexName     string
	datasetName   string
	numWorkers    int
	flushBytes    humanBytes
	numRuns       int
	numWarmupRuns int
	numItems      int
	numShards     int
	numReplicas   int
	wait          time.Duration
	useFasthttp   bool
	useEasyjson   bool
	mockserver    bool
	debug         bool
)

func init() {
	flag.StringVar(&indexName, "index", "test-bulk-benchmarks", "Index name")
	flag.StringVar(&datasetName, "dataset", "small", "Dataset to use for indexing")
	flag.IntVar(&numWorkers, "workers", runtime.NumCPU(), "Number of indexer workers")
	flag.Var(&flushBytes, "flush", "Flush threshold in bytes (default 3MB)")
	flag.IntVar(&numRuns, "runs", 10, "Number of runs")
	flag.IntVar(&numWarmupRuns, "warmup", 3, "Number of warmup runs")
	flag.IntVar(&numItems, "count", 100000, "Number of documents to generate")
	flag.IntVar(&numShards, "shards", 3, "Number of index shards")
	flag.IntVar(&numReplicas, "replicas", 0, "Number of index replicas (default 0)")
	flag.DurationVar(&wait, "wait", time.Second, "Wait duration between runs")
	flag.BoolVar(&useFasthttp, "fasthttp", false, "Use valyala/fasthttp for HTTP transport")
	flag.BoolVar(&useEasyjson, "easyjson", false, "Use mailru/easyjson for JSON decoding")
	flag.BoolVar(&mockserver, "mockserver", false, "Measure added, not flushed items")
	flag.BoolVar(&debug, "debug", false, "Enable logging output")
	flag.Parse()
}

func main() {
	log.SetFlags(0)

	indexName = indexName + "-" + datasetName
	if flushBytes < 1 {
		flushBytes = humanBytes(3e+6)
	}

	clientCfg := elasticsearch.Config{}

	if useFasthttp {
		clientCfg.Transport = &fasthttpTransport{}
	}

	if debug {
		clientCfg.Logger = &elastictransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true}
	}

	es, _ := elasticsearch.NewClient(clientCfg)

	runnerCfg := runner.Config{
		Client: es,

		IndexName:     indexName,
		DatasetName:   datasetName,
		NumShards:     numShards,
		NumReplicas:   numReplicas,
		NumItems:      numItems,
		NumRuns:       numRuns,
		NumWarmupRuns: numWarmupRuns,
		NumWorkers:    numWorkers,
		FlushBytes:    int(flushBytes),
		Wait:          wait,
		Mockserver:    mockserver,
	}

	if useEasyjson {
		runnerCfg.Decoder = easyjsonDecoder{}
	}

	runner, err := runner.NewRunner(runnerCfg)
	if err != nil {
		log.Fatalf("Error creating runner: %s", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)

	go func() { <-done; log.Println("\r" + strings.Repeat("▁", 110)); runner.Report(); os.Exit(0) }()
	defer func() { log.Println(strings.Repeat("▁", 110)); runner.Report() }()

	log.Printf(
		"%s: run [%sx] warmup [%dx] shards [%d] replicas [%d] workers [%d] flush [%s] wait [%s]%s%s",
		datasetName,
		humanize.Comma(int64(numRuns)),
		numWarmupRuns,
		numShards,
		numReplicas,
		numWorkers,
		humanize.Bytes(uint64(flushBytes)),
		wait,
		func() string {
			if useFasthttp {
				return " fasthttp"
			}
			return ""
		}(),
		func() string {
			if useEasyjson {
				return " easyjson"
			}
			return ""
		}())
	log.Println(strings.Repeat("▔", 110))

	runner.Run()
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

	dst.Body = io.NopCloser(strings.NewReader(string(src.Body())))

	return dst
}
