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
	"path/filepath"
	"text/template"
)

type apiPackage struct {
	Name        string
	Repo        string
	TypeName    string
	Methods     []*method
	Options     map[string]*param
	Enums       map[string]*param
	SubPackages map[string]*apiPackage
}

func newAPIPackage(m *method) (*apiPackage, error) {
	p := &apiPackage{
		Name:     m.PackageName,
		Repo:     m.Repo,
		TypeName: m.TypeName,
		Options:  map[string]*param{},
		Enums:    map[string]*param{},
	}
	err := p.addMethod(m)
	return p, err
}

func (p *apiPackage) addParam(op *param) error {
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

func (p *apiPackage) addMethod(m *method) error {
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

func (p *apiPackage) addSubpackage(sub *apiPackage) {
	if p.SubPackages == nil {
		p.SubPackages = map[string]*apiPackage{}
	}
	p.SubPackages[sub.Methods[0].PackageName] = sub
}

func (p *apiPackage) newWriter(outputDir, fileName string) (io.Writer, error) {
	return p.Methods[0].newWriter(outputDir, fileName)
}

func (p *apiPackage) generateAPI(templatesDir string, w io.Writer) error {
	templateFilePath := filepath.Join(templatesDir, "package.tmpl")
	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse template in %q: %s", templateFilePath, err)
	}
	err = t.Execute(w, p)
	if err != nil {
		return fmt.Errorf("failed to execute template in %q: %s", templateFilePath, err)
	}
	return err
}

func (p *apiPackage) generateOption(templatesDir string, w io.Writer) error {
	// TODO: implement options in template
	templateFilePath := filepath.Join(templatesDir, "option.tmpl")
	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse template in %q: %s", templateFilePath, err)
	}
	err = t.Execute(w, p)
	if err != nil {
		return fmt.Errorf("failed to execute template in %q: %s", templateFilePath, err)
	}
	return err
}

func (p *apiPackage) generate(templatesDir, outputDir string) error {
	w, err := p.newWriter(outputDir, "option.go")
	if err != nil {
		return err
	}
	err = p.generateOption(templatesDir, w)
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
	return p.generateAPI(templatesDir, w)
}
