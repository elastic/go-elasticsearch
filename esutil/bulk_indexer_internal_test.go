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
		bi := BulkIndexer{Client: es}
		bi.Add(nil)
		fmt.Println("NumPublished:", bi.NumPublished(), "NumFailed:", bi.NumFailed())
	})
}
