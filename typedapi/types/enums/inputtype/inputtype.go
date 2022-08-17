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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


// Package inputtype
package inputtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Input.ts#L102-L106
type InputType struct {
	name string
}

var (
	Http = InputType{"http"}

	Search = InputType{"search"}

	Simple = InputType{"simple"}
)

func (i InputType) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *InputType) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "http":
		*i = Http
	case "search":
		*i = Search
	case "simple":
		*i = Simple
	default:
		*i = InputType{string(text)}
	}

	return nil
}

func (i InputType) String() string {
	return i.name
}
