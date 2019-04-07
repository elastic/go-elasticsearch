package main_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/tidwall/gjson"

	"github.com/elastic/go-elasticsearch/v8/_examples/encoding/model"
)

func BenchmarkSearchResults(b *testing.B) {
	b.ReportAllocs()

	input := fixture("testdata/response_search.json")

	b.Run("json", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var res model.SearchResponse
			err := json.NewDecoder(bytes.NewReader(input.Bytes())).Decode(&res)
			if err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("easyjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var res model.SearchResponse
			err := easyjson.UnmarshalFromReader(bytes.NewReader(input.Bytes()), &res)
			if err != nil {
				b.Error(err)
			}
		}
	})
}

func BenchmarkClusterStats(b *testing.B) {
	b.ReportAllocs()

	input := fixture("testdata/response_cluster_stats.json")

	b.Run("json - map", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var out = make(map[string]interface{})
			err := json.NewDecoder(bytes.NewReader(input.Bytes())).Decode(&out)
			if err != nil {
				b.Error(err)
			}
		}
	})

	b.Run("json - struct", func(b *testing.B) {
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
			err := json.NewDecoder(bytes.NewReader(input.Bytes())).Decode(&out)
			if err != nil {
				b.Error(err)
			}
			if len(out.ClusterName) < 3 {
				b.Errorf("Unexpected len(%s)=%d", out.ClusterName, len(out.ClusterName))
			}
		}
	})

	b.Run("gjson", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var out []gjson.Result
			out = gjson.GetManyBytes(input.Bytes(), "cluster_name", "status", "indices.count", "indices.docs.count")
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
