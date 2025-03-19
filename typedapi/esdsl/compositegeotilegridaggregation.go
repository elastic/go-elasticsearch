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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/valuetype"
)

type _compositeGeoTileGridAggregation struct {
	v *types.CompositeGeoTileGridAggregation
}

func NewCompositeGeoTileGridAggregation() *_compositeGeoTileGridAggregation {

	return &_compositeGeoTileGridAggregation{v: types.NewCompositeGeoTileGridAggregation()}

}

func (s *_compositeGeoTileGridAggregation) Bounds(geobounds types.GeoBoundsVariant) *_compositeGeoTileGridAggregation {

	s.v.Bounds = *geobounds.GeoBoundsCaster()

	return s
}

// Either `field` or `script` must be present
func (s *_compositeGeoTileGridAggregation) Field(field string) *_compositeGeoTileGridAggregation {

	s.v.Field = &field

	return s
}

func (s *_compositeGeoTileGridAggregation) MissingBucket(missingbucket bool) *_compositeGeoTileGridAggregation {

	s.v.MissingBucket = &missingbucket

	return s
}

func (s *_compositeGeoTileGridAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_compositeGeoTileGridAggregation {

	s.v.MissingOrder = &missingorder
	return s
}

func (s *_compositeGeoTileGridAggregation) Order(order sortorder.SortOrder) *_compositeGeoTileGridAggregation {

	s.v.Order = &order
	return s
}

func (s *_compositeGeoTileGridAggregation) Precision(precision int) *_compositeGeoTileGridAggregation {

	s.v.Precision = &precision

	return s
}

// Either `field` or `script` must be present
func (s *_compositeGeoTileGridAggregation) Script(script types.ScriptVariant) *_compositeGeoTileGridAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_compositeGeoTileGridAggregation) ValueType(valuetype valuetype.ValueType) *_compositeGeoTileGridAggregation {

	s.v.ValueType = &valuetype
	return s
}

func (s *_compositeGeoTileGridAggregation) CompositeGeoTileGridAggregationCaster() *types.CompositeGeoTileGridAggregation {
	return s.v
}
