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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/serenize/snaker"
)

const (
	// DefaultSpecDir is the dir holding the JSON specs for the APIs
	DefaultSpecDir = "spec/elasticsearch/rest-api-spec/src/main/resources/rest-api-spec/api"
)

type url struct {
	Parts  map[string]*param `json:"parts"`
	Params map[string]*param `json:"params"`
}

type spec struct {
	Documentation string   `json:"documentation"`
	Methods       []string `json:"methods"`
	URL           url      `json:"url"`
	Body          param    `json:"body"`
}

type method struct {
	Name         string
	Spec         spec
	Repo         string
	PackageName  string
	FileName     string
	MethodName   string
	TypeName     string
	ReceiverName string
	HTTPMethod   string
}

func newMethod(specFilePath string) (*method, error) {
	bytes, err := ioutil.ReadFile(specFilePath)
	if err != nil {
		return nil, err
	}
	var m method
	var spec map[string]spec
	err = json.Unmarshal(bytes, &spec)
	if err != nil {
		return nil, err
	}
	if len(spec) > 1 {
		return nil, fmt.Errorf("Too many methods in %s", specFilePath)
	}
	for k := range spec {
		m.Name = k
		m.Spec = spec[k]
		break
	}
	m.normalizeParams(&m.Spec.URL.Parts)
	m.normalizeParams(&m.Spec.URL.Params)
	apiParts := strings.Split(m.Name, ".")
	m.PackageName = defaultPackage
	switch len(apiParts) {
	case 1:
		m.FileName = apiParts[0]
	case 2:
		m.PackageName = apiParts[0]
		m.FileName = apiParts[1]
	default:
		return nil, fmt.Errorf("Unexpected API name format: %s", m.Name)
	}
	m.Repo = defaultPackageRepo
	if m.PackageName != defaultPackage {
		m.Repo += "/" + m.PackageName
	}
	m.MethodName = snaker.SnakeToCamel(m.FileName)
	m.FileName += ".go"
	m.TypeName = snaker.SnakeToCamel(m.PackageName)
	m.ReceiverName = strings.ToLower(string(m.TypeName[0]))
	m.HTTPMethod = m.Spec.Methods[0]
	return &m, nil
}

func (m *method) normalizeParams(params *map[string]*param) {
	for name, p := range *params {
		p.normalize(name)
		if p.Name != name {
			delete(*params, name)
			(*params)[p.Name] = p
		}
	}
}

func (m *method) newWriter(outputDir, fileName string) (io.Writer, error) {
	if fileName == "" {
		fileName = m.FileName
	}
	goFileDir := outputDir
	if m.PackageName != defaultPackage {
		goFileDir = filepath.Join(goFileDir, defaultPackage)
	}
	goFileDir = filepath.Join(goFileDir, m.PackageName)
	os.MkdirAll(goFileDir, 0755)
	goFilePath := filepath.Join(goFileDir, fileName)
	goFile, err := os.Create(goFilePath)
	if err != nil {
		return nil, err
	}
	return goFile, nil
}

func (m *method) generate(templatesDir string, w io.Writer) error {
	templateFilePath := filepath.Join(templatesDir, "method.tmpl")
	t, err := template.New("method.tmpl").ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("Failed to parse template in %q: %s", templateFilePath, err)
	}
	err = t.Execute(w, m)
	if err != nil {
		return fmt.Errorf("Failed to execute template in %q: %s", templateFilePath, err)
	}
	return nil
}
