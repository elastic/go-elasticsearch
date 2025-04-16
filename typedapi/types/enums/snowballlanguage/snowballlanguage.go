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
// https://github.com/elastic/elasticsearch-specification/tree/f6a370d0fba975752c644fc730f7c45610e28f36

// Package snowballlanguage
package snowballlanguage

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/f6a370d0fba975752c644fc730f7c45610e28f36/specification/_types/analysis/languages.ts#L20-L43
type SnowballLanguage struct {
	Name string
}

var (
	Armenian = SnowballLanguage{"Armenian"}

	Basque = SnowballLanguage{"Basque"}

	Catalan = SnowballLanguage{"Catalan"}

	Danish = SnowballLanguage{"Danish"}

	Dutch = SnowballLanguage{"Dutch"}

	English = SnowballLanguage{"English"}

	Finnish = SnowballLanguage{"Finnish"}

	French = SnowballLanguage{"French"}

	German = SnowballLanguage{"German"}

	German2 = SnowballLanguage{"German2"}

	Hungarian = SnowballLanguage{"Hungarian"}

	Italian = SnowballLanguage{"Italian"}

	Kp = SnowballLanguage{"Kp"}

	Lovins = SnowballLanguage{"Lovins"}

	Norwegian = SnowballLanguage{"Norwegian"}

	Porter = SnowballLanguage{"Porter"}

	Portuguese = SnowballLanguage{"Portuguese"}

	Romanian = SnowballLanguage{"Romanian"}

	Russian = SnowballLanguage{"Russian"}

	Spanish = SnowballLanguage{"Spanish"}

	Swedish = SnowballLanguage{"Swedish"}

	Turkish = SnowballLanguage{"Turkish"}
)

func (s SnowballLanguage) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SnowballLanguage) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "armenian":
		*s = Armenian
	case "basque":
		*s = Basque
	case "catalan":
		*s = Catalan
	case "danish":
		*s = Danish
	case "dutch":
		*s = Dutch
	case "english":
		*s = English
	case "finnish":
		*s = Finnish
	case "french":
		*s = French
	case "german":
		*s = German
	case "german2":
		*s = German2
	case "hungarian":
		*s = Hungarian
	case "italian":
		*s = Italian
	case "kp":
		*s = Kp
	case "lovins":
		*s = Lovins
	case "norwegian":
		*s = Norwegian
	case "porter":
		*s = Porter
	case "portuguese":
		*s = Portuguese
	case "romanian":
		*s = Romanian
	case "russian":
		*s = Russian
	case "spanish":
		*s = Spanish
	case "swedish":
		*s = Swedish
	case "turkish":
		*s = Turkish
	default:
		*s = SnowballLanguage{string(text)}
	}

	return nil
}

func (s SnowballLanguage) String() string {
	return s.Name
}
