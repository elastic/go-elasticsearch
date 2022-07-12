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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// RankEvalMetricRecall type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_global/rank_eval/types.ts#L54-L58
type RankEvalMetricRecall struct {
	// K Sets the maximum number of documents retrieved per query. This value will act
	// in place of the usual size parameter in the query.
	K *int `json:"k,omitempty"`
	// RelevantRatingThreshold Sets the rating threshold above which documents are considered to be
	// "relevant".
	RelevantRatingThreshold *int `json:"relevant_rating_threshold,omitempty"`
}

// RankEvalMetricRecallBuilder holds RankEvalMetricRecall struct and provides a builder API.
type RankEvalMetricRecallBuilder struct {
	v *RankEvalMetricRecall
}

// NewRankEvalMetricRecall provides a builder for the RankEvalMetricRecall struct.
func NewRankEvalMetricRecallBuilder() *RankEvalMetricRecallBuilder {
	r := RankEvalMetricRecallBuilder{
		&RankEvalMetricRecall{},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalMetricRecall struct
func (rb *RankEvalMetricRecallBuilder) Build() RankEvalMetricRecall {
	return *rb.v
}

// K Sets the maximum number of documents retrieved per query. This value will act
// in place of the usual size parameter in the query.

func (rb *RankEvalMetricRecallBuilder) K(k int) *RankEvalMetricRecallBuilder {
	rb.v.K = &k
	return rb
}

// RelevantRatingThreshold Sets the rating threshold above which documents are considered to be
// "relevant".

func (rb *RankEvalMetricRecallBuilder) RelevantRatingThreshold(relevantratingthreshold int) *RankEvalMetricRecallBuilder {
	rb.v.RelevantRatingThreshold = &relevantratingthreshold
	return rb
}
