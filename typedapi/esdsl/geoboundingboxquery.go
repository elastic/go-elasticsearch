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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoexecution"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"
)

type _geoBoundingBoxQuery struct {
	v *types.GeoBoundingBoxQuery
}

// Matches geo_point and geo_shape values that intersect a bounding box.
func NewGeoBoundingBoxQuery() *_geoBoundingBoxQuery {

	return &_geoBoundingBoxQuery{v: types.NewGeoBoundingBoxQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_geoBoundingBoxQuery) Boost(boost float32) *_geoBoundingBoxQuery {

	s.v.Boost = &boost

	return s
}

func (s *_geoBoundingBoxQuery) GeoBoundingBoxQuery(geoboundingboxquery map[string]types.GeoBounds) *_geoBoundingBoxQuery {

	s.v.GeoBoundingBoxQuery = geoboundingboxquery
	return s
}

func (s *_geoBoundingBoxQuery) AddGeoBoundingBoxQuery(key string, value types.GeoBoundsVariant) *_geoBoundingBoxQuery {

	var tmp map[string]types.GeoBounds
	if s.v.GeoBoundingBoxQuery == nil {
		s.v.GeoBoundingBoxQuery = make(map[string]types.GeoBounds)
	} else {
		tmp = s.v.GeoBoundingBoxQuery
	}

	tmp[key] = *value.GeoBoundsCaster()

	s.v.GeoBoundingBoxQuery = tmp
	return s
}

// Set to `true` to ignore an unmapped field and not match any documents for
// this query.
// Set to `false` to throw an exception if the field is not mapped.
func (s *_geoBoundingBoxQuery) IgnoreUnmapped(ignoreunmapped bool) *_geoBoundingBoxQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_geoBoundingBoxQuery) QueryName_(queryname_ string) *_geoBoundingBoxQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_geoBoundingBoxQuery) Type(type_ geoexecution.GeoExecution) *_geoBoundingBoxQuery {

	s.v.Type = &type_
	return s
}

// Set to `IGNORE_MALFORMED` to accept geo points with invalid latitude or
// longitude.
// Set to `COERCE` to also try to infer correct latitude or longitude.
func (s *_geoBoundingBoxQuery) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *_geoBoundingBoxQuery {

	s.v.ValidationMethod = &validationmethod
	return s
}

func (s *_geoBoundingBoxQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.GeoBoundingBox = s.v

	return container
}

func (s *_geoBoundingBoxQuery) GeoBoundingBoxQueryCaster() *types.GeoBoundingBoxQuery {
	return s.v
}
