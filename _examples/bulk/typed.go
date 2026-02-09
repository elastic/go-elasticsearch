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

//go:build bulk_typed
// +build bulk_typed

// This example demonstrates indexing documents using the typed Bulk API.
//
// You can configure the number of documents and the batch size with command line flags:
//
//	go run typed.go -count=10000 -batch=1000
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operationtype"
)

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Published time.Time `json:"published"`
}

var (
	indexName string
	count     int
	batch     int
)

func init() {
	flag.StringVar(&indexName, "index", "articles-typed-bulk", "Index name")
	flag.IntVar(&count, "count", 10000, "Number of documents to generate")
	flag.IntVar(&batch, "batch", 1000, "Number of documents to send in one batch")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.SetFlags(0)

	log.Printf(
		"\x1b[1mTyped Bulk\x1b[0m: documents [%s] batch size [%s]",
		humanize.Comma(int64(count)),
		humanize.Comma(int64(batch)),
	)
	log.Println(strings.Repeat("▁", 65))

	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := client.Close(ctx); err != nil {
			log.Fatalf("Error closing the client: %s", err)
		}
	}()

	ctx := context.Background()
	if err := setupIndex(ctx, client, indexName); err != nil {
		log.Fatalf("Error setting up index: %s", err)
	}

	articles := make([]Article, 0, count)
	for i := 1; i <= count; i++ {
		articles = append(articles, Article{
			ID:        i,
			Title:     fmt.Sprintf("Title %d", i),
			Body:      "Lorem ipsum dolor sit amet...",
			Published: time.Now().Round(time.Second).UTC().AddDate(0, 0, rand.Intn(365)),
		})
	}
	log.Printf("→ Generated %s articles", humanize.Comma(int64(len(articles))))

	start := time.Now().UTC()

	var numFailed int
	for chunk := range slices.Chunk(articles, batch) {
		bulk := client.Bulk()
		for _, a := range chunk {
			id := strconv.Itoa(a.ID)
			if err := bulk.IndexOp(types.IndexOperation{Index_: &indexName, Id_: &id}, a); err != nil {
				log.Fatalf("Error adding bulk operation: %s", err)
			}
		}

		res, err := bulk.Do(ctx)
		if err != nil {
			log.Fatalf("Error executing bulk request: %s", err)
		}
		if res.Errors {
			numFailed += countFailedItems(res.Items)
		}
	}

	log.Println(strings.Repeat("▔", 65))

	dur := time.Since(start)
	if numFailed > 0 {
		log.Fatalf(
			"Indexed [%s] documents with [%s] errors in %s (%.0f docs/sec)",
			humanize.Comma(int64(len(articles))),
			humanize.Comma(int64(numFailed)),
			dur.Truncate(time.Millisecond),
			float64(len(articles))/dur.Seconds(),
		)
	}

	log.Printf(
		"Successfully indexed [%s] documents in %s (%.0f docs/sec)",
		humanize.Comma(int64(len(articles))),
		dur.Truncate(time.Millisecond),
		float64(len(articles))/dur.Seconds(),
	)
}

func setupIndex(ctx context.Context, client *elasticsearch.TypedClient, index string) error {
	if ok, _ := client.Indices.Exists(index).IsSuccess(ctx); ok {
		if _, err := client.Indices.Delete(index).Do(ctx); err != nil {
			return fmt.Errorf("delete index: %w", err)
		}
	}

	if _, err := client.Indices.Create(index).Do(ctx); err != nil {
		return fmt.Errorf("create index: %w", err)
	}

	return nil
}

func countFailedItems(items []map[operationtype.OperationType]types.ResponseItem) int {
	var failed int
	for _, item := range items {
		for _, res := range item {
			if res.Status > 299 {
				failed++
			}
		}
	}
	return failed
}
