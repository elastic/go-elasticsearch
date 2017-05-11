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
	"fmt"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

// apiTest is the test generator for an API namespace. It contains multiple tests, each of which is exercising one or
// more methods.
type apiTester struct {
	Specs       []*testSpec
	PackageName string
	template    *template.Template
}

// newAPITest instantiates a tester for a given API namespace.
func newAPITester(specDir, api string, methods map[string]*method, templates *template.Template) (*apiTester, error) {
	testSpecDir := filepath.Join(specDir, "test", api)
	files, err := ioutil.ReadDir(testSpecDir)
	if err != nil {
		return nil, err
	}
	a := &apiTester{
		Specs: []*testSpec{},
		// TODO: resolve package name
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".yaml" {
			continue
		}
		ts, err := newTestSpec(filepath.Join(testSpecDir, file.Name()), methods, templates)
		if err != nil {
			return nil, err
		}
		a.Specs = append(a.Specs, ts)
	}
	if a.template = templates.Lookup("test.tmpl"); a.template == nil {
		return nil, fmt.Errorf("cannot find template for tests")
	}
	return a, nil
}

func (a *apiTester) generate(outputDir string) error {
	// TODO: add boilerplate code to template (vars etc), implement generation
	return nil
}
