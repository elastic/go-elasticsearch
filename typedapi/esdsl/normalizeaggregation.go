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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalizemethod"
)

type _normalizeAggregation struct {
	v *types.NormalizeAggregation
}

// A parent pipeline aggregation which calculates the specific
// normalized/rescaled value for a specific bucket value.
func NewNormalizeAggregation() *_normalizeAggregation {

	return &_normalizeAggregation{v: types.NewNormalizeAggregation()}

}

func (s *_normalizeAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_normalizeAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_normalizeAggregation) Format(format string) *_normalizeAggregation {

	s.v.Format = &format

	return s
}

func (s *_normalizeAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_normalizeAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_normalizeAggregation) Method(method normalizemethod.NormalizeMethod) *_normalizeAggregation {

	s.v.Method = &method
	return s
}

func (s *_normalizeAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Normalize = s.v

	return container
}

func (s *_normalizeAggregation) NormalizeAggregationCaster() *types.NormalizeAggregation {
	return s.v
}
