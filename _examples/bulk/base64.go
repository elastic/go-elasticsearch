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

//go:build bulk_base64
// +build bulk_base64

// This example demonstrates indexing dense vectors using base64 encoding
// with the typed Bulk API.
//
//	go run base64.go -count=20000 -batch=500
package main

import (
	"bytes"
	"compress/bzip2"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/dustin/go-humanize"

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

var (
	indexName string
	count     int
	batch     int
)

func init() {
	flag.StringVar(&indexName, "index", "openai-vector-bulk", "Index name")
	flag.IntVar(&count, "count", 20000, "Number of documents to index")
	flag.IntVar(&batch, "batch", 500, "Number of documents to send in one batch")
	flag.Parse()
}

func main() {
	log.SetFlags(0)

	log.Printf(
		"\x1b[1mBulk Base64\x1b[0m: documents [%s] batch size [%s]",
		humanize.Comma(int64(count)),
		humanize.Comma(int64(batch)),
	)
	log.Println("▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁")

	cfg := elasticsearch.Config{
		CompressRequestBody: false,
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(ctx); err != nil {
			log.Fatalf("Error closing the client: %s", err)
		}
	}()

	docs, err := loadB64Docs(count)
	if err != nil {
		log.Fatalf("Error loading documents: %s", err)
	}
	log.Printf("→ Loaded %s documents", humanize.Comma(int64(len(docs))))

	ctx := context.Background()
	if err := setupIndex(ctx, es, indexName); err != nil {
		log.Fatalf("Error setting up index: %s", err)
	}

	start := time.Now().UTC()
	for chunk := range slices.Chunk(docs, batch) {
		if _, err := ingestDocs(ctx, es, indexName, chunk); err != nil {
			log.Fatalf("Error executing bulk request: %s", err)
		}
	}

	log.Println("▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔▔")

	dur := time.Since(start)
	rate := float64(len(docs)) / dur.Seconds()
	log.Printf(
		"Sucessfuly indexed [%s] documents in %s (%.0f docs/sec)",
		humanize.Comma(int64(len(docs))),
		dur.Truncate(time.Millisecond),
		rate,
	)
}

func loadB64Docs(count int) ([]B64Doc, error) {
	track := "https://rally-tracks.elastic.co/openai_vector/open_ai_corpus-initial-indexing-1k.json.bz2"
	res, err := http.Get(track)
	if err != nil {
		return nil, fmt.Errorf("download track file: %w", err)
	}
	defer func() {
		if cerr := res.Body.Close(); cerr != nil {
			log.Printf("Error closing track response body: %s", cerr)
		}
	}()
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf("download track file: %s", res.Status)
	}

	reader := bzip2.NewReader(res.Body)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("read track file: %w", err)
	}

	lines := bytes.Split(data, []byte("\n"))
	b64Docs := make([]B64Doc, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var doc B64Doc
		if err := json.Unmarshal(line, &doc); err != nil {
			return nil, fmt.Errorf("unmarshal document: %w", err)
		}
		b64Docs = append(b64Docs, doc)
	}
	if len(b64Docs) == 0 {
		return nil, fmt.Errorf("no documents found in track")
	}

	if count <= len(b64Docs) {
		return b64Docs[:count], nil
	}

	expanded := make([]B64Doc, 0, count)
	for len(expanded) < count {
		remaining := count - len(expanded)
		if remaining >= len(b64Docs) {
			expanded = append(expanded, b64Docs...)
			continue
		}
		expanded = append(expanded, b64Docs[:remaining]...)
	}

	return expanded, nil
}

func setupIndex(ctx context.Context, es *elasticsearch.TypedClient, index string) error {
	if ok, _ := es.Indices.Exists(index).IsSuccess(ctx); ok {
		if _, err := es.Indices.Delete(index).Do(ctx); err != nil {
			return fmt.Errorf("delete index: %w", err)
		}
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

	if _, err := es.Indices.
		Create(index).
		Mappings(mappings).
		Do(ctx); err != nil {
		return fmt.Errorf("create index: %w", err)
	}

	return nil
}

func ingestDocs(ctx context.Context, client *elasticsearch.TypedClient, index string, docs []B64Doc) (time.Duration, error) {
	start := time.Now()
	bulk := client.Bulk()
	for _, doc := range docs {
		if err := bulk.IndexOp(types.IndexOperation{Index_: &index}, doc); err != nil {
			return 0, fmt.Errorf("add bulk operation: %w", err)
		}
	}
	_, err := bulk.Do(ctx)
	return time.Since(start), err
}
