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
// https://github.com/elastic/elasticsearch-specification/tree/9fcf6a64c550d2e8090c8134867f200b56fd7fc7

// Package inputtype
package inputtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/9fcf6a64c550d2e8090c8134867f200b56fd7fc7/specification/watcher/_types/Input.ts#L100-L107
type InputType struct {
	Name string
}

var (
	Chain = InputType{"chain"}

	Http = InputType{"http"}

	None = InputType{"none"}

	Search = InputType{"search"}

	Simple = InputType{"simple"}

	Transform = InputType{"transform"}
)

func (i InputType) MarshalText() (text []byte, err error) {
	return []byte(i.String()), nil
}

func (i *InputType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "chain":
		*i = Chain
	case "http":
		*i = Http
	case "none":
		*i = None
	case "search":
		*i = Search
	case "simple":
		*i = Simple
	case "transform":
		*i = Transform
	default:
		*i = InputType{string(text)}
	}

	return nil
}

func (i InputType) String() string {
	return i.Name
}
