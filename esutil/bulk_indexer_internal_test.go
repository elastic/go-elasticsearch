// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package esutil

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// TODO(karmi): Benchmark in _examples with the Enron dataset?
// TODO(karmi): Benchmark in _examples with the "Geopoint" and "HTTP Logs" Rally tracks.

func TestBulkIndexer(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		var wg sync.WaitGroup
		bi, _ := NewBulkIndexer(BulkIndexerConfig{NumWorkers: 1, FlushBytes: 7})
		numItems := 3

		for i := 1; i <= numItems; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				err := bi.Add(context.Background(), BulkIndexerItem{
					Action:     "index",
					DocumentID: strconv.Itoa(i),
					Body:       strings.NewReader(fmt.Sprintf(`foo=%d`, i)),
				})
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}(i)
		}
		wg.Wait()

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

		fmt.Println("NumAdded:", stats.NumAdded, "NumFailed:", stats.NumFailed)
	})

	t.Run("Add() Timeout", func(t *testing.T) {
		bi, _ := NewBulkIndexer(BulkIndexerConfig{NumWorkers: 1})
		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
		defer cancel()
		var errs []error
		for i := 0; i < 10; i++ {
			errs = append(errs, bi.Add(ctx, BulkIndexerItem{Action: "index", DocumentID: "timeout"}))
		}
		if err := bi.Close(context.Background()); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		var gotError bool
		for _, err := range errs {
			if err != nil && err.Error() == "context deadline exceeded" {
				gotError = true
			}
		}
		if !gotError {
			t.Errorf("Expected timeout error, but none in: %s", errs)
		}
	})
}
