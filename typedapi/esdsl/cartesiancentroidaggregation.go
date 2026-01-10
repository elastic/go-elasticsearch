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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _cartesianCentroidAggregation struct {
	v *types.CartesianCentroidAggregation
}

// A metric aggregation that computes the weighted centroid from all coordinate
// values for point and shape fields.
func NewCartesianCentroidAggregation() *_cartesianCentroidAggregation {

	return &_cartesianCentroidAggregation{v: types.NewCartesianCentroidAggregation()}

}

func (s *_cartesianCentroidAggregation) Field(field string) *_cartesianCentroidAggregation {

	s.v.Field = &field

	return s
}

func (s *_cartesianCentroidAggregation) Missing(missing types.MissingVariant) *_cartesianCentroidAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_cartesianCentroidAggregation) Script(script types.ScriptVariant) *_cartesianCentroidAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_cartesianCentroidAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.CartesianCentroid = s.v

	return container
}

func (s *_cartesianCentroidAggregation) CartesianCentroidAggregationCaster() *types.CartesianCentroidAggregation {
	return s.v
}
