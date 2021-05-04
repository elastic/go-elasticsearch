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
