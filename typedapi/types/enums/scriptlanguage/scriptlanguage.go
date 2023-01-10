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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Package scriptlanguage
package scriptlanguage

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/Scripting.ts#L24-L33
type ScriptLanguage struct {
	Name string
}

var (
	Painless = ScriptLanguage{"painless"}

	Expression = ScriptLanguage{"expression"}

	Mustache = ScriptLanguage{"mustache"}

	Java = ScriptLanguage{"java"}
)

func (s ScriptLanguage) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *ScriptLanguage) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "painless":
		*s = Painless
	case "expression":
		*s = Expression
	case "mustache":
		*s = Mustache
	case "java":
		*s = Java
	default:
		*s = ScriptLanguage{string(text)}
	}

	return nil
}

func (s ScriptLanguage) String() string {
	return s.Name
}
