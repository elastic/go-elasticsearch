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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Package distanceunit
package distanceunit

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/Geo.ts#L30-L49
type DistanceUnit struct {
	Name string
}

var (
	In = DistanceUnit{"in"}

	Ft = DistanceUnit{"ft"}

	Yd = DistanceUnit{"yd"}

	Mi = DistanceUnit{"mi"}

	Nmi = DistanceUnit{"nmi"}

	Km = DistanceUnit{"km"}

	M = DistanceUnit{"m"}

	Cm = DistanceUnit{"cm"}

	Mm = DistanceUnit{"mm"}
)

func (d DistanceUnit) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DistanceUnit) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "in":
		*d = In
	case "ft":
		*d = Ft
	case "yd":
		*d = Yd
	case "mi":
		*d = Mi
	case "nmi":
		*d = Nmi
	case "km":
		*d = Km
	case "m":
		*d = M
	case "cm":
		*d = Cm
	case "mm":
		*d = Mm
	default:
		*d = DistanceUnit{string(text)}
	}

	return nil
}

func (d DistanceUnit) String() string {
	return d.Name
}
