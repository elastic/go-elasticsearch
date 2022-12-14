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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termvectoroption"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"
)

// DynamicProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/mapping/core.ts#L275-L306
type DynamicProperty struct {
	Analyzer            *string                        `json:"analyzer,omitempty"`
	Boost               *float64                       `json:"boost,omitempty"`
	Coerce              *bool                          `json:"coerce,omitempty"`
	CopyTo              []string                       `json:"copy_to,omitempty"`
	DocValues           *bool                          `json:"doc_values,omitempty"`
	Dynamic             *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals *bool                          `json:"eager_global_ordinals,omitempty"`
	Enabled             *bool                          `json:"enabled,omitempty"`
	Fields              map[string]Property            `json:"fields,omitempty"`
	Format              *string                        `json:"format,omitempty"`
	IgnoreAbove         *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed     *bool                          `json:"ignore_malformed,omitempty"`
	Index               *bool                          `json:"index,omitempty"`
	IndexOptions        *indexoptions.IndexOptions     `json:"index_options,omitempty"`
	IndexPhrases        *bool                          `json:"index_phrases,omitempty"`
	IndexPrefixes       *TextIndexPrefixes             `json:"index_prefixes,omitempty"`
	Locale              *string                        `json:"locale,omitempty"`
	// Meta Metadata about the field.
	Meta                 map[string]string                          `json:"meta,omitempty"`
	Norms                *bool                                      `json:"norms,omitempty"`
	NullValue            *FieldValue                                `json:"null_value,omitempty"`
	OnScriptError        *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	PositionIncrementGap *int                                       `json:"position_increment_gap,omitempty"`
	PrecisionStep        *int                                       `json:"precision_step,omitempty"`
	Properties           map[string]Property                        `json:"properties,omitempty"`
	Script               *Script                                    `json:"script,omitempty"`
	SearchAnalyzer       *string                                    `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer  *string                                    `json:"search_quote_analyzer,omitempty"`
	Similarity           *string                                    `json:"similarity,omitempty"`
	Store                *bool                                      `json:"store,omitempty"`
	TermVector           *termvectoroption.TermVectorOption         `json:"term_vector,omitempty"`
	TimeSeriesMetric     *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type                 string                                     `json:"type,omitempty"`
}

// NewDynamicProperty returns a DynamicProperty.
func NewDynamicProperty() *DynamicProperty {
	r := &DynamicProperty{
		Fields:     make(map[string]Property, 0),
		Meta:       make(map[string]string, 0),
		Properties: make(map[string]Property, 0),
	}

	r.Type = "{dynamic_property}"

	return r
}
