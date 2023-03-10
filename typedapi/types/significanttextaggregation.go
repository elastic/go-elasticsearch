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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

// SignificantTextAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/aggregations/bucket.ts#L360-L378
type SignificantTextAggregation struct {
	BackgroundFilter    *Query                                                       `json:"background_filter,omitempty"`
	ChiSquare           *ChiSquareHeuristic                                          `json:"chi_square,omitempty"`
	Exclude             []string                                                     `json:"exclude,omitempty"`
	ExecutionHint       *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field               *string                                                      `json:"field,omitempty"`
	FilterDuplicateText *bool                                                        `json:"filter_duplicate_text,omitempty"`
	Gnd                 *GoogleNormalizedDistanceHeuristic                           `json:"gnd,omitempty"`
	Include             []string                                                     `json:"include,omitempty"`
	Jlh                 *EmptyObject                                                 `json:"jlh,omitempty"`
	Meta                map[string]json.RawMessage                                   `json:"meta,omitempty"`
	MinDocCount         *int64                                                       `json:"min_doc_count,omitempty"`
	MutualInformation   *MutualInformationHeuristic                                  `json:"mutual_information,omitempty"`
	Name                *string                                                      `json:"name,omitempty"`
	Percentage          *PercentageScoreHeuristic                                    `json:"percentage,omitempty"`
	ScriptHeuristic     *ScriptedHeuristic                                           `json:"script_heuristic,omitempty"`
	ShardMinDocCount    *int64                                                       `json:"shard_min_doc_count,omitempty"`
	ShardSize           *int                                                         `json:"shard_size,omitempty"`
	Size                *int                                                         `json:"size,omitempty"`
	SourceFields        []string                                                     `json:"source_fields,omitempty"`
}

// NewSignificantTextAggregation returns a SignificantTextAggregation.
func NewSignificantTextAggregation() *SignificantTextAggregation {
	r := &SignificantTextAggregation{}

	return r
}
