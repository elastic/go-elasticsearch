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

// Package lowercasetokenfilterlanguages
package lowercasetokenfilterlanguages

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L344-L348
type LowercaseTokenFilterLanguages struct {
	Name string
}

var (
	Greek = LowercaseTokenFilterLanguages{"greek"}

	Irish = LowercaseTokenFilterLanguages{"irish"}

	Turkish = LowercaseTokenFilterLanguages{"turkish"}
)

func (l LowercaseTokenFilterLanguages) MarshalText() (text []byte, err error) {
	return []byte(l.String()), nil
}

func (l *LowercaseTokenFilterLanguages) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "greek":
		*l = Greek
	case "irish":
		*l = Irish
	case "turkish":
		*l = Turkish
	default:
		*l = LowercaseTokenFilterLanguages{string(text)}
	}

	return nil
}

func (l LowercaseTokenFilterLanguages) String() string {
	return l.Name
}
