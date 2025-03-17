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

type _samplerAggregation struct {
	v *types.SamplerAggregation
}

// A filtering aggregation used to limit any sub aggregations' processing to a
// sample of the top-scoring documents.
func NewSamplerAggregation() *_samplerAggregation {

	return &_samplerAggregation{v: types.NewSamplerAggregation()}

}

// Limits how many top-scoring documents are collected in the sample processed
// on each shard.
func (s *_samplerAggregation) ShardSize(shardsize int) *_samplerAggregation {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_samplerAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Sampler = s.v

	return container
}

func (s *_samplerAggregation) SamplerAggregationCaster() *types.SamplerAggregation {
	return s.v
}
