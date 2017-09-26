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
	"path/filepath"
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
	Vars    map[string]*action.Var
}

func newTestcase(testSpecFile, testSpec string, methods map[string]*api.Method, methodName string,
	templates *template.Template) (*testcase, error) {
	t := &testcase{
		Vars: map[string]*action.Var{},
	}
	if err := yaml.Unmarshal([]byte(testSpec), &t.Spec); err != nil {
		return nil, err
	}
	if len(t.Spec) > 3 {
		return nil, fmt.Errorf("document %#v contains more than one test", t.Spec)
	}
	testSpecFileExt := filepath.Ext(testSpecFile)
	specFileBase := filepath.Base(testSpecFile[:len(testSpecFile)-len(testSpecFileExt)])
	for name, actions := range t.Spec {
		switch name {
		case "setup":
			t.Type = testTypeSetup
		case "teardown":
			t.Type = testTypeTeardown
		default:
			t.RawName = name
			t.Type = testTypeTest
			name = ""
			for _, c := range t.RawName {
				switch c {
				case '[':
					continue
				case ']':
					continue
				case '\\':
					continue
				case '"':
					continue
				case '{':
					continue
				case '}':
					continue
				case ':':
					continue
				case '\'':
					continue
				case ' ':
					name += "_"
				case ',':
					name += "_"
				case '-':
					name += "_"
				case '=':
					name += "_equals"
				case '&':
					name += "_and"
				case '*':
					name += "_star"
				case '.':
					name += "_dot"
				case '/':
					name += "_slash"
				default:
					name += string(c)
				}
			}
			t.Name = "Test" + snaker.SnakeToCamel(specFileBase+"_"+methodName+"_"+name)
			// TODO: keep raw name as comment
			t.Actions = actions
		}
		for _, a := range actions {
			ctx, err := a.Resolve(testSpecFile, methods, templates)
			if err != nil {
				return nil, err
			}
			for _, v := range ctx.Vars {
				if existing, ok := t.Vars[v.Name]; !ok {
					// TODO: handle variables that are declared and not used (maybe it'll go away with the other TODOs?)
					t.Vars[v.Name] = v
				} else {
					if existing.Type != v.Type {
						return nil, fmt.Errorf("colliding vars in contexts from the same test")
					}
				}
			}
		}
	}
	return t, nil
}
