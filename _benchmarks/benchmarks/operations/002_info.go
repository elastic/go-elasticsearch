// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package operations

import (
	"io"
	"io/ioutil"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	"github.com/elastic/go-elasticsearch/v8/benchmarks"
	"github.com/elastic/go-elasticsearch/v8/benchmarks/runner"
)

func init() {
	benchmarks.Register(
		benchmarks.Operation{
			Action:         "info",
			Category:       "core",
			NumWarmups:     0,
			NumRepetitions: benchmarks.DefaultRepetitions,
			RunnerFunc: func(n int, c runner.Config) (*esapi.Response, error) {
				res, err := c.RunnerClient.Info()
				if err == nil && res != nil && res.Body != nil {
					defer res.Body.Close()
					io.Copy(ioutil.Discard, res.Body)
				}
				return res, err
			},
		},
	)
}
