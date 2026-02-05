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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _geoDistanceSort struct {
	v *types.GeoDistanceSort
}

func NewGeoDistanceSort() *_geoDistanceSort {

	return &_geoDistanceSort{v: types.NewGeoDistanceSort()}

}

func (s *_geoDistanceSort) DistanceType(distancetype geodistancetype.GeoDistanceType) *_geoDistanceSort {

	s.v.DistanceType = &distancetype
	return s
}

func (s *_geoDistanceSort) GeoDistanceSort(geodistancesort map[string][]types.GeoLocation) *_geoDistanceSort {

	s.v.GeoDistanceSort = geodistancesort
	return s
}

func (s *_geoDistanceSort) IgnoreUnmapped(ignoreunmapped bool) *_geoDistanceSort {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_geoDistanceSort) Mode(mode sortmode.SortMode) *_geoDistanceSort {

	s.v.Mode = &mode
	return s
}

func (s *_geoDistanceSort) Nested(nested types.NestedSortValueVariant) *_geoDistanceSort {

	s.v.Nested = nested.NestedSortValueCaster()

	return s
}

func (s *_geoDistanceSort) Order(order sortorder.SortOrder) *_geoDistanceSort {

	s.v.Order = &order
	return s
}

func (s *_geoDistanceSort) Unit(unit distanceunit.DistanceUnit) *_geoDistanceSort {

	s.v.Unit = &unit
	return s
}

func (s *_geoDistanceSort) SortOptionsCaster() *types.SortOptions {
	container := types.NewSortOptions()

	container.GeoDistance_ = s.v

	return container
}

func (s *_geoDistanceSort) GeoDistanceSortCaster() *types.GeoDistanceSort {
	return s.v
}
