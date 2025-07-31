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

type _geoBoundsAggregation struct {
	v *types.GeoBoundsAggregation
}

// A metric aggregation that computes the geographic bounding box containing all
// values for a Geopoint or Geoshape field.
func NewGeoBoundsAggregation() *_geoBoundsAggregation {

	return &_geoBoundsAggregation{v: types.NewGeoBoundsAggregation()}

}

func (s *_geoBoundsAggregation) Field(field string) *_geoBoundsAggregation {

	s.v.Field = &field

	return s
}

func (s *_geoBoundsAggregation) Missing(missing types.MissingVariant) *_geoBoundsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_geoBoundsAggregation) Script(script types.ScriptVariant) *_geoBoundsAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_geoBoundsAggregation) WrapLongitude(wraplongitude bool) *_geoBoundsAggregation {

	s.v.WrapLongitude = &wraplongitude

	return s
}

func (s *_geoBoundsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.GeoBounds = s.v

	return container
}

func (s *_geoBoundsAggregation) GeoBoundsAggregationCaster() *types.GeoBoundsAggregation {
	return s.v
}
