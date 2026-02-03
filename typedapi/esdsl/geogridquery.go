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

type _geoGridQuery struct {
	v *types.GeoGridQuery
}

func NewGeoGridQuery() *_geoGridQuery {
	return &_geoGridQuery{v: types.NewGeoGridQuery()}
}

func (s *_geoGridQuery) Geohash(geohash string) *_geoGridQuery {

	s.v.Geohash = &geohash

	return s
}

func (s *_geoGridQuery) Geohex(geohexcell string) *_geoGridQuery {

	s.v.Geohex = &geohexcell

	return s
}

func (s *_geoGridQuery) Geotile(geotile string) *_geoGridQuery {

	s.v.Geotile = &geotile

	return s
}

func (s *_geoGridQuery) Boost(boost float32) *_geoGridQuery {

	s.v.Boost = &boost

	return s
}

func (s *_geoGridQuery) QueryName_(queryname_ string) *_geoGridQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_geoGridQuery) GeoGridQueryCaster() *types.GeoGridQuery {
	return s.v
}
