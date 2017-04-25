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

package client

import (
	"net/http"
	"runtime"
)

// Index adds or updates a typed JSON document in a specific index, making it searchable. See http://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html for more info.
//
// documentType: the type of the document.
//
// index: the name of the index.
//
// body: the document.
//
// options: optional parameters. Supports the following functional options: WithID, WithOpType, WithParent, WithPipeline, WithRefresh, WithRouting, WithTimeout, WithTimestamp, WithTTL, WithVersion, WithVersionType, WithWaitForActiveShards, see the Option type in this package for more info.
func (c *Client) Index(documentType string, index string, body map[string]interface{}, options ...Option) (*http.Response, error) {
	supportedOptions := map[string]struct{}{
"WithID": struct{}{},
"WithOpType": struct{}{},
"WithParent": struct{}{},
"WithPipeline": struct{}{},
"WithRefresh": struct{}{},
"WithRouting": struct{}{},
"WithTimeout": struct{}{},
"WithTimestamp": struct{}{},
"WithTTL": struct{}{},
"WithVersion": struct{}{},
"WithVersionType": struct{}{},
"WithWaitForActiveShards": struct{}{},
}
	for _, option := range options{
		name := runtime.FuncForPC(reflect.ValueOf(option).Pointer()).Name()
		if _, ok := supportedOptions[name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", name)
		}
	}
	req := &http.Request{
		Method: "POST",
	}
	return c.client.Do(req)
}
`
	if d := diff(t, expectedCode, writer.String()); len(d) > 0 {
		t.Fail()
	}
}

func TestNormalizeParams(t *testing.T) {
	m, err := newMethod(filepath.Join("..", DefaultSpecDir, "cat.fielddata.json"), nil)
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
