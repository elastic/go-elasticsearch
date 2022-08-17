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

// RankEvalMetricRatingTreshold type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/rank_eval/types.ts#L34-L40
type RankEvalMetricRatingTreshold struct {
	// K Sets the maximum number of documents retrieved per query. This value will act
	// in place of the usual size parameter in the query.
	K *int `json:"k,omitempty"`
	// RelevantRatingThreshold Sets the rating threshold above which documents are considered to be
	// "relevant".
	RelevantRatingThreshold *int `json:"relevant_rating_threshold,omitempty"`
}

// RankEvalMetricRatingTresholdBuilder holds RankEvalMetricRatingTreshold struct and provides a builder API.
type RankEvalMetricRatingTresholdBuilder struct {
	v *RankEvalMetricRatingTreshold
}

// NewRankEvalMetricRatingTreshold provides a builder for the RankEvalMetricRatingTreshold struct.
func NewRankEvalMetricRatingTresholdBuilder() *RankEvalMetricRatingTresholdBuilder {
	r := RankEvalMetricRatingTresholdBuilder{
		&RankEvalMetricRatingTreshold{},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalMetricRatingTreshold struct
func (rb *RankEvalMetricRatingTresholdBuilder) Build() RankEvalMetricRatingTreshold {
	return *rb.v
}

// K Sets the maximum number of documents retrieved per query. This value will act
// in place of the usual size parameter in the query.

func (rb *RankEvalMetricRatingTresholdBuilder) K(k int) *RankEvalMetricRatingTresholdBuilder {
	rb.v.K = &k
	return rb
}

// RelevantRatingThreshold Sets the rating threshold above which documents are considered to be
// "relevant".

func (rb *RankEvalMetricRatingTresholdBuilder) RelevantRatingThreshold(relevantratingthreshold int) *RankEvalMetricRatingTresholdBuilder {
	rb.v.RelevantRatingThreshold = &relevantratingthreshold
	return rb
}
