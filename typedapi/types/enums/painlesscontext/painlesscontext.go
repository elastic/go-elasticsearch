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

// Package painlesscontext
package painlesscontext

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_global/scripts_painless_execute/types.ts#L57-L80
type PainlessContext struct {
	Name string
}

var (

	// Painlesstest The default context if no other context is specified.
	Painlesstest = PainlessContext{"painless_test"}

	// Filter Treats scripts as if they were run inside a script query.
	Filter = PainlessContext{"filter"}

	// Score Treats scripts as if they were run inside a `script_score` function in a
	// `function_score` query.
	Score = PainlessContext{"score"}

	// Booleanfield The context for boolean fields. The script returns a `true` or `false`
	// response.
	Booleanfield = PainlessContext{"boolean_field"}

	// Datefield The context for date fields. `emit` takes a long value and the script returns
	// a sorted list of dates.
	Datefield = PainlessContext{"date_field"}

	// Doublefield The context for double numeric fields. The script returns a sorted list of
	// double values.
	Doublefield = PainlessContext{"double_field"}

	// Geopointfield The context for geo-point fields. `emit` takes two double parameters, the
	// latitude and longitude values, and the script returns an object in GeoJSON
	// format containing the coordinates for the geo point.
	Geopointfield = PainlessContext{"geo_point_field"}

	// Ipfield The context for `ip` fields. The script returns a sorted list of IP
	// addresses.
	Ipfield = PainlessContext{"ip_field"}

	// Keywordfield The context for keyword fields. The script returns a sorted list of string
	// values.
	Keywordfield = PainlessContext{"keyword_field"}

	// Longfield The context for long numeric fields. The script returns a sorted list of long
	// values.
	Longfield = PainlessContext{"long_field"}

	// Compositefield The context for composite runtime fields. The script returns a map of values.
	Compositefield = PainlessContext{"composite_field"}
)

func (p PainlessContext) MarshalText() (text []byte, err error) {
	return []byte(p.String()), nil
}

func (p *PainlessContext) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "painless_test":
		*p = Painlesstest
	case "filter":
		*p = Filter
	case "score":
		*p = Score
	case "boolean_field":
		*p = Booleanfield
	case "date_field":
		*p = Datefield
	case "double_field":
		*p = Doublefield
	case "geo_point_field":
		*p = Geopointfield
	case "ip_field":
		*p = Ipfield
	case "keyword_field":
		*p = Keywordfield
	case "long_field":
		*p = Longfield
	case "composite_field":
		*p = Compositefield
	default:
		*p = PainlessContext{string(text)}
	}

	return nil
}

func (p PainlessContext) String() string {
	return p.Name
}
