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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoShapeQuery struct {
	v *types.GeoShapeQuery
}

// Filter documents indexed using either the `geo_shape` or the `geo_point`
// type.
func NewGeoShapeQuery() *_geoShapeQuery {

	return &_geoShapeQuery{v: types.NewGeoShapeQuery()}

}

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
