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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _dateDistanceFeatureQuery struct {
	v *types.DateDistanceFeatureQuery
}

// Boosts the relevance score of documents closer to a provided origin date or
// point.
// For example, you can use this query to give more weight to documents closer
// to a certain date or location.
func NewDateDistanceFeatureQuery() *_dateDistanceFeatureQuery {

	return &_dateDistanceFeatureQuery{v: types.NewDateDistanceFeatureQuery()}

}

func (s *_dateDistanceFeatureQuery) Boost(boost float32) *_dateDistanceFeatureQuery {

	s.v.Boost = &boost

	return s
}

func (s *_dateDistanceFeatureQuery) Field(field string) *_dateDistanceFeatureQuery {

	s.v.Field = field

	return s
}

func (s *_dateDistanceFeatureQuery) Origin(datemath string) *_dateDistanceFeatureQuery {

	s.v.Origin = datemath

	return s
}

func (s *_dateDistanceFeatureQuery) Pivot(duration types.DurationVariant) *_dateDistanceFeatureQuery {

	s.v.Pivot = *duration.DurationCaster()

	return s
}

func (s *_dateDistanceFeatureQuery) QueryName_(queryname_ string) *_dateDistanceFeatureQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_dateDistanceFeatureQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.DistanceFeature = s.v

	return container
}

func (s *_dateDistanceFeatureQuery) DateDistanceFeatureQueryCaster() *types.DateDistanceFeatureQuery {
	return s.v
}
