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
	"fmt"
	"io"
	"text/template"
)

const (
	// RootPackage is the root package, which maps to the default namespace
	RootPackage     = "api"
	rootPackageRepo = "github.com/elastic/go-elasticsearch/api"
)

// Package is a Go package mapping an API namespace.
type Package struct {
	Name        string
	Repo        string
	TypeName    string
	Methods     []*Method
	SubPackages map[string]*Package
	templates   *template.Template
}

// NewPackage creates a new package based on a method.
func NewPackage(m *Method, templates *template.Template) *Package {
	return &Package{
		Name:      m.PackageName,
		Repo:      m.Repo,
		TypeName:  m.TypeName,
		templates: templates,
		Methods:   []*Method{m},
	}
}

// AddSubpackage adds the given subpackage.
func (p *Package) AddSubpackage(sub *Package) {
	if p.SubPackages == nil {
		p.SubPackages = map[string]*Package{}
	}
	p.SubPackages[sub.Methods[0].PackageName] = sub
}

func (p *Package) newWriter(outputDir, fileName string) (io.Writer, error) {
	return p.Methods[0].NewWriter(outputDir, fileName)
}

func (p *Package) generateAPI(w io.Writer) error {
	t := p.templates.Lookup("package.tmpl")
	if t == nil {
		return fmt.Errorf("cannot fine template for package")
	}
	if err := t.Execute(w, p); err != nil {
		return fmt.Errorf("failed to execute package template: %s", err)
	}
	return nil
}

// Generate generates a Go file with package-level code (constructor etc.).
func (p *Package) Generate(outputDir string) error {
	fileName := p.Methods[0].PackageName
	if fileName == RootPackage {
		fileName = RootPackage
	}
	w, err := p.newWriter(outputDir, fileName+".go")
	if err != nil {
		return err
	}
	return p.generateAPI(w)
}
