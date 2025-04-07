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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sampleraggregationexecutionhint"
)

type _diversifiedSamplerAggregation struct {
	v *types.DiversifiedSamplerAggregation
}

// A filtering aggregation used to limit any sub aggregations' processing to a
// sample of the top-scoring documents.
// Similar to the `sampler` aggregation, but adds the ability to limit the
// number of matches that share a common value.
func NewDiversifiedSamplerAggregation() *_diversifiedSamplerAggregation {

	return &_diversifiedSamplerAggregation{v: types.NewDiversifiedSamplerAggregation()}

}

func (s *_diversifiedSamplerAggregation) ExecutionHint(executionhint sampleraggregationexecutionhint.SamplerAggregationExecutionHint) *_diversifiedSamplerAggregation {

	s.v.ExecutionHint = &executionhint
	return s
}

func (s *_diversifiedSamplerAggregation) Field(field string) *_diversifiedSamplerAggregation {

	s.v.Field = &field

	return s
}

func (s *_diversifiedSamplerAggregation) MaxDocsPerValue(maxdocspervalue int) *_diversifiedSamplerAggregation {

	s.v.MaxDocsPerValue = &maxdocspervalue

	return s
}

func (s *_diversifiedSamplerAggregation) Script(script types.ScriptVariant) *_diversifiedSamplerAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_diversifiedSamplerAggregation) ShardSize(shardsize int) *_diversifiedSamplerAggregation {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_diversifiedSamplerAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.DiversifiedSampler = s.v

	return container
}

func (s *_diversifiedSamplerAggregation) DiversifiedSamplerAggregationCaster() *types.DiversifiedSamplerAggregation {
	return s.v
}
