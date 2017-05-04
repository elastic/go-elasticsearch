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
	"net/url"
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
	// DefaultSpecDir is the dir holding the specs for the APIs and their tests
	DefaultSpecDir = "rest-api-spec"
)

type specURL struct {
	Path   string            `json:"path"`
	Paths  []string          `json:"paths"`
	Parts  map[string]*param `json:"parts"`
	Params map[string]*param `json:"params"`
}

type spec struct {
	Documentation string   `json:"documentation"`
	Methods       []string `json:"methods"`
	URL           specURL  `json:"url"`
	Body          *param   `json:"body"`
}

type method struct {
	rawName           string
	Spec              spec
	Repo              string
	PackageName       string
	FileName          string
	MethodName        string
	TypeName          string
	ReceiverName      string
	HTTPMethod        string
	RequiredURLParts  []*param
	OptionalURLParts  []*param
	RequiredURLParams []*param
	OptionalURLParams []*param
	allParams         map[string]*param
	HTTPCache         map[string]io.ReadCloser
}

func newMethod(specFilePath string, commonParams map[string]*param) (*method, error) {
	bytes, err := ioutil.ReadFile(specFilePath)
	if err != nil {
		return nil, err
	}
	m := &method{
		RequiredURLParts:  []*param{},
		OptionalURLParts:  []*param{},
		RequiredURLParams: []*param{},
		OptionalURLParams: []*param{},
		allParams:         map[string]*param{},
		HTTPCache:         map[string]io.ReadCloser{},
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
		m.rawName = k
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
	if err = m.sortParams(commonParams); err != nil {
		return nil, err
	}
	if m.Spec.Body != nil {
		if err = m.Spec.Body.resolve("body"); err != nil {
			return nil, err
		}
		m.Spec.Body.Type = "map[string]interface{}"
	}
	apiParts := strings.Split(m.rawName, ".")
	m.PackageName = defaultPackage
	switch len(apiParts) {
	case 1:
		m.FileName = apiParts[0]
	case 2:
		m.PackageName = apiParts[0]
		m.FileName = apiParts[1]
	default:
		return nil, fmt.Errorf("unexpected API name format: %s", m.rawName)
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
				return fmt.Errorf("failed to normalize params in %q: %s", m.rawName, err)
			}
			if _, ok := m.allParams[p.Name]; ok {
				p.addSuffix("Param")
				if _, ok := m.allParams[p.Name]; ok {
					return fmt.Errorf("param %q seen more than twice", name)
				}
			}
			m.allParams[p.Name] = p
		}
	}
	return nil
}

func (m *method) resolveDocumentation() error {
	docURL := m.Spec.Documentation
	u, err := url.Parse(docURL)
	if err != nil {
		return err
	}
	if u.Path == "/guide/" {
		m.Spec.Documentation = " - see " + docURL + " for more info."
		return nil
	}
	body, ok := m.HTTPCache[docURL]
	if !ok {
		glog.Infof("fetching %s", docURL)
		resp, err := http.Get(docURL)
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
				m.Spec.Documentation += docURL + " for more info."
				return nil
			}
		}
	}
}

func (m *method) sortParams(common map[string]*param) error {
	// Handle URL parts
	parts := strings.Split(m.Spec.URL.Path, "/")
	if len(parts) < 2 {
		return fmt.Errorf("unexpected URL format: %s", m.Spec.URL.Path)
	}
	// Sort required URL parts looking at the URL format
	for _, partName := range parts {
		// Strip the {}
		if !strings.HasPrefix(partName, "{") || !strings.HasSuffix(partName, "}") {
			continue
		}
		name := strings.Trim(partName, "{}")
		p, ok := m.Spec.URL.Parts[name]
		if !ok {
			return fmt.Errorf("cannot find URL part %q (from %q) in %#v", name, m.Spec.URL.Path, m.Spec.URL.Parts)
		}
		m.RequiredURLParts = append(m.RequiredURLParts, p)
	}
	// Sort optional URL parts by name
	names := []string{}
	for name, p := range m.Spec.URL.Parts {
		if !p.Required {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	for _, name := range names {
		m.OptionalURLParts = append(m.OptionalURLParts, m.Spec.URL.Parts[name])
	}

	// Handle params, body and common params
	names = []string{}
	for name := range m.Spec.URL.Params {
		names = append(names, name)
	}
	for name, p := range common {
		if p.Required {
			return fmt.Errorf("%q is required but common params are not supposed to be", p.Name)
		}
		names = append(names, name)
	}
	sort.Strings(names)
	// Sort required params by name
	for _, name := range names {
		if p, ok := m.Spec.URL.Params[name]; ok && p.Required {
			m.RequiredURLParams = append(m.RequiredURLParams, p)
		}
	}
	// Add body at the end of required params if applicable
	if m.Spec.Body != nil && m.Spec.Body.Required {
		m.RequiredURLParams = append(m.RequiredURLParams, m.Spec.Body)
	}
	// Sort optional params by name
	for _, name := range names {
		if p, ok := m.Spec.URL.Params[name]; ok && !p.Required {
			m.OptionalURLParams = append(m.OptionalURLParams, p)
		}
	}
	// Add body after optional params if applicable
	if m.Spec.Body != nil && !m.Spec.Body.Required {
		m.OptionalURLParams = append(m.OptionalURLParams, m.Spec.Body)
	}
	// Sort common params by name and add them at the end
	// TODO: we should sort these once and pass a list around
	for _, name := range names {
		if p, ok := common[name]; ok {
			m.OptionalURLParams = append(m.OptionalURLParams, p)
		}
	}
	return nil
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
