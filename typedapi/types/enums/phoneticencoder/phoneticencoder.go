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

// Package phoneticencoder
package phoneticencoder

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/analysis/phonetic-plugin.ts#L23-L36
type PhoneticEncoder struct {
	Name string
}

var (
	Metaphone = PhoneticEncoder{"metaphone"}

	Doublemetaphone = PhoneticEncoder{"double_metaphone"}

	Soundex = PhoneticEncoder{"soundex"}

	Refinedsoundex = PhoneticEncoder{"refined_soundex"}

	Caverphone1 = PhoneticEncoder{"caverphone1"}

	Caverphone2 = PhoneticEncoder{"caverphone2"}

	Cologne = PhoneticEncoder{"cologne"}

	Nysiis = PhoneticEncoder{"nysiis"}

	Koelnerphonetik = PhoneticEncoder{"koelnerphonetik"}

	Haasephonetik = PhoneticEncoder{"haasephonetik"}

	Beidermorse = PhoneticEncoder{"beider_morse"}

	Daitchmokotoff = PhoneticEncoder{"daitch_mokotoff"}
)

func (p PhoneticEncoder) MarshalText() (text []byte, err error) {
	return []byte(p.String()), nil
}

func (p *PhoneticEncoder) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "metaphone":
		*p = Metaphone
	case "double_metaphone":
		*p = Doublemetaphone
	case "soundex":
		*p = Soundex
	case "refined_soundex":
		*p = Refinedsoundex
	case "caverphone1":
		*p = Caverphone1
	case "caverphone2":
		*p = Caverphone2
	case "cologne":
		*p = Cologne
	case "nysiis":
		*p = Nysiis
	case "koelnerphonetik":
		*p = Koelnerphonetik
	case "haasephonetik":
		*p = Haasephonetik
	case "beider_morse":
		*p = Beidermorse
	case "daitch_mokotoff":
		*p = Daitchmokotoff
	default:
		*p = PhoneticEncoder{string(text)}
	}

	return nil
}

func (p PhoneticEncoder) String() string {
	return p.Name
}
