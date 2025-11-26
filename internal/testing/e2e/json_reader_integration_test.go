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

package e2e_test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
	"github.com/elastic/go-elasticsearch/v9/esutil"

	"testing/containertest"
)

func TestJSONReaderIntegration(t *testing.T) {
	stackVersion := elasticsearch.Version
	if v := os.Getenv("STACK_VERSION"); v != "" {
		stackVersion = v
	}

	elasticsearchSrv, err := containertest.NewElasticsearchService(stackVersion)
	if err != nil {
		t.Fatalf("Error setting up Elasticsearch container: %s", err)
	}
	defer func() {
		if err := elasticsearchSrv.Terminate(context.Background()); err != nil {
			t.Fatalf("Error terminating Elasticsearch container: %s", err)
		}
	}()

	tcCfg := elasticsearchSrv.ESConfig()

	t.Run("Index and search", func(t *testing.T) {
		var (
			res *esapi.Response
			err error
		)

		es, err := elasticsearch.NewClient(tcCfg)
		if err != nil {
			t.Fatalf("Error creating the client: %s\n", err)
		}
		defer func() {
			if err := es.Close(context.Background()); err != nil {
				t.Fatalf("Error closing the client: %s", err)
			}
		}()

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
