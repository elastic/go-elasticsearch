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

// This example demonstrates the built-in OpenTelemetry integration.
//
// elasticsearch.NewOpenTelemetryInstrumentation returns an
// elastictransport.Instrumentation that can be passed to the client via
// WithInstrumentation (or Config.Instrumentation). It produces spans that
// follow the Elasticsearch OpenTelemetry semantic conventions for every
// API call, with no per-call wiring required.
//
// Prefer this built-in hook over writing custom tracing interceptors. A
// manual, interceptor-based alternative (for when you need to extend the
// behaviour) lives at _examples/interceptor/cmd/custom_observability.
package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/esdsl"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func main() {
	// 1. Set up an OpenTelemetry TracerProvider that prints spans to stdout.
	//    In a real application this would export to an OTLP endpoint.
	tp := newTracerProvider()
	defer func() { _ = tp.Shutdown(context.Background()) }()

	// 2. Start a minimal fake Elasticsearch server so the example runs
	//    without a live cluster. Replace with WithAddresses("https://...")
	//    in production.
	srv := newFakeES()
	defer srv.Close()

	// 3. Create the client with the built-in OTel instrumentation.
	//    captureSearchBody includes the search query in the span
	//    attributes for the search-family endpoints.
	captureSearchBody := true
	es, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses(srv.URL),
		elasticsearch.WithInstrumentation(
			elasticsearch.NewOpenTelemetryInstrumentation(tp, captureSearchBody),
		),
	)
	if err != nil {
		log.Fatalf("error creating client: %s", err)
	}

	ctx := context.Background()

	// 4. Each call below emits a span: Info, Index, Search.
	if _, err := es.Info().Do(ctx); err != nil {
		log.Fatalf("info: %s", err)
	}

	if _, err := es.Index("example").
		Id("1").
		Document(map[string]any{"title": "Hello, OpenTelemetry"}).
		Do(ctx); err != nil {
		log.Fatalf("index: %s", err)
	}

	if _, err := es.Search().
		Index("example").
		Query(esdsl.NewMatchQuery("title", "opentelemetry")).
		Do(ctx); err != nil {
		log.Fatalf("search: %s", err)
	}

	// 5. Give the exporter a moment to flush before the process exits.
	time.Sleep(100 * time.Millisecond)
}

func newTracerProvider() *sdktrace.TracerProvider {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("stdouttrace: %s", err)
	}
	res, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("go-elasticsearch-otel-example"),
		),
	)
	return sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSyncer(exporter),
	)
}

func newFakeES() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/":
			_, _ = w.Write([]byte(`{"version":{"number":"9.0.0"},"tagline":"You Know, for Search"}`))
		case r.URL.Path == "/example/_search":
			_, _ = w.Write([]byte(`{"took":1,"timed_out":false,"hits":{"total":{"value":0,"relation":"eq"},"hits":[]}}`))
		default:
			_, _ = w.Write([]byte(`{"result":"created","_id":"1","_version":1}`))
		}
	}))
}
