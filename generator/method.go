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
	// DefaultSpecDir is the dir holding the specs for the APIs and their tests.
	DefaultSpecDir = "rest-api-spec"
	bodyParam      = "body"
)

type specURL struct {
	Path   string            `json:"path"`
	Paths  []string          `json:"paths"`
	Parts  map[string]*param `json:"parts"`
	Params map[string]*param `json:"params"`
}

func (u *specURL) clone() *specURL {
	c := &specURL{
		Path:   u.Path,
		Paths:  []string{},
		Parts:  map[string]*param{},
		Params: map[string]*param{},
	}
	for _, p := range u.Paths {
		c.Paths = append(c.Paths, p)
	}
	for name, p := range u.Parts {
		c.Parts[name] = p.clone()
	}
	for name, p := range u.Params {
		c.Params[name] = p.clone()
	}
	return c
}

type spec struct {
	Documentation string   `json:"documentation"`
	Methods       []string `json:"methods"`
	URL           *specURL `json:"url"`
	Body          *param   `json:"body"`
}

func (s *spec) clone() *spec {
	c := &spec{
		Documentation: s.Documentation,
		Methods:       []string{},
		URL:           s.URL.clone(),
	}
	if s.Body != nil {
		c.Body = s.Body.clone()
	}
	for _, m := range s.Methods {
		c.Methods = append(c.Methods, m)
	}
	return c
}

type method struct {
	rawName           string
	Spec              *spec
	Repo              string
	PackageName       string
	fileName          string
	testFileName      string
	MethodName        string
	TypeName          string
	ReceiverName      string
	ResponseName      string
	HTTPMethod        string
	RequiredURLParts  []*param
	OptionalURLParts  []*param
	RequiredURLParams []*param
	OptionalURLParams []*param
	commonParams      map[string]*param
	allParams         map[string]*param
	ParamsWithValues  []*param
	HTTPCache         map[string]io.ReadCloser
	templates         *template.Template
	offline           bool
}

func newMethod(specDir, specFileName string, commonParams map[string]*param, templates *template.Template,
	offline bool) (*method,
	error) {
	bytes, err := ioutil.ReadFile(filepath.Join(specDir, "api", specFileName))
	if err != nil {
		return nil, err
	}
	m := &method{
		RequiredURLParts:  []*param{},
		OptionalURLParts:  []*param{},
		RequiredURLParams: []*param{},
		OptionalURLParams: []*param{},
		commonParams:      commonParams,
		allParams:         map[string]*param{},
		HTTPCache:         map[string]io.ReadCloser{},
		templates:         templates,
		offline:           offline,
	}
	var spec map[string]*spec
	err = json.Unmarshal(bytes, &spec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s: %s", specFileName, err)
	}
	if len(spec) > 1 {
		return nil, fmt.Errorf("too many methods in %s", specFileName)
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
		if m.Spec.Body.SpecType == "" {
			m.Spec.Body.SpecType = specTypeDict
		}
		if err = m.Spec.Body.resolve(bodyParam, m.templates); err != nil {
			return nil, err
		}
	}
	apiParts := strings.Split(m.rawName, ".")
	m.PackageName = defaultPackage
	switch len(apiParts) {
	case 1:
		m.fileName = apiParts[0]
	case 2:
		m.PackageName = apiParts[0]
		m.fileName = apiParts[1]
	default:
		return nil, fmt.Errorf("unexpected API name format: %s", m.rawName)
	}
	m.Repo = defaultPackageRepo
	if m.PackageName != defaultPackage {
		m.Repo += "/" + m.PackageName
	}
	m.MethodName = snaker.SnakeToCamel(m.fileName)
	m.testFileName += m.fileName + "_test.go"
	m.fileName += ".go"
	m.TypeName = snaker.SnakeToCamel(m.PackageName)
	m.ReceiverName = strings.ToLower(string(m.TypeName[0]))
	m.ResponseName = snaker.SnakeToCamelLower(strings.ToLower(m.MethodName) + "_resp")
	m.HTTPMethod = m.Spec.Methods[0]
	return m, nil
}

func (m *method) normalizeParams(params map[string]*param, resolve bool) error {
	for name, p := range params {
		if resolve {
			err := p.resolve(name, m.templates)
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
		if m.offline {
			return nil
		}
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
					// Skip "The <method> API".
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
					// Overwrite everything, means we we're able un parse a full paragraph.
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
	// Sort required URL parts looking at the URL format.
	for _, partName := range parts {
		// Strip the {}
		if !strings.HasPrefix(partName, "{") || !strings.HasSuffix(partName, "}") {
			continue
		}
		name := strings.Trim(partName, "{}")
		p, ok := m.Spec.URL.Parts[name]
		if p == nil || !ok {
			return fmt.Errorf("cannot find URL part %q (from %q) in %#v", name, m.Spec.URL.Path, m.Spec.URL.Parts)
		}
		m.RequiredURLParts = append(m.RequiredURLParts, p)
	}
	// Sort optional URL parts by name.
	names := []string{}
	for name, p := range m.Spec.URL.Parts {
		if !p.Required {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	for _, name := range names {
		p, ok := m.Spec.URL.Parts[name]
		if p == nil || !ok {
			return fmt.Errorf("cannot find URL part %q (from %q) in %#v", name, m.Spec.URL.Path, m.Spec.URL.Parts)
		}
		m.OptionalURLParts = append(m.OptionalURLParts, p)
	}

	// Handle params, body and common params.
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
	// Sort required params by name.
	for _, name := range names {
		if p, ok := m.Spec.URL.Params[name]; ok && p.Required {
			m.RequiredURLParams = append(m.RequiredURLParams, p)
		}
	}
	// Add body at the end of required params if applicable.
	if m.Spec.Body != nil && m.Spec.Body.Required {
		m.RequiredURLParams = append(m.RequiredURLParams, m.Spec.Body)
	}
	// Sort optional params by name.
	for _, name := range names {
		if p, ok := m.Spec.URL.Params[name]; ok && !p.Required {
			m.OptionalURLParams = append(m.OptionalURLParams, p)
		}
	}
	// Add body after optional params if applicable.
	if m.Spec.Body != nil && !m.Spec.Body.Required {
		m.OptionalURLParams = append(m.OptionalURLParams, m.Spec.Body)
	}
	// Sort common params by name and add them at the end.
	// TODO: we should sort these once and pass a list around.
	for _, name := range names {
		if p, ok := common[name]; ok {
			m.OptionalURLParams = append(m.OptionalURLParams, p)
		}
	}
	return nil
}

func (m *method) newWriter(outputDir, fileName string) (io.Writer, error) {
	if fileName == "" {
		fileName = m.fileName
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

func (m *method) generate(w io.Writer) error {
	if err := m.resolveDocumentation(); err != nil {
		return err
	}
	if err := m.templates.Lookup("method.tmpl").Execute(w, m); err != nil {
		return fmt.Errorf("failed to execute method template: %s", err)
	}
	return nil
}

// clone clones the given method and most of its fields.
func (m *method) clone() (*method, error) {
	c := &method{
		rawName:           m.rawName,
		Spec:              m.Spec.clone(),
		Repo:              m.Repo,
		PackageName:       m.PackageName,
		fileName:          m.fileName,
		testFileName:      m.testFileName,
		MethodName:        m.MethodName,
		TypeName:          m.TypeName,
		ReceiverName:      m.ReceiverName,
		ResponseName:      m.ResponseName,
		HTTPMethod:        m.HTTPMethod,
		RequiredURLParts:  []*param{},
		OptionalURLParts:  []*param{},
		RequiredURLParams: []*param{},
		OptionalURLParams: []*param{},
		commonParams:      map[string]*param{},
		HTTPCache:         map[string]io.ReadCloser{},
	}
	for name, p := range m.commonParams {
		c.commonParams[name] = p.clone()
	}
	for _, p := range m.RequiredURLParts {
		up, ok := c.Spec.URL.Parts[p.rawName]
		if !ok {
			return nil, fmt.Errorf("invalid required part namewhile cloning %s: %s", m.rawName, p.rawName)
		}
		c.RequiredURLParts = append(c.RequiredURLParts, up)
	}
	for _, p := range m.RequiredURLParams {
		if p.rawName == bodyParam {
			c.RequiredURLParams = append(c.RequiredURLParams, c.Spec.Body)
			continue
		}
		up, ok := c.Spec.URL.Params[p.rawName]
		if !ok {
			return nil, fmt.Errorf("invalid required parameter name while cloning %s: %s", m.rawName, p.rawName)
		}
		c.RequiredURLParams = append(c.RequiredURLParams, up)
	}
	for _, p := range m.OptionalURLParts {
		up, ok := c.Spec.URL.Parts[p.rawName]
		if !ok {
			return nil, fmt.Errorf("invalid optional part name while cloning %s: %s", m.rawName, p.rawName)
		}
		c.OptionalURLParts = append(c.OptionalURLParts, up)
	}
	for _, p := range m.OptionalURLParams {
		var up *param
		var ok bool
		if up, ok = c.Spec.URL.Params[p.rawName]; !ok {
			if p.rawName == bodyParam {
				up = m.Spec.Body
			} else if up, ok = c.commonParams[p.rawName]; !ok {
				return nil, fmt.Errorf("invalid optional parameter name while cloning %s: %s", m.rawName, p.rawName)
			}
		}
		c.OptionalURLParams = append(c.OptionalURLParams, up)
	}
	return c, nil
}

// call returns a clone of the method with its parameters set to the specified values.
func (m *method) call(args map[string]interface{}) (*method, error) {
	c, err := m.clone()
	if err != nil {
		return nil, err
	}
	for name, value := range args {
		var p *param
		var ok bool
		if name == bodyParam && c.Spec.Body != nil {
			p = c.Spec.Body
		} else if p, ok = c.Spec.URL.Parts[name]; !ok {
			if p, ok = c.Spec.URL.Params[name]; !ok {
				if p, ok = c.commonParams[name]; !ok {
					return nil, fmt.Errorf("invalid parameter name: %s", name)
				}
			}
		}
		p.Value = value
	}
	c.ParamsWithValues = []*param{}
	for _, p := range c.RequiredURLParts {
		if p.Value != nil {
			c.ParamsWithValues = append(c.ParamsWithValues, p)
		}
	}
	for _, p := range c.RequiredURLParams {
		if p.Value != nil {
			c.ParamsWithValues = append(c.ParamsWithValues, p)
		}
	}
	for _, p := range c.OptionalURLParts {
		if p.Value != nil {
			c.ParamsWithValues = append(c.ParamsWithValues, p)
		}
	}
	for _, p := range c.OptionalURLParams {
		if p.Value != nil {
			c.ParamsWithValues = append(c.ParamsWithValues, p)
		}
	}
	return c, nil
}
