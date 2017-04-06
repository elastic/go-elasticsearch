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
)

const (
	specExt          = ".json"
	templatesDir     = "generator/templates"
	templateFile     = "method.tmpl"
	defaultNamespace = "client"
)

type method struct {
	TypeVar    string
	TypeName   string
	MethodName string
	DocURL     string
}

func executeTemplate(spec map[string]interface{}, templatesDir, outputDir string) error {
	if len(spec) > 1 {
		return fmt.Errorf("Each API spec is expected to contain a single API (found %d)", len(spec))
	}
	var api string
	for k := range spec {
		api = k
	}
	typeName := defaultNamespace
	apiParts := strings.Split(api, ".")
	switch len(apiParts) {
	case 1:
		api = apiParts[0]
	case 2:
		typeName = apiParts[0]
		api = apiParts[1]
	default:
		return fmt.Errorf("Unexpected API format: %s", api)
	}
	apiSpec, ok := spec[api].(map[string]interface{})
	if !ok {
		return fmt.Errorf("Unexpected type in spec: %T (expected map[string]interface{})", spec[api])
	}
	m := &method{
		TypeVar:    string(typeName[0]),
		TypeName:   typeName,
		MethodName: api,
		DocURL:     apiSpec["documentation"].(string),
	}
	t, err := template.ParseFiles(filepath.Join(templatesDir, templateFile))
	if err != nil {
		return fmt.Errorf("Failed to parse template in %q: %s", templateFile, err)
	}
	goFile, err := os.Create(filepath.Join(outputDir, api) + ".go")
	if err != nil {
		return err
	}
	defer goFile.Close()
	err = t.Execute(goFile, m)
	if err != nil {
		return fmt.Errorf("Failed to execute template in %q: %s", templateFile, err)
	}
	return err
}
