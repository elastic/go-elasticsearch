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

//go:build !integration
// +build !integration

package elasticsearch

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

var metaHeaderReValidation = regexp.MustCompile(`^[a-z]{1,}=[a-z0-9\.\-]{1,}(?:,[a-z]{1,}=[a-z0-9\.\-]+)*$`)
var called bool

type mockTransp struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

var defaultRoundTripFunc = func(req *http.Request) (*http.Response, error) {
	response := &http.Response{Header: http.Header{"X-Elastic-Product": []string{"Elasticsearch"}}}

	if req.URL.Path == "/" {
		response.Body = ioutil.NopCloser(strings.NewReader(`{
		  "version" : {
			"number" : "8.0.0-SNAPSHOT",
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

		u := c.Transport.(*elastictransport.Client).URLs()[0].String()

		if u != defaultURL {
			t.Errorf("Unexpected URL, want=%s, got=%s", defaultURL, u)
		}
	})

	t.Run("With URL from Addresses", func(t *testing.T) {
		c, err := NewClient(Config{Addresses: []string{"http://localhost:8080//"}})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		u := c.Transport.(*elastictransport.Client).URLs()[0].String()

		if u != "http://localhost:8080" {
			t.Errorf("Unexpected URL, want=http://localhost:8080, got=%s", u)
		}
	})

	t.Run("With URL from environment", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", "http://example.com")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		c, err := NewDefaultClient()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		u := c.Transport.(*elastictransport.Client).URLs()[0].String()

		if u != "http://example.com" {
			t.Errorf("Unexpected URL, want=http://example.com, got=%s", u)
		}
	})

	t.Run("With URL from environment and cfg.Addresses", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", "http://example.com")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		c, err := NewClient(Config{Addresses: []string{"http://localhost:8080//"}})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		u := c.Transport.(*elastictransport.Client).URLs()[0].String()

		if u != "http://localhost:8080" {
			t.Errorf("Unexpected URL, want=http://localhost:8080, got=%s", u)
		}
	})

	t.Run("With URL from environment and cfg.CloudID", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", "http://example.com")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		c, err := NewClient(Config{CloudID: "foo:YmFyLmNsb3VkLmVzLmlvJGFiYzEyMyRkZWY0NTY="})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		u := c.Transport.(*elastictransport.Client).URLs()[0].String()

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
		c, err := NewClient(Config{CloudID: "foo:YmFyLmNsb3VkLmVzLmlvJGFiYzEyMyRkZWY0NTY="})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		u := c.Transport.(*elastictransport.Client).URLs()[0].String()

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
	c, _ := NewClient(Config{EnableMetrics: true})

	m, err := c.Metrics()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if m.Requests != 0 {
		t.Errorf("Unexpected output: %s", m)
	}
}

func TestResponseCheckOnly(t *testing.T) {
	tests := []struct {
		name       string
		response   *http.Response
		requestErr error
		wantErr    bool
	}{
		{
			name: "Valid answer with header",
			response: &http.Response{
				Header: http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
				Body:   ioutil.NopCloser(strings.NewReader("{}")),
			},
			wantErr: false,
		},
		{
			name: "Valid answer without header",
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("{}")),
			},
			wantErr: true,
		},
		{
			name: "Valid answer with header and response check",
			response: &http.Response{
				Header: http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
				Body:   ioutil.NopCloser(strings.NewReader("{}")),
			},
			wantErr: false,
		},
		{
			name: "Valid answer without header and response check",
			response: &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("{}")),
			},
			wantErr: true,
		},
		{
			name:       "Request failed",
			response:   nil,
			requestErr: errors.New("request failed"),
			wantErr:    true,
		},
		{
			name: "Valid request, 500 response",
			response: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			requestErr: nil,
			wantErr:    false,
		},
		{
			name: "Valid request, 404 response",
			response: &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			requestErr: nil,
			wantErr:    false,
		},
		{
			name: "Valid request, 403 response",
			response: &http.Response{
				StatusCode: http.StatusForbidden,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			requestErr: nil,
			wantErr:    false,
		},
		{
			name: "Valid request, 401 response",
			response: &http.Response{
				StatusCode: http.StatusUnauthorized,
				Body:       ioutil.NopCloser(strings.NewReader("")),
			},
			requestErr: nil,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(Config{
				Transport: &mockTransp{RoundTripFunc: func(request *http.Request) (*http.Response, error) {
					return tt.response, tt.requestErr
				}},
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
	if _, err := c.Cat.Indices(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if c.productCheckSuccess {
		t.Fatalf("product check should be invalid, got %v", c.productCheckSuccess)
	}
	if _, err := c.Cat.Indices(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if n := len(requestPaths); n != 2 {
		t.Fatalf("expected 2 requests, got %d", n)
	}
	if !reflect.DeepEqual(requestPaths, []string{"/_cat/indices", "/_cat/indices"}) {
		t.Fatalf("unexpected request paths: %s", requestPaths)
	}
	if !c.productCheckSuccess {
		t.Fatalf("product check should be valid, got : %v", c.productCheckSuccess)
	}
}

func TestFingerprint(t *testing.T) {
	body := []byte(`{"body": true"}"`)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Elastic-Product", "Elasticsearch")
		w.Write(body)
	}))
	defer server.Close()

	config := Config{
		Addresses:    []string{server.URL},
		DisableRetry: true,
	}

	// Without certificate and authority, client should fail on TLS
	client, _ := NewClient(config)
	res, err := client.Info()
	if _, ok := err.(x509.UnknownAuthorityError); !ok {
		t.Fatalf("Uknown error, expected UnknownAuthorityError, got: %s", err)
	}

	// We add the fingerprint corresponding ton testcert.LocalhostCert
	//
	config.CertificateFingerprint = "448F628A8A65AA18560E53A80C53ACB38C51B427DF0334082349141147DC9BF6"
	client, _ = NewClient(config)
	res, err = client.Info()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	if bytes.Compare(data, body) != 0 {
		t.Fatalf("unexpected payload returned: expected: %s, got: %s", body, data)
	}
}

func TestCompatibilityHeader(t *testing.T) {
	tests := []struct {
		name                string
		envVar              bool
		configVar           bool
		compatibilityHeader bool
		bodyPresent         bool
		expectsHeader       []string
	}{
		{
			name:                "Compatibility header disabled",
			compatibilityHeader: false,
			bodyPresent:         false,
			envVar:              false,
			configVar:           false,
			expectsHeader:       []string{"application/json"},
		},
		{
			name:                "Compatibility header enabled with env var",
			compatibilityHeader: true,
			bodyPresent:         false,
			envVar:              true,
			configVar:           false,
			expectsHeader:       []string{"application/vnd.elasticsearch+json;compatible-with=8"},
		},
		{
			name:                "Compatibility header enabled with body with env var",
			compatibilityHeader: true,
			bodyPresent:         true,
			envVar:              true,
			configVar:           false,
			expectsHeader:       []string{"application/vnd.elasticsearch+json;compatible-with=8"},
		},
		{
			name:                "Compatibility header enabled in conf",
			compatibilityHeader: true,
			bodyPresent:         false,
			envVar:              false,
			configVar:           true,
			expectsHeader:       []string{"application/vnd.elasticsearch+json;compatible-with=8"},
		},
		{
			name:                "Compatibility header enabled with body in conf",
			compatibilityHeader: true,
			bodyPresent:         true,
			envVar:              false,
			configVar:           true,
			expectsHeader:       []string{"application/vnd.elasticsearch+json;compatible-with=8"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv(esCompatHeader, strconv.FormatBool(test.compatibilityHeader))

			c, _ := NewClient(Config{
				EnableCompatibilityMode: test.configVar,
				Addresses:               []string{},
				Transport: &mockTransp{
					RoundTripFunc: func(req *http.Request) (*http.Response, error) {
						if test.compatibilityHeader {
							if !reflect.DeepEqual(req.Header["Accept"], test.expectsHeader) {
								t.Errorf("Compatibility header enabled but header is, not in request headers, got: %s, want: %s", req.Header["Accept"], test.expectsHeader)
							}

							if test.bodyPresent {
								if !reflect.DeepEqual(req.Header["Content-Type"], test.expectsHeader) {
									t.Errorf("Compatibility header with Body enabled, not in request headers, got: %s, want: %s", req.Header["Content-Type"], test.expectsHeader)
								}
							}
						}

						return &http.Response{
							StatusCode: http.StatusOK,
							Status:     "MOCK",
							Body:       ioutil.NopCloser(strings.NewReader("{}")),
						}, nil
					},
				},
			})

			req := &http.Request{URL: &url.URL{}, Header: make(http.Header)}
			if test.bodyPresent {
				req.Body = ioutil.NopCloser(strings.NewReader("{}"))
			}

			_, _ = c.Perform(req)
		})
	}
}

func TestBuildStrippedVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Standard Go version",
			args: args{version: "go1.16"},
			want: "1.16",
		},
		{
			name: "Rc Go version",
			args: args{
				version: "go1.16rc1",
			},
			want: "1.16p",
		},
		{
			name: "Beta Go version (go1.16beta1 example)",
			args: args{
				version: "devel +2ff33f5e44 Thu Dec 17 16:03:19 2020 +0000",
			},
			want: "0.0p",
		},
		{
			name: "Random mostly good Go version",
			args: args{
				version: "1.16",
			},
			want: "1.16",
		},
		{
			name: "Client package version",
			args: args{
				version: "8.0.0",
			},
			want: "8.0.0",
		},
		{
			name: "Client pre release version",
			args: args{
				version: "8.0.0-SNAPSHOT",
			},
			want: "8.0.0p",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildStrippedVersion(tt.args.version); got != tt.want {
				t.Errorf("buildStrippedVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetaHeader(t *testing.T) {
	t.Run("MetaHeader with elastictransport", func(t *testing.T) {
		tp, _ := elastictransport.New(elastictransport.Config{
			URLs: []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: &mockTransp{
				RoundTripFunc: func(request *http.Request) (*http.Response, error) {
					h := request.Header.Get(HeaderClientMeta)
					if !metaHeaderReValidation.MatchString(h) {
						t.Errorf("expected client metaheader to validate regexp, got: %s", h)
					}

					return &http.Response{
						Header:     http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
						StatusCode: http.StatusOK,
						Status:     "OK",
						Body:       ioutil.NopCloser(strings.NewReader("")),
					}, nil
				},
			},
		})

		c, _ := NewDefaultClient()
		c.Transport = tp

		_, _ = c.Info()
	})
}
