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

// This example demonstrates how to use Interceptors to add custom
// observability to Elasticsearch requests using OpenTelemetry.
//
// Note: There is existing, built-in functionality to add observability to Elasticsearch requests.
// Prefer using the built-in functionality over this example where possible.
//
// It shows three interceptors for:
// - Logging: Request/response details using slog
// - Metrics: Request counter and duration histogram
// - Tracing: Distributed tracing with spans
//
// All telemetry is exported to stdout for demonstration purposes.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/_examples/interceptor/internal/fake"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	// Initialize OpenTelemetry providers
	shutdown := initOTel()
	defer shutdown()

	// Start a fake Elasticsearch server
	srv := fake.NewServer(
		fake.WithStatusCode(http.StatusOK),
		fake.WithResponseBody([]byte(`{"cluster_name":"example","version":{"number":"9.2.0"}}`)),
		fake.WithHeaders(func(h http.Header) {
			h.Set("X-Elastic-Product", "Elasticsearch")
			h.Set("Content-Type", "application/json")
		}),
	)
	defer srv.Close()

	// Create metrics instruments
	meter := otel.Meter("elasticsearch-client")
	requestCounter, _ := meter.Int64Counter("elasticsearch.client.requests",
		metric.WithDescription("Number of requests to Elasticsearch"),
		metric.WithUnit("{request}"),
	)
	requestDuration, _ := meter.Float64Histogram("elasticsearch.client.duration",
		metric.WithDescription("Duration of Elasticsearch requests"),
		metric.WithUnit("ms"),
	)

	// Create tracer
	tracer := otel.Tracer("elasticsearch-client")

	// Create an Elasticsearch client with observability interceptors
	es, err := elasticsearch.New(
		elasticsearch.WithAddresses(srv.URL()),
		elasticsearch.WithTransportOptions(elastictransport.WithInterceptors(
			LoggingInterceptor(),
			MetricsInterceptor(requestCounter, requestDuration),
			TracingInterceptor(tracer),
		)),
	)
	if err != nil {
		panic(err)
	}

	// Send some requests to demonstrate the observability
	fmt.Println(">>> Sending requests to demonstrate observability interceptors")
	fmt.Println()

	for i := 1; i <= 3; i++ {
		fmt.Printf("--- Request %d ---\n", i)
		_, _ = es.Info()
		fmt.Println()
		time.Sleep(100 * time.Millisecond) // Small delay between requests
	}

	// Give time for metrics to flush
	fmt.Println(">>> Waiting for metrics to flush...")
	time.Sleep(2 * time.Second)
}

// initOTel initializes OpenTelemetry with stdout exporters for both tracing and metrics.
func initOTel() func() {
	ctx := context.Background()

	// Create resource with service info
	res, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("elasticsearch-example"),
		),
	)

	// Initialize trace exporter (stdout with pretty print)
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	// Initialize tracer provider
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSyncer(traceExporter), // Use Syncer for immediate output
	)
	otel.SetTracerProvider(tracerProvider)

	// Initialize metric exporter (stdout)
	metricExporter, err := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	// Initialize meter provider with periodic reader
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter, sdkmetric.WithInterval(time.Second))),
	)
	otel.SetMeterProvider(meterProvider)

	// Return shutdown function
	return func() {
		_ = tracerProvider.Shutdown(ctx)
		_ = meterProvider.Shutdown(ctx)
	}
}

// LoggingInterceptor creates an interceptor that logs request and response details.
func LoggingInterceptor() elastictransport.InterceptorFunc {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	return func(next elastictransport.RoundTripFunc) elastictransport.RoundTripFunc {
		return func(req *http.Request) (*http.Response, error) {
			start := time.Now()

			logger.Info("elasticsearch request started",
				slog.String("method", req.Method),
				slog.String("url", req.URL.String()),
			)

			resp, err := next(req)

			duration := time.Since(start)

			if err != nil {
				logger.Error("elasticsearch request failed",
					slog.String("method", req.Method),
					slog.String("url", req.URL.String()),
					slog.Duration("duration", duration),
					slog.String("error", err.Error()),
				)
				return nil, err
			}

			logger.Info("elasticsearch request completed",
				slog.String("method", req.Method),
				slog.String("url", req.URL.String()),
				slog.Int("status_code", resp.StatusCode),
				slog.Duration("duration", duration),
			)

			return resp, nil
		}
	}
}

// MetricsInterceptor creates an interceptor that records request metrics.
func MetricsInterceptor(counter metric.Int64Counter, histogram metric.Float64Histogram) elastictransport.InterceptorFunc {
	return func(next elastictransport.RoundTripFunc) elastictransport.RoundTripFunc {
		return func(req *http.Request) (*http.Response, error) {
			start := time.Now()

			resp, err := next(req)

			duration := float64(time.Since(start).Milliseconds())

			// Record metrics with attributes
			attrs := []attribute.KeyValue{
				attribute.String("http.method", req.Method),
				attribute.String("url.path", req.URL.Path),
			}

			if err != nil {
				attrs = append(attrs, attribute.String("error.type", "request_error"))
			} else {
				attrs = append(attrs, attribute.Int("http.status_code", resp.StatusCode))
			}

			counter.Add(req.Context(), 1, metric.WithAttributes(attrs...))
			histogram.Record(req.Context(), duration, metric.WithAttributes(attrs...))

			return resp, err
		}
	}
}

// TracingInterceptor creates an interceptor that adds distributed tracing.
func TracingInterceptor(tracer trace.Tracer) elastictransport.InterceptorFunc {
	return func(next elastictransport.RoundTripFunc) elastictransport.RoundTripFunc {
		return func(req *http.Request) (*http.Response, error) {
			ctx, span := tracer.Start(req.Context(),
				fmt.Sprintf("elasticsearch %s", req.Method),
				trace.WithSpanKind(trace.SpanKindClient),
				trace.WithAttributes(
					semconv.HTTPRequestMethodKey.String(req.Method),
					semconv.URLFull(req.URL.String()),
					semconv.URLPath(req.URL.Path),
					semconv.ServerAddress(req.URL.Host),
				),
			)
			defer span.End()

			// Update request with traced context
			req = req.WithContext(ctx)

			resp, err := next(req)

			if err != nil {
				span.RecordError(err)
				span.SetAttributes(attribute.String("error.message", err.Error()))
				return nil, err
			}

			span.SetAttributes(semconv.HTTPResponseStatusCode(resp.StatusCode))

			return resp, nil
		}
	}
}
