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
	transportPackageRepo = "github.com/elastic/elasticsearch-go/client/http"
	transportFieldName   = "client"
)

var (
	toCamel = snaker.SnakeToCamel
)

type method struct {
	spec        map[string]interface{}
	packageName string
}

func newMethod(spec map[string]interface{}) (*method, error) {
	if len(spec) > 1 {
		return nil, fmt.Errorf("Unexpected format in API: %s", spec)
	}
	var methodName string
	for k := range spec {
		methodName = k
		break
	}
	return &method{
		spec: map[string]interface{}{
			"method":               methodName,
			"spec":                 spec[methodName],
			"transportPackageRepo": transportPackageRepo,
			"transportFieldName":   transportFieldName,
		},
	}, nil
}

func packageAndMethod(spec map[string]interface{}) (string, string, error) {
	api, ok := spec["method"].(string)
	if !ok {
		return "", "", fmt.Errorf("Unexpected type for method: %T (expected string)", spec["method"])
	}
	apiParts := strings.Split(api, ".")
	packageName := defaultPackage
	var methodName string
	switch len(apiParts) {
	case 1:
		methodName = apiParts[0]
	case 2:
		packageName = apiParts[0]
		methodName = apiParts[1]
	default:
		return "", "", fmt.Errorf("Unexpected API name format: %s", api)
	}
	return packageName, methodName, nil
}

func packageName(spec map[string]interface{}) (string, error) {
	p, _, err := packageAndMethod(spec)
	return p, err
}

func methodName(spec map[string]interface{}) (string, error) {
	_, m, err := packageAndMethod(spec)
	return toCamel(m), err
}

func typeName(spec map[string]interface{}) (string, error) {
	p, _, err := packageAndMethod(spec)
	if err != nil {
		return "", err
	}
	return toCamel(p), nil
}

func receiverName(spec map[string]interface{}) (string, error) {
	t, err := typeName(spec)
	if err != nil {
		return "", err
	}
	return strings.ToLower(string(t[0])), nil
}

func httpMethod(spec map[string]interface{}) (string, error) {
	m, ok := spec["spec"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("Unexpected type for method: %T (expected map[string]interface {})", spec["spec"])
	}
	methods, ok := m["methods"].([]interface{})
	if !ok {
		return "", fmt.Errorf("Unexpected type for methods: %T (expected []interface{})", m["methods"])
	}
	for _, m := range methods {
		methodName, ok := m.(string)
		if !ok {
			return "", fmt.Errorf("Unexpected type for method: %T (expected string)", m)
		}
		return methodName, nil
	}
	return "", fmt.Errorf("No HTTP methods in %q", spec)
}

func mkOutputDir(outputRootDir, packageName string) string {
	goFileDir := outputRootDir
	if packageName != defaultPackage {
		goFileDir = filepath.Join(goFileDir, defaultPackage)
	}
	goFileDir = filepath.Join(goFileDir, packageName)
	os.MkdirAll(goFileDir, 0755)
	return goFileDir
}

func (m *method) generate(templatesDir, outputDir string) error {
	templateFilePath := filepath.Join(templatesDir, "method.tmpl")
	funcMap := template.FuncMap{
		"packageName":  packageName,
		"methodName":   methodName,
		"typeName":     typeName,
		"receiverName": receiverName,
		"httpMethod":   httpMethod,
	}
	t, err := template.New("method.tmpl").Funcs(funcMap).ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("Failed to parse template in %q: %s", templateFilePath, err)
	}
	m.packageName, err = packageName(m.spec)
	if err != nil {
		return err
	}
	_, methodNameRaw, err := packageAndMethod(m.spec)
	if err != nil {
		return err
	}
	goFileDir := mkOutputDir(outputDir, m.packageName)
	goFilePath := filepath.Join(goFileDir, methodNameRaw) + ".go"
	goFile, err := os.Create(goFilePath)
	if err != nil {
		return err
	}
	defer goFile.Close()
	err = t.Execute(goFile, m.spec)
	if err != nil {
		return fmt.Errorf("Failed to execute template in %q: %s", templateFilePath, err)
	}
	return err
}
