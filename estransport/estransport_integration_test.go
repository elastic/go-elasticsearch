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

package estransport_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8/estransport"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

var (
	_ = fmt.Print
)

func TestTransportRetries(t *testing.T) {
	var counter int

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counter++

		body, _ := ioutil.ReadAll(r.Body)
		fmt.Println("req.Body:", string(body))

		http.Error(w, "FAKE 502", http.StatusBadGateway)
	}))
	serverURL, _ := url.Parse(server.URL)

	transport, _ := estransport.New(estransport.Config{URLs: []*url.URL{serverURL}})

	bodies := []io.Reader{
		strings.NewReader(`FAKE`),
		esutil.NewJSONReader(`FAKE`),
	}

	for _, body := range bodies {
		t.Run(fmt.Sprintf("Reset the %T request body", body), func(t *testing.T) {
			counter = 0

			req, err := http.NewRequest("GET", "/", body)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			res, err := transport.Perform(req)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			body, _ := ioutil.ReadAll(res.Body)

			fmt.Println("> GET", req.URL)
			fmt.Printf("< %s (tries: %d)\n", bytes.TrimSpace(body), counter)

			if counter != 4 {
				t.Errorf("Unexpected number of attempts, want=4, got=%d", counter)
			}
		})
	}
}

func TestTransportHeaders(t *testing.T) {
	u, _ := url.Parse("http://localhost:9200")

	hdr := http.Header{}
	hdr.Set("Accept", "application/yaml")

	tp, _ := estransport.New(estransport.Config{
		URLs:   []*url.URL{u},
		Header: hdr,
	})

	req, _ := http.NewRequest("GET", "/", nil)
	res, err := tp.Perform(req)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if !bytes.HasPrefix(body, []byte("---")) {
		t.Errorf("Unexpected response body:\n%s", body)
	}
}

func TestTransportCompression(t *testing.T) {
	var req *http.Request
	var res *http.Response
	var err error

	u, _ := url.Parse("http://localhost:9200")

	transport, _ := estransport.New(estransport.Config{
		URLs: []*url.URL{u},
		CompressRequestBody: true,
	})

	indexName := "/shiny_new_index"

	req, _ = http.NewRequest(http.MethodPut, indexName, nil)
	res, err = transport.Perform(req)
	if err != nil {
		t.Fatalf("Unexpected error, cannot create index: %v", err)
	}

	req, _ = http.NewRequest(http.MethodGet, indexName, nil)
	res, err = transport.Perform(req)
	if err != nil {
		t.Fatalf("Unexpected error, cannot find index: %v", err)
	}

	req, _ = http.NewRequest(
		http.MethodPost,
		strings.Join([]string{indexName, "/_doc"}, ""),
		strings.NewReader(`{"solidPayload": 1}`),
	)
	req.Header.Set("Content-Type", "application/json")
	res, err = transport.Perform(req)
	if err != nil {
		t.Fatalf("Unexpected error, cannot POST payload: %v", err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Unexpected StatusCode, expected 201, got: %v", res.StatusCode)
	}

	req, _ = http.NewRequest(http.MethodDelete, indexName, nil)
	_, err = transport.Perform(req)
	if err != nil {
		t.Fatalf("Unexpected error, cannot DELETE %s: %v", indexName, err)
	}
}