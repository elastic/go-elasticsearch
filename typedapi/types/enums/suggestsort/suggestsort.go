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

// Package suggestsort
package suggestsort

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_global/search/_types/suggester.ts#L496-L505
type SuggestSort struct {
	Name string
}

var (
	Score = SuggestSort{"score"}

	Frequency = SuggestSort{"frequency"}
)

func (s SuggestSort) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SuggestSort) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "score":
		*s = Score
	case "frequency":
		*s = Frequency
	default:
		*s = SuggestSort{string(text)}
	}

	return nil
}

func (s SuggestSort) String() string {
	return s.Name
}
