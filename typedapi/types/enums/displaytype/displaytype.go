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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package displaytype
package displaytype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/_types/Connector.ts#L35-L41
type DisplayType struct {
	Name string
}

var (
	Textbox = DisplayType{"textbox"}

	Textarea = DisplayType{"textarea"}

	Numeric = DisplayType{"numeric"}

	Toggle = DisplayType{"toggle"}

	Dropdown = DisplayType{"dropdown"}
)

func (d DisplayType) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DisplayType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "textbox":
		*d = Textbox
	case "textarea":
		*d = Textarea
	case "numeric":
		*d = Numeric
	case "toggle":
		*d = Toggle
	case "dropdown":
		*d = Dropdown
	default:
		*d = DisplayType{string(text)}
	}

	return nil
}

func (d DisplayType) String() string {
	return d.Name
}
