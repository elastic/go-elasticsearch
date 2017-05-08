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
	"text/template"
)

func TestGenerate(t *testing.T) {
	m := newIndexMethod(t)
	var writer bytes.Buffer
	if err := m.generate(&writer); err != nil {
		t.Fatal(err)
	}
	expectedCode := `// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/elastic/go-elasticsearch/transport"
)

// Index adds or updates a typed JSON document in a specific index, making it searchable. See http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html for more info.
//
// index: the name of the index.
//
// documentType: the type of the document.
//
// body: the document.
//
// options: optional parameters. Supports the following functional options: WithID, WithOpType, WithTimeout, WithVersion, WithWaitForActiveShards, see the Option type in this package for more info.
func (a *API) Index(index string, documentType string, body map[string]interface{}, options ...*Option) (*IndexResponse, error) {
	req := a.transport.NewRequest("POST")
	methodOptions := supportedOptions["Index"]
	for _, option := range options{
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := a.transport.Do(req)
	return &IndexResponse{resp}, err
}

// IndexResponse is the response for Index.
type IndexResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *IndexResponse) DecodeBody() (map[string]interface{}, error) {
	return transport.DecodeResponseBody(r.Response)
}
`
	if d := diff(t, expectedCode, writer.String()); len(d) > 0 {
		t.Fail()
	}
}

func TestNormalizeParams(t *testing.T) {
	templates, err := template.ParseFiles(filepath.Join("..", DefaultTemplatesDir, "option.tmpl"))
	if err != nil {
		t.Fatal(err)
	}
	m, err := newMethod(testSpecDir, "cat.fielddata.json", nil, templates)
	if err != nil {
		t.Fatal(err)
	}

	// Required params
	names := map[string]struct{}{}
	for _, p := range m.RequiredURLParts {
		if _, ok := names[p.Name]; ok {
			t.Fatalf("Param %s seen twice", p.Name)
		}
		names[p.Name] = struct{}{}
	}
	if len(names) != len(m.RequiredURLParts) {
		t.Fatalf("Not all URL parts had unique names")
	}

	// Optional params
	names = map[string]struct{}{}
	for _, p := range m.OptionalURLParts {
		if _, ok := names[p.Name]; ok {
			t.Fatalf("Param %s seen twice", p.Name)
		}
		names[p.Name] = struct{}{}
	}
	for _, p := range m.OptionalURLParams {
		if _, ok := names[p.Name]; ok {
			t.Fatalf("Param %s seen twice", p.Name)
		}
		names[p.Name] = struct{}{}
	}
	if len(names) != len(m.OptionalURLParts)+len(m.OptionalURLParams) {
		t.Fatalf("Not all URL parts had unique names")
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

func TestClone(t *testing.T) {
	m := newIndexMethod(t)
	c, err := m.clone()
	if err != nil {
		t.Fatal(err)
	}
	var tests = []struct {
		params        []*param
		clonedParams  []*param
		expectedNames []string
	}{
		{
			m.RequiredURLParts,
			c.RequiredURLParts,
			[]string{
				"index",
				"documentType",
			},
		},
		{
			m.RequiredURLParams,
			c.RequiredURLParams,
			[]string{
				"body",
			},
		},
		{
			m.OptionalURLParts,
			c.OptionalURLParts,
			[]string{
				"id",
			},
		},
		{
			m.OptionalURLParams,
			c.OptionalURLParams,
			[]string{
				"opType",
				"timeout",
				"version",
				"waitForActiveShards",
			},
		},
	}
	for _, test := range tests {
		if len(test.params) != len(test.expectedNames) {
			t.Fatalf("expected %s params but got %#v", test.expectedNames, test.params)
		}
		if len(test.params) != len(test.clonedParams) {
			t.Fatalf("expected %d params but got %#v", len(test.params), test.clonedParams)
		}
		for i, name := range test.expectedNames {
			p := test.params[i]
			if p.Name != name {
				t.Fatalf("expected %q but got %q", name, p.Name)
			}
			if !p.equals(test.clonedParams[i]) {
				t.Fatalf("cloned param %q differs", p.Name)
			}
		}
	}
}
