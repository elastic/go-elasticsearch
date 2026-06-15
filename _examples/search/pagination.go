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

//go:build search_pagination
// +build search_pagination

// This example demonstrates two ways to paginate through search results
// with the typed client (`TypedClient`):
//
//   - `from` + `size` for small result sets, and
//   - a point in time (PIT) plus `search_after` for deep pagination over a
//     consistent snapshot of the index.
package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

const (
	paginationIndex = "search-pagination-demo"
	totalDocs       = 25
	pageSize        = 10
)

type product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	log.SetFlags(0)

	es, err := elasticsearch.NewTyped()
	if err != nil {
		log.Fatalf("Error creating typed client: %s", err)
	}
	defer closePaginationClient(es)

	ctx := context.Background()
	setupPaginationData(ctx, es)

	log.Println("=== Pagination with from + size ===")
	runFromSizeExample(ctx, es)

	log.Println()
	log.Println("=== Pagination with PIT + search_after ===")
	runPITSearchAfterExample(ctx, es)
}

func closePaginationClient(es *elasticsearch.TypedClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = es.Close(ctx)
}

func setupPaginationData(ctx context.Context, es *elasticsearch.TypedClient) {
	if ok, _ := es.Indices.Exists(paginationIndex).IsSuccess(ctx); ok {
		if _, err := es.Indices.Delete(paginationIndex).Do(ctx); err != nil {
			log.Fatalf("Error deleting index: %s", err)
		}
	}
	if _, err := es.Indices.Create(paginationIndex).Do(ctx); err != nil {
		log.Fatalf("Error creating index: %s", err)
	}

	for i := 1; i <= totalDocs; i++ {
		doc := product{Name: "product-" + strconv.Itoa(i), Price: i * 10}
		if _, err := es.Index(paginationIndex).
			Id(strconv.Itoa(i)).
			Document(doc).
			Do(ctx); err != nil {
			log.Fatalf("Error indexing document %d: %s", i, err)
		}
	}

	if _, err := es.Indices.Refresh().Index(paginationIndex).Do(ctx); err != nil {
		log.Fatalf("Error refreshing index: %s", err)
	}

	log.Printf("Prepared demo index [%s] with %d documents", paginationIndex, totalDocs)
}

// runFromSizeExample walks pages using the `from` and `size` parameters.
//
// This is the simplest form of pagination and is the right tool for small
// result sets. Elasticsearch refuses to page past the 10,000th hit with
// `from` + `size` by default (the `index.max_result_window` setting), so
// for anything deeper use PIT + search_after (see runPITSearchAfterExample).
func runFromSizeExample(ctx context.Context, es *elasticsearch.TypedClient) {
	seen := 0
	for page := 0; ; page++ {
		res, err := es.Search().
			Index(paginationIndex).
			Query(esdsl.NewMatchAllQuery()).
			From(page * pageSize).
			Size(pageSize).
			Do(ctx)
		if err != nil {
			log.Fatalf("Error running search (from/size): %s", err)
		}
		log.Printf("page=%d hits=%d", page, len(res.Hits.Hits))
		seen += len(res.Hits.Hits)
		if len(res.Hits.Hits) < pageSize {
			break
		}
	}
	log.Printf("from/size: visited %d documents", seen)
}

// runPITSearchAfterExample walks pages using a point in time (PIT) and
// `search_after`. This is the recommended approach for deep pagination
// and for exports where a consistent view of the index matters across
// pages (PIT pins a snapshot; concurrent writes don't shift results).
//
// A deterministic sort is required for `search_after` to have stable
// cursor values. `_shard_doc` is the cheapest tie-breaker and is
// available whenever a PIT is in use.
func runPITSearchAfterExample(ctx context.Context, es *elasticsearch.TypedClient) {
	pit, err := es.OpenPointInTime(paginationIndex).KeepAlive("1m").Do(ctx)
	if err != nil {
		log.Fatalf("Error opening PIT: %s", err)
	}
	defer func() {
		if _, err := es.ClosePointInTime().Id(pit.Id).Do(ctx); err != nil {
			log.Printf("Error closing PIT: %s", err)
		}
	}()

	// Note: no Index(...) here. The PIT already pins the target indices;
	// setting Index together with Pit is rejected by the server.
	req := es.Search().
		Query(esdsl.NewMatchAllQuery()).
		Pit(esdsl.NewPointInTimeReference().
			Id(pit.Id).
			KeepAlive(esdsl.NewDuration().String("1m"))).
		Sort(esdsl.NewSortOptions().
			AddSortOption("_shard_doc", esdsl.NewFieldSort(sortorder.Asc))).
		Size(pageSize)

	seen := 0
	for page := 0; ; page++ {
		res, err := req.Do(ctx)
		if err != nil {
			log.Fatalf("Error running search (search_after): %s", err)
		}
		if len(res.Hits.Hits) == 0 {
			break
		}
		log.Printf("page=%d hits=%d", page, len(res.Hits.Hits))
		seen += len(res.Hits.Hits)

		req = nextPITPage(req, res)
	}
	log.Printf("PIT + search_after: visited %d documents", seen)
}

// nextPITPage advances the builder to fetch the page after res. It reuses
// the same builder so any Query, Size, or Sort settings carry over. The
// PIT id can rotate between requests, so prefer res.PitId when the server
// returned one.
func nextPITPage(req *search.Search, res *search.Response) *search.Search {
	last := res.Hits.Hits[len(res.Hits.Hits)-1]
	req = req.SearchAfterValues(last.Sort)
	if res.PitId != nil {
		req = req.Pit(esdsl.NewPointInTimeReference().
			Id(*res.PitId).
			KeepAlive(esdsl.NewDuration().String("1m")))
	}
	return req
}
