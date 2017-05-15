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
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/serenize/snaker"
)

const (
	// TODO: make type its own struct
	specTypeDict       = "dict"
	specTypeBoolean    = "boolean"
	specTypeEnum       = "enum"
	specTypeList       = "list"
	specTypeNumber     = "number"
	specTypeNumberList = "number_list"
	specTypeString     = "string"
	specTypeTime       = "time"
)

type enum struct {
	Name     string
	Type     string
	SpecName string
}

func (e *enum) clone() *enum {
	return &enum{
		Name:     e.Name,
		Type:     e.Type,
		SpecName: e.SpecName,
	}
}

type param struct {
	Name               string
	rawName            string
	SpecType           string `json:"type"`
	Type               string
	Description        string      `json:"description"`
	Required           bool        `json:"required"`
	Default            interface{} `json:"default"`
	Options            []string    `json:"options"`
	OptionName         string
	EnumValues         []*enum
	enumValuesRaw      map[string]*enum
	optionTemplateName string
	OptionTemplate     *template.Template
	// Value is used by the tester to temporarily assign values to the parameter for the purpose of generating tests.
	Value interface{}
}

func formatName(name string, required bool) (string, string) {
	// Make sure we don't mix _source and source
	if name == "source" {
		name = "source_param"
	}
	if strings.HasPrefix(name, "_") {
		name = name[1:]
	}
	var paramName string
	var optionName string
	switch name {
	case "type":
		paramName = "documentType"
	default:
		paramName = snaker.SnakeToCamelLower(name)
	}
	if !required {
		optionName = "With" + snaker.SnakeToCamel(name)
	}
	return paramName, optionName
}

func formatToken(index int, token string) string {
	// If the first word is all upper case let it be, otherwise convert to lower
	if index == 0 && token != "APIs" && (len(token) == 1 || strings.ToUpper(token) != token) {
		token = strings.ToLower(string(token[0])) + token[1:]
	}
	return strings.Replace(token, "`", "\"", -1)
}

func formatDescription(description string) string {
	formatted := ""
	for i, word := range strings.Split(description, " ") {
		token := formatToken(i, word)
		if i > 0 {
			formatted += " "
		}
		formatted += token
	}
	switch formatted[len(formatted)-1] {
	case '.':
		// TODO: investigate the "?" case
	case '?':
	default:
		formatted += "."
	}
	return formatted
}

func (p *param) resolve(name string, templates *template.Template) error {
	p.rawName = name
	p.Name, p.OptionName = formatName(name, p.Required)
	if p.OptionName != "" {
		if p.optionTemplateName == "" {
			p.optionTemplateName = "option.tmpl"
		}
		p.OptionTemplate = templates.Lookup(p.optionTemplateName)
		if p.OptionTemplate == nil {
			return fmt.Errorf("cannot find template for option %q", name)
		}
	}
	switch p.SpecType {
	case specTypeBoolean:
		p.Type = "bool"
	case specTypeEnum:
		p.Type = snaker.SnakeToCamel(p.Name)
		p.EnumValues = []*enum{}
		p.enumValuesRaw = map[string]*enum{}
		for _, o := range p.Options {
			if o == "" {
				o = "zero"
			}
			e := &enum{
				Name:     p.Type + snaker.SnakeToCamel(o),
				Type:     p.Type,
				SpecName: o,
			}
			p.EnumValues = append(p.EnumValues, e)
			p.enumValuesRaw[o] = e
		}
	case specTypeDict:
		p.Type = "map[string]interface{}"
	case specTypeList:
		p.Type = "[]string"
	case specTypeNumber:
		p.Type = "int"
	case specTypeNumberList:
		p.Type = "[]int"
	case specTypeString:
		p.Type = "string"
	case specTypeTime:
		p.Type = "time.Duration"
	case "":
		// TODO: we should remove this param
	default:
		return fmt.Errorf("invalid type for %s: %s", name, p.SpecType)
	}
	p.Description = formatDescription(p.Description)
	return nil
}

func (p *param) equals(other *param) bool {
	if !(p.Name == other.Name && p.SpecType == other.SpecType && p.Description == other.Description &&
		p.Required == other.Required && p.Default == other.Default && len(p.Options) == len(other.Options)) {
		return false
	}
	for i, o := range p.Options {
		if other.Options[i] != o {
			return false
		}
	}
	return true
}

func (p *param) addSuffix(suffix string) {
	p.Name += suffix
	p.OptionName += suffix
}

func (p *param) clone() *param {
	c := &param{
		Name:          p.Name,
		SpecType:      p.SpecType,
		Type:          p.Type,
		Description:   p.Description,
		Required:      p.Required,
		Default:       p.Default,
		Options:       []string{},
		OptionName:    p.OptionName,
		EnumValues:    []*enum{},
		enumValuesRaw: map[string]*enum{},
	}
	for _, o := range p.Options {
		c.Options = append(c.Options, o)
	}
	for _, e := range p.EnumValues {
		c.EnumValues = append(c.EnumValues, e.clone())
	}
	for name, value := range p.enumValuesRaw {
		c.enumValuesRaw[name] = value
	}
	return c
}

type invalidTypeError struct {
	p *param
}

func (i *invalidTypeError) Error() string {
	return fmt.Sprintf("invalid type for %s: %T (expected %s)", i.p.Name, i.p.Value, i.p.Type)
}

func (p *param) OptionString() (string, error) {
	var writer bytes.Buffer
	if err := p.OptionTemplate.Execute(&writer, p); err != nil {
		return "", err
	}
	return writer.String(), nil
}

func (p *param) String() (string, error) {
	if p.Value == nil {
		return "", nil
	}
	switch p.SpecType {
	case specTypeBoolean:
		v, ok := p.Value.(bool)
		if !ok {
			return "", &invalidTypeError{p}
		}
		return fmt.Sprint(v), nil
	case specTypeEnum:
		v, ok := p.Value.(string)
		if !ok {
			return "", &invalidTypeError{p}
		}
		e, ok := p.enumValuesRaw[v]
		if !ok {
			return "", fmt.Errorf("invalid value for enum %q: %s", p.Name, v)
		}
		return fmt.Sprint(e.Name), nil
	case specTypeNumber:
		v, ok := p.Value.(int)
		if !ok {
			return "", &invalidTypeError{p}
		}
		return fmt.Sprint(v), nil
	case specTypeDict:
		v, ok := p.Value.(map[interface{}]interface{})
		if !ok {
			return "", &invalidTypeError{p}
		}
		code := "map[string]interface{}{"
		for name, value := range v {
			// TODO: format value
			code += fmt.Sprintf("\n\"%s\" : \"%s\",", name.(string), value)
		}
		code += "\n}"
		return code, nil
	case specTypeString:
		v, ok := p.Value.(string)
		if !ok {
			i, ok := p.Value.(int)
			if !ok {
				return "", &invalidTypeError{p}
			}
			return fmt.Sprintf("\"%d\"", i), nil
		}
		return "\"" + v + "\"", nil
	case specTypeList:
		v, ok := p.Value.([]interface{})
		if !ok {
			stringValue, ok := p.Value.(string)
			if !ok {
				return "", &invalidTypeError{p}
			}
			return fmt.Sprint(stringValue), nil
		}
		value := "[]string{\n"
		for _, item := range v {
			value += item.(string) + ",\n"
		}
		return value + "\n}", nil
	case specTypeTime:
		v, ok := p.Value.(string)
		if !ok {
			return "", &invalidTypeError{p}
		}
		return fmt.Sprintf("time.ParseDuration(\"%s\")", v), nil
	}
	return "", nil
}
