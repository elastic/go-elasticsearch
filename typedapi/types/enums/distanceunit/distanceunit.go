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
// https://github.com/elastic/elasticsearch-specification/tree/19027dbdd366978ccae41842a040a636730e7c10

// Package distanceunit
package distanceunit

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/19027dbdd366978ccae41842a040a636730e7c10/specification/_types/Geo.ts#L30-L40
type DistanceUnit struct {
	Name string
}

var (
	Inches = DistanceUnit{"in"}

	Feet = DistanceUnit{"ft"}

	Yards = DistanceUnit{"yd"}

	Miles = DistanceUnit{"mi"}

	Nauticmiles = DistanceUnit{"nmi"}

	Kilometers = DistanceUnit{"km"}

	Meters = DistanceUnit{"m"}

	Centimeters = DistanceUnit{"cm"}

	Millimeters = DistanceUnit{"mm"}
)

func (d DistanceUnit) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DistanceUnit) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "in":
		*d = Inches
	case "ft":
		*d = Feet
	case "yd":
		*d = Yards
	case "mi":
		*d = Miles
	case "nmi":
		*d = Nauticmiles
	case "km":
		*d = Kilometers
	case "m":
		*d = Meters
	case "cm":
		*d = Centimeters
	case "mm":
		*d = Millimeters
	default:
		*d = DistanceUnit{string(text)}
	}

	return nil
}

func (d DistanceUnit) String() string {
	return d.Name
}
