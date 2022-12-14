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

// TopHitsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/_types/aggregations/metric.ts#L171-L184
type TopHitsAggregation struct {
	DocvalueFields   []string               `json:"docvalue_fields,omitempty"`
	Explain          *bool                  `json:"explain,omitempty"`
	Field            *string                `json:"field,omitempty"`
	From             *int                   `json:"from,omitempty"`
	Highlight        *Highlight             `json:"highlight,omitempty"`
	Missing          *Missing               `json:"missing,omitempty"`
	Script           *Script                `json:"script,omitempty"`
	ScriptFields     map[string]ScriptField `json:"script_fields,omitempty"`
	SeqNoPrimaryTerm *bool                  `json:"seq_no_primary_term,omitempty"`
	Size             *int                   `json:"size,omitempty"`
	Sort             []SortCombinations     `json:"sort,omitempty"`
	Source_          *SourceConfig          `json:"_source,omitempty"`
	StoredFields     []string               `json:"stored_fields,omitempty"`
	TrackScores      *bool                  `json:"track_scores,omitempty"`
	Version          *bool                  `json:"version,omitempty"`
}

// NewTopHitsAggregation returns a TopHitsAggregation.
func NewTopHitsAggregation() *TopHitsAggregation {
	r := &TopHitsAggregation{
		ScriptFields: make(map[string]ScriptField, 0),
	}

	return r
}
