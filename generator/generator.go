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

	"github.com/elastic/go-elasticsearch/generator/api"
	"github.com/elastic/go-elasticsearch/generator/test"
	"github.com/golang/glog"
)

const (
	// DefaultSpecDir is the dir holding the specs for the APIs and their tests.
	DefaultSpecDir = "rest-api-spec"
	outputDir      = "."
)

// Generator is a code generator based on JSON and YAML specs and Go template.
type Generator struct {
	// methods are API methods populated and map 1:1 with the JSON specs in the rest-api-spec/api dir.
	methods map[string]*api.Method
	// commonParams are common options populated based on rest-api-spec/api/_common.json.
	commonParams map[string]*api.Param
	// packages are Go packages that wrap the namespaces of the APIs. They also generate common objects such as
	// constructors, functional options and enums.
	packages map[string]*api.Package
	// testers are groups of tests for each API, populated with the YAML specs in the rest-api-spec/test dir. They map
	// 1:1 with the directories in rest-api-spec/test.
	testers map[string]*test.Method
}

// New creates a new generator
func New(specDir, templatesRoot string, offline bool) (*Generator, error) {
	g := &Generator{
		methods: map[string]*api.Method{},
		testers: map[string]*test.Method{},
	}
	glog.Info("parsing api templates")
	dir := filepath.Join(templatesRoot, "api/templates")
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	templateFiles := []string{}
	for _, file := range files {
		if !file.IsDir() {
			templateFiles = append(templateFiles, filepath.Join(dir, file.Name()))
		}
	}
	dir = filepath.Join(templatesRoot, "api/templates/options")
	files, err = ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		templateFiles = append(templateFiles, filepath.Join(dir, file.Name()))
	}
	templates, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}
	glog.Info("parsing common params")
	g.commonParams, err = api.NewCommonParams(specDir, templates)
	if err != nil {
		return nil, err
	}
	files, err = ioutil.ReadDir(filepath.Join(specDir, "api"))
	if err != nil {
		return nil, err
	}
	glog.Info("parsing method specs")
	for _, specFile := range files {
		if specFile.Name() == api.CommonParamsSpecFile {
			continue
		}
		var m *api.Method
		m, err = api.NewMethod(specDir, specFile.Name(), g.commonParams, templates, offline)
		if err != nil {
			return nil, err
		}
		g.methods[m.RawName] = m
	}
	glog.Info("grouping package info")
	g.packages, err = newPackages(g.methods, templates)
	if err != nil {
		return nil, err
	}
	testDirs, err := ioutil.ReadDir(filepath.Join(specDir, "test"))
	if err != nil {
		return nil, err
	}

	glog.Info("parsing test templates")
	dir = filepath.Join(templatesRoot, "test/templates")
	files, err = ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() {
			templateFiles = append(templateFiles, filepath.Join(dir, file.Name()))
		}
	}
	dir = filepath.Join(templatesRoot, "test/action/templates")
	files, err = ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		templateFiles = append(templateFiles, filepath.Join(dir, file.Name()))
	}
	templates, err = template.ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}
	glog.Info("parsing test specs")
	for _, dir := range testDirs {
		if !dir.IsDir() {
			continue
		}
		if g.testers[dir.Name()], err = test.NewMethod(specDir, dir.Name(), g.methods, templates); err != nil {
			return nil, err
		}
	}
	return g, nil
}

func newPackages(methods map[string]*api.Method, templates *template.Template) (map[string]*api.Package, error) {
	packages := map[string]*api.Package{}
	for _, m := range methods {
		if p, ok := packages[m.PackageName]; ok {
			p.AddMethod(m)
		} else {
			var err error
			packages[m.PackageName], err = api.NewPackage(m, templates)
			if err != nil {
				return nil, err
			}
		}
	}
	rootPackage := packages[api.RootPackage]
	for _, p := range packages {
		if p.Methods[0].PackageName != rootPackage.Methods[0].PackageName {
			rootPackage.AddSubpackage(p)
		}
	}
	return packages, nil
}

// Run runs the generator
func (g *Generator) Run() error {
	glog.Info("generating methods")
	for _, m := range g.methods {
		w, err := m.NewWriter(outputDir, "")
		if err != nil {
			return err
		}
		if err = m.Generate(w); err != nil {
			return err
		}
	}
	glog.Info("generating types, constructors and options")
	for _, p := range g.packages {
		if err := p.Generate(outputDir); err != nil {
			return err
		}
	}
	glog.Info("generating testers")
	for _, tester := range g.testers {
		if err := tester.Generate(outputDir); err != nil {
			return err
		}
	}
	return nil
}
