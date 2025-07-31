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

// Package elserservicetype
package elserservicetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L1301-L1303
type ElserServiceType struct {
	Name string
}

var (
	Elser = ElserServiceType{"elser"}
)

func (e ElserServiceType) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *ElserServiceType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "elser":
		*e = Elser
	default:
		*e = ElserServiceType{string(text)}
	}

	return nil
}

func (e ElserServiceType) String() string {
	return e.Name
}
