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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _rankEvalMetricDiscountedCumulativeGain struct {
	v *types.RankEvalMetricDiscountedCumulativeGain
}

func NewRankEvalMetricDiscountedCumulativeGain() *_rankEvalMetricDiscountedCumulativeGain {

	return &_rankEvalMetricDiscountedCumulativeGain{v: types.NewRankEvalMetricDiscountedCumulativeGain()}

}

// Sets the maximum number of documents retrieved per query. This value will act
// in place of the usual size parameter in the query.
func (s *_rankEvalMetricDiscountedCumulativeGain) K(k int) *_rankEvalMetricDiscountedCumulativeGain {

	s.v.K = &k

	return s
}

// If set to true, this metric will calculate the Normalized DCG.
func (s *_rankEvalMetricDiscountedCumulativeGain) Normalize(normalize bool) *_rankEvalMetricDiscountedCumulativeGain {

	s.v.Normalize = &normalize

	return s
}

func (s *_rankEvalMetricDiscountedCumulativeGain) RankEvalMetricDiscountedCumulativeGainCaster() *types.RankEvalMetricDiscountedCumulativeGain {
	return s.v
}
