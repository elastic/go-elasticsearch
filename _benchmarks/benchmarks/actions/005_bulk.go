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
	"strings"

	"github.com/tidwall/gjson"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	"github.com/elastic/go-elasticsearch/v8/benchmarks"
	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"
)

func init() {
	benchmarks.Register(
		benchmarks.Action{
			Name:     "bulk",
			Category: "core",

			NumWarmups:     10,
			NumRepetitions: 1000,
			NumOperations:  10000,

			SetupFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				var (
					res *esapi.Response
					err error

					indexName     = "test-bench-bulk"
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
					res *esapi.Response
					err error

					indexName = "test-bench-bulk"
					opMeta    = []byte("{ \"index\" : {} }\n")
					opBody    bytes.Buffer
				)

				opBody.Reset()

				docBody := bytes.NewBuffer(bytes.ReplaceAll(benchmarks.DataSources["small"].Bytes(), []byte("\n"), []byte("")))
				docBody.WriteRune('\n')

				for i := 0; i < c.NumOperations; i++ {
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

				var b bytes.Buffer
				if _, err := b.ReadFrom(res.Body); err != nil {
					return nil, err
				}
				res.Body.Close()

				output := gjson.GetBytes(b.Bytes(), "errors")
				if output.Bool() {
					return nil, fmt.Errorf("Unexpected errors in output: %s", b.String())
				}
				return res, err
			},
		},
	)
}
