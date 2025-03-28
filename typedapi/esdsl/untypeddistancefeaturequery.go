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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _untypedDistanceFeatureQuery struct {
	v *types.UntypedDistanceFeatureQuery
}

// Boosts the relevance score of documents closer to a provided origin date or
// point.
// For example, you can use this query to give more weight to documents closer
// to a certain date or location.
func NewUntypedDistanceFeatureQuery(origin json.RawMessage, pivot json.RawMessage) *_untypedDistanceFeatureQuery {

	tmp := &_untypedDistanceFeatureQuery{v: types.NewUntypedDistanceFeatureQuery()}

	tmp.Origin(origin)

	tmp.Pivot(pivot)

	return tmp

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_untypedDistanceFeatureQuery) Boost(boost float32) *_untypedDistanceFeatureQuery {

	s.v.Boost = &boost

	return s
}

// Name of the field used to calculate distances. This field must meet the
// following criteria:
// be a `date`, `date_nanos` or `geo_point` field;
// have an `index` mapping parameter value of `true`, which is the default;
// have an `doc_values` mapping parameter value of `true`, which is the default.
func (s *_untypedDistanceFeatureQuery) Field(field string) *_untypedDistanceFeatureQuery {

	s.v.Field = field

	return s
}

// Date or point of origin used to calculate distances.
// If the `field` value is a `date` or `date_nanos` field, the `origin` value
// must be a date.
// Date Math, such as `now-1h`, is supported.
// If the field value is a `geo_point` field, the `origin` value must be a
// geopoint.
func (s *_untypedDistanceFeatureQuery) Origin(origin json.RawMessage) *_untypedDistanceFeatureQuery {

	s.v.Origin = origin

	return s
}

// Distance from the `origin` at which relevance scores receive half of the
// `boost` value.
// If the `field` value is a `date` or `date_nanos` field, the `pivot` value
// must be a time unit, such as `1h` or `10d`. If the `field` value is a
// `geo_point` field, the `pivot` value must be a distance unit, such as `1km`
// or `12m`.
func (s *_untypedDistanceFeatureQuery) Pivot(pivot json.RawMessage) *_untypedDistanceFeatureQuery {

	s.v.Pivot = pivot

	return s
}

func (s *_untypedDistanceFeatureQuery) QueryName_(queryname_ string) *_untypedDistanceFeatureQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_untypedDistanceFeatureQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.DistanceFeature = s.v

	return container
}

func (s *_untypedDistanceFeatureQuery) UntypedDistanceFeatureQueryCaster() *types.UntypedDistanceFeatureQuery {
	return s.v
}
