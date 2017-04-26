/*
 * Licensed to Elasticsearch under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package generator

import (
	"bytes"
	"fmt"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	m := newIndexMethod(t)
	var writer bytes.Buffer
	err := m.generate(filepath.Join("..", templatesDir), &writer)
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `// generated by github.com/elastic/goelasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"
)

// Index adds or updates a typed JSON document in a specific index, making it searchable. See http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html for more info.
//
// documentType: the type of the document.
//
// index: the name of the index.
//
// body: the document.
//
// options: optional parameters. Supports the following functional options: WithID, WithOpType, WithTimeout, WithVersion, WithWaitForActiveShards, see the Option type in this package for more info.
func (a *API) Index(documentType string, index string, body map[string]interface{}, options ...*Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
		"WithID": struct{}{},
		"WithOpType": struct{}{},
		"WithTimeout": struct{}{},
		"WithVersion": struct{}{},
		"WithWaitForActiveShards": struct{}{},
	}
	req := &http.Request{
		URL: &url.URL{
			Scheme: a.transport.URL.Scheme,
			Host:   a.transport.URL.Host,
		},
		Method: "POST",
	}
	for _, option := range options{
		if _, ok := supportedOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	return a.transport.Do(req)
}
`
	if d := diff(t, expectedCode, writer.String()); len(d) > 0 {
		t.Fail()
	}
}

func TestNormalizeParams(t *testing.T) {
	m, err := newMethod("testdata/cat.fielddata.json", nil)
	if err != nil {
		t.Fatal(err)
	}
	names := map[string]struct{}{}
	for _, p := range m.OptionalParams {
		if _, ok := names[p.Name]; ok {
			t.Fatalf("Param %s seen twice", p.Name)
		}
		names[p.Name] = struct{}{}
	}
	if len(names) != len(m.OptionalParams) {
		t.Fatalf("Not all params had unique names")
	}
	for _, name := range []string{"fields", "fieldsParam"} {
		if _, ok := names[name]; !ok {
			t.Fatalf("Could not find %q in %s", name, names)
		}
	}
}

func TestResolveDocumentation(t *testing.T) {
	var tests = []struct {
		url         string
		expectedDoc string
	}{
		{
			url:         "http://www.elastic.co/guide/",
			expectedDoc: " - see %s for more info.",
		},
		{
			url:         "https://www.elastic.co/guide/",
			expectedDoc: " - see %s for more info.",
		},
	}
	m := newIndexMethod(t)
	for _, test := range tests {
		m.Spec.Documentation = test.url
		if err := m.resolveDocumentation(); err != nil {
			t.Fatal(err)
		}
		expectedDoc := fmt.Sprintf(test.expectedDoc, test.url)
		if m.Spec.Documentation != expectedDoc {
			t.Fatalf("expected %q to generate %q, got %q instead", test.url, expectedDoc, m.Spec.Documentation)
		}
	}
}
