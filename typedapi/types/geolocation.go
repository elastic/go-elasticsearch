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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// GeoLocation holds the union for the following types:
//
//	[]float64
//	GeoHashLocation
//	LatLonGeoLocation
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Geo.ts#L91-L105
type GeoLocation interface{}

// GeoLocationBuilder holds GeoLocation struct and provides a builder API.
type GeoLocationBuilder struct {
	v GeoLocation
}

// NewGeoLocation provides a builder for the GeoLocation struct.
func NewGeoLocationBuilder() *GeoLocationBuilder {
	return &GeoLocationBuilder{}
}

// Build finalize the chain and returns the GeoLocation struct
func (u *GeoLocationBuilder) Build() GeoLocation {
	return u.v
}

func (u *GeoLocationBuilder) Doubles(doubles ...float64) *GeoLocationBuilder {
	u.v = doubles
	return u
}

func (u *GeoLocationBuilder) GeoHashLocation(geohashlocation *GeoHashLocationBuilder) *GeoLocationBuilder {
	v := geohashlocation.Build()
	u.v = &v
	return u
}

func (u *GeoLocationBuilder) LatLonGeoLocation(latlongeolocation *LatLonGeoLocationBuilder) *GeoLocationBuilder {
	v := latlongeolocation.Build()
	u.v = &v
	return u
}

func (u *GeoLocationBuilder) String(string string) *GeoLocationBuilder {
	u.v = &string
	return u
}
