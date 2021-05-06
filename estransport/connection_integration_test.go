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

// +build integration

package estransport

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

func NewServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{Addr: addr, Handler: handler}
}

func TestStatusConnectionPool(t *testing.T) {
	defaultResurrectTimeoutInitial = time.Second
	defer func() { defaultResurrectTimeoutInitial = 60 * time.Second }()

	var (
		server      *http.Server
		servers     []*http.Server
		serverURLs  []*url.URL
		serverHosts []string
		numServers  = 3

		defaultHandler = func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "OK") }
	)

	for i := 1; i <= numServers; i++ {
		s := NewServer(fmt.Sprintf("localhost:1000%d", i), http.HandlerFunc(defaultHandler))

		go func(s *http.Server) {
			if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				t.Fatalf("Unable to start server: %s", err)
			}
		}(s)

		defer func(s *http.Server) { s.Close() }(s)

		servers = append(servers, s)
		time.Sleep(time.Millisecond)
	}

	for _, s := range servers {
		u, _ := url.Parse("http://" + s.Addr)
		serverURLs = append(serverURLs, u)
		serverHosts = append(serverHosts, u.String())
	}

	fmt.Printf("==> Started %d servers on %s\n", numServers, serverHosts)

	cfg := Config{URLs: serverURLs}

	if _, ok := os.LookupEnv("GITHUB_ACTIONS"); !ok {
		cfg.Logger = &TextLogger{Output: os.Stdout}
		cfg.EnableDebugLogger = true
	}

	transport, _ := New(cfg)

	pool := transport.pool.(*statusConnectionPool)

	for i := 1; i <= 9; i++ {
		req, _ := http.NewRequest("GET", "/", nil)
		res, err := transport.Perform(req)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res.StatusCode != 200 {
			t.Errorf("Unexpected status code, want=200, got=%d", res.StatusCode)
		}
	}

	pool.Lock()
	if len(pool.live) != 3 {
		t.Errorf("Unexpected number of live connections, want=3, got=%d", len(pool.live))
	}
	pool.Unlock()

	server = servers[1]
	fmt.Printf("==> Closing server: %s\n", server.Addr)
	if err := server.Close(); err != nil {
		t.Fatalf("Unable to close server: %s", err)
	}

	for i := 1; i <= 9; i++ {
		req, _ := http.NewRequest("GET", "/", nil)
		res, err := transport.Perform(req)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res.StatusCode != 200 {
			t.Errorf("Unexpected status code, want=200, got=%d", res.StatusCode)
		}
	}

	pool.Lock()
	if len(pool.live) != 2 {
		t.Errorf("Unexpected number of live connections, want=2, got=%d", len(pool.live))
	}
	pool.Unlock()

	pool.Lock()
	if len(pool.dead) != 1 {
		t.Errorf("Unexpected number of dead connections, want=1, got=%d", len(pool.dead))
	}
	pool.Unlock()

	server = NewServer("localhost:10002", http.HandlerFunc(defaultHandler))
	servers[1] = server
	fmt.Printf("==> Starting server: %s\n", server.Addr)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			t.Fatalf("Unable to start server: %s", err)
		}
	}()

	fmt.Println("==> Waiting 1.25s for resurrection...")
	time.Sleep(1250 * time.Millisecond)

	for i := 1; i <= 9; i++ {
		req, _ := http.NewRequest("GET", "/", nil)
		res, err := transport.Perform(req)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res.StatusCode != 200 {
			t.Errorf("Unexpected status code, want=200, got=%d", res.StatusCode)
		}
	}

	pool.Lock()
	if len(pool.live) != 3 {
		t.Errorf("Unexpected number of live connections, want=3, got=%d", len(pool.live))
	}
	pool.Unlock()

	pool.Lock()
	if len(pool.dead) != 0 {
		t.Errorf("Unexpected number of dead connections, want=0, got=%d", len(pool.dead))
	}
	pool.Unlock()
}
