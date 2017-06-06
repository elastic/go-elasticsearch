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
	"fmt"
	"strings"
	"text/template"

	"github.com/elastic/go-elasticsearch/generator/api"
	"github.com/elastic/go-elasticsearch/generator/test/action"
	"github.com/serenize/snaker"

	yaml "gopkg.in/yaml.v2"
)

type testType int

const (
	testTypeSetup = iota
	testTypeTest
	testTypeTeardown
)

type testcase struct {
	Spec    map[string][]*action.Router
	RawName string
	Name    string
	Type    testType
	Actions []*action.Router
}

func newTestcase(testSpecFile, testSpec string, methods map[string]*api.Method, templates *template.Template) (*testcase, error) {
	t := &testcase{}
	if err := yaml.Unmarshal([]byte(testSpec), &t.Spec); err != nil {
		return nil, err
	}
	if len(t.Spec) > 3 {
		return nil, fmt.Errorf("document %#v contains more than one test", t.Spec)
	}
	for name, actions := range t.Spec {
		switch name {
		case "setup":
			t.Type = testTypeSetup
		case "teardown":
			t.Type = testTypeTeardown
		default:
			t.RawName = name
			t.Type = testTypeTest
			// TODO: make these replacements more concise.
			name = strings.Replace(name, " ", "_", -1)
			name = strings.Replace(name, ",", "_", -1)
			name = strings.Replace(name, "=", "_equals_", -1)
			t.Name = "Test" + snaker.SnakeToCamel(name)
			t.Actions = actions
		}
		for _, a := range actions {
			if err := a.Resolve(testSpecFile, methods, templates); err != nil {
				return nil, err
			}
		}
	}
	return t, nil
}
