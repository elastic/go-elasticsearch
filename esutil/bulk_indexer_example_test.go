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

//go:build !integration
// +build !integration

package esutil_test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func ExampleNewBulkIndexer() {
	log.SetFlags(0)

	// Create the Elasticsearch client
	//
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		// Retry on 429 TooManyRequests statuses
		//
		RetryOnStatus: []int{502, 503, 504, 429},

		// A simple incremental backoff function
		//
		RetryBackoff: func(i int) time.Duration { return time.Duration(i) * 100 * time.Millisecond },

		// Retry up to 5 attempts
		//
		MaxRetries: 5,
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Create the indexer
	//
	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Client:     es,     // The Elasticsearch client
		Index:      "test", // The default index name
		NumWorkers: 4,      // The number of worker goroutines (default: number of CPUs)
		FlushBytes: 5e+6,   // The flush threshold in bytes (default: 5M)
	})
	if err != nil {
		log.Fatalf("Error creating the indexer: %s", err)
	}

	// Add an item to the indexer
	//
	err = indexer.Add(
		context.Background(),
		esutil.BulkIndexerItem{
			// Action field configures the operation to perform (index, create, delete, update)
			Action: "index",

			// DocumentID is the optional document ID
			DocumentID: "1",

			// Body is an `io.Reader` with the payload
			Body: strings.NewReader(`{"title":"Test"}`),

			// OnSuccess is the optional callback for each successful operation
			OnSuccess: func(
				ctx context.Context,
				item esutil.BulkIndexerItem,
				res esutil.BulkIndexerResponseItem,
			) {
				fmt.Printf("[%d] %s test/%s", res.Status, res.Result, item.DocumentID)
			},

			// OnFailure is the optional callback for each failed operation
			OnFailure: func(
				ctx context.Context,
				item esutil.BulkIndexerItem,
				res esutil.BulkIndexerResponseItem, err error,
			) {
				if err != nil {
					log.Printf("ERROR: %s", err)
				} else {
					log.Printf("ERROR: %s: %s", res.Error.Type, res.Error.Reason)
				}
			},
		},
	)
	if err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	// Close the indexer channel and flush remaining items
	//
	if err := indexer.Close(context.Background()); err != nil {
		log.Fatalf("Unexpected error: %s", err)
	}

	// Report the indexer statistics
	//
	stats := indexer.Stats()
	if stats.NumFailed > 0 {
		log.Fatalf("Indexed [%d] documents with [%d] errors", stats.NumFlushed, stats.NumFailed)
	} else {
		log.Printf("Successfully indexed [%d] documents", stats.NumFlushed)
	}

	// For optimal performance, consider using a third-party package for JSON decoding and HTTP transport.
	//
	// For more information, examples and benchmarks, see:
	//
	// --> https://github.com/elastic/go-elasticsearch/tree/master/_examples/bulk
}
