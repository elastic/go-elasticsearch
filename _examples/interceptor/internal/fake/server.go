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

package fake

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
)

// LogFn is a function called on each request for custom logging.
type LogFn func(*http.Request)

// Config holds the configuration for the fake server.
type Config struct {
	LogFn        LogFn
	ResponseBody []byte
	StatusCode   int
	Port         int // 0 = random (default)
	SetHeaders   func(http.Header)
	Middleware   func(http.Handler) http.Handler
}

// Option is a functional option for configuring the server.
type Option func(*Config)

// WithLogFn sets a function to be called on each request for custom logging.
func WithLogFn(fn LogFn) Option {
	return func(c *Config) {
		c.LogFn = fn
	}
}

// WithResponseBody sets the response body returned by the server.
func WithResponseBody(body []byte) Option {
	return func(c *Config) {
		c.ResponseBody = body
	}
}

// WithStatusCode sets the HTTP status code returned by the server.
func WithStatusCode(code int) Option {
	return func(c *Config) {
		c.StatusCode = code
	}
}

// WithPort sets a specific port for the server to listen on.
// Use 0 (default) for a random available port.
func WithPort(port int) Option {
	return func(c *Config) {
		c.Port = port
	}
}

// WithHeaders sets a function to customize response headers.
func WithHeaders(fn func(http.Header)) Option {
	return func(c *Config) {
		c.SetHeaders = fn
	}
}

// WithMiddleware wraps the default handler with custom middleware.
// This allows full control over request/response flow.
func WithMiddleware(mw func(http.Handler) http.Handler) Option {
	return func(c *Config) {
		c.Middleware = mw
	}
}

// Server wraps httptest.Server to emulate an Elasticsearch server.
type Server struct {
	*httptest.Server
	config *Config
}

// NewServer creates a new fake server with the given options.
// The server is started automatically and accepts requests to any endpoint.
func NewServer(opts ...Option) *Server {
	config := &Config{
		StatusCode: http.StatusOK,
		LogFn:      func(*http.Request) {}, // no-op by default
	}

	for _, opt := range opts {
		opt(config)
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		config.LogFn(r)

		if config.SetHeaders != nil {
			config.SetHeaders(w.Header())
		}

		w.WriteHeader(config.StatusCode)

		if config.ResponseBody != nil {
			_, _ = w.Write(config.ResponseBody)
		}
	})

	if config.Middleware != nil {
		handler = config.Middleware(handler)
	}

	srv := &Server{config: config}

	if config.Port != 0 {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
		if err != nil {
			panic(fmt.Sprintf("failed to listen on port %d: %v", config.Port, err))
		}
		srv.Server = &httptest.Server{
			Listener: listener,
			Config:   &http.Server{Handler: handler},
		}
		srv.Server.Start()
	} else {
		srv.Server = httptest.NewServer(handler)
	}

	return srv
}

// URL returns the server's URL, suitable for use as an Elasticsearch endpoint.
func (s *Server) URL() string {
	return s.Server.URL
}
