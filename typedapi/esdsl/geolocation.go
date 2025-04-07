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

// This is provide all the types that are part of the union.
type _geoLocation struct {
	v types.GeoLocation
}

func NewGeoLocation() *_geoLocation {
	return &_geoLocation{v: nil}
}

func (u *_geoLocation) LatLonGeoLocation(latlongeolocation types.LatLonGeoLocationVariant) *_geoLocation {

	u.v = &latlongeolocation

	return u
}

// Interface implementation for LatLonGeoLocation in GeoLocation union
func (u *_latLonGeoLocation) GeoLocationCaster() *types.GeoLocation {
	t := types.GeoLocation(u.v)
	return &t
}

func (u *_geoLocation) GeoHashLocation(geohashlocation types.GeoHashLocationVariant) *_geoLocation {

	u.v = &geohashlocation

	return u
}

// Interface implementation for GeoHashLocation in GeoLocation union
func (u *_geoHashLocation) GeoLocationCaster() *types.GeoLocation {
	t := types.GeoLocation(u.v)
	return &t
}

func (u *_geoLocation) Doubles(doubles ...types.Float64) *_geoLocation {

	u.v = make([]types.Float64, len(doubles))
	u.v = doubles

	return u
}

func (u *_geoLocation) String(string string) *_geoLocation {

	u.v = &string

	return u
}

func (u *_geoLocation) GeoLocationCaster() *types.GeoLocation {
	return &u.v
}
