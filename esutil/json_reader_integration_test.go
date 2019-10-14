// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build integration

package esutil_test

import (
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
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

		es.Indices.Delete([]string{"test"})

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
