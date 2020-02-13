// +build !integration

package esutil

import (
	"fmt"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

// TODO(karmi): Benchmark in _examples with the Enron dataset?
// TODO(karmi): Benchmark in _examples with the "Geopoint" and "HTTP Logs" Rally tracks.

func TestBulkIndexer(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		es, _ := elasticsearch.NewDefaultClient()
		bi := &BulkIndexer{Client: es}
		if err := bi.Add(BulkIndexerItem{Action: "index"}); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		fmt.Println("NumAdded:", bi.NumAdded(), "NumFailed:", bi.NumFailed())

		if bi.NumAdded() != 1 {
			t.Errorf("Unexpected NumAdded: %d", bi.NumAdded())
		}

		if err := bi.Wait(nil); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
	})
}
