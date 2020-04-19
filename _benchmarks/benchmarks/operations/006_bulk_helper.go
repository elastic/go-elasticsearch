// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package operations

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/estransport"
	"github.com/elastic/go-elasticsearch/v8/esutil"

	"github.com/elastic/go-elasticsearch/v8/benchmarks"
	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"
)

func init() {
	benchmarks.Register(
		benchmarks.Operation{
			Action:   "bulk-helper",
			Category: "helpers",

			NumWarmups:     1,
			NumRepetitions: 10,
			NumOperations:  1000000,

			SetupFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName     = "test-bench-bulk-helper"
					indexSettings = `{"settings": { "number_of_shards": 3, "refresh_interval":"5s"}}`
				)
				res, _ = c.RunnerClient.Indices.Delete([]string{indexName})
				if res != nil && res.Body != nil {
					res.Body.Close()
				}
				res, err = c.RunnerClient.Indices.Create(
					indexName,
					c.RunnerClient.Indices.Create.WithBody(strings.NewReader(indexSettings)),
					c.RunnerClient.Indices.Create.WithWaitForActiveShards("1"))
				if err != nil {
					return res, err
				}
				res.Body.Close()
				return res, err
			},

			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					err error

					indexName = "test-bench-bulk-helper"
					docBody   = bytes.NewBuffer([]byte(""))
				)

				docBody.Reset()

				var addresses []string
				for _, u := range c.RunnerClient.Transport.(*estransport.Client).URLs() {
					addresses = append(addresses, u.String())
				}

				es, err := elasticsearch.NewClient(elasticsearch.Config{
					Addresses:     addresses,
					RetryOnStatus: []int{502, 503, 504, 429}, // Retry on 429 TooManyRequests statuses
					MaxRetries:    5,
				})
				if err != nil {
					return nil, err
				}

				bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
					Index:      indexName,
					Client:     es,
					NumWorkers: 8,
					FlushBytes: 2e+6,
					// FlushInterval: 30 * time.Second,
				})
				if err != nil {
					return nil, err
				}

				f, err := os.Open(filepath.Join(benchmarks.Config["DATA_SOURCE"], "small", "document.json"))
				if err != nil {
					return nil, err
				}
				docBody.ReadFrom(f)
				docBody = bytes.NewBuffer(bytes.ReplaceAll(docBody.Bytes(), []byte("\n"), []byte("")))
				docBody.WriteRune('\n')

				for i := 0; i < c.NumOperations; i++ {
					err = bi.Add(
						context.Background(),
						esutil.BulkIndexerItem{
							Action: "index",
							Body:   bytes.NewReader(docBody.Bytes()),
						},
					)
					if err != nil {
						return nil, err
					}
				}

				err = bi.Close(context.Background())
				if err != nil {
					return nil, err
				}

				biStats := bi.Stats()
				if biStats.NumFailed > 0 {
					return nil, fmt.Errorf("Unexpected failures: %s", biStats)
				}
				if int(biStats.NumAdded) != c.NumOperations {
					return nil, fmt.Errorf("Unexpected failures: added=%d, expected=%d", biStats.NumAdded, c.NumOperations)
				}
				return &esapi.Response{StatusCode: 200}, nil
			},
		},
	)
}
