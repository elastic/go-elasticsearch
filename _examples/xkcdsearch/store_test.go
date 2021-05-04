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

package xkcdsearch_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/elastic/go-elasticsearch/v8/_examples/xkcdsearch"
)

var (
	fixtures = make(map[string]io.ReadCloser)
)

func init() {
	fixtureFiles, err := filepath.Glob("testdata/*.json")
	if err != nil {
		panic(fmt.Sprintf("Cannot glob fixture files: %s", err))
	}

	for _, fpath := range fixtureFiles {
		f, err := ioutil.ReadFile(fpath)
		if err != nil {
			panic(fmt.Sprintf("Cannot read fixture file: %s", err))
		}
		fixtures[filepath.Base(fpath)] = ioutil.NopCloser(bytes.NewReader(f))
	}
}

func fixture(fname string) io.ReadCloser {
	out := new(bytes.Buffer)
	b1 := bytes.NewBuffer([]byte{})
	b2 := bytes.NewBuffer([]byte{})
	tr := io.TeeReader(fixtures[fname], b1)

	defer func() { fixtures[fname] = ioutil.NopCloser(b1) }()
	io.Copy(b2, tr)
	out.ReadFrom(b2)

	return ioutil.NopCloser(out)
}

type MockTransport struct {
	Response    *http.Response
	RoundTripFn func(req *http.Request) (*http.Response, error)
}

func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.RoundTripFn(req)
}

func TestStore(t *testing.T) {
	t.Parallel()

	mocktrans := MockTransport{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(`{}`)),
		},
	}
	mocktrans.RoundTripFn = func(req *http.Request) (*http.Response, error) { return mocktrans.Response, nil }

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Transport: &mocktrans,
	})
	if err != nil {
		t.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	config := xkcdsearch.StoreConfig{Client: client}
	store, err := xkcdsearch.NewStore(config)
	if err != nil {
		t.Fatalf("Error creating the store: %s", err)
	}

	t.Run("Empty response", func(t *testing.T) {
		results, err := store.Search("foobar")
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if results.Total != 0 {
			t.Errorf("Unexpected total results, want=0, got=%d", results.Total)
		}
	})

	t.Run("Match all", func(t *testing.T) {
		mocktrans.Response = &http.Response{
			StatusCode: http.StatusOK,
			Body:       fixture("match_all.json"),
		}

		results, err := store.Search("")
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if results.Total != 42 {
			t.Errorf("Unexpected total results, want=42, got=%d", results.Total)
		}
		if results.Hits[0].Title != "Title One" {
			t.Errorf("Unexpected title, want=\"Title One\", got=%q", results.Hits[0].Title)
		}
	})

	t.Run("Match all with search_after", func(t *testing.T) {
		mocktrans.RoundTripFn = func(req *http.Request) (*http.Response, error) {
			var b map[string]interface{}
			if err := json.NewDecoder(req.Body).Decode(&b); err != nil {
				t.Fatalf("Error parsing search definition: %s", err)
			}
			if b["search_after"].([]interface{})[0].(float64) != 1 {
				t.Fatalf("Unexpected query: %s", req.Body)
			}

			return &http.Response{StatusCode: http.StatusOK, Body: fixture("match_all.json")}, nil
		}

		results, err := store.Search("", "1,2")
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if results.Total != 42 {
			t.Errorf("Unexpected total results, want=42, got=%d", results.Total)
		}
		if results.Hits[0].Title != "Title One" {
			t.Errorf("Unexpected title, want=\"Title One\", got=%q", results.Hits[0].Title)
		}
	})

	t.Run("Match query", func(t *testing.T) {
		mocktrans.RoundTripFn = func(req *http.Request) (*http.Response, error) {
			var b map[string]interface{}
			if err := json.NewDecoder(req.Body).Decode(&b); err != nil {
				t.Fatalf("Error parsing search definition: %s", err)
			}
			if b["sort"].([]interface{})[0].(map[string]interface{})["_score"] != "desc" {
				t.Fatalf("Unexpected query: %s", req.Body)
			}

			return &http.Response{StatusCode: http.StatusOK, Body: fixture("match_query.json")}, nil
		}

		results, err := store.Search("numbers")
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if results.Total != 47 {
			t.Errorf("Unexpected total results, want=47, got=%d", results.Total)
		}
		if results.Hits[0].Title != "Numbers" {
			t.Errorf("Unexpected title, want=\"Numbers\", got=%q", results.Hits[0].Title)
		}
		if results.Hits[0].Highlights.Title[0] != "<em>Numbers</em>" {
			t.Errorf("Unexpected title, want=\"<em>Numbers</em>\", got=%q", results.Hits[0].Title)
		}
	})
}
