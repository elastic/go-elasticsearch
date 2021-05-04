// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package actions

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
		benchmarks.Action{
			Name:           "get",
			Category:       "core",
			NumWarmups:     100,
			NumRepetitions: 10000,
			SetupFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName = "test-bench-get"
				)
				res, _ = c.RunnerClient.Indices.Delete([]string{indexName})
				if res != nil && res.Body != nil {
					io.Copy(ioutil.Discard, res.Body)
					res.Body.Close()
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
				io.Copy(ioutil.Discard, res.Body)
				res.Body.Close()
				return res, err
			},
			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var indexName = "test-bench-get"
				res, err := c.RunnerClient.Get(indexName, "1")
				if err != nil {
					return res, err
				}
				var b bytes.Buffer
				if _, err := b.ReadFrom(res.Body); err != nil {
					return nil, err
				}
				res.Body.Close()

				output := gjson.GetBytes(b.Bytes(), "_source.title")
				if output.Str != "Test" {
					return nil, fmt.Errorf("Unexpected output: %q", b.String())
				}
				return res, err
			},
		},
	)
}
