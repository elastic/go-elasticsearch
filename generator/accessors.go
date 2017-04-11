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

	"github.com/serenize/snaker"
)

const (
	accessorsTemplateFile = "accessors.tmpl"
)

type accessor struct {
	PackageRepo string
	Name        string
}

type accessors struct {
	Accessors map[string]*accessor
	filePath  string
}

func newAccessor(method *method) *accessor {
	return &accessor{
		PackageRepo: method.PackageRepo,
		Name:        snaker.SnakeToCamel(method.PackageName),
	}
}

func newAccessors(methods []*method, outputDir string) *accessors {
	a := map[string]*accessor{}
	for _, method := range methods {
		if _, ok := a[method.PackageName]; !ok {
			a[method.PackageName] = newAccessor(method)
		}
	}
	return &accessors{Accessors: a, filePath: filepath.Join(outputDir, "accessors.go")}
}

func (a accessors) executeTemplate(templateFilePath, outputFilePath string) error {
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

func (a *accessors) generate(templatesDir string) error {
	return a.executeTemplate(filepath.Join(templatesDir, accessorsTemplateFile), a.filePath)
}

func (a *accessors) clean() error {
	os.Remove(a.filePath)
	return nil
}
