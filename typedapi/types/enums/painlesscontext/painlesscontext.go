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

// Package painlesscontext
package painlesscontext

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/scripts_painless_execute/types.ts#L57-L80
type PainlessContext struct {
	Name string
}

var (
	Painlesstest = PainlessContext{"painless_test"}

	Filter = PainlessContext{"filter"}

	Score = PainlessContext{"score"}

	Booleanfield = PainlessContext{"boolean_field"}

	Datefield = PainlessContext{"date_field"}

	Doublefield = PainlessContext{"double_field"}

	Geopointfield = PainlessContext{"geo_point_field"}

	Ipfield = PainlessContext{"ip_field"}

	Keywordfield = PainlessContext{"keyword_field"}

	Longfield = PainlessContext{"long_field"}

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
