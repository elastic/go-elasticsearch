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

	"github.com/golang/glog"
)

const (
	outputDir            = "."
	templatesDir         = "generator/templates"
	commonParamsSpecFile = "_common.json"
)

// Generator is a code generator based on JSON specs and Go template
type Generator struct {
	methods      []*method
	commonParams map[string]*param
}

// New creates a new generator
func New(specDir string) (*Generator, error) {
	g := &Generator{
		methods: []*method{},
	}
	commonParamsSpecFilePath := filepath.Join(specDir, "api", commonParamsSpecFile)
	var err error
	glog.Infof("parsing %s", commonParamsSpecFile)
	g.commonParams, err = newCommonParams(commonParamsSpecFilePath)
	if err != nil {
		return nil, err
	}
	files, err := ioutil.ReadDir(filepath.Join(specDir, "api"))
	if err != nil {
		return nil, err
	}
	glog.Info("parsing specs")
	for _, specFile := range files {
		if specFile.Name() == commonParamsSpecFile {
			continue
		}
		m, err := newMethod(filepath.Join(specDir, "api", specFile.Name()), g.commonParams)
		if err != nil {
			return nil, err
		}
		g.methods = append(g.methods, m)
	}
	return g, nil
}

// Run runs the generator
func (g *Generator) Run() error {
	glog.Info("generating methods")
	for _, m := range g.methods {
		w, err := m.newWriter(outputDir, "")
		if err != nil {
			return err
		}
		if err = m.generate(templatesDir, w); err != nil {
			return err
		}
	}
	a, err := newAPIPackages(g.methods)
	if err != nil {
		return err
	}
	glog.Info("generating types, constructors and options")
	return a.generate(templatesDir, outputDir)
}
