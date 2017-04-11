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
	apiTemplateFile    = "api.tmpl"
	repo               = "github.com/elastic/elasticsearch-go"
	defaultPackage     = "api"
)

type method struct {
	PackageRepo  string
	PackageName  string
	ReceiverName string
	TypeName     string
	MethodName   string
	DocURL       string
	filePath     string
	apiFilePath  string
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
	receiverName := "a"
	typeName := "Api"
	switch len(apiParts) {
	case 1:
		fileName = apiParts[0]
	case 2:
		packageName = apiParts[0]
		packageRepo += "/" + packageName
		fileName = apiParts[1]
		receiverName = string(packageName[0])
		typeName = snaker.SnakeToCamel(packageName)
	default:
		return nil, fmt.Errorf("Unexpected API format: %s", api)
	}
	methodName = snaker.SnakeToCamel(fileName)
	m := &method{
		PackageRepo:  packageRepo,
		PackageName:  packageName,
		ReceiverName: receiverName,
		TypeName:     typeName,
		MethodName:   methodName,
		DocURL:       apiSpec["documentation"].(string),
	}
	goFileDir := outputDir
	if packageName != defaultPackage {
		goFileDir = filepath.Join(goFileDir, packageName)
	}
	m.apiFilePath = filepath.Join(goFileDir, packageName+".go")
	m.filePath = filepath.Join(goFileDir, fileName) + ".go"
	return m, nil
}

func (m *method) generateClient() error {
	return m.executeTemplate(filepath.Join(templatesDir, apiTemplateFile), m.apiFilePath)
}

func (m *method) executeTemplate(templateFilePath, outputFilePath string) error {
	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("Failed to parse template in %q: %s", templateFilePath, err)
	}
	os.MkdirAll(filepath.Dir(outputFilePath), 0755)
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
	return m.executeTemplate(filepath.Join(templatesDir, methodTemplateFile), m.filePath)
}
