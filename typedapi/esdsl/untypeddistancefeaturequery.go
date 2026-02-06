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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
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

func (s *_untypedDistanceFeatureQuery) Boost(boost float32) *_untypedDistanceFeatureQuery {

	s.v.Boost = &boost

	return s
}

func (s *_untypedDistanceFeatureQuery) Field(field string) *_untypedDistanceFeatureQuery {

	s.v.Field = field

	return s
}

func (s *_untypedDistanceFeatureQuery) Origin(origin json.RawMessage) *_untypedDistanceFeatureQuery {

	s.v.Origin = origin

	return s
}

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
