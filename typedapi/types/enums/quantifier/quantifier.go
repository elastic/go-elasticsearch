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


// Package quantifier
package quantifier

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/watcher/_types/Conditions.ts#L71-L74
type Quantifier struct {
	Name string
}

var (
	Some = Quantifier{"some"}

	All = Quantifier{"all"}
)

func (q Quantifier) MarshalText() (text []byte, err error) {
	return []byte(q.String()), nil
}

func (q *Quantifier) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "some":
		*q = Some
	case "all":
		*q = All
	default:
		*q = Quantifier{string(text)}
	}

	return nil
}

func (q Quantifier) String() string {
	return q.Name
}
