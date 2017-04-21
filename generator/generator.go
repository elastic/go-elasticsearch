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
)

const (
	outputDir    = "."
	templatesDir = "generator/templates"
	typesSpec    = "_common.json"
)

// Generator is a code generator based on JSON specs and Go template
type Generator struct {
	methods []*method
	types   *types
}

// New creates a new generator
func New(specDir string) (*Generator, error) {
	g := &Generator{
		methods: []*method{},
	}
	files, err := ioutil.ReadDir(specDir)
	if err != nil {
		return nil, err
	}
	for _, specFile := range files {
		specFilePath := filepath.Join(specDir, specFile.Name())
		if specFile.Name() == typesSpec {
			g.types = newTypes(specFilePath)
		} else {
			m, err := newMethod(specFilePath)
			if err != nil {
				return nil, err
			}
			g.methods = append(g.methods, m)
		}
	}
	return g, nil
}

// Run runs the generator
func (g *Generator) Run() error {
	for _, m := range g.methods {
		w, err := m.newWriter(outputDir, "")
		if err != nil {
			return err
		}
		if err = m.generate(templatesDir, w); err != nil {
			return err
		}
	}
	if err := g.types.generate(templatesDir, outputDir); err != nil {
		return err
	}
	a, err := newAPIPackages(g.methods)
	if err != nil {
		return err
	}
	return a.generate(templatesDir, outputDir)
}
