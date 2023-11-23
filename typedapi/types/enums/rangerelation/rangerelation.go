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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Package rangerelation
package rangerelation

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/term.ts#L170-L183
type RangeRelation struct {
	Name string
}

var (
	Within = RangeRelation{"within"}

	Contains = RangeRelation{"contains"}

	Intersects = RangeRelation{"intersects"}
)

func (r RangeRelation) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RangeRelation) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "within":
		*r = Within
	case "contains":
		*r = Contains
	case "intersects":
		*r = Intersects
	default:
		*r = RangeRelation{string(text)}
	}

	return nil
}

func (r RangeRelation) String() string {
	return r.Name
}
