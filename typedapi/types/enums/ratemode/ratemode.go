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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Package ratemode
package ratemode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/aggregations/metric.ts#L132-L135
type RateMode struct {
	Name string
}

var (
	Sum = RateMode{"sum"}

	Valuecount = RateMode{"value_count"}
)

func (r RateMode) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RateMode) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "sum":
		*r = Sum
	case "value_count":
		*r = Valuecount
	default:
		*r = RateMode{string(text)}
	}

	return nil
}

func (r RateMode) String() string {
	return r.Name
}
