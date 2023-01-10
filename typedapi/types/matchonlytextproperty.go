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


package types

// MatchOnlyTextProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/mapping/core.ts#L208-L233
type MatchOnlyTextProperty struct {
	// CopyTo Allows you to copy the values of multiple fields into a group
	// field, which can then be queried as a single field.
	CopyTo []string `json:"copy_to,omitempty"`
	// Fields Multi-fields allow the same string value to be indexed in multiple ways for
	// different purposes, such as one
	// field for search and a multi-field for sorting and aggregations, or the same
	// string value analyzed by different analyzers.
	Fields map[string]Property `json:"fields,omitempty"`
	// Meta Metadata about the field.
	Meta map[string]string `json:"meta,omitempty"`
	Type string            `json:"type,omitempty"`
}

// NewMatchOnlyTextProperty returns a MatchOnlyTextProperty.
func NewMatchOnlyTextProperty() *MatchOnlyTextProperty {
	r := &MatchOnlyTextProperty{
		Fields: make(map[string]Property, 0),
		Meta:   make(map[string]string, 0),
	}

	r.Type = "match_only_text"

	return r
}
