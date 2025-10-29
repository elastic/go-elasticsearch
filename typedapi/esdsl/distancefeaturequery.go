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

// This is provide all the types that are part of the union.
type _distanceFeatureQuery struct {
	v types.DistanceFeatureQuery
}

func NewDistanceFeatureQuery() *_distanceFeatureQuery {
	return &_distanceFeatureQuery{v: nil}
}

func (u *_distanceFeatureQuery) UntypedDistanceFeatureQuery(untypeddistancefeaturequery types.UntypedDistanceFeatureQueryVariant) *_distanceFeatureQuery {

	u.v = untypeddistancefeaturequery.UntypedDistanceFeatureQueryCaster()

	return u
}

// Interface implementation for UntypedDistanceFeatureQuery in DistanceFeatureQuery union
func (u *_untypedDistanceFeatureQuery) DistanceFeatureQueryCaster() *types.DistanceFeatureQuery {
	t := types.DistanceFeatureQuery(u.v)
	return &t
}

func (u *_distanceFeatureQuery) GeoDistanceFeatureQuery(geodistancefeaturequery types.GeoDistanceFeatureQueryVariant) *_distanceFeatureQuery {

	u.v = geodistancefeaturequery.GeoDistanceFeatureQueryCaster()

	return u
}

// Interface implementation for GeoDistanceFeatureQuery in DistanceFeatureQuery union
func (u *_geoDistanceFeatureQuery) DistanceFeatureQueryCaster() *types.DistanceFeatureQuery {
	t := types.DistanceFeatureQuery(u.v)
	return &t
}

func (u *_distanceFeatureQuery) DateDistanceFeatureQuery(datedistancefeaturequery types.DateDistanceFeatureQueryVariant) *_distanceFeatureQuery {

	u.v = datedistancefeaturequery.DateDistanceFeatureQueryCaster()

	return u
}

// Interface implementation for DateDistanceFeatureQuery in DistanceFeatureQuery union
func (u *_dateDistanceFeatureQuery) DistanceFeatureQueryCaster() *types.DistanceFeatureQuery {
	t := types.DistanceFeatureQuery(u.v)
	return &t
}

func (u *_distanceFeatureQuery) DistanceFeatureQueryCaster() *types.DistanceFeatureQuery {
	return &u.v
}
