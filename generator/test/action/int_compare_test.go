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

func TestIntCompare(t *testing.T) {
	spec := `
"Index with made up int comparison":

- lt: { foo: 10000 }
`
	var s map[string][]*Router
	if err := yaml.Unmarshal([]byte(spec), &s); err != nil {
		t.Fatal(err)
	}
	if len(s["Index with made up int comparison"]) != 1 {
		t.Fatalf("expected to see 1 action, found %d", len(s["Index with made up int comparison"]))
	}
	action := s["Index with made up int comparison"][0]
	templates, err := template.ParseFiles("templates/int_compare.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	if err = action.Resolve(nil, templates); err != nil {
		t.Fatal(err)
	}
	code, err := action.String()
	if err != nil {
		t.Fatal(err)
	}
	expectedCode := `if v, err := body.GetValue(foo); err == nil {
		if i, ok := v.(int); ok {
			if !(i < 10000) {
				t.Fatalf("unexpected value for %q: %d (expected: %d)", foo, i, 10000)
			}
		} else {
			t.Fatalf("unexpected type for %q: %T (expected int)", foo, v)
		}
} else {
	t.Fatalf("unable to find key %q: %s", foo, err)
}
`
	if d := common.Diff(t, expectedCode, code); len(d) > 0 {
		t.Fail()
	}
}
