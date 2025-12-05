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

// This example demonstrates how to use Interceptors to dynamically
// inject authentication credentials into requests.
//
// Interceptors allow you to modify requests before they are sent,
// making them ideal for scenarios where credentials may change at
// runtime (e.g., token refresh, credential rotation).
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"

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

	// Create a credential provider that can be updated at runtime
	authProvider := NewCredentialProvider("user1", "password1")

	// Create an Elasticsearch client with a custom auth interceptor
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{srv.URL()},
		Interceptors: []elastictransport.InterceptorFunc{
			DynamicAuthInterceptor(authProvider),
		},
	})
	if err != nil {
		panic(err)
	}

	// First request uses initial credentials
	fmt.Println(">>> Sending request with initial credentials (user1)")
	_, _ = es.Info()

	// Update credentials (simulating credential rotation)
	fmt.Println("\n>>> Rotating credentials to (user2)")
	authProvider.Update("user2", "password2")

	// Second request automatically uses the new credentials
	fmt.Println("\n>>> Sending request with rotated credentials (user2)")
	_, _ = es.Info()
}

// DynamicAuthInterceptor creates an interceptor that injects BasicAuth
// credentials from a CredentialProvider into each request.
func DynamicAuthInterceptor(provider *CredentialProvider) elastictransport.InterceptorFunc {
	return func(next elastictransport.RoundTripFunc) elastictransport.RoundTripFunc {
		return func(req *http.Request) (*http.Response, error) {
			username, password := provider.Get()
			req.SetBasicAuth(username, password)
			return next(req)
		}
	}
}

// CredentialProvider holds credentials that can be safely updated at runtime.
type CredentialProvider struct {
	mu       sync.RWMutex
	username string
	password string
}

func NewCredentialProvider(username, password string) *CredentialProvider {
	return &CredentialProvider{username: username, password: password}
}

func (p *CredentialProvider) Update(username, password string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.username = username
	p.password = password
}

func (p *CredentialProvider) Get() (username, password string) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.username, p.password
}
