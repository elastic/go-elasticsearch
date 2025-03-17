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

type _geoShapeQuery struct {
	v *types.GeoShapeQuery
}

// Filter documents indexed using either the `geo_shape` or the `geo_point`
// type.
func NewGeoShapeQuery() *_geoShapeQuery {

	return &_geoShapeQuery{v: types.NewGeoShapeQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_geoShapeQuery) Boost(boost float32) *_geoShapeQuery {

	s.v.Boost = &boost

	return s
}

func (s *_geoShapeQuery) GeoShapeQuery(geoshapequery map[string]types.GeoShapeFieldQuery) *_geoShapeQuery {

	s.v.GeoShapeQuery = geoshapequery
	return s
}

func (s *_geoShapeQuery) AddGeoShapeQuery(key string, value types.GeoShapeFieldQueryVariant) *_geoShapeQuery {

	var tmp map[string]types.GeoShapeFieldQuery
	if s.v.GeoShapeQuery == nil {
		s.v.GeoShapeQuery = make(map[string]types.GeoShapeFieldQuery)
	} else {
		tmp = s.v.GeoShapeQuery
	}

	tmp[key] = *value.GeoShapeFieldQueryCaster()

	s.v.GeoShapeQuery = tmp
	return s
}

// Set to `true` to ignore an unmapped field and not match any documents for
// this query.
// Set to `false` to throw an exception if the field is not mapped.
func (s *_geoShapeQuery) IgnoreUnmapped(ignoreunmapped bool) *_geoShapeQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_geoShapeQuery) QueryName_(queryname_ string) *_geoShapeQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_geoShapeQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.GeoShape = s.v

	return container
}

func (s *_geoShapeQuery) GeoShapeQueryCaster() *types.GeoShapeQuery {
	return s.v
}
