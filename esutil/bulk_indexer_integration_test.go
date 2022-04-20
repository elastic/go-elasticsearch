// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

//go:build integration
// +build integration

package esutil_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/estransport"
	"github.com/elastic/go-elasticsearch/v6/esutil"
)

func TestBulkIndexerIntegration(t *testing.T) {
	body := `{"body":"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."}`

	t.Run("Default", func(t *testing.T) {
		var countSuccessful uint64
		indexName := "test-bulk-integration"

		es, _ := elasticsearch.NewClient(elasticsearch.Config{
			Logger: &estransport.ColorLogger{Output: os.Stdout},
		})

		es.Indices.Delete([]string{indexName}, es.Indices.Delete.WithIgnoreUnavailable(true))
		es.Indices.Create(
			indexName,
			es.Indices.Create.WithBody(strings.NewReader(`{"settings": {"number_of_shards": 1, "number_of_replicas": 0, "refresh_interval":"5s"}}`)),
			es.Indices.Create.WithWaitForActiveShards("1"))

		bi, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:        indexName,
			DocumentType: "_doc",
			Client:       es,
			// FlushBytes: 3e+6,
		})

		numItems := 100000
		start := time.Now().UTC()

		for i := 1; i <= numItems; i++ {
			err := bi.Add(context.Background(), esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(i),
				Body:       strings.NewReader(body),
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},
			})
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		}

		if err := bi.Close(context.Background()); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		stats := bi.Stats()

		if stats.NumAdded != uint64(numItems) {
			t.Errorf("Unexpected NumAdded: want=%d, got=%d", numItems, stats.NumAdded)
		}

		if stats.NumIndexed != uint64(numItems) {
			t.Errorf("Unexpected NumIndexed: want=%d, got=%d", numItems, stats.NumIndexed)
		}

		if stats.NumFailed != 0 {
			t.Errorf("Unexpected NumFailed: want=0, got=%d", stats.NumFailed)
		}

		if countSuccessful != uint64(numItems) {
			t.Errorf("Unexpected countSuccessful: want=%d, got=%d", numItems, countSuccessful)
		}

		fmt.Printf("  Added %d documents to indexer. Succeeded: %d. Failed: %d. Requests: %d. Duration: %s (%.0f docs/sec)\n",
			stats.NumAdded,
			stats.NumFlushed,
			stats.NumFailed,
			stats.NumRequests,
			time.Since(start).Truncate(time.Millisecond),
			1000.0/float64(time.Since(start)/time.Millisecond)*float64(stats.NumFlushed))
	})

	t.Run("Multiple indices", func(t *testing.T) {
		es, _ := elasticsearch.NewClient(elasticsearch.Config{
			Logger: &estransport.ColorLogger{Output: os.Stdout, EnableRequestBody: true, EnableResponseBody: true},
		})

		bi, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:        "test-index-a",
			DocumentType: "_doc",
			Client:       es,
		})

		// Default index
		for i := 1; i <= 10; i++ {
			err := bi.Add(context.Background(), esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(i),
				Body:       strings.NewReader(body),
			})
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		}

		// Index 1
		for i := 1; i <= 10; i++ {
			err := bi.Add(context.Background(), esutil.BulkIndexerItem{
				Action: "index",
				Index:  "test-index-b",
				Body:   strings.NewReader(body),
			})
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		}

		// Index 2
		for i := 1; i <= 10; i++ {
			err := bi.Add(context.Background(), esutil.BulkIndexerItem{
				Action: "index",
				Index:  "test-index-c",
				Body:   strings.NewReader(body),
			})
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		}

		if err := bi.Close(context.Background()); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		stats := bi.Stats()

		expectedIndexed := 10 + 10 + 10
		if stats.NumIndexed != uint64(expectedIndexed) {
			t.Errorf("Unexpected NumIndexed: want=%d, got=%d", expectedIndexed, stats.NumIndexed)
		}

		res, err := es.Indices.Exists([]string{"test-index-a", "test-index-b", "test-index-c"})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if res.StatusCode != 200 {
			t.Errorf("Expected indices to exist, but got a [%s] response", res.Status())
		}
	})
	t.Run("External version", func(t *testing.T) {
		var index string = "test-index-a"

		es, _ := elasticsearch.NewClient(elasticsearch.Config{
			Logger: &estransport.ColorLogger{Output: os.Stdout},
		})

		es.Indices.Delete([]string{index}, es.Indices.Delete.WithIgnoreUnavailable(true))
		es.Indices.Create(index, es.Indices.Create.WithWaitForActiveShards("1"))

		bulkIndex := func(bulkIndexer esutil.BulkIndexer, baseVersion int) {
			var countTotal int = 500
			var countSuccessful uint64
			for i := 0; i < countTotal; i++ {
				version := int64(baseVersion + i)
				item := esutil.BulkIndexerItem{
					Action:      "index",
					Index:       index,
					DocumentID:  strconv.Itoa(i),
					Body:        strings.NewReader(body),
					Version:     &version,
					VersionType: "external",
					Routing:     `"{required": true}`,
					OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, item2 esutil.BulkIndexerResponseItem) {
						if version != item2.Version &&
							version != *item.Version &&
							item2.Result != "created" {
							t.Fatalf("version mismatch, expected: %d, got: %d && %d", version, item.Version, item2.Version)
						}
						atomic.AddUint64(&countSuccessful, 1)
					},
				}
				err := bulkIndexer.Add(context.Background(), item)
				if err != nil {
					t.Fatal(err)
				}
			}
			if err := bulkIndexer.Close(context.Background()); err != nil {
				t.Fatal(err)
			}

			if int(countSuccessful) != countTotal {
				t.Fatalf("Unexpected countSuccessful, expected %d, got: %d", countTotal, countSuccessful)
			}
		}

		bi, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:        index,
			Client:       es,
			DocumentType: "_doc",
		})
		bulkIndex(bi, 500)

		bi, _ = esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Index:        index,
			Client:       es,
			DocumentType: "_doc",
		})
		bulkIndex(bi, 900)
	})
}
