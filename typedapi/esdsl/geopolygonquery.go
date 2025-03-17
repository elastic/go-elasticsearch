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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"
)

type _geoPolygonQuery struct {
	v *types.GeoPolygonQuery
}

func NewGeoPolygonQuery() *_geoPolygonQuery {

	return &_geoPolygonQuery{v: types.NewGeoPolygonQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_geoPolygonQuery) Boost(boost float32) *_geoPolygonQuery {

	s.v.Boost = &boost

	return s
}

func (s *_geoPolygonQuery) GeoPolygonQuery(geopolygonquery map[string]types.GeoPolygonPoints) *_geoPolygonQuery {

	s.v.GeoPolygonQuery = geopolygonquery
	return s
}

func (s *_geoPolygonQuery) AddGeoPolygonQuery(key string, value types.GeoPolygonPointsVariant) *_geoPolygonQuery {

	var tmp map[string]types.GeoPolygonPoints
	if s.v.GeoPolygonQuery == nil {
		s.v.GeoPolygonQuery = make(map[string]types.GeoPolygonPoints)
	} else {
		tmp = s.v.GeoPolygonQuery
	}

	tmp[key] = *value.GeoPolygonPointsCaster()

	s.v.GeoPolygonQuery = tmp
	return s
}

func (s *_geoPolygonQuery) IgnoreUnmapped(ignoreunmapped bool) *_geoPolygonQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_geoPolygonQuery) QueryName_(queryname_ string) *_geoPolygonQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_geoPolygonQuery) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *_geoPolygonQuery {

	s.v.ValidationMethod = &validationmethod
	return s
}

func (s *_geoPolygonQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.GeoPolygon = s.v

	return container
}

func (s *_geoPolygonQuery) GeoPolygonQueryCaster() *types.GeoPolygonQuery {
	return s.v
}
