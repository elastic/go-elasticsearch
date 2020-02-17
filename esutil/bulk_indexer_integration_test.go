// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build integration

package esutil_test

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/estransport"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func TestBulkIndexerIntegration(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
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
			Index:  indexName,
			Client: es,
			// FlushBytes: 3e+6,
		})

		numItems := 100000
		body := `{"body":"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."}`
		start := time.Now().UTC()

		for i := 1; i <= numItems; i++ {
			err := bi.Add(context.Background(), esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: strconv.Itoa(i),
				Body:       strings.NewReader(body),
			})
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
		}

		stats := bi.Stats()

		if stats.NumAdded != uint(numItems) {
			t.Errorf("Unexpected NumAdded: %d", stats.NumAdded)
		}

		if stats.NumFailed != 0 {
			t.Errorf("Unexpected NumFailed: %d", stats.NumFailed)
		}

		if err := bi.Close(context.Background()); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		fmt.Printf("  Added %d documents to indexer. Succeeded: %d. Failed: %d. Duration: %s (%.0f docs/sec)\n",
			stats.NumAdded,
			stats.NumAdded-stats.NumFailed,
			stats.NumFailed,
			time.Since(start).Truncate(time.Millisecond),
			1000.0/float64(time.Since(start)/time.Millisecond)*float64(stats.NumAdded-stats.NumFailed))
	})
}
