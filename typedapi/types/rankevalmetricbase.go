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
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


package types

// RankEvalMetricBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/_global/rank_eval/types.ts#L26-L32
type RankEvalMetricBase struct {
	// K Sets the maximum number of documents retrieved per query. This value will act
	// in place of the usual size parameter in the query.
	K *int `json:"k,omitempty"`
}

// RankEvalMetricBaseBuilder holds RankEvalMetricBase struct and provides a builder API.
type RankEvalMetricBaseBuilder struct {
	v *RankEvalMetricBase
}

// NewRankEvalMetricBase provides a builder for the RankEvalMetricBase struct.
func NewRankEvalMetricBaseBuilder() *RankEvalMetricBaseBuilder {
	r := RankEvalMetricBaseBuilder{
		&RankEvalMetricBase{},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalMetricBase struct
func (rb *RankEvalMetricBaseBuilder) Build() RankEvalMetricBase {
	return *rb.v
}

// K Sets the maximum number of documents retrieved per query. This value will act
// in place of the usual size parameter in the query.

func (rb *RankEvalMetricBaseBuilder) K(k int) *RankEvalMetricBaseBuilder {
	rb.v.K = &k
	return rb
}
