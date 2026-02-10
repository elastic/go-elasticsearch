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

// This example demonstrates the MSearchTemplate API with:
//   - the functional options API (`esapi`) using raw NDJSON, and
//   - the typed API (`typedapi`) using `msearchtemplate.Request`.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/msearchtemplate"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const demoIndex = "msearch-template-demo"

func main() {
	log.SetFlags(0)

	cfg := elasticsearch.Config{}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}
	defer closeClient(es)

	typed, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error creating typed client: %s", err)
	}
	defer closeTypedClient(typed)

	ctx := context.Background()
	setupDemoData(ctx, es)

	log.Println("=== MSearchTemplate using esapi (raw NDJSON) ===")
	runEsapiExample(ctx, es)

	log.Println()
	log.Println("=== MSearchTemplate using typedapi ===")
	runTypedAPIExample(ctx, typed)
}

func closeClient(es *elasticsearch.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = es.Close(ctx)
}

func closeTypedClient(es *elasticsearch.TypedClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = es.Close(ctx)
}

func setupDemoData(ctx context.Context, es *elasticsearch.Client) {
	deleteRes, err := es.Indices.Delete([]string{demoIndex}, es.Indices.Delete.WithContext(ctx))
	if err == nil && deleteRes != nil {
		deleteRes.Body.Close()
	}

	createRes, err := es.Indices.Create(demoIndex, es.Indices.Create.WithContext(ctx))
	if err != nil {
		log.Fatalf("Error creating index: %s", err)
	}
	defer createRes.Body.Close()
	if createRes.IsError() {
		log.Fatalf("Error creating index [%s]: %s", demoIndex, createRes.String())
	}

	docs := []string{
		`{"title":"Go and Elasticsearch","category":"tech"}`,
		`{"title":"Elasticsearch templates in practice","category":"tech"}`,
		`{"title":"Travel plans for spring","category":"lifestyle"}`,
	}

	for i, doc := range docs {
		id := fmt.Sprintf("%d", i+1)
		res, err := es.Index(
			demoIndex,
			strings.NewReader(doc),
			es.Index.WithContext(ctx),
			es.Index.WithDocumentID(id),
		)
		if err != nil {
			log.Fatalf("Error indexing document %s: %s", id, err)
		}
		if res.IsError() {
			res.Body.Close()
			log.Fatalf("Error indexing document %s: %s", id, res.String())
		}
		res.Body.Close()
	}

	refreshRes, err := es.Indices.Refresh(
		es.Indices.Refresh.WithContext(ctx),
		es.Indices.Refresh.WithIndex(demoIndex),
	)
	if err != nil {
		log.Fatalf("Error refreshing index: %s", err)
	}
	defer refreshRes.Body.Close()
	if refreshRes.IsError() {
		log.Fatalf("Error refreshing index [%s]: %s", demoIndex, refreshRes.String())
	}

	log.Printf("Prepared demo index [%s] with %d documents", demoIndex, len(docs))
}

func pString(s string) *string {
	return &s
}

func runEsapiExample(ctx context.Context, es *elasticsearch.Client) {
	var ndjson strings.Builder
	ndjson.WriteString(fmt.Sprintf(`{"index":"%s"}`+"\n", demoIndex))
	ndjson.WriteString(`{"source":{"query":{"match":{"title":"{{q}}"}}},"params":{"q":"elasticsearch"}}` + "\n")
	ndjson.WriteString(fmt.Sprintf(`{"index":"%s"}`+"\n", demoIndex))
	ndjson.WriteString(`{"source":{"query":{"match":{"category":"{{cat}}"}}},"params":{"cat":"tech"}}` + "\n")

	res, err := es.MsearchTemplate(
		strings.NewReader(ndjson.String()),
		es.MsearchTemplate.WithContext(ctx),
		es.MsearchTemplate.WithHeader(map[string]string{
			"Content-Type": "application/x-ndjson",
		}),
	)
	if err != nil {
		log.Fatalf("Error running msearch template (esapi): %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("MSearchTemplate failed (esapi): %s", res.String())
	}

	var payload struct {
		Responses []json.RawMessage `json:"responses"`
	}
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		log.Fatalf("Error decoding msearch template response (esapi): %s", err)
	}

	log.Printf("responses=%d", len(payload.Responses))
}

func runTypedAPIExample(ctx context.Context, typed *elasticsearch.TypedClient) {
	paramsTech, err := json.Marshal("tech")
	if err != nil {
		log.Fatalf("Error marshaling params: %s", err)
	}
	paramsElastic, err := json.Marshal("elasticsearch")
	if err != nil {
		log.Fatalf("Error marshaling params: %s", err)
	}

	req := msearchtemplate.Request{
		types.MultisearchHeader{Index: []string{demoIndex}},
		types.TemplateConfig{
			Source: pString(`{"query":{"match":{"category":"{{cat}}"}}}`),
			Params: map[string]json.RawMessage{"cat": paramsTech},
		},
		types.MultisearchHeader{Index: []string{demoIndex}},
		types.TemplateConfig{
			Source: pString(`{"query":{"match":{"title":"{{q}}"}}}`),
			Params: map[string]json.RawMessage{"q": paramsElastic},
		},
	}

	resp, err := typed.MsearchTemplate().Request(&req).Do(ctx)
	if err != nil {
		log.Fatalf("Error running msearch template (typedapi): %s", err)
	}

	for i, item := range resp.Responses {
		switch v := item.(type) {
		case *types.MultiSearchItem:
			log.Printf("response[%d] status=OK hits=%d", i, len(v.Hits.Hits))
		case *types.ErrorResponseBase:
			log.Printf("response[%d] status=%d error=%v", i, v.Status, v.Error.Reason)
		default:
			log.Printf("response[%d] unexpected type %T", i, item)
		}
	}
}
