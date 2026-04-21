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

// This example demonstrates a basic usage of the Elasticsearch Go client
// with the typed client (`TypedClient`) and the `esdsl` query builders.
//
// It fetches the version information from the cluster, indexes a couple
// of documents concurrently, and performs a search request.

package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

type article struct {
	Title string `json:"title"`
}

func main() {
	log.SetFlags(0)

	// Initialize a typed client with the default settings.
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//
	es, err := elasticsearch.NewTyped()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(ctx); err != nil {
			log.Fatalf("Error closing the client: %s\n", err)
		}
	}()

	ctx := context.Background()

	// 1. Get cluster info
	//
	info, err := es.Info().Do(ctx)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", info.Version.Int)
	log.Println(strings.Repeat("~", 37))

	// 2. Index documents concurrently
	//
	var wg sync.WaitGroup
	for i, title := range []string{"Test One", "Test Two"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			res, err := es.Index("test").
				Id(strconv.Itoa(i + 1)).
				Document(article{Title: title}).
				Refresh(refresh.True).
				Do(ctx)
			if err != nil {
				log.Printf("Error indexing document ID=%d: %s", i+1, err)
				return
			}
			// Print the result and indexed document version.
			log.Printf("%s; version=%d", res.Result, res.Version_)
		}(i, title)
	}
	wg.Wait()

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents
	//
	res, err := es.Search().
		Index("test").
		Query(esdsl.NewMatchQuery("title", "test")).
		TrackTotalHits(esdsl.NewTrackHits().Bool(true)).
		Pretty(true).
		Do(ctx)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	// Print the number of results and request duration.
	total := int64(0)
	if res.Hits.Total != nil {
		total = res.Hits.Total.Value
	}
	log.Printf("%d hits; took: %dms", total, res.Took)
	// Print the ID and document source for each hit.
	for _, hit := range res.Hits.Hits {
		id := ""
		if hit.Id_ != nil {
			id = *hit.Id_
		}
		log.Printf(" * ID=%s, %s", id, string(hit.Source_))
	}

	log.Println(strings.Repeat("=", 37))
}
