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

// +build !integration

package elasticsearch_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var defaultResponse = http.Response{
	Status:        "200 OK",
	StatusCode:    200,
	ContentLength: 2,
	Header:        http.Header(map[string][]string{"Content-Type": {"application/json"}}),
	Body:          ioutil.NopCloser(strings.NewReader(`{}`)),
}

var infoResponse = http.Response{
	Status:     "200 OK",
	StatusCode: http.StatusOK,
	Header:     http.Header(map[string][]string{"Content-Type": {"application/json"}}),
}

type FakeTransport struct {
	InfoResponse *http.Response
	FakeResponse *http.Response
}

func (t *FakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.Path == "/" {
		response := t.InfoResponse
		response.Body = ioutil.NopCloser(strings.NewReader(`{
		  "name" : "es1",
		  "cluster_name" : "go-elasticsearch",
		  "cluster_uuid" : "clusteruuid",
		  "version" : {
			"number" : "7.14.0-SNAPSHOT",
			"build_flavor" : "default",
			"build_type" : "docker",
			"build_hash" : "somehash",
			"build_date" : "2021-06-09T06:34:20.411011746Z",
			"build_snapshot" : true,
			"lucene_version" : "8.9.0",
			"minimum_wire_compatibility_version" : "6.8.0",
			"minimum_index_compatibility_version" : "6.0.0-beta1"
		  },
		  "tagline" : "You Know, for Search"
		}`))
		return t.InfoResponse, nil
	}
	return t.FakeResponse, nil
}

func newFakeTransport(b *testing.B) *FakeTransport {
	return &FakeTransport{
		InfoResponse: &infoResponse,
		FakeResponse: &defaultResponse,
	}
}

func BenchmarkClient(b *testing.B) {
	b.ReportAllocs()

	b.Run("Create client with defaults", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := elasticsearch.NewClient(elasticsearch.Config{Transport: newFakeTransport(b)})

			if err != nil {
				b.Fatalf("Unexpected error when creating a client: %s", err)
			}
		}
	})
}

func BenchmarkClientAPI(b *testing.B) {
	b.ReportAllocs()

	ctx := context.Background()

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
		Transport: newFakeTransport(b),
	})
	if err != nil {
		b.Fatalf("ERROR: %s", err)
	}

	b.Run("InfoRequest{}.Do()", func(b *testing.B) {
		b.ResetTimer()

		req := esapi.InfoRequest{}

		for i := 0; i < b.N; i++ {
			if _, err := req.Do(ctx, client); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("IndexRequest{...}.Do()", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder

		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)

			body.Reset()
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)

			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: docID,
				Body:       strings.NewReader(body.String()),
				Refresh:    "true",
				Pretty:     true,
				Timeout:    100,
			}
			if _, err := req.Do(ctx, client); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("Index(...)", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder

		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)

			body.Reset()
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)

			_, err := client.Index(
				"test",
				strings.NewReader(body.String()),
				client.Index.WithDocumentID(docID),
				client.Index.WithRefresh("true"),
				client.Index.WithPretty(),
				client.Index.WithTimeout(100),
				client.Index.WithContext(ctx),
			)
			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("SearchRequest{...}.Do()", func(b *testing.B) {
		b.ResetTimer()

		body := `{"foo" : "bar"}`
		indx := []string{"test"}

		for i := 0; i < b.N; i++ {
			req := esapi.SearchRequest{
				Index:   indx,
				Body:    strings.NewReader(body),
				Size:    esapi.IntPtr(25),
				Pretty:  true,
				Timeout: 100,
			}
			if _, err := req.Do(ctx, client); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("Search(...)", func(b *testing.B) {
		b.ResetTimer()

		body := `{"foo" : "bar"}`
		indx := "test"

		for i := 0; i < b.N; i++ {
			_, err := client.Search(
				client.Search.WithIndex(indx),
				client.Search.WithBody(strings.NewReader(body)),
				client.Search.WithSize(25),
				client.Search.WithPretty(),
				client.Search.WithTimeout(100),
				client.Search.WithContext(ctx),
			)
			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("BulkRequest{...}.Do()", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder

		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)

			body.Reset()
			body.WriteString(`{"index" : { "_index" : "test", "_type" : "_doc", "_id" : "` + docID + `" }}`)
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)

			req := esapi.BulkRequest{
				Body:    strings.NewReader(body.String()),
				Refresh: "true",
				Pretty:  true,
				Timeout: 100,
			}
			if _, err := req.Do(ctx, client); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})

	b.Run("Bulk(...)", func(b *testing.B) {
		b.ResetTimer()
		var body strings.Builder

		for i := 0; i < b.N; i++ {
			docID := strconv.FormatInt(int64(i), 10)

			body.Reset()
			body.WriteString(`{"index" : { "_index" : "test", "_type" : "_doc", "_id" : "` + docID + `" }}`)
			body.WriteString(`{"foo" : "bar `)
			body.WriteString(docID)
			body.WriteString(`	" }`)

			_, err := client.Bulk(
				strings.NewReader(body.String()),
				client.Bulk.WithRefresh("true"),
				client.Bulk.WithPretty(),
				client.Bulk.WithTimeout(100),
				client.Bulk.WithContext(ctx),
			)
			if err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			}
		}
	})
}
