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

package test

import (
	"testing"
	"text/template"

	"github.com/elastic/go-elasticsearch/generator/api"
	"github.com/elastic/go-elasticsearch/generator/common"
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
	templates, err := template.ParseFiles("action/templates/do.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	tc, err := newTestcase("", spec, map[string]*api.Method{"index": m}, templates)
	if err != nil {
		t.Fatal(err)
	}
	expectedName := "TestIndexWithID"
	if tc.Name != expectedName {
		t.Fatalf("unexpected test name: %s (expected %q)", tc.Name, expectedName)
	}
	if len(tc.Spec["Index with ID"]) != 1 {
		t.Fatalf("expected to see 1 action, found %d", len(tc.Spec["Index with ID"]))
	}
	action := tc.Spec["Index with ID"][0]
	s, err := action.String()
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `indexResp, err = a.Index("test-weird-index-中文", "weird.type", map[string]interface{}{
		}, WithID("1"))
if err != nil {
	t.Fatal(err)
}
body, err = indexResp.DecodeBody()
if err != nil {
	t.Fatal(err)
}
`
	if d := common.Diff(t, expectedCode, s); len(d) > 0 {
		t.Fail()
	}
}

func TestDoWithMissingParam(t *testing.T) {
	spec := `
"check delete with blank index and blank alias":
 - do:
      catch: param
      indices.delete_alias:
          name: "alias1"`
	templates, err := template.ParseFiles("../api/templates/method.tmpl", "../api/templates/option.tmpl", "action/templates/do.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	m, err := api.NewMethod("testdata", "indices.delete_alias.json", nil, templates, true)
	if err != nil {
		t.Fatal(err)
	}
	tc, err := newTestcase("", spec, map[string]*api.Method{"indices.delete_alias": m}, templates)
	if err != nil {
		t.Fatal(err)
	}
	action := tc.Spec["check delete with blank index and blank alias"][0]
	s, err := action.String()
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `deleteAliasResp, err = i.DeleteAlias(nil, []string{
		"alias1",
})
if err != nil {
	t.Fatal(err)
}
body, err = deleteAliasResp.DecodeBody()
if err != nil {
	t.Fatal(err)
}
`
	if d := common.Diff(t, expectedCode, s); len(d) > 0 {
		t.Fail()
	}
}
