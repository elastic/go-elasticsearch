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
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"

	"golang.org/x/tools/imports"

	"github.com/elastic/go-elasticsearch/v8/internal/build/utils"
)

// Generator represents the "gensource" generator.
//
type Generator struct {
	b bytes.Buffer

	Endpoint *Endpoint
}

// Output returns the generator output.
//
func (g *Generator) Output() (io.Reader, error) {
	g.genHeader()
	g.genConstructor()
	g.genMethodDefinition()
	g.genRequestStruct()
	g.w("\n")
	g.genDoMethod()
	g.genWithOptions()

	return bytes.NewReader(g.b.Bytes()), nil
}

// OutputFormatted returns a formatted generator output.
//
func (g *Generator) OutputFormatted() (io.Reader, error) {
	out, err := g.Output()
	if err != nil {
		return bytes.NewReader(g.b.Bytes()), err
	}

	var b bytes.Buffer
	if _, err := io.Copy(&b, out); err != nil {
		return bytes.NewReader(g.b.Bytes()), err
	}

	fout, err := imports.Process(
		"esapi/api."+g.Endpoint.Name+".go",
		g.b.Bytes(),
		&imports.Options{
			AllErrors:  true,
			Comments:   true,
			FormatOnly: false,
			TabIndent:  true,
			TabWidth:   1,
		})
	if err != nil {
		return bytes.NewReader(b.Bytes()), err
	}

	g.b.Reset()
	g.b.Write(fout)

	return bytes.NewReader(fout), nil
}

func (g *Generator) w(s string) {
	g.b.WriteString(s)
}

func (g *Generator) genHeader() {
	g.w(`// Licensed to Elasticsearch B.V. under one or more contributor
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
// under the License.` + "\n")
	g.w("//\n// Code generated")
	if EsVersion != "" {
		g.w(fmt.Sprintf(" from specification version %s", EsVersion))
	}
	g.w(": DO NOT EDIT\n")
	g.w("\n")
	g.w("package esapi\n")
}

func (g *Generator) genConstructor() {
	g.w(`
func new` + g.Endpoint.MethodWithNamespace() + `Func(t Transport) ` + g.Endpoint.MethodWithNamespace() + ` {
	return func(`)
	g.genMethodArguments()
	g.w(`o ...func(*` + g.Endpoint.MethodWithNamespace() + `Request)) (*Response, error) {`)
	if len(g.Endpoint.RequiredArguments()) > 0 {
		g.w("\n\t\t" + `var r = ` + g.Endpoint.MethodWithNamespace() + `Request{`)
		for i, arg := range g.Endpoint.RequiredArguments() {
			if arg.Name == "type" {
				continue // Skip the type parameter, "_doc" is used by default
			}
			g.w(arg.GoName() + ": " + arg.Name)
			if i != len(g.Endpoint.RequiredArguments())-1 {
				g.w(", ")
			}
		}
		g.w("}\n")
	} else {
		g.w("\n\t\t" + `var r = ` + g.Endpoint.MethodWithNamespace() + `Request{}` + "\n")
	}
	g.w(`		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}
`)
}

func (g *Generator) genMethodDefinition() {
	g.w("\n// ----- API Definition -------------------------------------------------------\n\n")

	if g.Endpoint.Type == "xpack" {
		g.w(`// ` + g.Endpoint.MethodWithNamespace() + " - " + g.Endpoint.Documentation.Description)
	} else {
		if g.Endpoint.Documentation.Description != "" {
			words := strings.Split(g.Endpoint.Documentation.Description, " ")
			initial := strings.ToLower(words[0:1][0])
			description := initial + " " + strings.Join(words[1:], " ")
			lines := strings.Split(description, "\n")

			g.w(`// ` + g.Endpoint.MethodWithNamespace() + " " + lines[0:1][0])
			for _, line := range lines[1:] {
				g.w("\n// " + line)
			}
		}
	}

	if g.Endpoint.Stability != "stable" {
		g.w("\n//\n// This API is " + g.Endpoint.Stability + ".")
	}

	if g.Endpoint.Documentation.URL != "" {
		g.w("\n//\n" + `// See full documentation at ` + g.Endpoint.Documentation.URL + ".")
	}

	g.w(`
//
type ` + g.Endpoint.MethodWithNamespace() + ` func(`)
	g.genMethodArguments()
	g.w(`o ...func(*` + g.Endpoint.MethodWithNamespace() + `Request)) (*Response, error)
`)
}

func (g *Generator) genMethodArguments() {
	if len(g.Endpoint.RequiredArguments()) > 0 {
		for _, arg := range g.Endpoint.RequiredArguments() {
			if arg.Name == "type" {
				continue // Skip the type parameter, "_doc" is used by default
			}
			var argType string
			if arg.Name == "body" {
				argType = "io.Reader"
			} else {
				argType = arg.GoType()
			}
			var argName string
			if arg.Name == "type" {
				argName = "typ"
			} else {
				argName = arg.Name
			}
			g.w(argName + " " + argType)
			g.w(", ")
		}
	}
}

func (g *Generator) genRequestStruct() {
	g.w(`
// ` + g.Endpoint.MethodWithNamespace() + `Request configures the ` + g.Endpoint.HumanMethodWithNamespace() + ` API request.
//
type ` + g.Endpoint.MethodWithNamespace() + `Request struct {`)
	specialFields := []string{"index", "id"}
	for _, n := range specialFields {
		if param, ok := g.Endpoint.URL.AllParts[n]; ok {
			g.w("\n\t" + param.GoName())
			g.w("\t" + param.GoType(true))
		}
	}

	if len(g.Endpoint.URL.AllParts) > 0 {
		g.w("\n")
	}

	if g.Endpoint.Body != nil {
		g.w("\n\tBody io.Reader")
	}

	if len(g.Endpoint.URL.AllParts) > 0 || g.Endpoint.Body != nil {
		g.w("\n")
	}

	for _, name := range g.Endpoint.URL.PartNamesSorted {
		p, ok := g.Endpoint.URL.AllParts[name]
		if !ok {
			panic(fmt.Sprintf("Part %q not found", name))
		}

		skip := false
		for _, v := range specialFields {
			if p.Name == v {
				skip = true
			}
		}
		// Skip type only for selected APIs at this point
		if p.Name == "type" && p.Deprecated {
			if g.Endpoint.Name == "bulk" ||
				g.Endpoint.Name == "create" ||
				g.Endpoint.Name == "delete" ||
				g.Endpoint.Name == "index" ||
				g.Endpoint.Name == "update" {
				continue
			}
		}

		if skip {
			continue
		}
		g.w("\n\t" + p.GoName())
		g.w("\t" + p.GoType(true))

	}

	if len(g.Endpoint.URL.AllParts) > 0 {
		g.w("\n")
	}

	for _, name := range g.Endpoint.URL.ParamNamesSorted {
		p, ok := g.Endpoint.URL.Params[name]
		if !ok {
			panic(fmt.Sprintf("Parameter %q not found", name))
		}

		if _, ok := g.Endpoint.URL.AllParts[name]; ok {
			continue // skip params which are also parts
		}

		g.w("\n\t" + p.GoName())
		g.w("\t" + p.GoType(true))
	}

	g.w("\n\n\tPretty\tbool")
	g.w("\n\tHuman\tbool")
	g.w("\n\tErrorTrace\tbool")
	g.w("\n\tFilterPath\t[]string\n")

	g.w("\n\n\tHeader\thttp.Header\n")

	g.w("\n\tctx context.Context\n}\n")
}

func (g *Generator) genWithOptions() {
	// Generate WithContext first
	g.w(`
// WithContext sets the request context.
//
func (f ` + g.Endpoint.MethodWithNamespace() + `) WithContext(v context.Context) func(*` + g.Endpoint.MethodWithNamespace() + `Request) {
	return func(r *` + g.Endpoint.MethodWithNamespace() + `Request) {
		r.ctx = v
	}
}
`)

	// Skip adding With... options for arguments which are part of the method signature
	skipRequiredArgs := make(map[string]bool)
	for _, p := range g.Endpoint.RequiredArguments() {
		if p.Name != "type" {
			skipRequiredArgs[p.Name] = true
		}
	}

	var methodBody = func(e *Endpoint, a interface{}) string {
		var b strings.Builder

		switch a.(type) {
		case *Part, *Param: // pass
		default:
			panic(fmt.Sprintf("FAIL: %q: Unexpected type [%[1]T] for argument: %#[1]v", g.Endpoint.Name, a))
		}

		var (
			methodWithNamespace = e.MethodWithNamespace()

			typ        = reflect.TypeOf(a).String()
			pFieldName = reflect.ValueOf(a).MethodByName("GoName").Call([]reflect.Value{})[0].String()
			pGoType    = reflect.ValueOf(a).MethodByName("GoType").Call([]reflect.Value{reflect.ValueOf(false)})[0].String()
			// pType      = reflect.Indirect(reflect.ValueOf(a)).FieldByName("Type").String()
			// pName      = reflect.Indirect(reflect.ValueOf(a)).FieldByName("Name").String()
			// pGoName = reflect.ValueOf(a).MethodByName("GoName").Call([]reflect.Value{})[0].String()

			pDesc = utils.IDToUpper(strings.ToLower(reflect.Indirect(reflect.ValueOf(a)).FieldByName("Description").String()))
		)

		// Adjust descriptions
		if strings.Contains(pDesc, "a comma-separated list") {
			pDesc = strings.Replace(pDesc, "a comma-separated list", "a list", -1)
		}
		if strings.Contains(pDesc, "use `_all` or empty string") {
			pDesc = strings.Replace(pDesc, "use `_all` or empty string", "use _all", -1)
		}

		// Generate annotation
		b.WriteString("\n// With" + pFieldName)
		if typ == "*gensource.Part" {
			b.WriteString(` - ` + pDesc)
		} else {
			b.WriteString(` - ` + pDesc)
		}
		b.WriteString(`.`)

		// Generate function
		b.WriteString("\n//\nfunc (f " + methodWithNamespace + `) With` + pFieldName + `(`)

		switch pGoType {
		case "bool":
			// empty function argument
		case "*bool":
			b.WriteString(`v bool`)
		case "*int":
			b.WriteString(`v int`)
		case "[]string":
			b.WriteString(`v ...string`)
		default:
			b.WriteString(`v ` + pGoType)
		}

		b.WriteString(`) func(*` + methodWithNamespace + `Request) {
	return func(r *` + methodWithNamespace + `Request) {` + "\n")

		switch pGoType {
		case "bool":
			b.WriteString("\t\t" + `r.` + pFieldName + ` = true`)
		case "*bool", "*int":
			b.WriteString("\t\t" + `r.` + pFieldName + ` = &v`)
		default:
			b.WriteString("\t\t" + `r.` + pFieldName + ` = v`)
		}

		b.WriteString("\n\t}\n}\n")

		return b.String()
	}

	// Generate WithBody method
	if b := g.Endpoint.Body; b != nil {
		// Do not add the option when body is part of the method signature
		if !skipRequiredArgs["body"] {
			g.w(`
// WithBody` + ` - ` + b.Description + `.
//
func (f ` + g.Endpoint.MethodWithNamespace() + `) WithBody(v io.Reader) func(*` + g.Endpoint.MethodWithNamespace() + `Request) {
	return func(r *` + g.Endpoint.MethodWithNamespace() + `Request) {
		r.Body = v
	}
}
`)
		}
	}

	// Generate With... methods for parts
	for _, pName := range g.Endpoint.URL.PartNamesSorted {
		if p, ok := g.Endpoint.URL.AllParts[pName]; ok {
			if skipRequiredArgs[p.Name] {
				continue
			}

			g.w(methodBody(g.Endpoint, p))
		} else {
			g.w(`// TODO: ` + p.Name)
		}
	}

	// Generate With... methods for params
	for _, pName := range g.Endpoint.URL.ParamNamesSorted {
		if _, ok := g.Endpoint.URL.AllParts[pName]; ok {
			continue // skip params which are also parts
		}
		if p, ok := g.Endpoint.URL.Params[pName]; ok {
			g.w(methodBody(g.Endpoint, p))
		} else {
			g.w(`// TODO: ` + p.Name)
		}
	}

	// Generate methods for common parameters
	if g.Endpoint.Name != "info" {
		g.w(`
// WithPretty makes the response body pretty-printed.
//
func (f ` + g.Endpoint.MethodWithNamespace() + `) WithPretty() func(*` + g.Endpoint.MethodWithNamespace() + `Request) {
	return func(r *` + g.Endpoint.MethodWithNamespace() + `Request) {
		r.Pretty = true
	}
}
`)
	}

	g.w(`
// WithHuman makes statistical values human-readable.
//
func (f ` + g.Endpoint.MethodWithNamespace() + `) WithHuman() func(*` + g.Endpoint.MethodWithNamespace() + `Request) {
	return func(r *` + g.Endpoint.MethodWithNamespace() + `Request) {
		r.Human = true
	}
}
`)

	g.w(`
// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f ` + g.Endpoint.MethodWithNamespace() + `) WithErrorTrace() func(*` + g.Endpoint.MethodWithNamespace() + `Request) {
	return func(r *` + g.Endpoint.MethodWithNamespace() + `Request) {
		r.ErrorTrace = true
	}
}
`)

	g.w(`
// WithFilterPath filters the properties of the response body.
//
func (f ` + g.Endpoint.MethodWithNamespace() + `) WithFilterPath(v ...string) func(*` + g.Endpoint.MethodWithNamespace() + `Request) {
	return func(r *` + g.Endpoint.MethodWithNamespace() + `Request) {
		r.FilterPath = v
	}
}
`)

	// Generate methods for HTTP headers
	g.w(`
// WithHeader adds the headers to the HTTP request.
//
func (f ` + g.Endpoint.MethodWithNamespace() + `) WithHeader(h map[string]string) func(*` + g.Endpoint.MethodWithNamespace() + `Request) {
	return func(r *` + g.Endpoint.MethodWithNamespace() + `Request) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}
`)

	// Generate methods for the X-Opaque-ID header
	g.w(`
// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f ` + g.Endpoint.MethodWithNamespace() + `) WithOpaqueID(s string) func(*` + g.Endpoint.MethodWithNamespace() + `Request) {
	return func(r *` + g.Endpoint.MethodWithNamespace() + `Request) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
`)
}

func (g *Generator) genDoMethod() {
	g.w(`// Do executes the request and returns response or error.
//
func (r ` + g.Endpoint.MethodWithNamespace() + `Request) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method  string
		path    strings.Builder
		params  map[string]string
	)` + "\n\n")

	switch g.Endpoint.Name {
	case "index":
		g.w("\t")
		g.w(`if r.DocumentID != "" {
		method = "PUT"
	} else {
		method = "POST"
	}`)
		g.w("\n\n")
	default:
		var httpMethod string
		// If endpoint has both GET and POST available
		// Prefer POST usage in order to prevent go routine leak
		// See https://github.com/golang/go/issues/29246
		if g.Endpoint.URL.ContainsMethods("GET", "POST") {
			httpMethod = "POST"
		} else {
			httpMethod = g.Endpoint.URL.Paths[0].Methods[0]
		}
		g.w("\t" + `method = "` + httpMethod + `"` + "\n\n")
	}

	// Get default part values for specific APIs
	// TODO: Move to overrides
	var defparts bool
	switch g.Endpoint.Name {
	case "index", "create", "delete", "explain", "exists", "get", "get_source", "update", "termvectors":
		for _, p := range g.Endpoint.URL.AllParts {
			if p.Default != nil {
				var fieldName string
				var fieldValue string
				var fieldCondition string

				fieldName = p.GoName()
				switch p.Type {
				case "string", "enum":
					fieldCondition = `r.` + fieldName + ` == ""`
					fieldValue = `"` + p.Default.(string) + `"`
				case "number":
					fieldCondition = `r.` + fieldName + ` == 0`
					fieldValue = p.Default.(string)
				case "list":
					fieldCondition = ` len(r.` + fieldName + `) < 1`
					fieldValue = `[]string{"` + p.Default.(string) + `"}`
				default:
					panic(fmt.Sprintf("FAIL: %q: unexpected parameter type %q for URL part %q", g.Endpoint.Name, p.Type, p.Name))
				}
				g.w("\t")
				g.w(`if ` + fieldCondition + ` {
		r.` + fieldName + ` = ` + fieldValue + `
	}` + "\n")
				defparts = true
			}
		}
	}
	if defparts {
		g.w("\n")
	}

	// Generate the URL path
	//
	if f := g.GetOverride("url", g.Endpoint.Name); f != nil {
		g.w(f(g.Endpoint))
	} else {
		var (
			pathGrow    strings.Builder
			pathContent strings.Builder
		)

		pathGrow.WriteString(`	path.Grow(`)

		// FIXME: Select longest path based on number of template entries, not string length
		longestPath := g.Endpoint.URL.Paths[0]
		for _, p := range g.Endpoint.URL.Paths {
			// Skip deprecated parts only for selected APIs at this point
			if p.Deprecated.Version != "" && p.Deprecated.Version < EsVersion {
				if g.Endpoint.Name == "bulk" ||
					g.Endpoint.Name == "create" ||
					g.Endpoint.Name == "delete" ||
					g.Endpoint.Name == "get" ||
					g.Endpoint.Name == "index" ||
					g.Endpoint.Name == "update" {
					continue
				}
			}
			if len(p.Path) > len(longestPath.Path) {
				longestPath = p
			}
		}

		if len(longestPath.Parts) < 1 {
			if len(g.Endpoint.URL.Paths) < 1 {
				panic(fmt.Sprintf("FAIL: %q: empty endpoint\n", g.Endpoint.Name))
			}
			pathGrow.WriteString(`len("` + longestPath.Path + `")`)
			pathContent.WriteString(`	path.WriteString("` + longestPath.Path + `")` + "\n")

		} else {
			pathParts := make([]string, 0)
			apiArgs := g.Endpoint.RequiredArguments()
			for _, v := range strings.Split(longestPath.Path, "/") {
				if v != "" {
					pathParts = append(pathParts, v)
				}
			}

			r := strings.NewReplacer("{", "", "}", "")

			for _, v := range pathParts {
				var p string

				// Required arguments
				for _, a := range apiArgs {
					if strings.HasPrefix(v, "{") && a.Name == r.Replace(v) {
						p = a.GoName()
						pathGrow.WriteString(`1 + `)
						pathContent.WriteString(`	path.WriteString("/")` + "\n")
						switch a.Type {
						case "int":
							pathGrow.WriteString(`len(strconv.Itoa(*r.` + p + `)) + `)
							pathContent.WriteString(`	path.WriteString(strconv.Itoa(*r.` + p + `))` + "\n")
						case "string":
							pathGrow.WriteString(`len(r.` + p + `) + `)
							pathContent.WriteString(`	path.WriteString(r.` + p + `)` + "\n")
						case "list":
							pathGrow.WriteString(`len(strings.Join(r.` + p + `, ",")) + `)
							pathContent.WriteString(`	path.WriteString(strings.Join(r.` + p + `, ","))` + "\n")
						case "long":
							pathGrow.WriteString(`len(strconv.Itoa(*r.` + p + `)) + `)
							pathContent.WriteString(`	path.WriteString(strconv.Itoa(*r.` + p + `))` + "\n")
						default:
							panic(fmt.Sprintf("FAIL: %q: unexpected type %q for URL part %q\n", g.Endpoint.Name, a.Type, a.Name))
						}
						break
					}
				}

				// Optional arguments
				if p == "" {
					for _, a := range longestPath.Parts {
						// fmt.Printf("%s: %+v\n", a.Name, a)
						if strings.HasPrefix(v, "{") && a.Name == r.Replace(v) {
							p = a.GoName()

							switch a.Type {
							case "string":
								pathGrow.WriteString(`1 + len(r.` + p + `) + `)
								pathContent.WriteString(`	if r.` + p + ` != "" {` + "\n")
								pathContent.WriteString(`		path.WriteString("/")` + "\n")
								pathContent.WriteString(`		path.WriteString(r.` + p + `)` + "\n")
								pathContent.WriteString(`	}` + "\n")
							case "list":
								pathGrow.WriteString(`1 + len(strings.Join(r.` + p + `, ",")) + `)
								pathContent.WriteString(`	if len(r.` + p + `) > 0 {` + "\n")
								pathContent.WriteString(`		path.WriteString("/")` + "\n")
								pathContent.WriteString(`		path.WriteString(strings.Join(r.` + p + `, ","))` + "\n")
								pathContent.WriteString(`	}` + "\n")
							case "int", "long":
								pathContent.WriteString(`	if r.` + p + ` != nil {` + "\n")
								pathContent.WriteString(`		value := strconv.FormatInt(int64(*r.` + p + `), 10)` + "\n")
								pathContent.WriteString(`		path.Grow(1 + len(value))` + "\n")
								pathContent.WriteString(`		path.WriteString("/")` + "\n")
								pathContent.WriteString(`		path.WriteString(value)` + "\n")
								pathContent.WriteString(`	}` + "\n")
							default:
								panic(fmt.Sprintf("FAIL: %q: unexpected type %q for URL part %q\n", g.Endpoint.Name, a.Type, a.Name))
							}

							break
						}
					}
				}

				// Optional arguments
				if p == "" {
					for _, a := range g.Endpoint.URL.Params {
						if strings.HasPrefix(v, "{") && a.Name == r.Replace(v) {
							p = a.GoName()
							pathGrow.WriteString("1 +")
							pathContent.WriteString(`	path.WriteString("/")` + "\n")
							switch a.Type {
							case "string":
								pathGrow.WriteString(`len(r.` + p + `)`)
								pathContent.WriteString(`	path.WriteString(r.` + p + `)` + "\n")
							case "list":
								pathGrow.WriteString(`len(strings.Join(r.` + p + `, ","))`)
								pathContent.WriteString(`	path.WriteString(strings.Join(r.` + p + `, ","))` + "\n")
							default:
								panic(fmt.Sprintf("FAIL: %q: unexpected type %q for URL param %q\n", g.Endpoint.Name, a.Type, a.Name))
							}
							break
						}
					}
				}

				// Static parts
				if p == "" {
					pathGrow.WriteString(`1 + len("` + v + `") + `)
					pathContent.WriteString(`	path.WriteString("/")` + "\n")
					pathContent.WriteString(`	path.WriteString("` + v + `")` + "\n")
				}
			}
		}

		// Write out the content
		pathGrow.WriteString(`)`)
		g.w(strings.Replace(pathGrow.String(), " + )", ")", 1) + "\n")
		g.w(pathContent.String() + "\n")
	}

	// Generate the URL params
	g.w(`
	params = make(map[string]string)` + "\n")
	for _, n := range g.Endpoint.URL.ParamNamesSorted {
		if p, ok := g.Endpoint.URL.Params[n]; ok {
			var (
				fieldName      string
				fieldType      string
				fieldValue     string
				fieldCondition string
			)

			fieldName = p.GoName()
			fieldType = p.GoType()
			switch fieldType {
			case "bool":
				fieldCondition = `r.` + fieldName
				fieldValue = `strconv.FormatBool(r.` + fieldName + `)`
			case "*bool":
				fieldCondition = `r.` + fieldName + ` != nil`
				fieldValue = `strconv.FormatBool(*r.` + fieldName + `)`
			case "string":
				fieldCondition = `r.` + fieldName + ` != ""`
				fieldValue = `r.` + fieldName
			case "int":
				fieldCondition = `r.` + fieldName + ` != 0`
				fieldValue = `strconv.FormatInt(int64(r.` + fieldName + `), 10)`
			case "*int":
				fieldCondition = `r.` + fieldName + ` != nil`
				fieldValue = `strconv.FormatInt(int64(*r.` + fieldName + `), 10)`
			case "uint":
				fieldCondition = `r.` + fieldName + ` != 0`
				fieldValue = `strconv.FormatUint(uint64(r.` + fieldName + `), 10)`
			case "*uint":
				fieldCondition = `r.` + fieldName + ` != 0`
				fieldValue = `strconv.FormatUint(uint64(*r.` + fieldName + `), 10)`
			case "[]string":
				fieldCondition = ` len(r.` + fieldName + `) > 0`
				fieldValue = `strings.Join(r.` + fieldName + `, ",")`
			case "time.Duration":
				fieldCondition = `r.` + fieldName + ` != 0`
				fieldValue = `formatDuration(r.` + fieldName + `)`
			default: // interface{}
				fieldCondition = `r.` + fieldName + ` != nil`
				// TODO: Use type switching instead?
				fieldValue = `fmt.Sprintf("%v", r.` + fieldName + `)`
			}

			g.w(`
	if ` + fieldCondition + ` {
		params["` + p.Name + `"] = ` + fieldValue + `
	}` + "\n")

		} else {
			panic(fmt.Sprintf("FAIL: %q: Unknown parameter %q in URL parameters", g.Endpoint.Name, n))
		}
	}

	// Common parameters
	g.w(`
	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}
	`)
	g.w("\n\n")

	// Generate the HTTP request options
	var httpBody string
	if g.Endpoint.Body != nil {
		httpBody = "r.Body"
	} else {
		httpBody = "nil"
	}

	g.w(`req, err := newRequest(method, path.String(), ` + httpBody + `)` + "\n")
	g.w(`if err != nil {
		return nil, err
	}` + "\n\n")

	g.w(`if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}` + "\n\n")

	if g.Endpoint.Body != nil {
		g.w(`if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
	}` + "\n\n")
	}

	g.w(`if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}` + "\n\n")

	g.w(`if ctx != nil {
		req = req.WithContext(ctx)
	}` + "\n\n")

	g.w(`
	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}` + "\n\n")

	// Generate the return value
	g.w(`
	response := Response{
		StatusCode:	res.StatusCode,
		Body:				res.Body,
		Header:			res.Header,
	}` + "\n")
	g.w("\n\treturn &response, nil\n")

	g.w("}\n")
}
