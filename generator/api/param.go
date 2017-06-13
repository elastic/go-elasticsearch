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

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/serenize/snaker"
)

const (
	bodyParam = "body"
	nilString = "nil"
	// TODO: make type its own struct
	specTypeBoolean    = "boolean"
	specTypeEnum       = "enum"
	specTypeList       = "list"
	specTypeNumber     = "number"
	specTypeNumberList = "number_list"
	specTypeString     = "string"
	specTypeTime       = "time"
	// These 2 type names are made up to match "body" and the argument to Bulk().
	specTypeDict = "dict"
	specTypeBulk = "bulk"
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

// Param is a parameter for a method.
type Param struct {
	Name               string
	rawName            string
	SpecType           string `json:"type"`
	Type               string
	Description        string      `json:"description"`
	Required           bool        `json:"required"`
	Serialize          string      `json:"serialize"`
	Default            interface{} `json:"default"`
	Options            []string    `json:"options"`
	PackageName        string
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

func (p *Param) resolve(name, packageName string, templates *template.Template) error {
	p.rawName = name
	p.PackageName = packageName
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
	if p.SpecType == "" && p.rawName == bodyParam {
		if p.Serialize == "bulk" {
			p.SpecType = specTypeBulk
		} else {
			p.SpecType = specTypeDict
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
	case specTypeBulk:
		p.Type = "[]interface{}"
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
		return &noTypeError{p}
	default:
		return &invalidTypeError{p}
	}
	p.Description = formatDescription(p.Description)
	return nil
}

func (p *Param) diff(other *Param) map[string]struct{} {
	diffs := map[string]struct{}{}
	if p.Name != other.Name {
		diffs["Name"] = struct{}{}
	}
	if p.SpecType != other.SpecType {
		diffs["SpecType"] = struct{}{}
	}
	if p.Required != other.Required {
		diffs["Required"] = struct{}{}
	}
	if p.Default != other.Default {
		diffs["Default"] = struct{}{}
	}
	if len(p.Options) != len(other.Options) {
		diffs["Options"] = struct{}{}
		return diffs
	}
	for i, o := range p.Options {
		if other.Options[i] != o {
			diffs["Options"] = struct{}{}
			return diffs
		}
	}
	return diffs
}

func (p *Param) addSuffix(suffix string) error {
	if strings.HasSuffix(p.Name, suffix) {
		return fmt.Errorf("attempting to re-add suffix to %s", p.Name)
	}
	p.Name += suffix
	p.OptionName += suffix
	return nil
}

func (p *Param) deduplicate(other *Param, swap bool) error {
	diffs := p.diff(other)
	if len(diffs) == 0 {
		return nil
	}
	if _, ok := diffs["SpecType"]; !ok {
		// If the types are different and only the defaults are different, apply the same the default to both if one of them
		// is not set.
		if _, ok = diffs["Default"]; ok && len(diffs) == 1 {
			if p.Default == nil {
				p.Default = other.Default
			} else if other.Default == nil {
				other.Default = p.Default
			} else {
				return fmt.Errorf("found two versions of %q with different defaults", p.Name)
			}
		} else {
			return fmt.Errorf("found two different versions of %q differing in: %s", p.Name, diffs)
		}
	}
	switch p.SpecType {
	case specTypeBoolean:
		if other.SpecType != specTypeBoolean {
			p.addSuffix("Flag")
		}
	case specTypeList:
		if other.SpecType != specTypeList {
			p.addSuffix("List")
		}
	default:
		if swap {
			return other.deduplicate(p, false)
		}
		return fmt.Errorf("found two different versions of %q (%s and %s)", p.Name, p.SpecType, other.SpecType)
	}
	return nil
}

func (p *Param) clone() *Param {
	c := &Param{
		Name:          p.Name,
		SpecType:      p.SpecType,
		Type:          p.Type,
		Description:   p.Description,
		Required:      p.Required,
		Default:       p.Default,
		Options:       []string{},
		PackageName:   p.PackageName,
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

// OptionString renders a functional option.
func (p *Param) OptionString() (string, error) {
	var writer bytes.Buffer
	if err := p.OptionTemplate.Execute(&writer, p); err != nil {
		return "", err
	}
	return writer.String(), nil
}

func (p *Param) String() (string, error) {
	switch p.SpecType {
	case specTypeBoolean:
		if p.Value == nil {
			if p.Required {
				return "false", nil
			}
			return "", nil
		}
		v, ok := p.Value.(bool)
		if !ok {
			return "", &invalidTypeError{p}
		}
		return fmt.Sprint(v), nil
	case specTypeEnum:
		if p.Value == nil {
			if p.Required {
				return p.EnumValues[0].Name, nil
			}
			return "", nil
		}
		v, ok := p.Value.(string)
		if !ok {
			boolValue, isBool := p.Value.(bool)
			if !isBool {
				listValue, isList := p.Value.([]interface{})
				if !isList {
					return "", &invalidTypeError{p}
				}
				if len(listValue) > 1 {
					if len(listValue) == 2 && listValue[0].(string) == "open" && listValue[1].(string) == "closed" {
						listValue = []interface{}{"all"}
					} else {
						return "", &multipleEnumValuesError{p}
					}
				}
				if v, ok = listValue[0].(string); !ok {
					return "", &invalidTypeError{p}
				}
			} else if boolValue {
				v = "true"
			} else {
				v = "false"
			}
		}
		// Use the enum name if we find one, otherwise assume we're testing a negative case.
		if e, valid := p.enumValuesRaw[v]; valid {
			v = p.PackageName + "." + e.Name
		} else {
			v = strconv.Itoa(len(p.EnumValues))
		}
		return v, nil
	case specTypeNumber:
		if p.Value == nil {
			if p.Required {
				return "0", nil
			}
			return "", nil
		}
		v, ok := p.Value.(int)
		if !ok {
			return "", &invalidTypeError{p}
		}
		return fmt.Sprint(v), nil
	case specTypeDict:
		if p.Value == nil {
			if p.Required {
				return nilString, nil
			}
			return "", nil
		}
		v, ok := p.Value.(map[interface{}]interface{})
		if !ok {
			stringValue, isString := p.Value.(string)
			if !isString {
				return "", &invalidTypeError{p}
			}
			var value map[string]interface{}
			if err := json.Unmarshal([]byte(stringValue), &value); err != nil {
				return "", &invalidDictValueError{p: p, err: err}
			}
			v = map[interface{}]interface{}{}
			for key, item := range value {
				v[key] = item
			}
		}
		code := "map[string]interface{}{"
		for _ = range v {
			// TODO: implement regex matching here
		}
		code += "\n}"
		return code, nil
	case specTypeBulk:
		if p.Value == nil {
			if p.Required {
				return nilString, nil
			}
			return "", nil
		}
		v, ok := p.Value.([]interface{})
		if !ok {
			stringValue, isString := p.Value.(string)
			if !isString {
				return "", &invalidTypeError{p}
			}
			for _, line := range strings.Split(stringValue, "\n") {
				var value map[string]interface{}
				line = strings.Trim(line, " ")
				if line == "" {
					break
				}
				if err := json.Unmarshal([]byte(line), &value); err != nil {
					return "", &invalidBulkValueError{p: p, err: err}
				}
				v = append(v, value)
			}
		}
		code := "[]map[string]interface{}{"
		for _ = range v {
			// TODO: implement
			code += "{},"
		}
		code += "\n}"
		return code, nil
	case specTypeString:
		if p.Value == nil {
			if p.Required {
				return `""`, nil
			}
			return "", nil
		}
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
		if p.Value == nil {
			if p.Required {
				return nilString, nil
			}
			return "", nil
		}
		v, ok := p.Value.([]interface{})
		if !ok {
			stringValues, ok := p.Value.(string)
			if !ok {
				boolValue, ok := p.Value.(bool)
				if !ok {
					return "", &invalidTypeError{p}
				}
				// Translate true to nil and false to an empty list (meaning "no fields")
				if boolValue {
					return nilString, nil
				}
				return "[]interface{}{}", nil
			}
			values := strings.Split(stringValues, ",")
			v = []interface{}{}
			for _, stringValue := range values {
				v = append(v, strings.Trim(stringValue, " "))
			}
		}
		value := "[]string{\n"
		for _, item := range v {
			value += "\"" + item.(string) + "\",\n"
		}
		return value + "}", nil
	case specTypeTime:
		if p.Value == nil {
			if p.Required {
				return nilString, nil
			}
			return "", nil
		}
		v, ok := p.Value.(string)
		if !ok {
			return "", &invalidTypeError{p}
		}
		return fmt.Sprintf("time.ParseDuration(\"%s\")", v), nil
	}
	return "", nil
}
