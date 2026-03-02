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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

// Package stopwordlanguage
package stopwordlanguage

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/e196f9953fa743572ee46884835f1934bce9a16b/specification/_types/analysis/StopWords.ts#L20-L58
type StopWordLanguage struct {
	Name string
}

var (
	Arabic = StopWordLanguage{"_arabic_"}

	Armenian = StopWordLanguage{"_armenian_"}

	Basque = StopWordLanguage{"_basque_"}

	Bengali = StopWordLanguage{"_bengali_"}

	Brazilian = StopWordLanguage{"_brazilian_"}

	Bulgarian = StopWordLanguage{"_bulgarian_"}

	Catalan = StopWordLanguage{"_catalan_"}

	Cjk = StopWordLanguage{"_cjk_"}

	Czech = StopWordLanguage{"_czech_"}

	Danish = StopWordLanguage{"_danish_"}

	Dutch = StopWordLanguage{"_dutch_"}

	English = StopWordLanguage{"_english_"}

	Estonian = StopWordLanguage{"_estonian_"}

	Finnish = StopWordLanguage{"_finnish_"}

	French = StopWordLanguage{"_french_"}

	Galician = StopWordLanguage{"_galician_"}

	German = StopWordLanguage{"_german_"}

	Greek = StopWordLanguage{"_greek_"}

	Hindi = StopWordLanguage{"_hindi_"}

	Hungarian = StopWordLanguage{"_hungarian_"}

	Indonesian = StopWordLanguage{"_indonesian_"}

	Irish = StopWordLanguage{"_irish_"}

	Italian = StopWordLanguage{"_italian_"}

	Latvian = StopWordLanguage{"_latvian_"}

	Lithuanian = StopWordLanguage{"_lithuanian_"}

	Norwegian = StopWordLanguage{"_norwegian_"}

	Persian = StopWordLanguage{"_persian_"}

	Portuguese = StopWordLanguage{"_portuguese_"}

	Romanian = StopWordLanguage{"_romanian_"}

	Russian = StopWordLanguage{"_russian_"}

	Serbian = StopWordLanguage{"_serbian_"}

	Sorani = StopWordLanguage{"_sorani_"}

	Spanish = StopWordLanguage{"_spanish_"}

	Swedish = StopWordLanguage{"_swedish_"}

	Thai = StopWordLanguage{"_thai_"}

	Turkish = StopWordLanguage{"_turkish_"}

	None = StopWordLanguage{"_none_"}
)

func (s StopWordLanguage) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *StopWordLanguage) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "_arabic_":
		*s = Arabic
	case "_armenian_":
		*s = Armenian
	case "_basque_":
		*s = Basque
	case "_bengali_":
		*s = Bengali
	case "_brazilian_":
		*s = Brazilian
	case "_bulgarian_":
		*s = Bulgarian
	case "_catalan_":
		*s = Catalan
	case "_cjk_":
		*s = Cjk
	case "_czech_":
		*s = Czech
	case "_danish_":
		*s = Danish
	case "_dutch_":
		*s = Dutch
	case "_english_":
		*s = English
	case "_estonian_":
		*s = Estonian
	case "_finnish_":
		*s = Finnish
	case "_french_":
		*s = French
	case "_galician_":
		*s = Galician
	case "_german_":
		*s = German
	case "_greek_":
		*s = Greek
	case "_hindi_":
		*s = Hindi
	case "_hungarian_":
		*s = Hungarian
	case "_indonesian_":
		*s = Indonesian
	case "_irish_":
		*s = Irish
	case "_italian_":
		*s = Italian
	case "_latvian_":
		*s = Latvian
	case "_lithuanian_":
		*s = Lithuanian
	case "_norwegian_":
		*s = Norwegian
	case "_persian_":
		*s = Persian
	case "_portuguese_":
		*s = Portuguese
	case "_romanian_":
		*s = Romanian
	case "_russian_":
		*s = Russian
	case "_serbian_":
		*s = Serbian
	case "_sorani_":
		*s = Sorani
	case "_spanish_":
		*s = Spanish
	case "_swedish_":
		*s = Swedish
	case "_thai_":
		*s = Thai
	case "_turkish_":
		*s = Turkish
	case "_none_":
		*s = None
	default:
		*s = StopWordLanguage{string(text)}
	}

	return nil
}

func (s StopWordLanguage) String() string {
	return s.Name
}
