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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _inferenceAggregation struct {
	v *types.InferenceAggregation
}

// A parent pipeline aggregation which loads a pre-trained model and performs
// inference on the collated result fields from the parent bucket aggregation.
func NewInferenceAggregation() *_inferenceAggregation {

	return &_inferenceAggregation{v: types.NewInferenceAggregation()}

}

func (s *_inferenceAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_inferenceAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_inferenceAggregation) Format(format string) *_inferenceAggregation {

	s.v.Format = &format

	return s
}

func (s *_inferenceAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_inferenceAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_inferenceAggregation) InferenceConfig(inferenceconfig types.InferenceConfigContainerVariant) *_inferenceAggregation {

	s.v.InferenceConfig = inferenceconfig.InferenceConfigContainerCaster()

	return s
}

func (s *_inferenceAggregation) ModelId(name string) *_inferenceAggregation {

	s.v.ModelId = name

	return s
}

func (s *_inferenceAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Inference = s.v

	return container
}

func (s *_inferenceAggregation) InferenceAggregationCaster() *types.InferenceAggregation {
	return s.v
}
