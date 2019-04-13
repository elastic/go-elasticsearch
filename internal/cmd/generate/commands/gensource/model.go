package gensource

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/elastic/go-elasticsearch/v8/internal/cmd/generate/utils"
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

	if err := json.NewDecoder(f).Decode(&spec); err != nil {
		return nil, err
	}
	for name, e := range spec {
		endpoint = e
		endpoint.Name = name
	}
	for partName, p := range endpoint.URL.Parts {
		p.Endpoint = &endpoint
		p.Name = partName
	}
	for paramName, p := range endpoint.URL.Params {
		p.Endpoint = &endpoint
		p.Name = paramName
	}

	var partNames []string
	for _, param := range endpoint.URL.Parts {
		partNames = append(partNames, param.Name)
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

	for _, param := range endpoint.URL.Parts {
		if param.Name == "type" {
			param.Default = "_doc"
		}
	}

	if info, ok := apiDescriptions[endpoint.Name]; ok {
		if desc, ok := info.(map[interface{}]interface{})["description"].(string); ok {
			endpoint.Description = desc
		}
	}

	return &endpoint, nil
}

// Endpoint represents an API endpoint.
//
type Endpoint struct {
	Name string `json:"-"`

	Description   string   `json:"-"`
	Documentation string   `json:"documentation"`
	Methods       []string `json:"methods"`
	URL           *URL     `json:"url"`
	Body          *Body    `json:"body"`
}

// URL represents API endpoint URL.
//
type URL struct {
	Endpoint *Endpoint `json:"-"`

	Path   string            `json:"path"`
	Paths  []string          `json:"paths"`
	Parts  map[string]*Part  `json:"parts"`
	Params map[string]*Param `json:"params"`

	PartNamesSorted  []string
	ParamNamesSorted []string
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
	return strings.Title(ep[0])
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
	ep := strings.Split(e.Name, ".")
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

// HumanMethodWithNamespace returns the API endpoint method name in humanized form.
//
func (e *Endpoint) HumanMethodWithNamespace() string {
	ep := strings.Split(e.Name, ".")
	ns := make([]string, len(ep))
	for _, v := range ep {
		m := strings.Split(v, "_")
		mn := make([]string, len(m))
		for _, vv := range m {
			mn = append(mn, strings.Title(vv))
		}
		ns = append(ns, strings.Join(mn, " "))
	}
	return strings.TrimSpace(strings.Join(ns, ""))
}

// RequiredArguments return the list of required method arguments.
//
func (e *Endpoint) RequiredArguments() []MethodArgument {
	var args = make([]MethodArgument, 0)
	var prominentArgs = []string{"index", "type", "id", "repository", "snapshot"}

	var contains = func(s string) bool {
		for _, v := range args {
			if v.Name == s {
				return true
			}
		}
		return false
	}

	// Return the prominent arguments first
	for _, d := range prominentArgs {
		for _, p := range e.URL.Parts {
			if p.Name != d {
				continue
			}
			if p.Required {
				args = append(args, MethodArgument{
					Name:        p.Name,
					Type:        p.Type,
					Description: p.Description,
					Required:    true,
				})
			}
			// Make the 'type' required for selected APIs
			if p.Name == "type" && (e.Name == "index" || e.Name == "create") {
				args = append(args, MethodArgument{
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
			Name:        "body",
			Description: e.Body.Description,
		})
	}

	// Return rest of the URL parts
	for _, p := range e.URL.Parts {
		if contains(p.Name) {
			continue
		}
		if p.Required {
			args = append(args, MethodArgument{
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
		return utils.NameToGo(p.Name)
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
		return utils.NameToGo(p.Name)
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
	return utils.NameToGo(p.Name)
}

// GoType returns a Go type for method argument.
//
func (p *MethodArgument) GoType(comment ...bool) string {
	return utils.TypeToGo(p.Type)
}
