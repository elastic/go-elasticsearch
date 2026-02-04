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

//go:build integration && !multinode
// +build integration,!multinode

package e2e_test

import (
	"bytes"
	"compress/bzip2"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"
	"testing"
	"testing/containertest"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorindexoptionstype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/densevectorsimilarity"
)

type Doc struct {
	DocID string    `json:"docid"`
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Emb   []float32 `json:"emb"`
}

type B64Doc struct {
	DocID string               `json:"docid"`
	Title string               `json:"title"`
	Text  string               `json:"text"`
	Emb   types.DenseVectorF32 `json:"emb"` // base64 encoded
}

type Run struct {
	DatasetSize int
	ChunkSize   int
}

type RunResult struct {
	DatasetSize int `json:"dataset_size"`
	ChunkSize   int `json:"chunk_size"`
	Float32     struct {
		Duration int `json:"duration"`
	} `json:"float32"`
	Base64 struct {
		Duration int `json:"duration"`
	} `json:"base64"`
}

func TestBase64BulkIndexing(t *testing.T) {
	stackVersion := elasticsearch.Version
	if v := os.Getenv("STACK_VERSION"); v != "" {
		stackVersion = v
	}

	// If ELASTICSEARCH_URL is set use the external cluster instead of starting a testcontainer.
	var (
		esClient         *elasticsearch.TypedClient
		err              error
		elasticsearchSrv *containertest.ElasticsearchService
	)

	if envURL, ok := os.LookupEnv("ELASTICSEARCH_URL"); ok && envURL != "" {
		t.Logf("ELASTICSEARCH_URL is set, using external Elasticsearch at %s", envURL)
		// Create a config that will make the client pick addresses from the environment.
		cfg := elasticsearch.Config{}
		cfg.CompressRequestBody = false
		esClient, err = elasticsearch.NewTypedClient(cfg)
		if err != nil {
			t.Fatalf("Error creating the client from environment: %s", err)
		}
	} else {
		// Start a testcontainer Elasticsearch instance as before.
		elasticsearchSrv, err = containertest.NewElasticsearchService(stackVersion, containertest.WithResolveLatestPatch(true))
		if err != nil {
			t.Fatalf("Error setting up Elasticsearch container: %s", err)
		}
		defer func() {
			if err := elasticsearchSrv.Terminate(context.Background()); err != nil {
				t.Fatalf("Error terminating Elasticsearch container: %s", err)
			}
		}()

		cfg := elasticsearchSrv.ESConfig()
		cfg.CompressRequestBody = false
		esClient, err = elasticsearch.NewTypedClient(cfg)
		if err != nil {
			t.Fatalf("Error creating the client: %s", err)
		}
	}
	defer func() {
		if esClient != nil {
			if err := esClient.Close(context.Background()); err != nil {
				t.Fatalf("Error closing the client: %s", err)
			}
		}
	}()

	// we get the file from https://rally-tracks.elastic.co/openai_vector/open_ai_corpus-initial-indexing-1k.json.bz2
	// store it and load it in memory, then use it for bulk indexing
	// json format is a bunch of documents separated by newlines compressed in bzip2
	track := "https://rally-tracks.elastic.co/openai_vector/open_ai_corpus-initial-indexing-1k.json.bz2"
	res, err := http.Get(track)
	if err != nil {
		t.Fatalf("Error downloading track file: %s", err)
	}
	defer func() {
		if cerr := res.Body.Close(); cerr != nil {
			t.Logf("Error closing track response body: %s", cerr)
		}
	}()

	docs := []Doc{}
	b64Docs := []B64Doc{}
	// we need to bunzip2 the data to read the json lines
	reader := bzip2.NewReader(res.Body)
	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("Error reading track file: %s", err)
	}
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		var doc Doc
		var b64Doc B64Doc
		if len(line) == 0 {
			continue
		}
		if err := json.Unmarshal(line, &doc); err != nil {
			t.Fatalf("Error unmarshaling document: %s", err)
		}
		if err := json.Unmarshal(line, &b64Doc); err != nil {
			t.Fatalf("Error unmarshaling B64 document: %s", err)
		}
		docs = append(docs, doc)
		b64Docs = append(b64Docs, b64Doc)
	}

	repeats := 20
	iterations := 3
	runChunkSizes := []int{100, 250, 500, 1000}
	if testing.Short() {
		repeats = 2
		iterations = 1
		runChunkSizes = []int{200}
	}

	docsAll := make([]Doc, 0, len(docs)*repeats)
	b64DocsAll := make([]B64Doc, 0, len(b64Docs)*repeats)
	// repeat to have a larger dataset
	for i := 0; i < repeats; i++ {
		docsAll = append(docsAll, docs...)
		b64DocsAll = append(b64DocsAll, b64Docs...)
	}

	t.Run("Testing vectors as []float32 and base64", func(t *testing.T) {
		var index = "vec-test"

		if testing.Short() {
			t.Logf("Short mode enabled: repeats=%d iterations=%d chunkSizes=%v", repeats, iterations, runChunkSizes)
		}

		runs := make([]Run, 0, len(runChunkSizes))
		for _, chunkSize := range runChunkSizes {
			runs = append(runs, Run{DatasetSize: len(docsAll), ChunkSize: chunkSize})
		}

		results := []RunResult{}

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered. Error:\n", r)
			}
		}()

		for _, run := range runs {
			rr := RunResult{
				DatasetSize: run.DatasetSize,
				ChunkSize:   run.ChunkSize,
			}

			floatDuration := []int{}
			base64Duration := []int{}
			for i := 0; i < iterations; i++ {
				// float32
				floatDurationAsMillis := 0
				setupElasticsearch(t, esClient, index)

				for chunk := range slices.Chunk(docsAll, run.ChunkSize) {
					d, err := ingestDocs(t, esClient, index, chunk)
					if err != nil {
						t.Fatalf("Error executing bulk request for float32 with %d docs: %s", run.DatasetSize, err)
					}
					floatDurationAsMillis += int(d.Milliseconds())
				}
				floatDuration = append(floatDuration, floatDurationAsMillis)

				// base64
				base64DurationAsMillis := 0
				setupElasticsearch(t, esClient, index)
				for chunk := range slices.Chunk(b64DocsAll, run.ChunkSize) {
					d, err := ingestDocs(t, esClient, index, chunk)
					if err != nil {
						t.Fatalf("Error executing bulk request for base64 with %d docs: %s", run.DatasetSize, err)
					}
					base64DurationAsMillis += int(d.Milliseconds())
				}
				base64Duration = append(base64Duration, base64DurationAsMillis)
			}
			// calculate average duration
			totalFloat := 0
			for _, d := range floatDuration {
				totalFloat += d
			}
			avgFloat := totalFloat / len(floatDuration)

			totalBase64 := 0
			for _, d := range base64Duration {
				totalBase64 += d
			}
			avgBase64 := totalBase64 / len(base64Duration)

			rr.Float32.Duration = avgFloat
			rr.Base64.Duration = avgBase64

			// append result
			results = append(results, rr)
		}

		for _, result := range results {
			t.Logf("Dataset Size: %d, Chunk Size: %d, Float32 Duration: %d ms, Base64 Duration: %d ms",
				result.DatasetSize,
				result.ChunkSize,
				result.Float32.Duration,
				result.Base64.Duration,
			)
			t.Logf("Documents per second (float32): %.2f", float32(result.DatasetSize*1000)/float32(result.Float32.Duration))
			t.Logf("Documents per second (base64): %.2f", float32(result.DatasetSize*1000)/float32(result.Base64.Duration))
			t.Logf("Base64 is %.2fx times faster than float32",
				float32(result.Float32.Duration)/float32(result.Base64.Duration),
			)
			t.Log("\n")
		}

		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			t.Fatalf("Error marshaling results: %s", err)
		}
		t.Log(string(data))
	})
}

func setupElasticsearch(t *testing.T, es *elasticsearch.TypedClient, index string) {
	t.Helper()
	if ok, _ := es.Indices.Exists(index).IsSuccess(context.Background()); ok {
		_, _ = es.Indices.Delete(index).Do(context.Background())
	}

	mappings := esdsl.NewTypeMapping().
		AddProperty("docid", esdsl.NewKeywordProperty()).
		AddProperty("title", esdsl.NewTextProperty()).
		AddProperty("text", esdsl.NewTextProperty().AddField("keyword", esdsl.NewKeywordProperty().IgnoreAbove(256))).
		AddProperty("emb", esdsl.NewDenseVectorProperty().
			Dims(1536).
			Index(true).
			Similarity(densevectorsimilarity.Cosine).
			IndexOptions(esdsl.NewDenseVectorIndexOptions(densevectorindexoptionstype.Flat)))

	_, err := es.Indices.
		Create(index).
		Mappings(mappings).
		Do(context.Background())
	if err != nil {
		t.Fatalf("Error creating index: %s", err)
	}
}

func ingestDocs[T any](t *testing.T, client *elasticsearch.TypedClient, index string, docs []T) (time.Duration, error) {
	t.Helper()
	start := time.Now()
	bulk := client.Bulk()
	for _, doc := range docs {
		if err := bulk.IndexOp(types.IndexOperation{Index_: &index}, doc); err != nil {
			t.Fatalf("Error adding bulk operation: %s", err)
		}
	}
	_, err := bulk.Do(context.Background())
	elapsed := time.Since(start)
	return elapsed, err
}
