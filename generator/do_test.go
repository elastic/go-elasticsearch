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
	"path/filepath"
	"testing"
	"text/template"
)

func TestDo(t *testing.T) {
	spec := `
"Index with ID":

 - do:
      index:
          index:  test-weird-index-中文
          type:   weird.type
          id:     1
          body:   { foo: bar }
`
	m := newIndexMethod(t)
	dir := filepath.Join("..", DefaultTemplatesDir)
	templates, err := template.ParseFiles(filepath.Join(dir, "do.tmpl"))
	if err != nil {
		t.Fatal(err)
	}
	yt, err := newTest(spec, map[string]*method{"index": m}, templates)
	if err != nil {
		t.Fatal(err)
	}
	expectedName := "TestIndexWithID"
	if yt.Name != expectedName {
		t.Fatalf("unexpected test name: %s (expected %q)", yt.Name, expectedName)
	}
	if len(yt.Spec["Index with ID"]) != 1 {
		t.Fatalf("expected to see 1 action, found %d", len(yt.Spec["Index with ID"]))
	}
	action := yt.Spec["Index with ID"][0]
	s, err := action.String()
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `indexResp, err := Index(
		"test-weird-index-中文", "weird.type", map[string]interface{}{
			"foo" : "bar",
		}, WithID("1")
	)
	if err != nil {
		t.Fatal(err)
	}
	body, err = indexResp.DecodeBody()
	if err != nil {
		t.Fatal(err)
	}
`
	if d := diff(t, expectedCode, s); len(d) > 0 {
		t.Fail()
	}
}
