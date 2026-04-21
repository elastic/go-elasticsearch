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

//go:build search_default
// +build search_default

// This example demonstrates the Search API with the typed client
// (`TypedClient`) in two equivalent styles:
//
//   - building the query with the `esdsl` helpers (recommended), and
//   - building the query with raw `types.*` structs.
package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const demoIndex = "search-demo"

type article struct {
	Title    string `json:"title"`
	Category string `json:"category"`
}

func main() {
	log.SetFlags(0)

	typed, err := elasticsearch.NewTyped()
	if err != nil {
		log.Fatalf("Error creating typed client: %s", err)
	}
	defer closeClient(typed)

	ctx := context.Background()
	setupDemoData(ctx, typed)

	log.Println("=== Search using esdsl builders ===")
	runEsdslExample(ctx, typed)

	log.Println()
	log.Println("=== Search using raw types.* structs ===")
	runTypesExample(ctx, typed)
}

func closeClient(es *elasticsearch.TypedClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = es.Close(ctx)
}

func setupDemoData(ctx context.Context, es *elasticsearch.TypedClient) {
	if ok, _ := es.Indices.Exists(demoIndex).IsSuccess(ctx); ok {
		if _, err := es.Indices.Delete(demoIndex).Do(ctx); err != nil {
			log.Fatalf("Error deleting index: %s", err)
		}
	}
	if _, err := es.Indices.Create(demoIndex).Do(ctx); err != nil {
		log.Fatalf("Error creating index: %s", err)
	}

	docs := []article{
		{Title: "Go and Elasticsearch", Category: "tech"},
		{Title: "Elasticsearch typed client in practice", Category: "tech"},
		{Title: "Travel plans for spring", Category: "lifestyle"},
	}
	for i, doc := range docs {
		if _, err := es.Index(demoIndex).
			Id(strconv.Itoa(i + 1)).
			Document(doc).
			Do(ctx); err != nil {
			log.Fatalf("Error indexing document %d: %s", i+1, err)
		}
	}

	if _, err := es.Indices.Refresh().Index(demoIndex).Do(ctx); err != nil {
		log.Fatalf("Error refreshing index: %s", err)
	}

	log.Printf("Prepared demo index [%s] with %d documents", demoIndex, len(docs))
}

// runEsdslExample shows the recommended query-builder style using the
// helpers in the `esdsl` package. Each builder satisfies the matching
// `types.*Variant` interface, which the fluent `Search().Query(...)`
// setter accepts directly.
func runEsdslExample(ctx context.Context, es *elasticsearch.TypedClient) {
	res, err := es.Search().
		Index(demoIndex).
		Query(esdsl.NewBoolQuery().
			Must(esdsl.NewMatchQuery("title", "elasticsearch")).
			Filter(esdsl.NewTermQuery("category", esdsl.NewFieldValue().String("tech")))).
		Do(ctx)
	if err != nil {
		log.Fatalf("Error running search (esdsl): %s", err)
	}

	printHits(res)
}

// runTypesExample builds the same search directly from the raw
// `types.*` structs. Use this style when a builder is missing or when
// you prefer struct-literal composition.
func runTypesExample(ctx context.Context, es *elasticsearch.TypedClient) {
	category := "tech"
	titleQuery := "elasticsearch"

	req := &search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Must: []types.Query{
					{Match: map[string]types.MatchQuery{
						"title": {Query: titleQuery},
					}},
				},
				Filter: []types.Query{
					{Term: map[string]types.TermQuery{
						"category": {Value: &category},
					}},
				},
			},
		},
	}

	res, err := es.Search().
		Index(demoIndex).
		Request(req).
		Do(ctx)
	if err != nil {
		log.Fatalf("Error running search (types): %s", err)
	}

	printHits(res)
}

func printHits(res *search.Response) {
	total := int64(0)
	if res.Hits.Total != nil {
		total = res.Hits.Total.Value
	}
	log.Printf("total=%d hits=%d", total, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		id := ""
		if hit.Id_ != nil {
			id = *hit.Id_
		}
		log.Printf(" * [%s] %s", id, string(hit.Source_))
	}
}
