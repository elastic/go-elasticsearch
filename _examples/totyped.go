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

//go:build ignore
// +build ignore

// This example demonstrates incremental migration from the functional
// (low-level) API to the typed API.
//
// (*Client).ToTyped() returns a *TypedClient that shares the source
// client's transport, connection pool, compatibility header, and
// instrumentation. There is no second transport and no second product
// check once both clients have exchanged a first request: the existing
// *Client keeps working for code that has not been migrated, while new
// code uses the typed builder API through the returned *TypedClient.
//
// For the full migration guide see docs/reference/low-level-api/migration.md.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
)

func main() {
	log.SetFlags(0)

	// Existing application code: a functional (low-level) client.
	// ELASTICSEARCH_URL is used when set, otherwise localhost:9200.
	es, err := elasticsearch.New()
	if err != nil {
		log.Fatalf("error creating the client: %s", err)
	}
	// The functional and typed clients share the same transport (see
	// ToTyped below), so closing es also closes the pool used by the
	// typed client.
	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(closeCtx); err != nil {
			log.Fatalf("error closing the client: %s", err)
		}
	}()

	ctx := context.Background()

	// Unchanged call site using the functional API. Existing code that
	// depends on esapi response shapes keeps compiling as-is.
	infoRes, err := es.Info()
	if err != nil {
		log.Fatalf("functional Info: %s", err)
	}
	defer infoRes.Body.Close()

	var info struct {
		Version struct {
			Number string `json:"number"`
		} `json:"version"`
	}
	if err := json.NewDecoder(infoRes.Body).Decode(&info); err != nil {
		log.Fatalf("decode: %s", err)
	}
	log.Printf("[functional] server version: %s", info.Version.Number)

	// Migrate one call at a time. ToTyped reuses the same transport,
	// connection pool, and configuration, so no second client setup is
	// required. Call it once and reuse the result.
	typed := es.ToTyped()

	// Migrated call site using the typed builder API. The response is
	// a decoded Go struct; no body.Close, no manual JSON parsing.
	typedInfo, err := typed.Info().Do(ctx)
	if err != nil {
		log.Fatalf("typed Info: %s", err)
	}
	log.Printf("[typed] server version: %s", typedInfo.Version.Int)

	log.Println(strings.Repeat("-", 37))

	// Mix freely: keep using the functional client for endpoints that
	// have not been migrated yet, and the typed client for the ones
	// that have.
	if _, err := typed.Index("example").
		Id("1").
		Document(map[string]any{"title": "Hello, typed API"}).
		Do(ctx); err != nil {
		log.Fatalf("typed Index: %s", err)
	}

	searchRes, err := typed.Search().
		Index("example").
		Query(esdsl.NewMatchQuery("title", "typed")).
		Do(ctx)
	if err != nil {
		log.Fatalf("typed Search: %s", err)
	}
	for _, hit := range searchRes.Hits.Hits {
		fmt.Printf(" * id=%s source=%s\n", *hit.Id_, hit.Source_)
	}
}
