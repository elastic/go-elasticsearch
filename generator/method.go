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
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/serenize/snaker"
)

const (
	// SpecExt is the extension for spec files for the REST API
	SpecExt            = ".json"
	templatesDir       = "generator/templates"
	methodTemplateFile = "method.tmpl"
	clientTemplateFile = "client.tmpl"
	repo               = "github.com/elastic/elasticsearch-go"
	defaultPackage     = "client"
)

type method struct {
	PackageRepo    string
	PackageName    string
	MethodName     string
	DocURL         string
	filePath       string
	clientFilePath string
}

func newMethod(spec map[string]interface{}, outputDir string) (*method, error) {
	if len(spec) > 1 {
		return nil, fmt.Errorf("Each API spec is expected to contain a single API (found %d)", len(spec))
	}
	var api string
	for k := range spec {
		api = k
	}
	if api == "" {
		return nil, fmt.Errorf("Unable to find API in %s", spec)
	}
	apiSpec, ok := spec[api].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Unexpected type in spec: %T (expected map[string]interface{})", spec[api])
	}
	packageName := defaultPackage
	packageRepo := repo + "/" + defaultPackage
	apiParts := strings.Split(api, ".")
	var methodName, fileName string
	switch len(apiParts) {
	case 1:
		fileName = apiParts[0]
	case 2:
		packageName = apiParts[0]
		packageRepo += "/" + packageName
		fileName = apiParts[1]
	default:
		return nil, fmt.Errorf("Unexpected API format: %s", api)
	}
	methodName = snaker.SnakeToCamel(fileName)
	m := &method{
		PackageRepo: packageRepo,
		PackageName: packageName,
		MethodName:  methodName,
		DocURL:      apiSpec["documentation"].(string),
	}
	goFileDir := outputDir
	if packageName != defaultPackage {
		goFileDir = filepath.Join(goFileDir, packageName)
		m.clientFilePath = filepath.Join(goFileDir, "client.go")
	}
	m.filePath = filepath.Join(goFileDir, fileName) + ".go"
	return m, nil
}

func (m *method) generateClient() error {
	if m.clientFilePath == "" {
		return nil
	}
	err := os.Mkdir(filepath.Dir(m.filePath), 0755)
	// If the dir exists assume that so does the client
	if err != nil {
		return nil
	}
	return m.executeTemplate(filepath.Join(templatesDir, clientTemplateFile), m.clientFilePath)
}

func (m *method) executeTemplate(templateFilePath, outputFilePath string) error {
	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("Failed to parse template in %q: %s", templateFilePath, err)
	}
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	err = t.Execute(outputFile, m)
	if err != nil {
		return fmt.Errorf("Failed to execute template in %q: %s", templateFilePath, err)
	}
	return err
}

func (m *method) generate(templatesDir string) error {
	if err := m.generateClient(); err != nil {
		return err
	}
	return m.executeTemplate(filepath.Join(templatesDir, methodTemplateFile), m.filePath)
}

func clean(spec map[string]interface{}, outputDir string) error {
	m, err := newMethod(spec, outputDir)
	if err != nil {
		return err
	}
	os.Remove(m.filePath)
	if m.clientFilePath != "" {
		os.RemoveAll(filepath.Dir(m.clientFilePath))
	}
	return nil
}
