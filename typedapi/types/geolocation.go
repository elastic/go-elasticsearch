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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package types

// GeoLocation holds the union for the following types:
//
//	LatLonGeoLocation
//	GeoHashLocation
//	[]Float64
//	string
//
// A latitude/longitude as a 2 dimensional point. It can be represented in
// various ways:
//
//   - as a `{lat, long}` object
//   - as a geo hash value
//   - as a `[lon, lat]` array
//   - as a string in `"<lat>, <lon>"` or WKT point formats
//
// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/Geo.ts#L98-L112
type GeoLocation any

type GeoLocationVariant interface {
	GeoLocationCaster() *GeoLocation
}
