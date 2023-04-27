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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

// Package synonymformat
package synonymformat

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/analysis/token_filters.ts#L105-L108
type SynonymFormat struct {
	Name string
}

var (
	Solr = SynonymFormat{"solr"}

	Wordnet = SynonymFormat{"wordnet"}
)

func (s SynonymFormat) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SynonymFormat) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "solr":
		*s = Solr
	case "wordnet":
		*s = Wordnet
	default:
		*s = SynonymFormat{string(text)}
	}

	return nil
}

func (s SynonymFormat) String() string {
	return s.Name
}
