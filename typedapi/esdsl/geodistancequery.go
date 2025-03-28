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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"
)

type _geoDistanceQuery struct {
	v *types.GeoDistanceQuery
}

// Matches `geo_point` and `geo_shape` values within a given distance of a
// geopoint.
func NewGeoDistanceQuery() *_geoDistanceQuery {

	return &_geoDistanceQuery{v: types.NewGeoDistanceQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_geoDistanceQuery) Boost(boost float32) *_geoDistanceQuery {

	s.v.Boost = &boost

	return s
}

// The radius of the circle centred on the specified location.
// Points which fall into this circle are considered to be matches.
func (s *_geoDistanceQuery) Distance(distance string) *_geoDistanceQuery {

	s.v.Distance = distance

	return s
}

// How to compute the distance.
// Set to `plane` for a faster calculation that's inaccurate on long distances
// and close to the poles.
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

// Set to `true` to ignore an unmapped field and not match any documents for
// this query.
// Set to `false` to throw an exception if the field is not mapped.
func (s *_geoDistanceQuery) IgnoreUnmapped(ignoreunmapped bool) *_geoDistanceQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_geoDistanceQuery) QueryName_(queryname_ string) *_geoDistanceQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Set to `IGNORE_MALFORMED` to accept geo points with invalid latitude or
// longitude.
// Set to `COERCE` to also try to infer correct latitude or longitude.
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
