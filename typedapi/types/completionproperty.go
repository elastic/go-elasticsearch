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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// CompletionProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/mapping/specialized.ts#L27-L35
type CompletionProperty struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Contexts       []SuggestContext               `json:"contexts,omitempty"`
	CopyTo         []string                       `json:"copy_to,omitempty"`
	DocValues      *bool                          `json:"doc_values,omitempty"`
	Dynamic        *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields         map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove    *int                           `json:"ignore_above,omitempty"`
	MaxInputLength *int                           `json:"max_input_length,omitempty"`
	// Meta Metadata about the field.
	Meta                       map[string]string   `json:"meta,omitempty"`
	PreservePositionIncrements *bool               `json:"preserve_position_increments,omitempty"`
	PreserveSeparators         *bool               `json:"preserve_separators,omitempty"`
	Properties                 map[string]Property `json:"properties,omitempty"`
	SearchAnalyzer             *string             `json:"search_analyzer,omitempty"`
	Similarity                 *string             `json:"similarity,omitempty"`
	Store                      *bool               `json:"store,omitempty"`
	Type                       string              `json:"type,omitempty"`
}

// NewCompletionProperty returns a CompletionProperty.
func NewCompletionProperty() *CompletionProperty {
	r := &CompletionProperty{
		Fields:     make(map[string]Property, 0),
		Meta:       make(map[string]string, 0),
		Properties: make(map[string]Property, 0),
	}

	r.Type = "completion"

	return r
}
