// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package clusterstatus_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/elastic/go-elasticsearch/v8/_examples/clusterstatus"
)

// Mock transport replaces the HTTP transport for tests
type MockTransport struct{}

// RoundTrip returns a mock response.
func (t *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{Body: ioutil.NopCloser(strings.NewReader(`{"status":"mocked"}`))}, nil
}

func TestHealth(t *testing.T) {
	clusterstatus.ES, _ = elasticsearch.NewClient(elasticsearch.Config{Transport: &MockTransport{}})

	w := httptest.NewRecorder()
	clusterstatus.Health(w, &http.Request{})

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("> %s %s\n> %s\n", resp.Status, resp.Header.Get("Content-Type"), body)

	if string(body) != `{"status":"mocked"}` {
		t.Errorf("Unexpected body: %s", body)
	}
}
