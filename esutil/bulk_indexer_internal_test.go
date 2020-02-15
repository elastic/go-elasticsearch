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
		bi, _ := NewBulkIndexer(BulkIndexerConfig{NumWorkers: 1, FlushThresholdBytes: 7})
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

		if bi.NumAdded() != numItems {
			t.Errorf("Unexpected NumAdded: %d", bi.NumAdded())
		}

		if bi.NumFailed() != 0 {
			t.Errorf("Unexpected NumFailed: %d", bi.NumFailed())
		}

		if err := bi.Wait(context.Background()); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		fmt.Println("NumAdded:", bi.NumAdded(), "NumFailed:", bi.NumFailed())
	})

	t.Run("Add() Timeout", func(t *testing.T) {
		bi, _ := NewBulkIndexer(BulkIndexerConfig{NumWorkers: 1})
		ctx, _ := context.WithTimeout(context.Background(), time.Nanosecond)
		var errs []error
		for i := 0; i < 10; i++ {
			errs = append(errs, bi.Add(ctx, BulkIndexerItem{Action: "index"}))
		}
		if err := bi.Wait(context.Background()); err != nil {
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
