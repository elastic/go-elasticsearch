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
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/elastic/go-elasticsearch/generator/api"
)

// testSpec maps a spec file for one or more tests.
type testSpec struct {
	Tests []*testcase
}

func newTestSpec(specFilePath string, methods map[string]*api.Method, methodName string,
	templates *template.Template) (*testSpec, error) {
	bytes, err := ioutil.ReadFile(specFilePath)
	if err != nil {
		return nil, err
	}
	mt := &testSpec{
		Tests: []*testcase{},
	}
	for _, s := range strings.Split(string(bytes), "---") {
		t, err := newTestcase(specFilePath, s, methods, methodName, templates)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %q: %s", specFilePath, err)
		}
		mt.Tests = append(mt.Tests, t)
	}
	return mt, nil
}
