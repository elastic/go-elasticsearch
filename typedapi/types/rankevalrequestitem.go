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
)

// RankEvalRequestItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_global/rank_eval/types.ts#L98-L109
type RankEvalRequestItem struct {
	// Id The search request’s ID, used to group result details later.
	Id string `json:"id"`
	// Params The search template parameters.
	Params map[string]json.RawMessage `json:"params,omitempty"`
	// Ratings List of document ratings
	Ratings []DocumentRating `json:"ratings"`
	// Request The query being evaluated.
	Request *RankEvalQuery `json:"request,omitempty"`
	// TemplateId The search template Id
	TemplateId *string `json:"template_id,omitempty"`
}

// NewRankEvalRequestItem returns a RankEvalRequestItem.
func NewRankEvalRequestItem() *RankEvalRequestItem {
	r := &RankEvalRequestItem{
		Params: make(map[string]json.RawMessage, 0),
	}

	return r
}
