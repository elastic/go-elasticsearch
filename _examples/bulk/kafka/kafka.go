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

package main

import (
	"context"
	"expvar"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	_ "net/http/pprof"

	"github.com/dustin/go-humanize"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmelasticsearch"

	"github.com/elastic/go-elasticsearch/v8/_examples/bulk/kafka/consumer"
	"github.com/elastic/go-elasticsearch/v8/_examples/bulk/kafka/producer"
)

var (
	brokerURL string

	topicName  = "stocks"
	topicParts = 4
	msgRate    int

	indexName    = "stocks"
	numProducers = 1
	numConsumers = 4
	numIndexers  = 1
	flushBytes   = 0 // Default
	numWorkers   = 0 // Default
	indexerError error

	mapping = `{
  "mappings": {
    "properties": {
    	"time":     { "type": "date"    },
    	"symbol":   { "type": "keyword" },
    	"side":     { "type": "keyword" },
      "account":  { "type": "keyword" },
      "quantity": { "type": "long"    },
      "price":    { "type": "long"    },
      "amount":   { "type": "long"    }
      }
    }}}`
)

func init() {
	if v := os.Getenv("KAFKA_URL"); v != "" {
		brokerURL = v
	} else {
		brokerURL = "localhost:9092"
	}
	flag.IntVar(&msgRate, "rate", 1000, "Producer rate (msg/sec)")
	flag.IntVar(&numProducers, "producers", numProducers, "Number of producers")
	flag.IntVar(&numConsumers, "consumers", numConsumers, "Number of consumers")
	flag.IntVar(&numIndexers, "indexers", numIndexers, "Number of indexers")
	flag.Parse()
}

func main() {
	log.SetFlags(0)

	// Serve the "/debug/pprof/" and "/debug/vars" pages
	//
	go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()

	var (
		wg  sync.WaitGroup
		ctx = context.Background()

		producers []*producer.Producer
		consumers []*consumer.Consumer
		indexers  []esutil.BulkIndexer
	)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	go func() { <-done; log.Println(""); os.Exit(0) }()

	// Set up producers
	//
	for i := 1; i <= numProducers; i++ {
		producers = append(producers,
			&producer.Producer{
				BrokerURL:   brokerURL,
				TopicName:   topicName,
				TopicParts:  topicParts,
				MessageRate: msgRate})
	}

	// Create an Elasticsearch client
	//
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		RetryOnStatus: []int{502, 503, 504, 429}, // Add 429 to the list of retryable statuses
		RetryBackoff:  func(i int) time.Duration { return time.Duration(i) * 100 * time.Millisecond },
		MaxRetries:    5,
		EnableMetrics: true,
		Transport:     apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
	})
	if err != nil {
		log.Fatalf("Error: NewClient(): %s", err)
	}
	// Export client metrics to the "expvar" package
	expvar.Publish("go-elasticsearch", expvar.Func(func() interface{} { m, _ := es.Metrics(); return m }))

	// Create the "stocks" index with correct mappings
	//
	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		log.Fatalf("Error: Indices.Exists: %s", err)
	}
	res.Body.Close()
	if res.StatusCode == 404 {
		res, err := es.Indices.Create(
			indexName,
			es.Indices.Create.WithBody(strings.NewReader(mapping)),
			es.Indices.Create.WithWaitForActiveShards("1"),
		)
		if err != nil {
			log.Fatalf("Error: Indices.Create: %s", err)
		}
		if res.IsError() {
			log.Fatalf("Error: Indices.Create: %s", res)
		}
	}

	// Set up indexers
	//
	for i := 1; i <= numIndexers; i++ {
		idx, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:      indexName,
			Client:     es,
			NumWorkers: numWorkers,
			FlushBytes: int(flushBytes),
			// Elastic APM: Instrument the flush operations and capture errors
			OnFlushStart: func(ctx context.Context) context.Context {
				txn := apm.DefaultTracer.StartTransaction("Bulk", "indexing")
				return apm.ContextWithTransaction(ctx, txn)
			},
			OnFlushEnd: func(ctx context.Context) {
				apm.TransactionFromContext(ctx).End()
			},
			OnError: func(ctx context.Context, err error) {
				indexerError = err
				apm.CaptureError(ctx, err).Send()
			},
		})
		if err != nil {
			log.Fatalf("ERROR: NewBulkIndexer(): %s", err)
		}
		indexers = append(indexers, idx)
	}

	// Set up consumers
	//
	for i := 1; i <= numConsumers; i++ {
		consumers = append(consumers,
			&consumer.Consumer{
				BrokerURL: brokerURL,
				TopicName: topicName,
				Indexer:   indexers[i%numIndexers]})
	}

	// Set up reporting output
	//
	reporter := time.NewTicker(500 * time.Millisecond)
	defer reporter.Stop()
	go func() {
		fmt.Printf("Initializing... producers=%d consumers=%d indexers=%d\n", numProducers, numConsumers, numIndexers)
		for {
			select {
			case <-reporter.C:
				fmt.Print(report(producers, consumers, indexers))
			}
		}
	}()
	errcleaner := time.NewTicker(10 * time.Second)
	defer errcleaner.Stop()
	go func() {
		for {
			select {
			case <-errcleaner.C:
				indexerError = nil
			}
		}
	}()

	// Create the Kafka topic
	//
	if len(producers) > 0 {
		if err := producers[0].CreateTopic(ctx); err != nil {
			log.Fatalf("ERROR: Producer: %s", err)
		}
	}

	// Launch consumers
	//
	for _, c := range consumers {
		wg.Add(1)
		go func(c *consumer.Consumer) {
			defer wg.Done()
			if err := c.Run(ctx); err != nil {
				log.Fatalf("ERROR: Consumer: %s", err)
			}
		}(c)
	}

	// Launch producers
	//
	time.Sleep(5 * time.Second) // Leave some room for consumers to connect
	for _, p := range producers {
		wg.Add(1)
		go func(p *producer.Producer) {
			defer wg.Done()
			if err := p.Run(ctx); err != nil {
				log.Fatalf("ERROR: Producer: %s", err)
			}
		}(p)
	}

	wg.Wait()

	fmt.Print(report(producers, consumers, indexers))
}

func report(
	producers []*producer.Producer,
	consumers []*consumer.Consumer,
	indexers []esutil.BulkIndexer,
) string {
	var (
		b strings.Builder

		value    string
		currRow  = 1
		numCols  = 6
		colWidth = 20

		divider = func(last bool) {
			fmt.Fprintf(&b, "\033[%d;0H", currRow)
			fmt.Fprint(&b, "┣")
			for i := 1; i <= numCols; i++ {
				fmt.Fprint(&b, strings.Repeat("━", colWidth))
				if last && i == 5 {
					fmt.Fprint(&b, "┷")
					continue
				}
				if i < numCols {
					fmt.Fprint(&b, "┿")
				}
			}
			fmt.Fprint(&b, "┫")
			currRow++
		}
	)

	fmt.Print("\033[2J\033[K")
	fmt.Printf("\033[%d;0H", currRow)

	fmt.Fprint(&b, "┏")
	for i := 1; i <= numCols; i++ {
		fmt.Fprint(&b, strings.Repeat("━", colWidth))
		if i < numCols {
			fmt.Fprint(&b, "┯")
		}
	}
	fmt.Fprint(&b, "┓")
	currRow++

	for i, p := range producers {
		fmt.Fprintf(&b, "\033[%d;0H", currRow)
		value = fmt.Sprintf("Producer %d", i+1)
		fmt.Fprintf(&b, "┃ %-*s│", colWidth-1, value)
		s := p.Stats()
		value = fmt.Sprintf("duration=%s", s.Duration.Truncate(time.Second))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("msg/sec=%s", humanize.FtoaWithDigits(s.Throughput, 2))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("sent=%s", humanize.Comma(int64(s.TotalMessages)))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("bytes=%s", humanize.Bytes(uint64(s.TotalBytes)))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("errors=%s", humanize.Comma(int64(s.TotalErrors)))
		fmt.Fprintf(&b, " %-*s┃", colWidth-1, value)
		currRow++
		divider(i == len(producers)-1)
	}

	for i, c := range consumers {
		fmt.Fprintf(&b, "\033[%d;0H", currRow)
		value = fmt.Sprintf("Consumer %d", i+1)
		fmt.Fprintf(&b, "┃ %-*s│", colWidth-1, value)
		s := c.Stats()
		value = fmt.Sprintf("lagging=%s", humanize.Comma(s.TotalLag))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("msg/sec=%s", humanize.FtoaWithDigits(s.Throughput, 2))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("received=%s", humanize.Comma(s.TotalMessages))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("bytes=%s", humanize.Bytes(uint64(s.TotalBytes)))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("errors=%s", humanize.Comma(s.TotalErrors))
		fmt.Fprintf(&b, " %-*s┃", colWidth-1, value)
		currRow++
		divider(i == len(consumers)-1)
	}

	for i, x := range indexers {
		fmt.Fprintf(&b, "\033[%d;0H", currRow)
		value = fmt.Sprintf("Indexer %d", i+1)
		fmt.Fprintf(&b, "┃ %-*s│", colWidth-1, value)
		s := x.Stats()
		value = fmt.Sprintf("added=%s", humanize.Comma(int64(s.NumAdded)))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("flushed=%s", humanize.Comma(int64(s.NumFlushed)))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		value = fmt.Sprintf("failed=%s", humanize.Comma(int64(s.NumFailed)))
		fmt.Fprintf(&b, " %-*s│", colWidth-1, value)
		if indexerError != nil {
			value = "err=" + indexerError.Error()
			if len(value) > 2*colWidth {
				value = value[:2*colWidth]
			}
		} else {
			value = ""
		}
		fmt.Fprintf(&b, " %-*s┃", 2*colWidth, value)
		currRow++
		if i < len(indexers)-1 {
			divider(true)
		}
	}

	fmt.Fprintf(&b, "\033[%d;0H", currRow)
	fmt.Fprint(&b, "┗")
	for i := 1; i <= numCols; i++ {
		fmt.Fprint(&b, strings.Repeat("━", colWidth))
		if i == 5 {
			fmt.Fprint(&b, "━")
			continue
		}
		if i < numCols {
			fmt.Fprint(&b, "┷")
		}
	}
	fmt.Fprint(&b, "┛")
	currRow++

	return b.String()
}
