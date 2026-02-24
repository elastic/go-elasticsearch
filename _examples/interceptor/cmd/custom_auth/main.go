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

// This example demonstrates how to use Interceptors to implement
// Kerberos/SPNEGO authentication with challenge-response handling.
//
// Note: This example is a mock implementation and does not use true Kerberos/SPNEGO authentication.
// In production, you would use a Kerberos library like gokrb5 to obtain a service ticket.
//
// The interceptor handles 401 responses with WWW-Authenticate: Negotiate
// by obtaining a token and retrying the request with the Authorization header.
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/_examples/interceptor/internal/fake"
)

func main() {
	// Start a fake Elasticsearch server with SPNEGO auth middleware
	srv := fake.NewServer(
		fake.WithLogFn(func(r *http.Request) {
			auth := r.Header.Get("Authorization")
			slog.Info("server received request",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("authorization", auth),
			)
		}),
		fake.WithResponseBody([]byte(`{"cluster_name":"example"}`)),
		fake.WithHeaders(func(h http.Header) {
			h.Set("X-Elastic-Product", "Elasticsearch")
			h.Set("Content-Type", "application/json")
		}),
		fake.WithMiddleware(SPNEGOAuthMiddleware),
	)
	defer srv.Close()

	// Create an Elasticsearch client with the Kerberos interceptor
	es, err := elasticsearch.New(
		elasticsearch.WithAddresses(srv.URL()),
		elasticsearch.WithTransportOptions(elastictransport.WithInterceptors(
			KerberosInterceptor(MockTokenProvider),
		)),
	)
	if err != nil {
		panic(err)
	}

	// Send a request - the interceptor handles the 401 challenge automatically
	fmt.Println(">>> Sending request (interceptor will handle SPNEGO challenge)")
	resp, err := es.Info()
	if err != nil {
		panic(err)
	}
	fmt.Printf(">>> Response status: %s\n", resp.Status())
}

// KerberosInterceptor creates an interceptor that handles Kerberos/SPNEGO authentication.
// When a 401 response with WWW-Authenticate: Negotiate is received, it obtains a token
// using the provided tokenProvider and retries the request.
func KerberosInterceptor(tokenProvider func() (string, error)) elastictransport.InterceptorFunc {
	return func(next elastictransport.RoundTripFunc) elastictransport.RoundTripFunc {
		return func(req *http.Request) (*http.Response, error) {
			// Send the initial request
			resp, err := next(req)
			if err != nil {
				return nil, err
			}

			// Check if server requires SPNEGO authentication
			if resp.StatusCode == http.StatusUnauthorized {
				authHeader := resp.Header.Get("WWW-Authenticate")
				if strings.HasPrefix(authHeader, "Negotiate") {
					slog.Info("received SPNEGO challenge, obtaining token and retrying")

					// Close the original response body
					resp.Body.Close()

					// Obtain Kerberos token
					token, err := tokenProvider()
					if err != nil {
						return nil, fmt.Errorf("failed to obtain Kerberos token: %w", err)
					}

					// Clone the request and add the Authorization header
					retryReq := req.Clone(req.Context())
					retryReq.Header.Set("Authorization", "Negotiate "+token)

					// Retry with the token
					return next(retryReq)
				}
			}

			return resp, nil
		}
	}
}

// SPNEGOAuthMiddleware simulates a server that requires SPNEGO authentication.
// It returns 401 with WWW-Authenticate: Negotiate if no valid token is provided,
// or passes through to the next handler if authentication succeeds.
func SPNEGOAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		// Check for valid Negotiate token
		if strings.HasPrefix(auth, "Negotiate ") {
			token := strings.TrimPrefix(auth, "Negotiate ")
			if token != "" {
				// Token provided - authentication successful
				slog.Info("SPNEGO authentication successful")
				next.ServeHTTP(w, r)
				return
			}
		}

		// No valid token - send 401 challenge
		slog.Info("no valid token, sending SPNEGO challenge")
		w.Header().Set("WWW-Authenticate", "Negotiate")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":"authentication required"}`))
	})
}

// MockTokenProvider simulates obtaining a Kerberos token.
// In a real implementation, this would use a Kerberos library like gokrb5.
func MockTokenProvider() (string, error) {
	// In production, this would:
	// 1. Load Kerberos credentials (keytab or credential cache)
	// 2. Obtain a service ticket for the Elasticsearch SPN
	// 3. Return the base64-encoded SPNEGO token
	return "MOCK_KERBEROS_TOKEN_BASE64", nil
}
