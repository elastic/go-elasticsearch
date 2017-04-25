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
	"strings"

	"github.com/golang/glog"
	"github.com/serenize/snaker"
)

type enum struct {
	Name     string
	Type     string
	SpecName string
}

type param struct {
	Name        string
	SpecType    string `json:"type"`
	Type        string
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Default     interface{} `json:"default"`
	Options     []string    `json:"options"`
	OptionName  string
	EnumValues  []*enum
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

func (p *param) resolve(name string) error {
	if strings.HasPrefix(name, "_") {
		name = name[1:]
	}
	switch name {
	case "type":
		p.Name = "documentType"
	default:
		p.Name = snaker.SnakeToCamelLower(name)
	}
	if !p.Required {
		p.OptionName = "With" + snaker.SnakeToCamel(name)
	}
	switch p.SpecType {
	case "":
		glog.Warningf("%s has no type", name)
	case "boolean":
		p.Type = "bool"
	case "enum":
		p.Type = snaker.SnakeToCamel(p.Name)
		p.EnumValues = []*enum{}
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
		}
	case "list":
		p.Type = "[]string"
	case "number":
		p.Type = "int"
	case "string":
		p.Type = "string"
	case "time":
		p.Type = "time.Time"
	default:
		return fmt.Errorf("Invalid type for %s: %s", name, p.SpecType)
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
