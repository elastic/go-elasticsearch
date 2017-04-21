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
	switch name {
	case "type":
		p.Name = "DocumentType"
	default:
		p.Name = snaker.SnakeToCamel(name)
	}
	if !p.Required {
		p.OptionName = "With" + p.Name
	}
	p.Name = strings.ToLower(string(p.Name[0])) + p.Name[1:]
	switch p.SpecType {
	case "boolean":
		p.Type = "bool"
	case "enum":
		// TODO: implement
	case "list":
		p.Type = "[]string"
	case "number":
		p.Type = "int"
	case "string":
		p.Type = "string"
	case "time":
		p.Type = "time.Time"
	default:
		return fmt.Errorf("Invalid type for %s: %s", name, p.Type)
	}
	return nil
}
