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

// Package gappolicy
package gappolicy

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/aggregations/pipeline.ts#L52-L67
type GapPolicy struct {
	Name string
}

var (
	Skip = GapPolicy{"skip"}

	Insertzeros = GapPolicy{"insert_zeros"}

	Keepvalues = GapPolicy{"keep_values"}
)

func (g GapPolicy) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GapPolicy) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "skip":
		*g = Skip
	case "insert_zeros":
		*g = Insertzeros
	case "keep_values":
		*g = Keepvalues
	default:
		*g = GapPolicy{string(text)}
	}

	return nil
}

func (g GapPolicy) String() string {
	return g.Name
}
