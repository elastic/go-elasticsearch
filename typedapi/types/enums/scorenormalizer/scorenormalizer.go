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
// https://github.com/elastic/elasticsearch-specification/tree/3a94b6715915b1e9311724a2614c643368eece90

// Package scorenormalizer
package scorenormalizer

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3a94b6715915b1e9311724a2614c643368eece90/specification/_types/Retriever.ts#L88-L91
type ScoreNormalizer struct {
	Name string
}

var (
	None = ScoreNormalizer{"none"}

	Minmax = ScoreNormalizer{"minmax"}
)

func (s ScoreNormalizer) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *ScoreNormalizer) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "none":
		*s = None
	case "minmax":
		*s = Minmax
	default:
		*s = ScoreNormalizer{string(text)}
	}

	return nil
}

func (s ScoreNormalizer) String() string {
	return s.Name
}
