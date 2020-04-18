// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package operations

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/tidwall/gjson"

	"github.com/elastic/go-elasticsearch/v8/benchmarks"
	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"
)

func init() {
	benchmarks.Register(
		benchmarks.Operation{
			Action:   "bulk",
			Category: "core",

			NumWarmups:     0,
			NumRepetitions: 1000,
			NumOperations:  10000,

			SetupFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName = "test-bench-bulk"
				)
				res, _ = c.RunnerClient.Indices.Delete([]string{indexName})
				if res != nil && res.Body != nil {
					res.Body.Close()
				}
				res, err = c.RunnerClient.Indices.Create(indexName)
				if err != nil {
					return res, err
				}
				res.Body.Close()
				return res, err
			},

			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName = "test-bench-bulk"
					opNum     = 10000
					opMeta    = []byte("{ \"index\" : {} }\n")
					opBody    bytes.Buffer
					docBody   = bytes.NewBuffer([]byte(""))
				)

				opBody.Reset()
				docBody.Reset()

				f, err := os.Open(filepath.Join(benchmarks.Config["DATA_SOURCE"], "small", "document.json"))
				if err != nil {
					return nil, err
				}
				docBody.ReadFrom(f)
				docBody = bytes.NewBuffer(bytes.ReplaceAll(docBody.Bytes(), []byte("\n"), []byte("")))
				docBody.WriteRune('\n')

				for i := 0; i < opNum; i++ {
					var copyDocBody bytes.Buffer
					tr := io.TeeReader(docBody, &copyDocBody)
					opBody.Write(opMeta)
					opBody.ReadFrom(tr)
					docBody = &copyDocBody
				}

				res, err = c.RunnerClient.Bulk(&opBody, c.RunnerClient.Bulk.WithIndex(indexName))
				if err != nil {
					return res, err
				}
				defer res.Body.Close()

				var b bytes.Buffer
				if _, err := b.ReadFrom(res.Body); err != nil {
					return nil, err
				}
				output := gjson.GetBytes(b.Bytes(), "errors")
				if output.Bool() {
					return nil, fmt.Errorf("Unexpected errors in output: %s", b.String())
				}
				return res, err
			},
		},
	)
}
