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

//go:build integration
// +build integration

package esutil_test

import (
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func TestJSONReaderIntegration(t *testing.T) {
	t.Run("Index and search", func(t *testing.T) {
		var (
			res *esapi.Response
			err error
		)

		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			t.Fatalf("Error creating the client: %s\n", err)
		}

		es.Indices.Delete([]string{"test"}, es.Indices.Delete.WithIgnoreUnavailable(true))

		doc := struct {
			Title string `json:"title"`
		}{Title: "Foo Bar"}

		res, err = es.Index("test", esutil.NewJSONReader(&doc), es.Index.WithRefresh("true"))
		if err != nil {
			t.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			t.Fatalf("Error response: %s", res.String())
		}

		query := map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"title": "foo",
				},
			},
		}

		res, err = es.Search(
			es.Search.WithIndex("test"),
			es.Search.WithBody(esutil.NewJSONReader(&query)),
			es.Search.WithPretty(),
		)
		if err != nil {
			t.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			t.Errorf("Error response: %s", res)
		}

		if !strings.Contains(res.String(), "Foo Bar") {
			t.Errorf("Unexpected response: %s", res)
		}
	})
}
