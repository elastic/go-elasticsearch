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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rankFeatureQuery struct {
	v *types.RankFeatureQuery
}

// Boosts the relevance score of documents based on the numeric value of a
// `rank_feature` or `rank_features` field.
func NewRankFeatureQuery() *_rankFeatureQuery {

	return &_rankFeatureQuery{v: types.NewRankFeatureQuery()}

}

func (s *_rankFeatureQuery) Field(field string) *_rankFeatureQuery {

	s.v.Field = field

	return s
}

func (s *_rankFeatureQuery) Linear(linear types.RankFeatureFunctionLinearVariant) *_rankFeatureQuery {

	s.v.Linear = linear.RankFeatureFunctionLinearCaster()

	return s
}

func (s *_rankFeatureQuery) Log(log types.RankFeatureFunctionLogarithmVariant) *_rankFeatureQuery {

	s.v.Log = log.RankFeatureFunctionLogarithmCaster()

	return s
}

func (s *_rankFeatureQuery) Saturation(saturation types.RankFeatureFunctionSaturationVariant) *_rankFeatureQuery {

	s.v.Saturation = saturation.RankFeatureFunctionSaturationCaster()

	return s
}

func (s *_rankFeatureQuery) Sigmoid(sigmoid types.RankFeatureFunctionSigmoidVariant) *_rankFeatureQuery {

	s.v.Sigmoid = sigmoid.RankFeatureFunctionSigmoidCaster()

	return s
}

func (s *_rankFeatureQuery) Boost(boost float32) *_rankFeatureQuery {

	s.v.Boost = &boost

	return s
}

func (s *_rankFeatureQuery) QueryName_(queryname_ string) *_rankFeatureQuery {

	s.v.QueryName_ = &queryname_

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
