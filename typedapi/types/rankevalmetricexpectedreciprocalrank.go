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

// RankEvalMetricExpectedReciprocalRank type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/rank_eval/types.ts#L79-L88
type RankEvalMetricExpectedReciprocalRank struct {
	// K Sets the maximum number of documents retrieved per query. This value will act
	// in place of the usual size parameter in the query.
	K *int `json:"k,omitempty"`
	// MaximumRelevance The highest relevance grade used in the user-supplied relevance judgments.
	MaximumRelevance int `json:"maximum_relevance"`
}

// RankEvalMetricExpectedReciprocalRankBuilder holds RankEvalMetricExpectedReciprocalRank struct and provides a builder API.
type RankEvalMetricExpectedReciprocalRankBuilder struct {
	v *RankEvalMetricExpectedReciprocalRank
}

// NewRankEvalMetricExpectedReciprocalRank provides a builder for the RankEvalMetricExpectedReciprocalRank struct.
func NewRankEvalMetricExpectedReciprocalRankBuilder() *RankEvalMetricExpectedReciprocalRankBuilder {
	r := RankEvalMetricExpectedReciprocalRankBuilder{
		&RankEvalMetricExpectedReciprocalRank{},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalMetricExpectedReciprocalRank struct
func (rb *RankEvalMetricExpectedReciprocalRankBuilder) Build() RankEvalMetricExpectedReciprocalRank {
	return *rb.v
}

// K Sets the maximum number of documents retrieved per query. This value will act
// in place of the usual size parameter in the query.

func (rb *RankEvalMetricExpectedReciprocalRankBuilder) K(k int) *RankEvalMetricExpectedReciprocalRankBuilder {
	rb.v.K = &k
	return rb
}

// MaximumRelevance The highest relevance grade used in the user-supplied relevance judgments.

func (rb *RankEvalMetricExpectedReciprocalRankBuilder) MaximumRelevance(maximumrelevance int) *RankEvalMetricExpectedReciprocalRankBuilder {
	rb.v.MaximumRelevance = maximumrelevance
	return rb
}
