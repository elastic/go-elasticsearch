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
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/net/html"

	"github.com/golang/glog"
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
	Body          *param   `json:"body"`
}

type method struct {
	Name           string
	Spec           spec
	Repo           string
	PackageName    string
	FileName       string
	MethodName     string
	TypeName       string
	ReceiverName   string
	HTTPMethod     string
	RequiredParams []*param
	OptionalParams []*param
	ParamNames     map[string]struct{}
	HTTPCache      map[string]io.ReadCloser
}

func newMethod(specFilePath string, commonParams map[string]*param) (*method, error) {
	bytes, err := ioutil.ReadFile(specFilePath)
	if err != nil {
		return nil, err
	}
	m := &method{
		RequiredParams: []*param{},
		OptionalParams: []*param{},
		ParamNames:     map[string]struct{}{},
		HTTPCache:      map[string]io.ReadCloser{},
	}
	var spec map[string]spec
	err = json.Unmarshal(bytes, &spec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %s", specFilePath, err)
	}
	if len(spec) > 1 {
		return nil, fmt.Errorf("too many methods in %s", specFilePath)
	}
	for k := range spec {
		m.Name = k
		m.Spec = spec[k]
		break
	}
	if err = m.normalizeParams(m.Spec.URL.Parts, true); err != nil {
		return nil, err
	}
	if err = m.normalizeParams(m.Spec.URL.Params, true); err != nil {
		return nil, err
	}
	if err = m.normalizeParams(commonParams, false); err != nil {
		return nil, err
	}
	m.sortParams(commonParams)
	if m.Spec.Body != nil {
		if err = m.Spec.Body.resolve("body"); err != nil {
			return nil, err
		}
	}
	apiParts := strings.Split(m.Name, ".")
	m.PackageName = defaultPackage
	switch len(apiParts) {
	case 1:
		m.FileName = apiParts[0]
	case 2:
		m.PackageName = apiParts[0]
		m.FileName = apiParts[1]
	default:
		return nil, fmt.Errorf("unexpected API name format: %s", m.Name)
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
	return m, nil
}

func (m *method) normalizeParams(params map[string]*param, resolve bool) error {
	for name, p := range params {
		if resolve {
			err := p.resolve(name)
			if err != nil {
				return fmt.Errorf("failed to normalize params in %q: %s", m.Name, err)
			}
		}
		if _, ok := m.ParamNames[p.Name]; ok {
			p.addSuffix("Param")
			if _, ok := m.ParamNames[p.Name]; ok {
				return fmt.Errorf("param %q specified more than twice in %s", name, m.Name)
			}
		}
		m.ParamNames[p.Name] = struct{}{}
	}
	return nil
}

func (m *method) resolveDocumentation() error {
	url := m.Spec.Documentation
	if url == "http://www.elastic.co/guide/" {
		m.Spec.Documentation = " - see " + url + "."
		return nil
	}
	body, ok := m.HTTPCache[url]
	if !ok {
		glog.Infof("fetching %s", url)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		body = resp.Body
	}
	defer body.Close()
	tokenizer := html.NewTokenizer(body)
	for {
		token := tokenizer.Next()
		if token == html.StartTagToken {
			t := tokenizer.Token()
			if t.Data == "p" {
				tokenizer.Next()
				text := strings.Replace(string(tokenizer.Text()), "\n", " ", -1)
				m.Spec.Documentation = ""
				var offset int
				if strings.HasPrefix(text, "The "+strings.ToLower(m.MethodName)+" API") {
					offset = 3
				} else {
					m.Spec.Documentation = " -"
				}
				period := -1
				for i, word := range strings.Split(text, " ") {
					// Skip "The <method> API"
					if i < offset {
						continue
					}
					word = formatToken(i-offset, word)
					m.Spec.Documentation += " " + word
					if strings.HasSuffix(word, ".") {
						period = i
						break
					}
				}
				if period > 0 {
					m.Spec.Documentation += " See "
				} else {
					// Overwrite everything, means we we're able un parse a full paragraph
					m.Spec.Documentation = " - see "
				}
				m.Spec.Documentation += url + " for more info."
				return nil
			}
		}
	}
}

func (m *method) sortParams(common map[string]*param) {
	requiredNames := []string{}
	optionalNames := []string{}
	allParams := map[string]*param{}
	for _, p := range m.Spec.URL.Parts {
		if p.Required {
			requiredNames = append(requiredNames, p.Name)
		} else {
			optionalNames = append(optionalNames, p.Name)
		}
		allParams[p.Name] = p
	}
	for _, p := range m.Spec.URL.Params {
		if p.Required {
			requiredNames = append(requiredNames, p.Name)
		} else {
			optionalNames = append(optionalNames, p.Name)
		}
		allParams[p.Name] = p
	}
	for _, p := range common {
		if p.Required {
			requiredNames = append(requiredNames, p.Name)
		} else {
			optionalNames = append(optionalNames, p.Name)
		}
		allParams[p.Name] = p
	}
	sort.Strings(requiredNames)
	for _, name := range requiredNames {
		m.RequiredParams = append(m.RequiredParams, allParams[name])
	}
	sort.Strings(optionalNames)
	for _, name := range optionalNames {
		m.OptionalParams = append(m.OptionalParams, allParams[name])
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
	if err := m.resolveDocumentation(); err != nil {
		return err
	}
	templateFilePath := filepath.Join(templatesDir, "method.tmpl")
	t, err := template.New("method.tmpl").ParseFiles(templateFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse template in %q: %s", templateFilePath, err)
	}
	err = t.Execute(w, m)
	if err != nil {
		return fmt.Errorf("failed to execute template in %q: %s", templateFilePath, err)
	}
	return nil
}
