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

// This example demonstrates how to use Interceptors to override
// authentication credentials on a per-request basis using context.Context.
//
// This pattern is useful when different requests need different credentials,
// such as multi-tenant applications or impersonation scenarios.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/_examples/interceptor/internal/fake"
	"github.com/elastic/go-elasticsearch/v9/_examples/interceptor/internal/redact"
)

func main() {
	// Start a fake Elasticsearch server that logs incoming auth credentials
	srv := fake.NewServer(
		fake.WithLogFn(func(r *http.Request) {
			username, password, _ := redact.BasicAuth(r)
			slog.Info("server received request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("username", username),
				slog.String("password", password),
			)
		}),
		fake.WithStatusCode(http.StatusOK),
		fake.WithResponseBody([]byte(`{"cluster_name":"example"}`)),
		fake.WithHeaders(func(h http.Header) {
			h.Set("X-Elastic-Product", "Elasticsearch")
			h.Set("Content-Type", "application/json")
		}),
	)
	defer srv.Close()

	// Create an Elasticsearch client with default credentials and context auth interceptor
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{srv.URL()},
		Username:  "default_user",
		Password:  "default_password",
		Interceptors: []elastictransport.InterceptorFunc{
			ContextAuthInterceptor(),
		},
	})
	if err != nil {
		panic(err)
	}

	// Request without context override uses default credentials
	fmt.Println(">>> Sending request with default credentials")
	_, _ = es.Info()

	// Request with context override uses the specified credentials
	fmt.Println("\n>>> Sending request with context override (tenant_a)")
	ctx := WithBasicAuth(context.Background(), "tenant_a", "tenant_a_secret")
	_, _ = es.Info(es.Info.WithContext(ctx))

	// Another request with different context credentials
	fmt.Println("\n>>> Sending request with context override (tenant_b)")
	ctx = WithBasicAuth(context.Background(), "tenant_b", "tenant_b_secret")
	_, _ = es.Info(es.Info.WithContext(ctx))

	// Request without context override still uses default credentials
	fmt.Println("\n>>> Sending request with default credentials again")
	_, _ = es.Info()
}

// basicAuthKey is the context key for storing basic auth credentials.
type basicAuthKey struct{}

type basicAuthValue struct {
	username string
	password string
}

// WithBasicAuth returns a context with basic auth credentials attached.
// Use this to override the default client credentials for a specific request.
func WithBasicAuth(ctx context.Context, username, password string) context.Context {
	return context.WithValue(ctx, basicAuthKey{}, basicAuthValue{username, password})
}

// ContextAuthInterceptor creates an interceptor that overrides BasicAuth
// credentials if they are present in the request's context.
// If no credentials are in the context, the request proceeds unchanged.
func ContextAuthInterceptor() elastictransport.InterceptorFunc {
	return func(next elastictransport.RoundTripFunc) elastictransport.RoundTripFunc {
		return func(req *http.Request) (*http.Response, error) {
			if auth, ok := req.Context().Value(basicAuthKey{}).(basicAuthValue); ok {
				req.SetBasicAuth(auth.username, auth.password)
			}
			return next(req)
		}
	}
}
