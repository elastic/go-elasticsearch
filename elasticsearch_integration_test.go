// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build integration,!multinode

package elasticsearch_test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/estransport"
)

func TestClientTransport(t *testing.T) {
	t.Run("Persistent", func(t *testing.T) {
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			t.Fatalf("Error creating the client: %s", err)
		}

		var total int

		for i := 0; i < 101; i++ {
			var curTotal int

			res, err := es.Nodes.Stats(es.Nodes.Stats.WithMetric("http"))
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			defer res.Body.Close()

			r := struct {
				Nodes map[string]struct {
					HTTP struct {
						TotalOpened int `json:"total_opened"`
					}
				}
			}{}

			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				t.Fatalf("Error parsing the response body: %s", err)
			}

			for _, v := range r.Nodes {
				curTotal = v.HTTP.TotalOpened
				break
			}

			if curTotal < 1 {
				t.Errorf("Unexpected total_opened: %d", curTotal)
			}

			if total == 0 {
				total = curTotal
			}

			if total != curTotal {
				t.Errorf("Expected total_opened=%d, got: %d", total, curTotal)
			}
		}

		log.Printf("total_opened: %d", total)
	})

	t.Run("Concurrent", func(t *testing.T) {
		var wg sync.WaitGroup

		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			t.Fatalf("Error creating the client: %s", err)
		}

		for i := 0; i < 101; i++ {
			wg.Add(1)
			time.Sleep(10 * time.Millisecond)

			go func(i int) {
				defer wg.Done()
				res, err := es.Info()
				if err != nil {
					t.Errorf("Unexpected error: %s", err)
				} else {
					defer res.Body.Close()
				}
			}(i)
		}
		wg.Wait()
	})

	t.Run("WithContext", func(t *testing.T) {
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			t.Fatalf("Error creating the client: %s", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
		defer cancel()

		res, err := es.Info(es.Info.WithContext(ctx))
		if err == nil {
			res.Body.Close()
			t.Fatal("Expected 'context deadline exceeded' error")
		}

		log.Printf("Request cancelled with %T", err)
	})

	t.Run("Configured", func(t *testing.T) {
		cfg := elasticsearch.Config{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   10,
				ResponseHeaderTimeout: time.Second,
				DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
				TLSClientConfig: &tls.Config{
					MinVersion:         tls.VersionTLS11,
					InsecureSkipVerify: true,
				},
			},
		}

		es, err := elasticsearch.NewClient(cfg)
		if err != nil {
			t.Fatalf("Error creating the client: %s", err)
		}

		res, err := es.Info()
		if err == nil {
			t.Fatalf("Expected error, but got: %v", err)
		}
		if _, ok := err.(*net.OpError); !ok {
			t.Fatalf("Expected net.OpError, but got: %T", err)
		}

		if res != nil {
			t.Fatalf("Unexpected response: %+v", res)
		}
	})
}

type CustomTransport struct {
	client *http.Client
}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Foo", "bar")
	log.Printf("> %s %s %s\n", req.Method, req.URL.String(), req.Header)
	return t.client.Do(req)
}

func TestClientCustomTransport(t *testing.T) {
	t.Run("Customized", func(t *testing.T) {
		cfg := elasticsearch.Config{
			Transport: &CustomTransport{
				client: http.DefaultClient,
			},
		}

		es, err := elasticsearch.NewClient(cfg)
		if err != nil {
			t.Fatalf("Error creating the client: %s", err)
		}

		for i := 0; i < 10; i++ {
			res, err := es.Info()
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			defer res.Body.Close()
		}
	})

	t.Run("Manual", func(t *testing.T) {
		tr := estransport.New(estransport.Config{
			URLs: []*url.URL{
				{Scheme: "http", Host: "localhost:9200"},
			},
			Transport: http.DefaultTransport,
		})

		es := elasticsearch.Client{
			Transport: tr, API: esapi.New(tr),
		}

		for i := 0; i < 10; i++ {
			res, err := es.Info()
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			defer res.Body.Close()
		}
	})
}

type ReplacedTransport struct {
	counter uint64
}

func (t *ReplacedTransport) Perform(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	req.URL.Host = "localhost:9200"

	atomic.AddUint64(&t.counter, 1)

	return http.DefaultTransport.RoundTrip(req)
}

func (t *ReplacedTransport) Count() uint64 {
	return atomic.LoadUint64(&t.counter)
}

func TestClientReplaceTransport(t *testing.T) {
	t.Run("Replaced", func(t *testing.T) {
		tr := &ReplacedTransport{}
		es := elasticsearch.Client{
			Transport: tr, API: esapi.New(tr),
		}

		for i := 0; i < 10; i++ {
			res, err := es.Info()
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			defer res.Body.Close()
		}

		if tr.Count() != 10 {
			t.Errorf("Expected 10 requests, got=%d", tr.Count())
		}
	})
}

func TestClientAPI(t *testing.T) {
	t.Run("Info", func(t *testing.T) {
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			log.Fatalf("Error creating the client: %s\n", err)
		}

		res, err := es.Info()
		if err != nil {
			log.Fatalf("Error getting the response: %s\n", err)
		}
		defer res.Body.Close()

		var d map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&d)
		if err != nil {
			log.Fatalf("Error parsing the response: %s\n", err)
		}

		fmt.Println(d["tagline"])
	})
}
