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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Package scriptsorttype
package scriptsorttype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/sort.ts#L80-L84
type ScriptSortType struct {
	Name string
}

var (
	String = ScriptSortType{"string"}

	Number = ScriptSortType{"number"}

	Version = ScriptSortType{"version"}
)

func (s ScriptSortType) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *ScriptSortType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "string":
		*s = String
	case "number":
		*s = Number
	case "version":
		*s = Version
	default:
		*s = ScriptSortType{string(text)}
	}

	return nil
}

func (s ScriptSortType) String() string {
	return s.Name
}
