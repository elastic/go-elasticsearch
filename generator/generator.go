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

// Package generator allows to generate the Elasticsearch APIs by fitting their JSON spec into the templates stored in
// the templates directory.
package generator

import (
	"io/ioutil"
	"path/filepath"
	"text/template"

	"github.com/golang/glog"
)

const (
	outputDir = "."
	// DefaultTemplatesDir is the directory containing the templates for the code generation.
	DefaultTemplatesDir = "generator/templates"
	defaultPackage      = "api"
	templateFileName    = "method.tmpl"
	defaultPackageRepo  = "github.com/elastic/go-elasticsearch/api"
)

// Generator is a code generator based on JSON and YAML specs and Go template.
type Generator struct {
	// methods are API methods populated and map 1:1 with the JSON specs in the rest-api-spec/api dir.
	methods map[string]*method
	// commonParams are common options populated based on rest-api-spec/api/_common.json.
	commonParams map[string]*param
	// packages are Go packages that wrap the namespaces of the APIs. They also generate common objects such as
	// constructors, functional options and enums.
	packages map[string]*goPackage
	// tests are groups of tests for each API, populated with the YAML specs in the rest-api-spec/test dir. They map
	// 1:1 with the directories in rest-api-spec/test.
	tests map[string]*apiTester
}

// New creates a new generator
func New(specDir, templatesDir string) (*Generator, error) {
	g := &Generator{
		methods: map[string]*method{},
		tests:   map[string]*apiTester{},
	}
	glog.Info("parsing templates")
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return nil, err
	}
	templateFiles := []string{}
	for _, file := range files {
		if !file.IsDir() {
			templateFiles = append(templateFiles, filepath.Join(templatesDir, file.Name()))
		}
	}
	files, err = ioutil.ReadDir(filepath.Join(templatesDir, "options"))
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		templateFiles = append(templateFiles, filepath.Join(templatesDir, "options", file.Name()))
	}
	templates, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}
	glog.Info("parsing common params")
	g.commonParams, err = newCommonParams(specDir, templates)
	if err != nil {
		return nil, err
	}
	files, err = ioutil.ReadDir(filepath.Join(specDir, "api"))
	if err != nil {
		return nil, err
	}
	glog.Info("parsing specs")
	for _, specFile := range files {
		if specFile.Name() == commonParamsSpecFile {
			continue
		}
		var m *method
		m, err = newMethod(specDir, specFile.Name(), g.commonParams, templates)
		if err != nil {
			return nil, err
		}
		g.methods[m.rawName] = m
	}
	g.packages, err = newPackages(g.methods, templates)
	if err != nil {
		return nil, err
	}
	testDirs, err := ioutil.ReadDir(filepath.Join(specDir, "test"))
	if err != nil {
		return nil, err
	}
	for _, dir := range testDirs {
		if !dir.IsDir() {
			continue
		}
		if g.tests[dir.Name()], err = newAPITester(specDir, dir.Name(), g.methods, templates); err != nil {
			return nil, err
		}
	}
	return g, nil
}

func newPackages(methods map[string]*method, templates *template.Template) (map[string]*goPackage, error) {
	packages := map[string]*goPackage{}
	for _, m := range methods {
		if p, ok := packages[m.PackageName]; ok {
			p.addMethod(m)
		} else {
			var err error
			packages[m.PackageName], err = newGoPackage(m, templates)
			if err != nil {
				return nil, err
			}
		}
	}
	rootPackage := packages[defaultPackage]
	for _, p := range packages {
		if p.Methods[0].PackageName != rootPackage.Methods[0].PackageName {
			rootPackage.addSubpackage(p)
		}
	}
	return packages, nil
}

// Run runs the generator
func (g *Generator) Run() error {
	glog.Info("generating methods")
	for _, m := range g.methods {
		w, err := m.newWriter(outputDir, "")
		if err != nil {
			return err
		}
		if err = m.generate(w); err != nil {
			return err
		}
	}
	glog.Info("generating types, constructors and options")
	for _, p := range g.packages {
		if err := p.generate(outputDir); err != nil {
			return err
		}
	}
	glog.Info("generating tests")
	for _, test := range g.tests {
		err := test.generate(outputDir)
		if err != nil {
			return err
		}
	}
	return nil
}
