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
	"sort"

	"github.com/serenize/snaker"
)

const (
	rootAPITemplateFile = "root_api.tmpl"
)

type api struct {
	PackageRepo string
	PackageName string
	APIName     string
	TypeName    string
}

type apis struct {
	Apis     []*api
	filePath string
}

func newAPI(method *method) *api {
	// Refuse to generate accessors for the top level package
	if method.PackageName == defaultPackage {
		return nil
	}
	name := snaker.SnakeToCamel(method.PackageName)
	return &api{
		PackageRepo: method.PackageRepo,
		PackageName: method.PackageName,
		TypeName:    method.TypeName,
		APIName:     name,
	}
}

func newApis(methods []*method, outputDir string) *apis {
	aMap := map[string]*api{}
	var repos []string
	for _, method := range methods {
		if _, ok := aMap[method.PackageRepo]; !ok {
			a := newAPI(method)
			if a == nil {
				continue
			}
			aMap[method.PackageRepo] = a
			repos = append(repos, method.PackageRepo)
		}
	}
	sort.Strings(repos)
	var a []*api
	for _, r := range repos {
		a = append(a, aMap[r])
	}
	return &apis{Apis: a, filePath: filepath.Join(outputDir, "api.go")}
}

func (a *apis) executeTemplate(templateFilePath, outputFilePath string) error {
	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("Failed to parse template in %q: %s", templateFilePath, err)
	}
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	err = t.Execute(outputFile, a)
	if err != nil {
		return fmt.Errorf("Failed to execute template in %q: %s", templateFilePath, err)
	}
	return err
}

func (a *apis) generate(templatesDir string) error {
	return a.executeTemplate(filepath.Join(templatesDir, rootAPITemplateFile), a.filePath)
}
