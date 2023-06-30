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

// Package phoneticlanguage
package phoneticlanguage

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/analysis/phonetic-plugin.ts#L38-L51
type PhoneticLanguage struct {
	Name string
}

var (
	Any = PhoneticLanguage{"any"}

	Common = PhoneticLanguage{"common"}

	Cyrillic = PhoneticLanguage{"cyrillic"}

	English = PhoneticLanguage{"english"}

	French = PhoneticLanguage{"french"}

	German = PhoneticLanguage{"german"}

	Hebrew = PhoneticLanguage{"hebrew"}

	Hungarian = PhoneticLanguage{"hungarian"}

	Polish = PhoneticLanguage{"polish"}

	Romanian = PhoneticLanguage{"romanian"}

	Russian = PhoneticLanguage{"russian"}

	Spanish = PhoneticLanguage{"spanish"}
)

func (p PhoneticLanguage) MarshalText() (text []byte, err error) {
	return []byte(p.String()), nil
}

func (p *PhoneticLanguage) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "any":
		*p = Any
	case "common":
		*p = Common
	case "cyrillic":
		*p = Cyrillic
	case "english":
		*p = English
	case "french":
		*p = French
	case "german":
		*p = German
	case "hebrew":
		*p = Hebrew
	case "hungarian":
		*p = Hungarian
	case "polish":
		*p = Polish
	case "romanian":
		*p = Romanian
	case "russian":
		*p = Russian
	case "spanish":
		*p = Spanish
	default:
		*p = PhoneticLanguage{string(text)}
	}

	return nil
}

func (p PhoneticLanguage) String() string {
	return p.Name
}
