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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _geoGridQuery struct {
	v *types.GeoGridQuery
}

func NewGeoGridQuery() *_geoGridQuery {
	return &_geoGridQuery{v: types.NewGeoGridQuery()}
}

// AdditionalGeoGridQueryProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_geoGridQuery) AdditionalGeoGridQueryProperty(key string, value json.RawMessage) *_geoGridQuery {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalGeoGridQueryProperty = tmp
	return s
}

func (s *_geoGridQuery) Boost(boost float32) *_geoGridQuery {

	s.v.Boost = &boost

	return s
}

func (s *_geoGridQuery) Geogrid(geotile string) *_geoGridQuery {

	s.v.Geogrid = &geotile

	return s
}

func (s *_geoGridQuery) Geohash(geohash string) *_geoGridQuery {

	s.v.Geohash = &geohash

	return s
}

func (s *_geoGridQuery) Geohex(geohexcell string) *_geoGridQuery {

	s.v.Geohex = &geohexcell

	return s
}

func (s *_geoGridQuery) QueryName_(queryname_ string) *_geoGridQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_geoGridQuery) GeoGridQueryCaster() *types.GeoGridQuery {
	return s.v
}
