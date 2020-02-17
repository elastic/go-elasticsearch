// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package esutil_test

import (
	"context"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func BenchmarkBulkIndexer(b *testing.B) {
	b.ReportAllocs()

	b.Run("Basic", func(b *testing.B) {
		b.ResetTimer()

		bi, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{FlushBytes: 1024})
		defer bi.Close(context.Background())

		for i := 0; i < b.N; i++ {
			bi.Add(context.Background(), esutil.BulkIndexerItem{
				Action:     "foo",
				DocumentID: "bar",
			})
		}
	})
}
