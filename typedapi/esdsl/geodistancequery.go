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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geovalidationmethod"
)

type _geoDistanceQuery struct {
	v *types.GeoDistanceQuery
}

// Matches `geo_point` and `geo_shape` values within a given distance of a
// geopoint.
func NewGeoDistanceQuery() *_geoDistanceQuery {

	return &_geoDistanceQuery{v: types.NewGeoDistanceQuery()}

}

func (s *_geoDistanceQuery) Boost(boost float32) *_geoDistanceQuery {

	s.v.Boost = &boost

	return s
}

func (s *_geoDistanceQuery) Distance(distance string) *_geoDistanceQuery {

	s.v.Distance = distance

	return s
}

func (s *_geoDistanceQuery) DistanceType(distancetype geodistancetype.GeoDistanceType) *_geoDistanceQuery {

	s.v.DistanceType = &distancetype
	return s
}

func (s *_geoDistanceQuery) GeoDistanceQuery(geodistancequery map[string]types.GeoLocation) *_geoDistanceQuery {

	s.v.GeoDistanceQuery = geodistancequery
	return s
}

func (s *_geoDistanceQuery) AddGeoDistanceQuery(key string, value types.GeoLocationVariant) *_geoDistanceQuery {

	var tmp map[string]types.GeoLocation
	if s.v.GeoDistanceQuery == nil {
		s.v.GeoDistanceQuery = make(map[string]types.GeoLocation)
	} else {
		tmp = s.v.GeoDistanceQuery
	}

	tmp[key] = *value.GeoLocationCaster()

	s.v.GeoDistanceQuery = tmp
	return s
}

func (s *_geoDistanceQuery) IgnoreUnmapped(ignoreunmapped bool) *_geoDistanceQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_geoDistanceQuery) QueryName_(queryname_ string) *_geoDistanceQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_geoDistanceQuery) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *_geoDistanceQuery {

	s.v.ValidationMethod = &validationmethod
	return s
}

func (s *_geoDistanceQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.GeoDistance = s.v

	return container
}

func (s *_geoDistanceQuery) GeoDistanceQueryCaster() *types.GeoDistanceQuery {
	return s.v
}
