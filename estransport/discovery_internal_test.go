// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package estransport

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestDiscovery(t *testing.T) {
	defaultHandler := func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("testdata/nodes.info.json")
		if err != nil {
			http.Error(w, fmt.Sprintf("Fixture error: %s", err), 500)
			return
		}
		io.Copy(w, f)
	}

	srv := &http.Server{Addr: "localhost:10100", Handler: http.HandlerFunc(defaultHandler)}
	srvTLS := &http.Server{Addr: "localhost:10200", Handler: http.HandlerFunc(defaultHandler)}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Fatalf("Unable to start server: %s", err)
		}
	}()
	go func() {
		if err := srvTLS.ListenAndServeTLS("testdata/cert.pem", "testdata/key.pem"); err != nil && err != http.ErrServerClosed {
			t.Fatalf("Unable to start server: %s", err)
		}
	}()
	defer func() { srv.Close() }()
	defer func() { srvTLS.Close() }()

	time.Sleep(time.Millisecond)

	t.Run("getNodesInfo()", func(t *testing.T) {
		u, _ := url.Parse("http://" + srv.Addr)
		tp := New(Config{URLs: []*url.URL{u}})

		nodes, err := tp.getNodesInfo()
		if err != nil {
			t.Fatalf("ERROR: %s", err)
		}
		fmt.Printf("NodesInfo: %+v\n", nodes)

		if len(nodes) != 3 {
			t.Errorf("Unexpected number of nodes, want=3, got=%d", len(nodes))
		}

		for _, node := range nodes {
			switch node.Name {
			case "es1":
				if node.URL.String() != "http://127.0.0.1:9200" {
					t.Errorf("Unexpected URL: %s", node.URL.String())
				}
			case "es2":
				if node.URL.String() != "http://localhost:9201" {
					t.Errorf("Unexpected URL: %s", node.URL.String())
				}
			case "es3":
				if node.URL.String() != "http://127.0.0.1:9202" {
					t.Errorf("Unexpected URL: %s", node.URL.String())
				}
			}
		}
	})

	t.Run("DiscoverNodes()", func(t *testing.T) {
		u, _ := url.Parse("http://" + srv.Addr)
		tp := New(Config{URLs: []*url.URL{u}})

		tp.DiscoverNodes()

		pool, ok := tp.pool.(*statusConnectionPool)
		if !ok {
			t.Fatalf("Unexpected pool, want=statusConnectionPool, got=%T", tp.pool)
		}

		if len(pool.live) != 2 {
			t.Errorf("Unexpected number of nodes, want=2, got=%d", len(pool.live))
		}

		for _, conn := range pool.live {
			switch conn.Name {
			case "es1":
				if conn.URL.String() != "http://127.0.0.1:9200" {
					t.Errorf("Unexpected URL: %s", conn.URL.String())
				}
			case "es2":
				if conn.URL.String() != "http://localhost:9201" {
					t.Errorf("Unexpected URL: %s", conn.URL.String())
				}
			default:
				t.Errorf("Unexpected node: %s", conn.Name)
			}
		}
	})

	t.Run("DiscoverNodes() with SSL and authorization", func(t *testing.T) {
		u, _ := url.Parse("https://" + srvTLS.Addr)
		tp := New(Config{
			URLs:     []*url.URL{u},
			Username: "foo",
			Password: "bar",
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		})

		tp.DiscoverNodes()

		pool, ok := tp.pool.(*statusConnectionPool)
		if !ok {
			t.Fatalf("Unexpected pool, want=statusConnectionPool, got=%T", tp.pool)
		}

		if len(pool.live) != 2 {
			t.Errorf("Unexpected number of nodes, want=2, got=%d", len(pool.live))
		}

		for _, conn := range pool.live {
			switch conn.Name {
			case "es1":
				if conn.URL.String() != "https://127.0.0.1:9200" {
					t.Errorf("Unexpected URL: %s", conn.URL.String())
				}
			case "es2":
				if conn.URL.String() != "https://localhost:9201" {
					t.Errorf("Unexpected URL: %s", conn.URL.String())
				}
			default:
				t.Errorf("Unexpected node: %s", conn.Name)
			}
		}
	})

	t.Run("scheduleDiscoverNodes()", func(t *testing.T) {
		u, _ := url.Parse("http://" + srv.Addr)

		tp := New(Config{URLs: []*url.URL{u}, DiscoverNodesInterval: time.Millisecond})

		if len(tp.pool.URLs()) != 1 {
			t.Errorf("Unexpected number of nodes, want=1, got=%d", len(tp.pool.URLs()))
		}

		time.Sleep(5 * time.Millisecond)
		if len(tp.pool.URLs()) != 2 {
			t.Errorf("Unexpected number of nodes, want=2, got=%d", len(tp.pool.URLs()))
		}
	})
}
