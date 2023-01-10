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


// Package language
package language

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/analysis/languages.ts#L20-L55
type Language struct {
	Name string
}

var (
	Arabic = Language{"Arabic"}

	Armenian = Language{"Armenian"}

	Basque = Language{"Basque"}

	Brazilian = Language{"Brazilian"}

	Bulgarian = Language{"Bulgarian"}

	Catalan = Language{"Catalan"}

	Chinese = Language{"Chinese"}

	Cjk = Language{"Cjk"}

	Czech = Language{"Czech"}

	Danish = Language{"Danish"}

	Dutch = Language{"Dutch"}

	English = Language{"English"}

	Estonian = Language{"Estonian"}

	Finnish = Language{"Finnish"}

	French = Language{"French"}

	Galician = Language{"Galician"}

	German = Language{"German"}

	Greek = Language{"Greek"}

	Hindi = Language{"Hindi"}

	Hungarian = Language{"Hungarian"}

	Indonesian = Language{"Indonesian"}

	Irish = Language{"Irish"}

	Italian = Language{"Italian"}

	Latvian = Language{"Latvian"}

	Norwegian = Language{"Norwegian"}

	Persian = Language{"Persian"}

	Portuguese = Language{"Portuguese"}

	Romanian = Language{"Romanian"}

	Russian = Language{"Russian"}

	Sorani = Language{"Sorani"}

	Spanish = Language{"Spanish"}

	Swedish = Language{"Swedish"}

	Turkish = Language{"Turkish"}

	Thai = Language{"Thai"}
)

func (l Language) MarshalText() (text []byte, err error) {
	return []byte(l.String()), nil
}

func (l *Language) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "Arabic":
		*l = Arabic
	case "Armenian":
		*l = Armenian
	case "Basque":
		*l = Basque
	case "Brazilian":
		*l = Brazilian
	case "Bulgarian":
		*l = Bulgarian
	case "Catalan":
		*l = Catalan
	case "Chinese":
		*l = Chinese
	case "Cjk":
		*l = Cjk
	case "Czech":
		*l = Czech
	case "Danish":
		*l = Danish
	case "Dutch":
		*l = Dutch
	case "English":
		*l = English
	case "Estonian":
		*l = Estonian
	case "Finnish":
		*l = Finnish
	case "French":
		*l = French
	case "Galician":
		*l = Galician
	case "German":
		*l = German
	case "Greek":
		*l = Greek
	case "Hindi":
		*l = Hindi
	case "Hungarian":
		*l = Hungarian
	case "Indonesian":
		*l = Indonesian
	case "Irish":
		*l = Irish
	case "Italian":
		*l = Italian
	case "Latvian":
		*l = Latvian
	case "Norwegian":
		*l = Norwegian
	case "Persian":
		*l = Persian
	case "Portuguese":
		*l = Portuguese
	case "Romanian":
		*l = Romanian
	case "Russian":
		*l = Russian
	case "Sorani":
		*l = Sorani
	case "Spanish":
		*l = Spanish
	case "Swedish":
		*l = Swedish
	case "Turkish":
		*l = Turkish
	case "Thai":
		*l = Thai
	default:
		*l = Language{string(text)}
	}

	return nil
}

func (l Language) String() string {
	return l.Name
}
