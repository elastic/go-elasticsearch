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

// +build !integration

package elasticsearch

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7/estransport"
)

var called bool

type mockTransp struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

var defaultRoundTripFunc = func(req *http.Request) (*http.Response, error) {
	response := &http.Response{Header: http.Header{"X-Elastic-Product": []string{"Elasticsearch"}}}

	if req.URL.Path == "/" {
		response.Body = ioutil.NopCloser(strings.NewReader(`{
		  "version" : {
			"number" : "7.14.0-SNAPSHOT",
			"build_flavor" : "default"
		  },
		  "tagline" : "You Know, for Search"
		}`))
		response.Header.Add("Content-Type", "application/json")
	} else {
		called = true
	}

	return response, nil
}

func (t *mockTransp) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.RoundTripFunc == nil {
		return defaultRoundTripFunc(req)
	}
	return t.RoundTripFunc(req)
}

func TestClientConfiguration(t *testing.T) {
	t.Parallel()

	t.Run("With empty", func(t *testing.T) {
		c, err := NewDefaultClient()

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		u := c.Transport.(*estransport.Client).URLs()[0].String()

		if u != defaultURL {
			t.Errorf("Unexpected URL, want=%s, got=%s", defaultURL, u)
		}
	})

	t.Run("With URL from Addresses", func(t *testing.T) {
		c, err := NewClient(Config{Addresses: []string{"http://localhost:8080//"}, Transport: &mockTransp{}})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		u := c.Transport.(*estransport.Client).URLs()[0].String()

		if u != "http://localhost:8080" {
			t.Errorf("Unexpected URL, want=http://localhost:8080, got=%s", u)
		}
	})

	t.Run("With URL from environment", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", "http://example.com")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		c, err := NewClient(Config{Transport: &mockTransp{}})
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		u := c.Transport.(*estransport.Client).URLs()[0].String()

		if u != "http://example.com" {
			t.Errorf("Unexpected URL, want=http://example.com, got=%s", u)
		}
	})

	t.Run("With URL from environment and cfg.Addresses", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", "http://example.com")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		c, err := NewClient(Config{Addresses: []string{"http://localhost:8080//"}, Transport: &mockTransp{}})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		u := c.Transport.(*estransport.Client).URLs()[0].String()

		if u != "http://localhost:8080" {
			t.Errorf("Unexpected URL, want=http://localhost:8080, got=%s", u)
		}
	})

	t.Run("With URL from environment and cfg.CloudID", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", "http://example.com")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		c, err := NewClient(Config{CloudID: "foo:YmFyLmNsb3VkLmVzLmlvJGFiYzEyMyRkZWY0NTY=", Transport: &mockTransp{}})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		u := c.Transport.(*estransport.Client).URLs()[0].String()

		if u != "https://abc123.bar.cloud.es.io" {
			t.Errorf("Unexpected URL, want=https://abc123.bar.cloud.es.io, got=%s", u)
		}
	})

	t.Run("With cfg.Addresses and cfg.CloudID", func(t *testing.T) {
		_, err := NewClient(Config{Addresses: []string{"http://localhost:8080//"}, CloudID: "foo:ABC="})
		if err == nil {
			t.Fatalf("Expected error, got: %v", err)
		}
		match, _ := regexp.MatchString("both .* are set", err.Error())
		if !match {
			t.Errorf("Expected error when addresses from environment and configuration are used together, got: %v", err)
		}
	})

	t.Run("With CloudID", func(t *testing.T) {
		// bar.cloud.es.io$abc123$def456
		c, err := NewClient(Config{CloudID: "foo:YmFyLmNsb3VkLmVzLmlvJGFiYzEyMyRkZWY0NTY=", Transport: &mockTransp{}})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		u := c.Transport.(*estransport.Client).URLs()[0].String()

		if u != "https://abc123.bar.cloud.es.io" {
			t.Errorf("Unexpected URL, want=https://abc123.bar.cloud.es.io, got=%s", u)
		}
	})

	t.Run("With invalid CloudID", func(t *testing.T) {
		var err error

		_, err = NewClient(Config{CloudID: "foo:ZZZ==="})
		if err == nil {
			t.Errorf("Expected error for CloudID, got: %v", err)
		}

		_, err = NewClient(Config{CloudID: "foo:Zm9v"})
		if err == nil {
			t.Errorf("Expected error for CloudID, got: %v", err)
		}

		_, err = NewClient(Config{CloudID: "foo:"})
		if err == nil {
			t.Errorf("Expected error for CloudID, got: %v", err)
		}
	})

	t.Run("With invalid URL", func(t *testing.T) {
		u := ":foo"
		_, err := NewClient(Config{Addresses: []string{u}})

		if err == nil {
			t.Errorf("Expected error for URL %q, got %v", u, err)
		}
	})

	t.Run("With invalid URL from environment", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", ":foobar")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		c, err := NewDefaultClient()
		if err == nil {
			t.Errorf("Expected error, got: %+v", c)
		}
	})

	t.Run("With skip check", func(t *testing.T) {
		_, err := NewClient(
			Config{
				Transport: &mockTransp{
					RoundTripFunc: func(request *http.Request) (*http.Response, error) {
						return &http.Response{
							Header: http.Header{
								"X-Elastic-Product": []string{"X-Elastic-Product"},
							},
							Body: ioutil.NopCloser(strings.NewReader("")),
						}, nil
					},
				},
			})
		if err != nil {
			t.Errorf("Unexpected error, got: %+v", err)
		}
	})
}

func TestClientInterface(t *testing.T) {
	t.Run("Transport", func(t *testing.T) {
		c, err := NewClient(Config{Transport: &mockTransp{}})

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if called != false { // megacheck ignore
			t.Errorf("Unexpected call to transport by client")
		}

		c.Perform(&http.Request{URL: &url.URL{}, Header: make(http.Header)}) // errcheck ignore

		if called != true { // megacheck ignore
			t.Errorf("Expected client to call transport")
		}
	})
}

func TestAddrsToURLs(t *testing.T) {
	tt := []struct {
		name  string
		addrs []string
		urls  []*url.URL
		err   error
	}{
		{
			name: "valid",
			addrs: []string{
				"http://example.com",
				"https://example.com",
				"http://192.168.255.255",
				"http://example.com:8080",
			},
			urls: []*url.URL{
				{Scheme: "http", Host: "example.com"},
				{Scheme: "https", Host: "example.com"},
				{Scheme: "http", Host: "192.168.255.255"},
				{Scheme: "http", Host: "example.com:8080"},
			},
			err: nil,
		},
		{
			name:  "trim trailing slash",
			addrs: []string{"http://example.com/", "http://example.com//"},
			urls: []*url.URL{
				{Scheme: "http", Host: "example.com", Path: ""},
				{Scheme: "http", Host: "example.com", Path: ""},
			},
		},
		{
			name:  "keep suffix",
			addrs: []string{"http://example.com/foo"},
			urls:  []*url.URL{{Scheme: "http", Host: "example.com", Path: "/foo"}},
		},
		{
			name:  "invalid url",
			addrs: []string{"://invalid.com"},
			urls:  nil,
			err:   errors.New("missing protocol scheme"),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := addrsToURLs(tc.addrs)

			if tc.err != nil {
				if err == nil {
					t.Errorf("Expected error, got: %v", err)
				}
				match, _ := regexp.MatchString(tc.err.Error(), err.Error())
				if !match {
					t.Errorf("Expected err [%s] to match: %s", err.Error(), tc.err.Error())
				}
			}

			for i := range tc.urls {
				if res[i].Scheme != tc.urls[i].Scheme {
					t.Errorf("%s: Unexpected scheme, want=%s, got=%s", tc.name, tc.urls[i].Scheme, res[i].Scheme)
				}
			}
			for i := range tc.urls {
				if res[i].Host != tc.urls[i].Host {
					t.Errorf("%s: Unexpected host, want=%s, got=%s", tc.name, tc.urls[i].Host, res[i].Host)
				}
			}
			for i := range tc.urls {
				if res[i].Path != tc.urls[i].Path {
					t.Errorf("%s: Unexpected path, want=%s, got=%s", tc.name, tc.urls[i].Path, res[i].Path)
				}
			}
		})
	}
}

func TestCloudID(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		var testdata = []struct {
			in  string
			out string
		}{
			{
				in:  "name:" + base64.StdEncoding.EncodeToString([]byte("host$es_uuid$kibana_uuid")),
				out: "https://es_uuid.host",
			},
			{
				in:  "name:" + base64.StdEncoding.EncodeToString([]byte("host:9243$es_uuid$kibana_uuid")),
				out: "https://es_uuid.host:9243",
			},
			{
				in:  "name:" + base64.StdEncoding.EncodeToString([]byte("host$es_uuid$")),
				out: "https://es_uuid.host",
			},
			{
				in:  "name:" + base64.StdEncoding.EncodeToString([]byte("host$es_uuid")),
				out: "https://es_uuid.host",
			},
		}

		for _, tt := range testdata {
			actual, err := addrFromCloudID(tt.in)
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}
			if actual != tt.out {
				t.Errorf("Unexpected output, want=%q, got=%q", tt.out, actual)
			}
		}

	})

	t.Run("Invalid format", func(t *testing.T) {
		input := "foobar"
		_, err := addrFromCloudID(input)
		if err == nil {
			t.Errorf("Expected error for input %q, got %v", input, err)
		}
		match, _ := regexp.MatchString("unexpected format", err.Error())
		if !match {
			t.Errorf("Unexpected error string: %s", err)
		}
	})

	t.Run("Invalid base64 value", func(t *testing.T) {
		input := "foobar:xxxxx"
		_, err := addrFromCloudID(input)
		if err == nil {
			t.Errorf("Expected error for input %q, got %v", input, err)
		}
		match, _ := regexp.MatchString("illegal base64 data", err.Error())
		if !match {
			t.Errorf("Unexpected error string: %s", err)
		}
	})
}

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version is empty")
	}
}

func TestClientMetrics(t *testing.T) {
	c, _ := NewClient(Config{EnableMetrics: true, Transport: &mockTransp{}})

	m, err := c.Metrics()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if m.Requests > 1 {
		t.Errorf("Unexpected output: %s", m)
	}
}

func TestParseElasticsearchVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
		major   int64
		minor   int64
		patch   int64
		wantErr bool
	}{
		{
			name:    "Nominal version parsing",
			version: "7.14.0",
			major:   7,
			minor:   14,
			patch:   0,
			wantErr: false,
		},
		{
			name:    "Snapshot version parsing",
			version: "7.14.0-SNAPSHOT",
			major:   7,
			minor:   14,
			patch:   0,
			wantErr: false,
		},
		{
			name:    "Previous major parsing",
			version: "6.15.1",
			major:   6,
			minor:   15,
			patch:   1,
			wantErr: false,
		},
		{
			name:    "Error parsing version",
			version: "6.15",
			major:   0,
			minor:   0,
			patch:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := ParseElasticsearchVersion(tt.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseElasticsearchVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.major {
				t.Errorf("ParseElasticsearchVersion() got = %v, want %v", got, tt.major)
			}
			if got1 != tt.minor {
				t.Errorf("ParseElasticsearchVersion() got1 = %v, want %v", got1, tt.minor)
			}
			if got2 != tt.patch {
				t.Errorf("ParseElasticsearchVersion() got2 = %v, want %v", got2, tt.patch)
			}
		})
	}
}

func TestGenuineCheckInfo(t *testing.T) {
	tests := []struct {
		name    string
		info    info
		wantErr bool
		err 	error
	}{
		{
			name: "Genuine Elasticsearch 7.14.0",
			info: info{
				Version: esVersion{
					Number:      "7.14.0",
					BuildFlavor: "default",
				},
				Tagline: "You Know, for Search",
			},
			wantErr: false,
			err: nil,
		},
		{
			name: "Genuine Elasticsearch 6.15.1",
			info: info{
				Version: esVersion{
					Number:      "6.15.1",
					BuildFlavor: "default",
				},
				Tagline: "You Know, for Search",
			},
			wantErr: false,
			err: nil,
		},
		{
			name: "Not so genuine Elasticsearch 7 major",
			info: info{
				Version: esVersion{
					Number:      "7.12.0",
					BuildFlavor: "newer",
				},
				Tagline: "You Know, for Search",
			},
			wantErr: true,
			err: errors.New(unknownProduct),
		},
		{
			name: "Not so genuine Elasticsearch 6 major",
			info: info{
				Version: esVersion{
					Number:      "6.12.0",
					BuildFlavor: "default",
				},
				Tagline: "You Know, for Fun",
			},
			wantErr: true,
			err: errors.New(unknownProduct),
		},
		{
			name: "Way older Elasticsearch major",
			info: info{
				Version: esVersion{
					Number:      "5.12.0",
					BuildFlavor: "default",
				},
				Tagline: "You Know, for Fun",
			},
			wantErr: true,
			err: errors.New(unknownProduct),
		},
		{
			name: "Elasticsearch oss",
			info: info{
				Version: esVersion{
					Number:      "7.10.0",
					BuildFlavor: "oss",
				},
				Tagline: "You Know, for Search",
			},
			wantErr: true,
			err: errors.New(unsupportedProduct),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := genuineCheckInfo(tt.info); (err != nil) != tt.wantErr && err != tt.err {
				t.Errorf("genuineCheckInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenuineCheckHeader(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		wantErr bool
	}{
		{
			name: "Available and good product header",
			headers: http.Header{
				"X-Elastic-Product": []string{"Elasticsearch"},
			},
			wantErr: false,
		},
		{
			name: "Available and bad product header",
			headers: http.Header{
				"X-Elastic-Product": []string{"Elasticmerch"},
			},
			wantErr: true,
		},
		{
			name:    "Unavailable product header",
			headers: http.Header{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := genuineCheckHeader(tt.headers); (err != nil) != tt.wantErr {
				t.Errorf("genuineCheckHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponseCheckOnly(t *testing.T) {
	tests := []struct {
		name                 string
		useResponseCheckOnly bool
		response             *http.Response
		requestErr           error
		wantErr              bool
	}{
		{
			name:                 "Valid answer with header",
			useResponseCheckOnly: false,
			response: &http.Response{
				Header: http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
				Body:   ioutil.NopCloser(strings.NewReader("{}")),
			},
			wantErr: false,
		},
		{
			name:                 "Valid answer without header",
			useResponseCheckOnly: false,
			response: &http.Response{
				Body: ioutil.NopCloser(strings.NewReader("{}")),
			},
			wantErr: true,
		},
		{
			name:                 "Valid answer with header and response check",
			useResponseCheckOnly: true,
			response: &http.Response{
				Header: http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
				Body:   ioutil.NopCloser(strings.NewReader("{}")),
			},
			wantErr: false,
		},
		{
			name:                 "Valid answer without header and response check",
			useResponseCheckOnly: true,
			response: &http.Response{
				Body: ioutil.NopCloser(strings.NewReader("{}")),
			},
			wantErr: true,
		},
		{
			name:                 "Request failed",
			useResponseCheckOnly: true,
			response:             nil,
			requestErr:           errors.New("request failed"),
			wantErr:              true,
		},
		{
			name:                 "Valid request, 500 response",
			useResponseCheckOnly: false,
			response:             &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body: ioutil.NopCloser(strings.NewReader("")),
			},
			requestErr:           nil,
			wantErr:              true,
		},
		{
			name:                 "Valid request, 404 response",
			useResponseCheckOnly: false,
			response:             &http.Response{
				StatusCode: http.StatusNotFound,
				Body: ioutil.NopCloser(strings.NewReader("")),
			},
			requestErr:           nil,
			wantErr:              true,
		},
		{
			name:                 "Valid request, 403 response",
			useResponseCheckOnly: false,
			response:             &http.Response{
				StatusCode: http.StatusForbidden,
				Body: ioutil.NopCloser(strings.NewReader("")),
			},
			requestErr:           nil,
			wantErr:              false,
		},
		{
			name:                 "Valid request, 401 response",
			useResponseCheckOnly: false,
			response:             &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body: ioutil.NopCloser(strings.NewReader("")),
			},
			requestErr:           nil,
			wantErr:              false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(Config{
				Transport: &mockTransp{RoundTripFunc: func(request *http.Request) (*http.Response, error) {
					return tt.response, tt.requestErr
				}},
				UseResponseCheckOnly: tt.useResponseCheckOnly,
			})
			_, err := c.Cat.Indices()
			if (err != nil) != tt.wantErr {
				t.Errorf("Unexpected error, got %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductCheckError(t *testing.T) {
	var requestPaths []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPaths = append(requestPaths, r.URL.Path)
		if len(requestPaths) == 1 {
			// Simulate transient error from a proxy on the first request.
			// This must not be cached by the client.
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.Write([]byte("{}"))
	}))
	defer server.Close()

	c, _ := NewClient(Config{Addresses: []string{server.URL}, DisableRetry: true})
	if _, err := c.Cat.Indices(); err == nil {
		t.Fatal("expected error")
	}
	if c.productCheckSuccess {
		t.Fatalf("product check should be invalid, got %v", c.productCheckSuccess)
	}
	if _, err := c.Cat.Indices(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if n := len(requestPaths); n != 3 {
		t.Fatalf("expected 3 requests, got %d", n)
	}
	if !reflect.DeepEqual(requestPaths, []string{"/", "/", "/_cat/indices"}) {
		t.Fatalf("unexpected request paths: %s", requestPaths)
	}
	if !c.productCheckSuccess {
		t.Fatalf("product check should be valid, got : %v", c.productCheckSuccess)
	}
}
