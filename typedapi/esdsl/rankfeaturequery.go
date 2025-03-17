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

type _rankFeatureQuery struct {
	v *types.RankFeatureQuery
}

// Boosts the relevance score of documents based on the numeric value of a
// `rank_feature` or `rank_features` field.
func NewRankFeatureQuery() *_rankFeatureQuery {

	return &_rankFeatureQuery{v: types.NewRankFeatureQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_rankFeatureQuery) Boost(boost float32) *_rankFeatureQuery {

	s.v.Boost = &boost

	return s
}

// `rank_feature` or `rank_features` field used to boost relevance scores.
func (s *_rankFeatureQuery) Field(field string) *_rankFeatureQuery {

	s.v.Field = field

	return s
}

// Linear function used to boost relevance scores based on the value of the rank
// feature `field`.
func (s *_rankFeatureQuery) Linear(linear types.RankFeatureFunctionLinearVariant) *_rankFeatureQuery {

	s.v.Linear = linear.RankFeatureFunctionLinearCaster()

	return s
}

// Logarithmic function used to boost relevance scores based on the value of the
// rank feature `field`.
func (s *_rankFeatureQuery) Log(log types.RankFeatureFunctionLogarithmVariant) *_rankFeatureQuery {

	s.v.Log = log.RankFeatureFunctionLogarithmCaster()

	return s
}

func (s *_rankFeatureQuery) QueryName_(queryname_ string) *_rankFeatureQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Saturation function used to boost relevance scores based on the value of the
// rank feature `field`.
func (s *_rankFeatureQuery) Saturation(saturation types.RankFeatureFunctionSaturationVariant) *_rankFeatureQuery {

	s.v.Saturation = saturation.RankFeatureFunctionSaturationCaster()

	return s
}

// Sigmoid function used to boost relevance scores based on the value of the
// rank feature `field`.
func (s *_rankFeatureQuery) Sigmoid(sigmoid types.RankFeatureFunctionSigmoidVariant) *_rankFeatureQuery {

	s.v.Sigmoid = sigmoid.RankFeatureFunctionSigmoidCaster()

	return s
}

func (s *_rankFeatureQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.RankFeature = s.v

	return container
}

func (s *_rankFeatureQuery) RankFeatureQueryCaster() *types.RankFeatureQuery {
	return s.v
}
