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

// Package cohereinputtype
package cohereinputtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/f6a370d0fba975752c644fc730f7c45610e28f36/specification/inference/_types/CommonTypes.ts#L631-L636
type CohereInputType struct {
	Name string
}

var (
	Classification = CohereInputType{"classification"}

	Clustering = CohereInputType{"clustering"}

	Ingest = CohereInputType{"ingest"}

	Search = CohereInputType{"search"}
)

func (c CohereInputType) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CohereInputType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "classification":
		*c = Classification
	case "clustering":
		*c = Clustering
	case "ingest":
		*c = Ingest
	case "search":
		*c = Search
	default:
		*c = CohereInputType{string(text)}
	}

	return nil
}

func (c CohereInputType) String() string {
	return c.Name
}
