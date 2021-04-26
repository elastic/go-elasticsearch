// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package estransport

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

	srv := &http.Server{Addr: "localhost:10001", Handler: http.HandlerFunc(defaultHandler)}
	srvTLS := &http.Server{Addr: "localhost:12001", Handler: http.HandlerFunc(defaultHandler)}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("Unable to start server: %s", err)
			return
		}
	}()
	go func() {
		if err := srvTLS.ListenAndServeTLS("testdata/cert.pem", "testdata/key.pem"); err != nil && err != http.ErrServerClosed {
			t.Errorf("Unable to start server: %s", err)
			return
		}
	}()
	defer func() { srv.Close() }()
	defer func() { srvTLS.Close() }()

	time.Sleep(50 * time.Millisecond)

	t.Run("getNodesInfo()", func(t *testing.T) {
		u, _ := url.Parse("http://" + srv.Addr)
		tp, _ := New(Config{URLs: []*url.URL{u}})

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
				if node.URL.String() != "http://127.0.0.1:10001" {
					t.Errorf("Unexpected URL: %s", node.URL.String())
				}
			case "es2":
				if node.URL.String() != "http://localhost:10002" {
					t.Errorf("Unexpected URL: %s", node.URL.String())
				}
			case "es3":
				if node.URL.String() != "http://127.0.0.1:10003" {
					t.Errorf("Unexpected URL: %s", node.URL.String())
				}
			}
		}
	})

	t.Run("DiscoverNodes()", func(t *testing.T) {
		u, _ := url.Parse("http://" + srv.Addr)
		tp, _ := New(Config{URLs: []*url.URL{u}})

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
				if conn.URL.String() != "http://127.0.0.1:10001" {
					t.Errorf("Unexpected URL: %s", conn.URL.String())
				}
			case "es2":
				if conn.URL.String() != "http://localhost:10002" {
					t.Errorf("Unexpected URL: %s", conn.URL.String())
				}
			default:
				t.Errorf("Unexpected node: %s", conn.Name)
			}
		}
	})

	t.Run("DiscoverNodes() with SSL and authorization", func(t *testing.T) {
		u, _ := url.Parse("https://" + srvTLS.Addr)
		tp, _ := New(Config{
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
				if conn.URL.String() != "https://127.0.0.1:10001" {
					t.Errorf("Unexpected URL: %s", conn.URL.String())
				}
			case "es2":
				if conn.URL.String() != "https://localhost:10002" {
					t.Errorf("Unexpected URL: %s", conn.URL.String())
				}
			default:
				t.Errorf("Unexpected node: %s", conn.Name)
			}
		}
	})

	t.Run("scheduleDiscoverNodes()", func(t *testing.T) {
		t.Skip("Skip") // TODO(karmi): Investigate the intermittent failures of this test

		var numURLs int
		u, _ := url.Parse("http://" + srv.Addr)

		tp, _ := New(Config{URLs: []*url.URL{u}, DiscoverNodesInterval: 10 * time.Millisecond})

		tp.Lock()
		numURLs = len(tp.pool.URLs())
		tp.Unlock()
		if numURLs != 1 {
			t.Errorf("Unexpected number of nodes, want=1, got=%d", numURLs)
		}

		time.Sleep(18 * time.Millisecond) // Wait until (*Client).scheduleDiscoverNodes()
		tp.Lock()
		numURLs = len(tp.pool.URLs())
		tp.Unlock()
		if numURLs != 2 {
			t.Errorf("Unexpected number of nodes, want=2, got=%d", numURLs)
		}
	})

	t.Run("Role based nodes discovery", func(t *testing.T) {
		type Node struct {
			Name string
			Url string
			Roles []string
		}

		type fields struct {
			Nodes []Node
		}
		type wants struct {
			wantErr bool
			wantsNConn int
		}
		tests := []struct {
			name    string
			args  fields
			want  wants
		}{
			{
				"Default roles should allow every node to be selected",
				fields{
					Nodes: []Node{
						{
							Name: "es1",
							Url: "http://es1:9200",
							Roles: []string{
								"data",
								"data_cold",
								"data_content",
								"data_frozen",
								"data_hot",
								"data_warm",
								"ingest",
								"master",
								"ml",
								"remote_cluster_client",
								"transform",
							},
						},{
							Name: "es2",
							Url: "http://es2:9200",
							Roles: []string{
								"data",
								"data_cold",
								"data_content",
								"data_frozen",
								"data_hot",
								"data_warm",
								"ingest",
								"master",
								"ml",
								"remote_cluster_client",
								"transform",
							},
						},{
							Name: "es3",
							Url: "http://es3:9200",
							Roles: []string{
								"data",
								"data_cold",
								"data_content",
								"data_frozen",
								"data_hot",
								"data_warm",
								"ingest",
								"master",
								"ml",
								"remote_cluster_client",
								"transform",
							},
						},
					},
				},
				wants{
					false, 3,
				},
			},
			{
				"Master only node should not be selected",
				fields{
					Nodes: []Node{
						{
							Name: "es1",
							Url: "http://es1:9200",
							Roles: []string{
								"master",
							},
						},
						{
							Name: "es2",
							Url: "http://es2:9200",
							Roles: []string{
								"data",
								"data_cold",
								"data_content",
								"data_frozen",
								"data_hot",
								"data_warm",
								"ingest",
								"master",
								"ml",
								"remote_cluster_client",
								"transform",
							},
						},
						{
							Name: "es3",
							Url: "http://es3:9200",
							Roles: []string{
								"data",
								"data_cold",
								"data_content",
								"data_frozen",
								"data_hot",
								"data_warm",
								"ingest",
								"master",
								"ml",
								"remote_cluster_client",
								"transform",
							},
						},
					},
				},

				wants{
					false, 2,
				},
			},
			{
				"Master and data only nodes should be selected",
				fields{
					Nodes: []Node{
						{
							Name: "es1",
							Url: "http://es1:9200",
							Roles: []string{
								"data",
								"master",
							},
						},{
							Name: "es2",
							Url: "http://es2:9200",
							Roles: []string{
								"data",
								"master",
							},
						},
					},
				},

				wants{
					false, 2,
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var names []string
				var urls []*url.URL
				for _, v := range tt.args.Nodes {
					u, _ := url.Parse(v.Url)
					urls = append(urls, u)
					names = append(names, v.Name)
				}

				newRoundTripper := func() http.RoundTripper {
					return &mockTransp{
						RoundTripFunc: func(req *http.Request) (*http.Response, error) {
							nodes := make(map[string]map[string]nodeInfo)
							nodes["nodes"] = make(map[string]nodeInfo)
							for _, v := range tt.args.Nodes {
								nodes["nodes"][v.Name] = nodeInfo{Roles: v.Roles}
							}

							b, _ := json.Marshal(nodes)

							return &http.Response{
								Status:        "200 OK",
								StatusCode:    200,
								ContentLength: int64(len(b)),
								Header:        http.Header(map[string][]string{"Content-Type": {"application/json"}}),
								Body:          ioutil.NopCloser(bytes.NewReader(b)),
							}, nil
						},
					}
				}

				c, _ := New(Config{
					URLs:     urls,
					Transport: newRoundTripper(),
				})
				c.DiscoverNodes()

				pool, ok := c.pool.(*statusConnectionPool)
				if !ok {
					t.Fatalf("Unexpected pool, want=statusConnectionPool, got=%T", c.pool)
				}

				if len(pool.live) != tt.want.wantsNConn {
					t.Errorf("Unexpected number of nodes, want=%d, got=%d", tt.want.wantsNConn, len(pool.live))
				}

				if err := c.DiscoverNodes(); (err != nil) != tt.want.wantErr {
					t.Errorf("DiscoverNodes() error = %v, wantErr %v", err, tt.want.wantErr)
				}
			})
		}
	})
}
