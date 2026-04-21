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

// This example demonstrates using a single typed endpoint against an
// existing functional (low-level) *elasticsearch.Client, without
// constructing a full *TypedClient.
//
// Every typed endpoint package (for example typedapi/core/search,
// typedapi/indices/create, typedapi/core/bulk) exports a `New`
// constructor that takes an elastictransport.Interface:
//
//	type Interface interface {
//	    Perform(*http.Request) (*http.Response, error)
//	}
//
// *elasticsearch.Client and *elasticsearch.TypedClient both embed
// BaseClient, which implements Perform. That means the existing
// low-level client can be passed straight into any typed endpoint
// constructor, sharing its transport, connection pool, retry policy,
// and instrumentation. No second client, no ToTyped() call, no
// MethodAPI tree: just the one endpoint you want to migrate.
//
// Use this when you want to try out the typed API for a single call
// site (for example, migrate search first) without committing to a
// full typed client yet. For a full typed surface over the same
// transport, see ./totyped.go.
//
// For fresh code that only ever needs a handful of typed endpoints,
// elasticsearch.NewBase(opts ...Option) returns a *BaseClient directly
// and skips allocating the esapi tree on *Client and the typedapi
// MethodAPI tree on *TypedClient. Same transport and configuration,
// no unused API surface — see the NewBase(...) block in main below.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
)

func main() {
	log.SetFlags(0)

	// Existing functional client. ELASTICSEARCH_URL is used when set,
	// otherwise localhost:9200.
	es, err := elasticsearch.New()
	if err != nil {
		log.Fatalf("error creating the client: %s", err)
	}
	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := es.Close(closeCtx); err != nil {
			log.Fatalf("error closing the client: %s", err)
		}
	}()

	ctx := context.Background()

	// Unchanged: existing call sites keep using the functional API.
	infoRes, err := es.Info()
	if err != nil {
		log.Fatalf("functional Info: %s", err)
	}
	infoRes.Body.Close()

	// Migrated: build just the typed Search endpoint against the
	// existing low-level client. search.New accepts anything that
	// implements elastictransport.Interface, and *elasticsearch.Client
	// satisfies it via the embedded BaseClient.Perform method.
	typedSearch := search.New(es)

	res, err := typedSearch.
		Index("example").
		Query(esdsl.NewMatchQuery("title", "typed")).
		Do(ctx)
	if err != nil {
		log.Fatalf("typed Search: %s", err)
	}

	fmt.Printf("hits: %d\n", len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		fmt.Printf(" * id=%s source=%s\n", *hit.Id_, hit.Source_)
	}

	// Alternative: if you are writing fresh code and only ever want
	// one or a few typed endpoints, skip both *Client and *TypedClient
	// entirely. elasticsearch.NewBase returns the shared *BaseClient
	// (transport, product check, meta headers) without allocating the
	// esapi tree or the typedapi MethodAPI tree. Pass it to any typed
	// endpoint constructor exactly like *Client above.
	base, err := elasticsearch.NewBase()
	if err != nil {
		log.Fatalf("error creating the base client: %s", err)
	}
	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := base.Close(closeCtx); err != nil {
			log.Fatalf("error closing the base client: %s", err)
		}
	}()

	if _, err := search.New(base).
		Index("example").
		Query(esdsl.NewMatchQuery("title", "typed")).
		Do(ctx); err != nil {
		log.Fatalf("typed Search (NewBase): %s", err)
	}
}
