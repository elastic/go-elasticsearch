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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// RankEvalMetricPrecision type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/rank_eval/types.ts#L42-L52
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

// RankEvalMetricPrecisionBuilder holds RankEvalMetricPrecision struct and provides a builder API.
type RankEvalMetricPrecisionBuilder struct {
	v *RankEvalMetricPrecision
}

// NewRankEvalMetricPrecision provides a builder for the RankEvalMetricPrecision struct.
func NewRankEvalMetricPrecisionBuilder() *RankEvalMetricPrecisionBuilder {
	r := RankEvalMetricPrecisionBuilder{
		&RankEvalMetricPrecision{},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalMetricPrecision struct
func (rb *RankEvalMetricPrecisionBuilder) Build() RankEvalMetricPrecision {
	return *rb.v
}

// IgnoreUnlabeled Controls how unlabeled documents in the search results are counted. If set to
// true, unlabeled documents are ignored and neither count as relevant or
// irrelevant. Set to false (the default), they are treated as irrelevant.

func (rb *RankEvalMetricPrecisionBuilder) IgnoreUnlabeled(ignoreunlabeled bool) *RankEvalMetricPrecisionBuilder {
	rb.v.IgnoreUnlabeled = &ignoreunlabeled
	return rb
}

// K Sets the maximum number of documents retrieved per query. This value will act
// in place of the usual size parameter in the query.

func (rb *RankEvalMetricPrecisionBuilder) K(k int) *RankEvalMetricPrecisionBuilder {
	rb.v.K = &k
	return rb
}

// RelevantRatingThreshold Sets the rating threshold above which documents are considered to be
// "relevant".

func (rb *RankEvalMetricPrecisionBuilder) RelevantRatingThreshold(relevantratingthreshold int) *RankEvalMetricPrecisionBuilder {
	rb.v.RelevantRatingThreshold = &relevantratingthreshold
	return rb
}
