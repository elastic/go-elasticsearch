// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package operations

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/tidwall/gjson"

	"github.com/elastic/go-elasticsearch/v8/benchmarks"
	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"
)

func init() {
	benchmarks.Register(
		benchmarks.Operation{
			Action:         "get",
			Category:       "core",
			NumWarmups:     100,
			NumRepetitions: benchmarks.DefaultRepetitions,
			SetupFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName = "test-bench-get"
				)
				res, _ = c.RunnerClient.Indices.Delete([]string{indexName})
				if res != nil && res.Body != nil {
					defer res.Body.Close()
					io.Copy(ioutil.Discard, res.Body)
				}
				res, err = c.RunnerClient.Index(indexName, strings.NewReader(`{"title":"Test"}`), c.RunnerClient.Index.WithDocumentID("1"))
				if err != nil {
					return res, err
				}
				defer res.Body.Close()
				io.Copy(ioutil.Discard, res.Body)
				res, err = c.RunnerClient.Indices.Refresh(c.RunnerClient.Indices.Refresh.WithIndex(indexName))
				if err != nil {
					return res, err
				}
				defer res.Body.Close()
				io.Copy(ioutil.Discard, res.Body)
				return res, err
			},
			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var indexName = "test-bench-get"
				res, err := c.RunnerClient.Get(indexName, "1")
				if err != nil {
					return res, err
				}
				defer res.Body.Close()
				var b bytes.Buffer
				if _, err := b.ReadFrom(res.Body); err != nil {
					return nil, err
				}
				output := gjson.GetBytes(b.Bytes(), "_source.title")
				if output.Str != "Test" {
					return nil, fmt.Errorf("Unexpected output: %q", b.String())
				}
				return res, err
			},
		},
	)
}
