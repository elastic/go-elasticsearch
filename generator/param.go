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

type param struct {
	Name        string
	SpecType    string `json:"type"`
	Type        string
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Default     interface{} `json:"default"`
	Options     []string    `json:"options"`
	OptionName  string
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
		// TODO: implement
		p.Type = "struct{}"
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
	description := strings.Split(p.Description, " ")
	firstWord := description[0]
	// If the first word is all upper case let it be, otherwise convert to lower
	if strings.ToUpper(firstWord) != firstWord {
		firstWord = strings.ToLower(string(firstWord[0])) + firstWord[1:]
	}
	p.Description = firstWord + " " + strings.Replace(strings.Join(description[1:], " "), "`", "", -1)
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
