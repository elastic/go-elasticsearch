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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Package phoneticnametype
package phoneticnametype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/analysis/phonetic-plugin.ts#L53-L57
type PhoneticNameType struct {
	Name string
}

var (
	Generic = PhoneticNameType{"generic"}

	Ashkenazi = PhoneticNameType{"ashkenazi"}

	Sephardic = PhoneticNameType{"sephardic"}
)

func (p PhoneticNameType) MarshalText() (text []byte, err error) {
	return []byte(p.String()), nil
}

func (p *PhoneticNameType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "generic":
		*p = Generic
	case "ashkenazi":
		*p = Ashkenazi
	case "sephardic":
		*p = Sephardic
	default:
		*p = PhoneticNameType{string(text)}
	}

	return nil
}

func (p PhoneticNameType) String() string {
	return p.Name
}
