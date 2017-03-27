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
	"io"
	"text/template"
)

type goPackage struct {
	Name        string
	Repo        string
	TypeName    string
	Methods     []*method
	Options     map[string]*param
	Enums       map[string]*param
	SubPackages map[string]*goPackage
}

func newGoPackage(m *method) (*goPackage, error) {
	p := &goPackage{
		Name:     m.PackageName,
		Repo:     m.Repo,
		TypeName: m.TypeName,
		Options:  map[string]*param{},
		Enums:    map[string]*param{},
	}
	err := p.addMethod(m)
	return p, err
}

func (p *goPackage) addParam(op *param) error {
	if existingParam, ok := p.Options[op.Name]; ok {
		if !existingParam.equals(op) {
			return fmt.Errorf("found two different versions of %q in %q", op.Name, p.Name)
		}
	} else {
		p.Options[op.Name] = op
	}
	if op.SpecType == "enum" {
		if existingEnum, ok := p.Enums[op.Name]; ok {
			if !existingEnum.equals(op) {
				return fmt.Errorf("found two different versions of %q in %q", op.Name, p.Name)
			}
		} else {
			p.Enums[op.Name] = op
		}
	}
	return nil
}

func (p *goPackage) addMethod(m *method) error {
	p.Methods = append(p.Methods, m)
	for _, op := range m.OptionalURLParts {
		if err := p.addParam(op); err != nil {
			return fmt.Errorf("failed to add method %s to package %s: %s", m.rawName, p.Name, err)
		}
	}
	for _, op := range m.OptionalURLParams {
		if err := p.addParam(op); err != nil {
			return fmt.Errorf("failed to add method %s to package %s: %s", m.rawName, p.Name, err)
		}
	}
	return nil
}

func (p *goPackage) addSubpackage(sub *goPackage) {
	if p.SubPackages == nil {
		p.SubPackages = map[string]*goPackage{}
	}
	p.SubPackages[sub.Methods[0].PackageName] = sub
}

func (p *goPackage) newWriter(outputDir, fileName string) (io.Writer, error) {
	return p.Methods[0].newWriter(outputDir, fileName)
}

func (p *goPackage) generateAPI(templates *template.Template, w io.Writer) error {
	if err := templates.Lookup("package.tmpl").Execute(w, p); err != nil {
		return fmt.Errorf("failed to execute package template: %s", err)
	}
	return nil
}

func (p *goPackage) generateOption(templates *template.Template, w io.Writer) error {
	if err := templates.Lookup("option.tmpl").Execute(w, p); err != nil {
		return fmt.Errorf("failed to execute option template: %s", err)
	}
	return nil
}

func (p *goPackage) generate(templates *template.Template, outputDir string) error {
	w, err := p.newWriter(outputDir, "option.go")
	if err != nil {
		return err
	}
	err = p.generateOption(templates, w)
	if err != nil {
		return err
	}
	fileName := p.Methods[0].PackageName
	if fileName == defaultPackage {
		fileName = "api"
	}
	w, err = p.newWriter(outputDir, fileName+".go")
	if err != nil {
		return err
	}
	return p.generateAPI(templates, w)
}
