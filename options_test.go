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
	"context"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

func TestNew_DefaultClient(t *testing.T) {
	t.Parallel()
	c, err := New()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	u := c.Transport.(*elastictransport.Client).URLs()[0].String()
	if u != defaultURL {
		t.Errorf("Unexpected URL, want=%s, got=%s", defaultURL, u)
	}
}

func TestNew_WithAddresses(t *testing.T) {
	t.Parallel()
	c, err := New(WithAddresses("http://localhost:8080//"))
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	u := c.Transport.(*elastictransport.Client).URLs()[0].String()
	if u != "http://localhost:8080" {
		t.Errorf("Unexpected URL, want=http://localhost:8080, got=%s", u)
	}
}

func TestNew_WithMultipleAddresses(t *testing.T) {
	t.Parallel()
	c, err := New(WithAddresses("http://es01:9200", "http://es02:9200"))
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	urls := c.Transport.(*elastictransport.Client).URLs()
	if len(urls) != 2 {
		t.Fatalf("Expected 2 URLs, got %d", len(urls))
	}
	if urls[0].String() != "http://es01:9200" {
		t.Errorf("Unexpected URL[0], want=http://es01:9200, got=%s", urls[0])
	}
	if urls[1].String() != "http://es02:9200" {
		t.Errorf("Unexpected URL[1], want=http://es02:9200, got=%s", urls[1])
	}
}

func TestNew_WithAddressesFromEnvironment(t *testing.T) {
	t.Setenv("ELASTICSEARCH_URL", "http://example.com")
	c, err := New()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	u := c.Transport.(*elastictransport.Client).URLs()[0].String()
	if u != "http://example.com" {
		t.Errorf("Unexpected URL, want=http://example.com, got=%s", u)
	}
}

func TestNew_WithAddressesOverridesEnvironment(t *testing.T) {
	t.Setenv("ELASTICSEARCH_URL", "http://example.com")
	c, err := New(WithAddresses("http://localhost:8080"))
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	u := c.Transport.(*elastictransport.Client).URLs()[0].String()
	if u != "http://localhost:8080" {
		t.Errorf("Unexpected URL, want=http://localhost:8080, got=%s", u)
	}
}

func TestNew_WithCloudID(t *testing.T) {
	t.Parallel()
	c, err := New(WithCloudID("foo:YmFyLmNsb3VkLmVzLmlvJGFiYzEyMyRkZWY0NTY="))
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	u := c.Transport.(*elastictransport.Client).URLs()[0].String()
	if u != "https://abc123.bar.cloud.es.io" {
		t.Errorf("Unexpected URL, want=https://abc123.bar.cloud.es.io, got=%s", u)
	}
}

func TestNew_CloudIDOverridesEnvironment(t *testing.T) {
	t.Setenv("ELASTICSEARCH_URL", "http://example.com")
	c, err := New(WithCloudID("foo:YmFyLmNsb3VkLmVzLmlvJGFiYzEyMyRkZWY0NTY="))
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	u := c.Transport.(*elastictransport.Client).URLs()[0].String()
	if u != "https://abc123.bar.cloud.es.io" {
		t.Errorf("Unexpected URL, want=https://abc123.bar.cloud.es.io, got=%s", u)
	}
}

func TestNew_AddressesAndCloudID_Error(t *testing.T) {
	t.Parallel()
	_, err := New(
		WithAddresses("http://localhost:8080"),
		WithCloudID("foo:ABC="),
	)
	if err == nil {
		t.Fatal("Expected error when both Addresses and CloudID are set")
	}
	match, _ := regexp.MatchString("both .* are set", err.Error())
	if !match {
		t.Errorf("Unexpected error message: %s", err)
	}
}

func TestNew_InvalidCloudID(t *testing.T) {
	t.Parallel()
	for _, id := range []string{"foo:ZZZ===", "foo:Zm9v", "foo:"} {
		_, err := New(WithCloudID(id))
		if err == nil {
			t.Errorf("Expected error for CloudID %q", id)
		}
	}
}

func TestNew_InvalidURL(t *testing.T) {
	t.Parallel()
	_, err := New(WithAddresses(":foo"))
	if err == nil {
		t.Fatal("Expected error for invalid URL")
	}
}

func TestNew_InvalidURLFromEnvironment(t *testing.T) {
	t.Setenv("ELASTICSEARCH_URL", ":foobar")
	_, err := New()
	if err == nil {
		t.Fatal("Expected error for invalid URL from environment")
	}
}

func TestNew_WithBasicAuth(t *testing.T) {
	t.Parallel()

	var capturedReq *http.Request
	c, err := New(
		WithBasicAuth("testuser", "testpass"),
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					capturedReq = req
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
						Body:       io.NopCloser(strings.NewReader("{}")),
					}, nil
				},
			}),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	_, err = c.Info()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	user, pass, ok := capturedReq.Header.Get("Authorization"), "", true
	_ = pass
	if !ok || user == "" {
		t.Error("Expected Authorization header to be set")
	}
}

func TestNew_WithAPIKey(t *testing.T) {
	t.Parallel()

	var capturedReq *http.Request
	c, err := New(
		WithAPIKey("dGVzdGtleQ=="),
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					capturedReq = req
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
						Body:       io.NopCloser(strings.NewReader("{}")),
					}, nil
				},
			}),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	_, err = c.Info()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	auth := capturedReq.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "APIKey ") {
		t.Errorf("Expected ApiKey authorization, got: %s", auth)
	}
}

func TestNew_WithTransportOptions(t *testing.T) {
	t.Parallel()

	c, err := New(
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{}),
			elastictransport.WithMaxRetries(5),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	_, err = c.Perform(&http.Request{URL: &url.URL{}, Header: make(http.Header)})
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

func TestNew_WithCompatibilityMode(t *testing.T) {
	t.Parallel()

	var capturedReq *http.Request
	c, err := New(
		WithCompatibilityMode(),
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					capturedReq = req
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
						Body:       io.NopCloser(strings.NewReader("{}")),
					}, nil
				},
			}),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Body:   io.NopCloser(strings.NewReader("{}")),
	}
	_, err = c.Perform(req)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if capturedReq.Header.Get("Accept") != compatibilityHeader {
		t.Errorf("Expected compatibility Accept header, got: %s", capturedReq.Header.Get("Accept"))
	}
	if capturedReq.Header.Get("Content-Type") != compatibilityHeader {
		t.Errorf("Expected compatibility Content-Type header, got: %s", capturedReq.Header.Get("Content-Type"))
	}
}

func TestNew_WithCompatibilityModeEnv(t *testing.T) {
	t.Setenv(esCompatHeader, "true")

	c, err := New(
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{
				RoundTripFunc: func(_ *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
						Body:       io.NopCloser(strings.NewReader("{}")),
					}, nil
				},
			}),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	req := &http.Request{
		URL:    &url.URL{},
		Header: make(http.Header),
		Body:   io.NopCloser(strings.NewReader("{}")),
	}
	_, err = c.Perform(req)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

func TestNew_WithDisableMetaHeader(t *testing.T) {
	t.Parallel()

	var capturedReq *http.Request
	c, err := New(
		WithDisableMetaHeader(),
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					capturedReq = req
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
						Body:       io.NopCloser(strings.NewReader("{}")),
					}, nil
				},
			}),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	_, err = c.Perform(&http.Request{URL: &url.URL{}, Header: make(http.Header)})
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if capturedReq.Header.Get(HeaderClientMeta) != "" {
		t.Errorf("Expected meta header to be absent, got: %s", capturedReq.Header.Get(HeaderClientMeta))
	}
}

func TestNew_MetaHeader(t *testing.T) {
	t.Parallel()

	c, err := New(
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					h := req.Header.Get(HeaderClientMeta)
					if !metaHeaderReValidation.MatchString(h) {
						t.Errorf("expected meta header to match regexp, got: %s", h)
					}
					if strings.Contains(h, "hl=1") {
						t.Errorf("low-level client should not have hl=1, got: %s", h)
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
						Body:       io.NopCloser(strings.NewReader("{}")),
					}, nil
				},
			}),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	_, err = c.Info()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

func TestNewTyped_Default(t *testing.T) {
	t.Parallel()
	c, err := NewTyped()
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	u := c.Transport.(*elastictransport.Client).URLs()[0].String()
	if u != defaultURL {
		t.Errorf("Unexpected URL, want=%s, got=%s", defaultURL, u)
	}
}

func TestNewTyped_MetaHeader(t *testing.T) {
	t.Parallel()

	c, err := NewTyped(
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					h := req.Header.Get(HeaderClientMeta)
					if !metaHeaderReValidation.MatchString(h) {
						t.Errorf("expected meta header to match regexp, got: %s", h)
					}
					if !strings.Contains(h, "hl=1") {
						t.Errorf("typed client should have hl=1, got: %s", h)
					}
					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"X-Elastic-Product": []string{"Elasticsearch"}},
						Body:       io.NopCloser(strings.NewReader("{}")),
					}, nil
				},
			}),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	_, _ = c.Info().Do(context.Background())
}

func TestNewTyped_WithAddresses(t *testing.T) {
	t.Parallel()
	c, err := NewTyped(WithAddresses("http://localhost:8080"))
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	u := c.Transport.(*elastictransport.Client).URLs()[0].String()
	if u != "http://localhost:8080" {
		t.Errorf("Unexpected URL, want=http://localhost:8080, got=%s", u)
	}
}

func TestNew_Close(t *testing.T) {
	t.Parallel()
	c, err := New(
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{}),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if closeErr := c.Close(context.Background()); closeErr != nil {
		t.Errorf("Unexpected error: %s", closeErr)
	}
	_, err = c.Perform(&http.Request{URL: &url.URL{}, Header: make(http.Header)})
	if err == nil {
		t.Fatal("Expected error after close")
	}
}

func TestNewTyped_Close(t *testing.T) {
	t.Parallel()
	c, err := NewTyped(
		WithTransportOptions(
			elastictransport.WithTransport(&mockTransp{}),
		),
	)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
	if closeErr := c.Close(context.Background()); closeErr != nil {
		t.Errorf("Unexpected error: %s", closeErr)
	}
	_, err = c.Perform(&http.Request{URL: &url.URL{}, Header: make(http.Header)})
	if err == nil {
		t.Fatal("Expected error after close")
	}
}
