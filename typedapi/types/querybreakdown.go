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

// QueryBreakdown type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_global/search/_types/profile.ts#L97-L116
type QueryBreakdown struct {
	Advance                     int64 `json:"advance"`
	AdvanceCount                int64 `json:"advance_count"`
	BuildScorer                 int64 `json:"build_scorer"`
	BuildScorerCount            int64 `json:"build_scorer_count"`
	ComputeMaxScore             int64 `json:"compute_max_score"`
	ComputeMaxScoreCount        int64 `json:"compute_max_score_count"`
	CreateWeight                int64 `json:"create_weight"`
	CreateWeightCount           int64 `json:"create_weight_count"`
	Match                       int64 `json:"match"`
	MatchCount                  int64 `json:"match_count"`
	NextDoc                     int64 `json:"next_doc"`
	NextDocCount                int64 `json:"next_doc_count"`
	Score                       int64 `json:"score"`
	ScoreCount                  int64 `json:"score_count"`
	SetMinCompetitiveScore      int64 `json:"set_min_competitive_score"`
	SetMinCompetitiveScoreCount int64 `json:"set_min_competitive_score_count"`
	ShallowAdvance              int64 `json:"shallow_advance"`
	ShallowAdvanceCount         int64 `json:"shallow_advance_count"`
}

// NewQueryBreakdown returns a QueryBreakdown.
func NewQueryBreakdown() *QueryBreakdown {
	r := &QueryBreakdown{}

	return r
}
