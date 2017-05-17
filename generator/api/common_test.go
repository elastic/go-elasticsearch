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

package api

import (
	"testing"
	"text/template"
)

func TestResolve(t *testing.T) {
	templates, err := template.ParseFiles("templates/option.tmpl", "templates/options/ignore.tmpl")
	if err != nil {
		t.Fatal(err)
	}
	params, err := NewCommonParams(testSpecDir, templates)
	if err != nil {
		t.Fatal(err)
	}
	if len(params) == 0 {
		t.Fatalf("Didn't find any params in the common params spec file")
	}
	human, ok := params["human"]
	if !ok {
		t.Fatalf("Expected to see 'human' in common params")
	}
	expected := &Param{
		Name:        "human",
		SpecType:    "boolean",
		Type:        "bool",
		Default:     true,
		Description: "return human readable values for statistics.",
		OptionName:  "WithHuman",
	}
	if len(human.diff(expected)) > 0 {
		t.Fatalf("Expected %#v, got %#v", expected, human)
	}
}
