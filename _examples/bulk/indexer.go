// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build ignore

package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"
	"github.com/elastic/go-elasticsearch/v8/esutil"

	"github.com/elastic/go-elasticsearch/v8/_examples/bulk/model"
)

var (
	numRuns  = 5
	numItems = 100000
)

func init() {
	if os.Getenv("RUNS") != "" {
		numRuns, _ = strconv.Atoi(os.Getenv("RUNS"))
	}
	if os.Getenv("COUNT") != "" {
		numItems, _ = strconv.Atoi(os.Getenv("COUNT"))
	}
}

func main() {
	log.SetFlags(0)

	var countSuccessful uint64
	indexName := "test-bulk-integration"

	cfg := elasticsearch.Config{}
	cfg.Transport = &fasthttpTransport{}
	if os.Getenv("DEBUG") != "" {
		cfg.Logger = &estransport.ColorLogger{Output: os.Stdout}
	}

	es, _ := elasticsearch.NewClient(cfg)

	for n := 1; n <= numRuns; n++ {
		bi, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:      indexName,
			Client:     es,
			Decoder:    easyjsonDecoder{},
			NumWorkers: 8,
			FlushBytes: 1e+6,
		})

		es.Indices.Delete([]string{indexName}, es.Indices.Delete.WithIgnoreUnavailable(true))
		es.Indices.Create(
			indexName,
			es.Indices.Create.WithBody(strings.NewReader(`{"settings": {"number_of_shards": 3, "number_of_replicas": 0, "refresh_interval":"5s"}}`)),
			es.Indices.Create.WithWaitForActiveShards("1"))

		body := `{"body":"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."}`
		start := time.Now().UTC()

		for i := 1; i <= numItems; i++ {
			err := bi.Add(context.Background(), esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(i),
				Body:       strings.NewReader(body),
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

		stats := bi.Stats()

		if stats.NumAdded != uint(numItems) {
			log.Fatalf("Unexpected NumAdded: %d", stats.NumAdded)
		}

		if stats.NumFailed != 0 {
			log.Fatalf("Unexpected NumFailed: %d", stats.NumFailed)
		}

		if countSuccessful != uint64(n*numItems) {
			log.Fatalf("Unexpected countSuccessful: %d", countSuccessful)
		}

		log.Printf("%4d) added=%s flushed=%s failed=%s duration=%s throughput=%s docs/sec\n",
			n,
			humanize.Comma(int64(stats.NumAdded)),
			humanize.Comma(int64(stats.NumFlushed)),
			humanize.Comma(int64(stats.NumFailed)),
			time.Since(start).Truncate(time.Millisecond),
			humanize.Comma(int64(1000.0/float64(time.Since(start)/time.Millisecond)*float64(stats.NumFlushed))))
	}
}

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
