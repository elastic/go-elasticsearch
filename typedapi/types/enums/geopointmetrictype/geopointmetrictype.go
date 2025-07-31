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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package geopointmetrictype
package geopointmetrictype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/geo.ts#L35-L39
type GeoPointMetricType struct {
	Name string
}

var (
	Gauge = GeoPointMetricType{"gauge"}

	Counter = GeoPointMetricType{"counter"}

	Position = GeoPointMetricType{"position"}
)

func (g GeoPointMetricType) MarshalText() (text []byte, err error) {
	return []byte(g.String()), nil
}

func (g *GeoPointMetricType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "gauge":
		*g = Gauge
	case "counter":
		*g = Counter
	case "position":
		*g = Position
	default:
		*g = GeoPointMetricType{string(text)}
	}

	return nil
}

func (g GeoPointMetricType) String() string {
	return g.Name
}
