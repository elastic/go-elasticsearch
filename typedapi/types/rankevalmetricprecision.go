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
// https://github.com/elastic/elasticsearch-specification/tree/ec3159eb31c62611202a4fb157ea88fa6ff78e1a


package types

// RankEvalMetricPrecision type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ec3159eb31c62611202a4fb157ea88fa6ff78e1a/specification/_global/rank_eval/types.ts#L42-L52
type RankEvalMetricPrecision struct {
	// IgnoreUnlabeled Controls how unlabeled documents in the search results are counted. If set to
	// true, unlabeled documents are ignored and neither count as relevant or
	// irrelevant. Set to false (the default), they are treated as irrelevant.
	IgnoreUnlabeled *bool `json:"ignore_unlabeled,omitempty"`
	// K Sets the maximum number of documents retrieved per query. This value will act
	// in place of the usual size parameter in the query.
	K *int `json:"k,omitempty"`
	// RelevantRatingThreshold Sets the rating threshold above which documents are considered to be
	// "relevant".
	RelevantRatingThreshold *int `json:"relevant_rating_threshold,omitempty"`
}

// NewRankEvalMetricPrecision returns a RankEvalMetricPrecision.
func NewRankEvalMetricPrecision() *RankEvalMetricPrecision {
	r := &RankEvalMetricPrecision{}

	return r
}
