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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _randomSamplerAggregation struct {
	v *types.RandomSamplerAggregation
}

// A single bucket aggregation that randomly includes documents in the
// aggregated results.
// Sampling provides significant speed improvement at the cost of accuracy.
func NewRandomSamplerAggregation(probability types.Float64) *_randomSamplerAggregation {

	tmp := &_randomSamplerAggregation{v: types.NewRandomSamplerAggregation()}

	tmp.Probability(probability)

	return tmp

}

func (s *_randomSamplerAggregation) Probability(probability types.Float64) *_randomSamplerAggregation {

	s.v.Probability = probability

	return s
}

func (s *_randomSamplerAggregation) Seed(seed int) *_randomSamplerAggregation {

	s.v.Seed = &seed

	return s
}

func (s *_randomSamplerAggregation) ShardSeed(shardseed int) *_randomSamplerAggregation {

	s.v.ShardSeed = &shardseed

	return s
}

func (s *_randomSamplerAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.RandomSampler = s.v

	return container
}

func (s *_randomSamplerAggregation) RandomSamplerAggregationCaster() *types.RandomSamplerAggregation {
	return s.v
}
