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


// Package functionscoremode
package functionscoremode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/query_dsl/compound.ts#L129-L136
type FunctionScoreMode struct {
	Name string
}

var (
	Multiply = FunctionScoreMode{"multiply"}

	Sum = FunctionScoreMode{"sum"}

	Avg = FunctionScoreMode{"avg"}

	First = FunctionScoreMode{"first"}

	Max = FunctionScoreMode{"max"}

	Min = FunctionScoreMode{"min"}
)

func (f FunctionScoreMode) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FunctionScoreMode) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "multiply":
		*f = Multiply
	case "sum":
		*f = Sum
	case "avg":
		*f = Avg
	case "first":
		*f = First
	case "max":
		*f = Max
	case "min":
		*f = Min
	default:
		*f = FunctionScoreMode{string(text)}
	}

	return nil
}

func (f FunctionScoreMode) String() string {
	return f.Name
}
