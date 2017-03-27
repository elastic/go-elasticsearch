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
	"io/ioutil"
	"path/filepath"
	"text/template"
)

// tester is the test generator for an API method
type tester struct {
	testers []*methodTester
}

// New instantiates a tester for a given API
func newTester(specDir, api string, methods map[string]*method, templates *template.Template) (*tester, error) {
	testSpecDir := filepath.Join(specDir, "test", api)
	files, err := ioutil.ReadDir(testSpecDir)
	if err != nil {
		return nil, err
	}
	t := &tester{
		testers: []*methodTester{},
	}
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".yaml" {
			continue
		}
		mt, err := newMethodTester(filepath.Join(testSpecDir, file.Name()), methods, templates)
		if err != nil {
			return nil, err
		}
		t.testers = append(t.testers, mt)
	}
	return t, nil
}

func (t *tester) generate(templates *template.Template, outputDir string) error {
	// TODO: implement
	return nil
}
