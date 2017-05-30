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

package action

import (
	"testing"
	"text/template"

	"github.com/elastic/go-elasticsearch/generator/common"

	yaml "gopkg.in/yaml.v2"
)

func TestMatch(t *testing.T) {
	spec := `
"Index with ID":

 - match:   { _index:   test-weird-index-中文 }
`
	var s map[string][]*Router
	if err := yaml.Unmarshal([]byte(spec), &s); err != nil {
		t.Fatal(err)
	}
	if len(s["Index with ID"]) != 1 {
		t.Fatalf("expected to see 1 action, found %d", len(s["Index with ID"]))
	}
	action := s["Index with ID"][0]
	templates, err := template.ParseFiles("templates/match.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	if err = action.Resolve("", nil, templates); err != nil {
		t.Fatal(err)
	}
	code, err := action.String()
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `for name, expectedValue := range map[string]string{
	"_index" : "test-weird-index-中文",
} {
	value, ok := body[name]
	if !ok {
		t.Fatalf("response does not contain %q", name)
	}
	if value != expectedValue {
		t.Fatalf("expected %q to be %q, got %q", name, expectedValue, value)
	}
}
`
	if d := common.Diff(t, expectedCode, code); len(d) > 0 {
		t.Fail()
	}
}
