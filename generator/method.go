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

func (m *method) params() (map[string]interface{}, error) {
	p := map[string]interface{}{}
	url, err := methodUrl(m.spec)
	if err != nil {
		return nil, err
	}
	for k, v := range url["parts"].(map[string]interface{}) {
		spec := v.(map[string]interface{})
		if !spec["required"].(bool) {
			p[k] = spec
		}
	}
	for k, v := range url["params"].(map[string]interface{}) {
		p[k] = v
	}
	return p, nil
}
