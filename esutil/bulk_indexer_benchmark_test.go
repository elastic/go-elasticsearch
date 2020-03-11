// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package esutil_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

var mockResponseBody = `{
  "took": 30,
  "errors": false,
  "items": [
    {
      "index": {
        "_index": "test",
        "_id": "1",
        "_version": 1,
        "result": "created",
        "_shards": { "total": 2, "successful": 1, "failed": 0 },
        "status": 201,
        "_seq_no": 0,
        "_primary_term": 1
      }
    }
  ]
}`

type mockTransp struct{}

func (t *mockTransp) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{Body: ioutil.NopCloser(strings.NewReader(mockResponseBody))}, nil // 1x alloc
}

func BenchmarkBulkIndexer(b *testing.B) {
	b.ReportAllocs()

	b.Run("Basic", func(b *testing.B) {
		b.ResetTimer()

		es, _ := elasticsearch.NewClient(elasticsearch.Config{Transport: &mockTransp{}})
		bi, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
			Client:     es,
			FlushBytes: 1024,
		})
		defer bi.Close(context.Background())

		docID := make([]byte, 0, 16)
		var docIDBuf bytes.Buffer
		docIDBuf.Grow(cap(docID))

		for i := 0; i < b.N; i++ {
			docID = strconv.AppendInt(docID, int64(i), 10)
			docIDBuf.Write(docID)
			bi.Add(context.Background(), esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: docIDBuf.String(),                  // 1x alloc
				Body:       strings.NewReader(`{"foo":"bar"}`), // 1x alloc
			})
			docID = docID[:0]
			docIDBuf.Reset()
		}
	})
}
