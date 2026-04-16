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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

// Package customserviceinputtype
package customserviceinputtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/836fca874204ca4173ae5c36fb6b5107d28d2fc0/specification/inference/_types/CommonTypes.ts#L1173-L1178
type CustomServiceInputType struct {
	Name string
}

var (
	Classification = CustomServiceInputType{"classification"}

	Clustering = CustomServiceInputType{"clustering"}

	Ingest = CustomServiceInputType{"ingest"}

	Search = CustomServiceInputType{"search"}
)

func (c CustomServiceInputType) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *CustomServiceInputType) UnmarshalText(text []byte) error {
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
		*c = CustomServiceInputType{string(text)}
	}

	return nil
}

func (c CustomServiceInputType) String() string {
	return c.Name
}
