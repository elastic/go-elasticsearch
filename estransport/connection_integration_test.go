// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

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

	transport := New(
		Config{
			URLs:   serverURLs,
			Logger: &TextLogger{Output: os.Stdout},

			EnableDebugLogger: true,
		})

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

	if len(pool.live) != 3 {
		t.Errorf("Unexpected number of live connections, want=3, got=%d", len(pool.live))
	}

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

	if len(pool.live) != 2 {
		t.Errorf("Unexpected number of live connections, want=2, got=%d", len(pool.live))
	}

	if len(pool.dead) != 1 {
		t.Errorf("Unexpected number of dead connections, want=1, got=%d", len(pool.dead))
	}

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

	if len(pool.live) != 3 {
		t.Errorf("Unexpected number of live connections, want=3, got=%d", len(pool.live))
	}

	if len(pool.dead) != 0 {
		t.Errorf("Unexpected number of dead connections, want=0, got=%d", len(pool.dead))
	}
}
