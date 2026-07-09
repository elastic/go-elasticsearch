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
// https://github.com/elastic/elasticsearch-specification/tree/17fab0b2c19030e59a2eac3e1dab8fba7a8acb8f

// Package nvidiainputtype
package nvidiainputtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/17fab0b2c19030e59a2eac3e1dab8fba7a8acb8f/specification/inference/_types/CommonTypes.ts#L1885-L1888
type NvidiaInputType struct {
	Name string
}

var (
	Ingest = NvidiaInputType{"ingest"}

	Search = NvidiaInputType{"search"}
)

func (n NvidiaInputType) MarshalText() (text []byte, err error) {
	return []byte(n.String()), nil
}

func (n *NvidiaInputType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "ingest":
		*n = Ingest
	case "search":
		*n = Search
	default:
		*n = NvidiaInputType{string(text)}
	}

	return nil
}

func (n NvidiaInputType) String() string {
	return n.Name
}
