// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package gensource

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"unicode"

	"gopkg.in/yaml.v2"

	"github.com/elastic/go-elasticsearch/v8/internal/build/utils"
)

var (
	apiDescriptions map[interface{}]interface{}
)

func init() {
	err := yaml.NewDecoder(strings.NewReader(apiDescriptionsYAML)).Decode(&apiDescriptions)
	if err != nil {
		panic(fmt.Sprintf("ERROR: %v", err))
	}
}

// NewEndpoint creates a new API endpoint.
//
func NewEndpoint(f io.Reader) (*Endpoint, error) {
	var endpoint Endpoint
	var spec map[string]Endpoint

	var xpackNamespaces = []string{
		"async_search",
		"autoscaling",
		"cat.ml_data_frame_analytics",
		"cat.ml_datafeeds",
		"cat.ml_jobs",
		"cat.ml_trained_models",
		"cat.transforms",
		"ccr",
		"close_point_in_time",
		"data_frame_transform_deprecated",
		"enrich",
		"eql",
		"graph",
		"ilm",
		"indices.create_data_stream",
		"indices.data_streams_stats",
		"indices.delete_data_stream",
		"indices.freeze",
		"indices.get_data_stream",
		"indices.migrate_to_data_stream",
		"indices.promote_data_stream",
		"indices.reload_search_analyzers",
		"indices.unfreeze",
		"license",
		"logstash",
		"migration",
		"ml",
		"monitoring",
		"open_point_in_time",
		"rollup",
		"searchable_snapshots",
		"security",
		"slm",
		"sql",
		"ssl",
		"text_structure",
		"transform",
		"watcher",
		"xpack.info",
		"xpack.usage",
	}

	if err := json.NewDecoder(f).Decode(&spec); err != nil {
		return nil, err
	}

	for name, e := range spec {
		endpoint = e
		endpoint.Name = name
		endpoint.URL.Params = endpoint.Params
	}

	if fpath, ok := f.(*os.File); ok {
		if strings.Contains(fpath.Name(), "x-pack") {
			endpoint.Type = "xpack"
		}
		for _, namespace := range xpackNamespaces {
			if strings.Contains(fpath.Name(), namespace) {
				endpoint.Type = "xpack"
			}
		}
	}
	if endpoint.Type == "" {
		endpoint.Type = "core"
	}

	endpoint.URL.AllParts = make(map[string]*Part)

	for _, path := range endpoint.URL.Paths {
		for partName, part := range path.Parts {
			part.Endpoint = &endpoint
			part.Name = partName
		}
	}

	var params [][]string
	for _, path := range endpoint.URL.Paths {
		var parts []string
		for partName, part := range path.Parts {
			// Skip type only for selected APIs at this point
			if part.Name == "type" && part.Deprecated {
				if endpoint.Name == "bulk" ||
					endpoint.Name == "create" ||
					endpoint.Name == "delete" ||
					endpoint.Name == "get" ||
					endpoint.Name == "index" ||
					endpoint.Name == "update" {
					continue
				}
			}
			parts = append(parts, partName)
		}
		params = append(params, parts)
	}

	var paramsCounter = make(map[string]int)
	for _, pp := range params {
		for _, p := range pp {
			paramsCounter[p] += 1
		}
	}

	// Update the required field of path parts
	for _, path := range endpoint.URL.Paths {
		for partName, part := range path.Parts {
			if p, ok := paramsCounter[partName]; ok {
				if p == len(params) {
					part.Required = true
				}
			}
		}
	}

	for paramName, p := range endpoint.URL.Params {
		p.Endpoint = &endpoint
		p.Name = paramName
	}

	// Update the AllParts field
	var partSeen = make(map[string]bool)
	for _, path := range endpoint.URL.Paths {
		for partName, part := range path.Parts {
			if !partSeen[partName] {
				endpoint.URL.AllParts[partName] = part
			}
			partSeen[partName] = true
		}
	}

	// Fix up the documentation property (X-Pack spec related); TODO: PR
	if !strings.HasPrefix(endpoint.Documentation.URL, "http") {
		if endpoint.Type == "xpack" && endpoint.Documentation.Description == "" && endpoint.Documentation.URL != "" {
			endpoint.Documentation.Description = endpoint.Documentation.URL

			if endpoint.Documentation.URL == "TODO" {
				endpoint.Documentation.URL = ""
			}
			if endpoint.Documentation.Description == "TODO" {
				endpoint.Documentation.Description = ""
			}
		}
		endpoint.Documentation.URL = ""
	}

	var partNames []string
	for param := range paramsCounter {
		partNames = append(partNames, param)
	}
	sort.Slice(partNames, func(i, j int) bool {
		return strings.Replace(partNames[i], "_", "", 1) < strings.Replace(partNames[j], "_", "", 1)
	})
	endpoint.URL.PartNamesSorted = partNames

	var paramNames []string
	for _, param := range endpoint.URL.Params {
		paramNames = append(paramNames, param.Name)
	}
	sort.Slice(paramNames, func(i, j int) bool {
		return strings.Replace(paramNames[i], "_", "", 1) < strings.Replace(paramNames[j], "_", "", 1)
	})
	endpoint.URL.ParamNamesSorted = paramNames

	// if info, ok := apiDescriptions[endpoint.Name]; ok {
	// 	if desc, ok := info.(map[interface{}]interface{})["description"].(string); ok {
	// 		endpoint.Documentation.Description = desc
	// 	}
	// }

	return &endpoint, nil
}

// Endpoint represents an API endpoint.
//
type Endpoint struct {
	Name string `json:"-"`
	Type string `json:"-"`

	Documentation struct {
		URL         string `json:"url"`
		Description string `json:"description"`
	} `json:"documentation"`

	Stability string `json:"stability"`

	URL    *URL              `json:"url"`
	Params map[string]*Param `json:"params"`
	Body   *Body             `json:"body"`
}

// URL represents API endpoint URL.
//
type URL struct {
	Endpoint *Endpoint `json:"-"`

	Paths           []Path `json:"paths"`
	DeprecatedPaths []struct {
		Path        string `json:"path"`
		Version     string `json:"version"`
		Description string `json:"description"`
	} `json:"deprecated_paths"`
	Params map[string]*Param `json:"params"`

	AllParts         map[string]*Part
	PartNamesSorted  []string
	ParamNamesSorted []string
}

// Path represents URL path
type Path struct {
	Path       string           `json:"path"`
	Methods    []string         `json:"methods"`
	Parts      map[string]*Part `json:"parts"`
	Deprecated struct {
		Version     string `json:"version"`
		Description string `json:"description"`
	}
}

// Part represents part of the API endpoint URL.
//
type Part struct {
	Endpoint *Endpoint `json:"-"`

	Name    string      `json:"-"`
	Default interface{} `json:"-"`

	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`

	Deprecated bool `json:"deprecated"`
}

// Param represents API endpoint parameter.
//
type Param struct {
	Endpoint *Endpoint `json:"-"`

	Name string `json:"-"`

	Type        string      `json:"type"`
	Description string      `json:"description"`
	Options     []string    `json:"options"`
	Default     interface{} `json:"default"`
	Required    bool        `json:"required"`
}

// Body represents API endpoint body.
//
type Body struct {
	Endpoint *Endpoint `json:"-"`

	Description string `json:"description"`
	Required    bool   `json:"required"`
	ContentType string `json:"serialize"`
}

// MethodArgument represents a method argument for API endpoint.
//
type MethodArgument struct {
	Endpoint *Endpoint

	Name        string
	Type        string
	Description string
	Options     []string
	Default     interface{}
	Required    bool
}

// Namespace returns the API endpoint namespace.
//
func (e *Endpoint) Namespace() string {
	ep := strings.Split(e.Name, ".")
	return utils.NameToGo(ep[0])
}

// MethodName returns the API endpoint method name.
//
func (e *Endpoint) MethodName() string {
	ep := strings.Split(e.Name, ".")
	ep = append(ep[:0], ep[1:]...)
	ns := make([]string, len(ep))
	for _, v := range ep {
		m := strings.Split(v, "_")
		mn := make([]string, len(m))
		for _, vv := range m {
			mn = append(mn, strings.Title(vv))
		}
		ns = append(ns, strings.Join(mn, ""))
	}
	return strings.Join(ns, "")
}

// MethodWithNamespace returns the API endpoint method name with namespace.
//
func (e *Endpoint) MethodWithNamespace() string {
	return utils.APIToGo(e.Name)
}

// HumanMethodWithNamespace returns the API endpoint method name in humanized form.
//
func (e *Endpoint) HumanMethodWithNamespace() string {
	var (
		src = e.MethodWithNamespace()
		out string
	)
	for i, l := range src {
		if unicode.IsUpper(l) && i > 0 {
			if i+2 <= len(src) && unicode.IsUpper([]rune(src[i+1 : i+2])[0]) {
				out += string(l)
			} else {
				out += " " + string(l)
			}
		} else {
			out += string(l)
		}
	}
	return out
}

// RequiredArguments return the list of required method arguments.
//
func (e *Endpoint) RequiredArguments() []MethodArgument {
	var args = make([]MethodArgument, 0)
	var prominentArgs = []string{
		"index",
		"id",
		"repository",
		"snapshot",
		"snapshot_id",
		"calendar_id",
		"job_id",
		"username",
		"name",
	}

	var contains = func(s string) bool {
		for _, v := range args {
			if v.Name == s {
				return true
			}
		}
		return false
	}

	var keys []string
	for k := range e.URL.Paths[0].Parts{
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Return the prominent arguments first
	for _, d := range prominentArgs {
		for _, k := range keys {
			p := e.URL.Paths[0].Parts[k]
			if p.Name != d {
				continue
			}
			if p.Required {
				args = append(args, MethodArgument{
					Endpoint:    e,
					Name:        p.Name,
					Type:        p.Type,
					Description: p.Description,
					Required:    true,
				})
			}
		}
	}

	// Return the body after prominent arguments
	if e.Body != nil && e.Body.Required {
		args = append(args, MethodArgument{
			Endpoint:    e,
			Name:        "body",
			Description: e.Body.Description,
		})
	}

	// Return rest of the URL parts
	for _, k := range keys {
		p := e.URL.Paths[0].Parts[k]
		if contains(p.Name) {
			continue
		}
		if p.Required {
			args = append(args, MethodArgument{
				Endpoint:    e,
				Name:        p.Name,
				Type:        p.Type,
				Description: p.Description,
				Required:    true,
			})
		}
	}

	// Return URL params
	for _, p := range e.URL.Params {
		if contains(p.Name) {
			continue
		}
		if p.Required {
			args = append(args, MethodArgument{
				Endpoint:    e,
				Name:        p.Name,
				Type:        p.Type,
				Description: p.Description,
				Options:     p.Options,
				Default:     p.Default,
				Required:    p.Required,
			})
		}
	}

	return args
}

// GoName returns a Go name for part.
//
func (p *Part) GoName() string {
	switch {
	case p.Name == "context":
		return "ScriptContext"
	default:
		return utils.NameToGo(p.Name, p.Endpoint.MethodWithNamespace())
	}
}

// GoType returns a Go type for part.
//
func (p *Part) GoType(comment ...bool) string {
	return utils.TypeToGo(p.Type)
}

// GoName returns a Go name for parameter.
//
func (p *Param) GoName() string {
	switch {
	case p.Name == "context":
		return "ScriptContext"
	case p.Name == "q":
		return "Query"
	default:
		return utils.NameToGo(p.Name, p.Endpoint.MethodWithNamespace())
	}
}

// GoType returns a Go type for parameter.
//
func (p *Param) GoType(comment ...bool) string {
	if f := (&Generator{Endpoint: p.Endpoint}).GetOverride("polymorphic-param", p.Endpoint.Name); f != nil {
		if out := f(p.Endpoint, p.Name); out != "" {
			return out
		}
	}
	return utils.TypeToGo(p.Type)
}

// GoName returns a Go name for method argument.
//
func (p *MethodArgument) GoName() string {
	return utils.NameToGo(p.Name, p.Endpoint.MethodWithNamespace())
}

// GoType returns a Go type for method argument.
//
func (p *MethodArgument) GoType(comment ...bool) string {
	return utils.TypeToGo(p.Type)
}

// ContainsMethods return true if every method passed in argument is present in every Path
func (u *URL) ContainsMethods(methods ...string) bool {
	for _, path := range u.Paths {
		// Fast exit if both collections are not the same size
		if len(methods) != len(path.Methods) {
			return false
		}

		// We iterate over every items
		items := make(map[string]struct{})
		for _, v := range path.Methods {
			items[v] = struct{}{}
		}

		for _, method := range methods {
			if _, ok := items[method]; ok {
				continue
			}
			return false
		}
		continue
	}
	return true
}