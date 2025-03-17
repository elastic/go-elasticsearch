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

type _rankEvalMetricExpectedReciprocalRank struct {
	v *types.RankEvalMetricExpectedReciprocalRank
}

func NewRankEvalMetricExpectedReciprocalRank(maximumrelevance int) *_rankEvalMetricExpectedReciprocalRank {

	tmp := &_rankEvalMetricExpectedReciprocalRank{v: types.NewRankEvalMetricExpectedReciprocalRank()}

	tmp.MaximumRelevance(maximumrelevance)

	return tmp

}

// Sets the maximum number of documents retrieved per query. This value will act
// in place of the usual size parameter in the query.
func (s *_rankEvalMetricExpectedReciprocalRank) K(k int) *_rankEvalMetricExpectedReciprocalRank {

	s.v.K = &k

	return s
}

// The highest relevance grade used in the user-supplied relevance judgments.
func (s *_rankEvalMetricExpectedReciprocalRank) MaximumRelevance(maximumrelevance int) *_rankEvalMetricExpectedReciprocalRank {

	s.v.MaximumRelevance = maximumrelevance

	return s
}

func (s *_rankEvalMetricExpectedReciprocalRank) RankEvalMetricExpectedReciprocalRankCaster() *types.RankEvalMetricExpectedReciprocalRank {
	return s.v
}
