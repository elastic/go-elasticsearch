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

// RankEvalMetric type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/rank_eval/types.ts#L90-L96
type RankEvalMetric struct {
	Dcg                    *RankEvalMetricDiscountedCumulativeGain `json:"dcg,omitempty"`
	ExpectedReciprocalRank *RankEvalMetricExpectedReciprocalRank   `json:"expected_reciprocal_rank,omitempty"`
	MeanReciprocalRank     *RankEvalMetricMeanReciprocalRank       `json:"mean_reciprocal_rank,omitempty"`
	Precision              *RankEvalMetricPrecision                `json:"precision,omitempty"`
	Recall                 *RankEvalMetricRecall                   `json:"recall,omitempty"`
}

// RankEvalMetricBuilder holds RankEvalMetric struct and provides a builder API.
type RankEvalMetricBuilder struct {
	v *RankEvalMetric
}

// NewRankEvalMetric provides a builder for the RankEvalMetric struct.
func NewRankEvalMetricBuilder() *RankEvalMetricBuilder {
	r := RankEvalMetricBuilder{
		&RankEvalMetric{},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalMetric struct
func (rb *RankEvalMetricBuilder) Build() RankEvalMetric {
	return *rb.v
}

func (rb *RankEvalMetricBuilder) Dcg(dcg *RankEvalMetricDiscountedCumulativeGainBuilder) *RankEvalMetricBuilder {
	v := dcg.Build()
	rb.v.Dcg = &v
	return rb
}

func (rb *RankEvalMetricBuilder) ExpectedReciprocalRank(expectedreciprocalrank *RankEvalMetricExpectedReciprocalRankBuilder) *RankEvalMetricBuilder {
	v := expectedreciprocalrank.Build()
	rb.v.ExpectedReciprocalRank = &v
	return rb
}

func (rb *RankEvalMetricBuilder) MeanReciprocalRank(meanreciprocalrank *RankEvalMetricMeanReciprocalRankBuilder) *RankEvalMetricBuilder {
	v := meanreciprocalrank.Build()
	rb.v.MeanReciprocalRank = &v
	return rb
}

func (rb *RankEvalMetricBuilder) Precision(precision *RankEvalMetricPrecisionBuilder) *RankEvalMetricBuilder {
	v := precision.Build()
	rb.v.Precision = &v
	return rb
}

func (rb *RankEvalMetricBuilder) Recall(recall *RankEvalMetricRecallBuilder) *RankEvalMetricBuilder {
	v := recall.Build()
	rb.v.Recall = &v
	return rb
}
