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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termvectoroption"
)

// TextProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/mapping/core.ts#L247-L263
type TextProperty struct {
	Analyzer                 *string                        `json:"analyzer,omitempty"`
	Boost                    *float64                       `json:"boost,omitempty"`
	CopyTo                   []string                       `json:"copy_to,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals      *bool                          `json:"eager_global_ordinals,omitempty"`
	Fielddata                *bool                          `json:"fielddata,omitempty"`
	FielddataFrequencyFilter *FielddataFrequencyFilter      `json:"fielddata_frequency_filter,omitempty"`
	Fields                   map[string]Property            `json:"fields,omitempty"`
	IgnoreAbove              *int                           `json:"ignore_above,omitempty"`
	Index                    *bool                          `json:"index,omitempty"`
	IndexOptions             *indexoptions.IndexOptions     `json:"index_options,omitempty"`
	IndexPhrases             *bool                          `json:"index_phrases,omitempty"`
	IndexPrefixes            *TextIndexPrefixes             `json:"index_prefixes,omitempty"`
	// Meta Metadata about the field.
	Meta                 map[string]string                  `json:"meta,omitempty"`
	Norms                *bool                              `json:"norms,omitempty"`
	PositionIncrementGap *int                               `json:"position_increment_gap,omitempty"`
	Properties           map[string]Property                `json:"properties,omitempty"`
	SearchAnalyzer       *string                            `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer  *string                            `json:"search_quote_analyzer,omitempty"`
	Similarity           *string                            `json:"similarity,omitempty"`
	Store                *bool                              `json:"store,omitempty"`
	TermVector           *termvectoroption.TermVectorOption `json:"term_vector,omitempty"`
	Type                 string                             `json:"type,omitempty"`
}

// NewTextProperty returns a TextProperty.
func NewTextProperty() *TextProperty {
	r := &TextProperty{
		Fields:     make(map[string]Property, 0),
		Meta:       make(map[string]string, 0),
		Properties: make(map[string]Property, 0),
	}

	r.Type = "text"

	return r
}
