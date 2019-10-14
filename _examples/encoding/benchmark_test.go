// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package main_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/tidwall/gjson"

	"github.com/elastic/go-elasticsearch/v8/_examples/encoding/model"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

func BenchmarkEncode(b *testing.B) {
	b.ReportAllocs()

	var (
		buf bytes.Buffer

		article = &model.Article{
			ID:        1,
			Title:     "Test",
			Body:      "Test",
			Published: time.Now(),
			Author: &model.Author{
				FirstName: "Alice",
				LastName:  "Smith",
			},
		}

		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"title": "test",
				},
			},
		}
	)

	b.Run("Article - json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := json.NewEncoder(&buf).Encode(article)
			if err != nil {
				b.Error(err)
			}
			buf.Reset()
		}
	})

	b.Run("Article - JSONReader", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := io.Copy(&buf, esutil.NewJSONReader(article))
			if err != nil {
				b.Error(err)
			}
			buf.Reset()
		}
	})

	b.Run("Article - easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := easyjson.MarshalToWriter(article, &buf)
			if err != nil {
				b.Error(err)
			}
			buf.Reset()
		}
	})

	b.Run("map - json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := json.NewEncoder(&buf).Encode(query)
			if err != nil {
				b.Error(err)
			}
			buf.Reset()
		}
	})

	b.Run("map - JSONReader", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := io.Copy(&buf, esutil.NewJSONReader(query))
			if err != nil {
				b.Error(err)
			}
			buf.Reset()
		}
	})
}

func BenchmarkDecode(b *testing.B) {
	b.ReportAllocs()

	resSearch := fixture("testdata/response_search.json")
	resClusterStats := fixture("testdata/response_cluster_stats.json")

	b.Run("Search - json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var res model.SearchResponse
			err := json.NewDecoder(bytes.NewReader(resSearch.Bytes())).Decode(&res)
			if err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Search - easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var res model.SearchResponse
			err := easyjson.UnmarshalFromReader(bytes.NewReader(resSearch.Bytes()), &res)
			if err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Cluster - json - map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var out = make(map[string]interface{})
			err := json.NewDecoder(bytes.NewReader(resClusterStats.Bytes())).Decode(&out)
			if err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("Cluster - json - stc", func(b *testing.B) {
		type ClusterHealthResponse struct {
			ClusterName string `json:"cluster_name"`
			Status      string
			Indices     struct {
				Count int
				Docs  struct {
					Count int
				}
			}
		}

		for i := 0; i < b.N; i++ {
			var out ClusterHealthResponse
			err := json.NewDecoder(bytes.NewReader(resClusterStats.Bytes())).Decode(&out)
			if err != nil {
				b.Error(err)
			}
			if len(out.ClusterName) < 3 {
				b.Errorf("Unexpected len(%s)=%d", out.ClusterName, len(out.ClusterName))
			}
		}
	})

	b.Run("Cluster - gjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var out []gjson.Result
			out = gjson.GetManyBytes(resClusterStats.Bytes(), "cluster_name", "status", "indices.count", "indices.docs.count")
			if len(out[0].String()) < 3 {
				b.Errorf("Unexpected len(%s)=%d", out[0], len(out[0].String()))
			}
		}
	})
}

func fixture(fname string) *bytes.Buffer {
	payload, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	return bytes.NewBuffer(payload)
}
