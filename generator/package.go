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
	"html/template"
	"os"
	"path/filepath"
)

type apiPackage struct {
	Method      *method
	Repo        string
	SubPackages map[string]*apiPackage
}

func newAPIPackage(m *method) *apiPackage {
	return &apiPackage{
		Method: m,
	}
}

func (p *apiPackage) addSubpackage(sub *apiPackage) {
	if p.SubPackages == nil {
		p.SubPackages = map[string]*apiPackage{}
	}
	p.SubPackages[sub.Method.PackageName] = sub
}

func (p *apiPackage) generate(templatesDir, outputDir string) error {
	templateFilePath := filepath.Join(templatesDir, "api.tmpl")
	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("Failed to parse template in %q: %s", templateFilePath, err)
	}
	goFileDir := mkOutputDir(outputDir, p.Method.PackageName)
	goFilePath := filepath.Join(goFileDir, p.Method.PackageName+".go")
	goFile, err := os.Create(goFilePath)
	if err != nil {
		return err
	}
	defer goFile.Close()
	err = t.Execute(goFile, p)
	if err != nil {
		return fmt.Errorf("Failed to execute template in %q: %s", templateFilePath, err)
	}
	return err
}
