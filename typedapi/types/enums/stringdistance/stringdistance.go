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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Package stringdistance
package stringdistance

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/suggester.ts#L473-L494
type StringDistance struct {
	Name string
}

var (
	Internal = StringDistance{"internal"}

	Dameraulevenshtein = StringDistance{"damerau_levenshtein"}

	Levenshtein = StringDistance{"levenshtein"}

	Jarowinkler = StringDistance{"jaro_winkler"}

	Ngram = StringDistance{"ngram"}
)

func (s StringDistance) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *StringDistance) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "internal":
		*s = Internal
	case "damerau_levenshtein":
		*s = Dameraulevenshtein
	case "levenshtein":
		*s = Levenshtein
	case "jaro_winkler":
		*s = Jarowinkler
	case "ngram":
		*s = Ngram
	default:
		*s = StringDistance{string(text)}
	}

	return nil
}

func (s StringDistance) String() string {
	return s.Name
}
