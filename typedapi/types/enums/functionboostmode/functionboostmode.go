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

// Package functionboostmode
package functionboostmode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/query_dsl/compound.ts#L303-L329
type FunctionBoostMode struct {
	Name string
}

var (

	// Multiply Query score and function score are multiplied
	Multiply = FunctionBoostMode{"multiply"}

	// Replace Only the function score is used. The query score is ignored.
	Replace = FunctionBoostMode{"replace"}

	// Sum Query score and function score are added
	Sum = FunctionBoostMode{"sum"}

	// Avg Query score and function score are averaged
	Avg = FunctionBoostMode{"avg"}

	// Max Max of query score and function score
	Max = FunctionBoostMode{"max"}

	// Min Min of query score and function score
	Min = FunctionBoostMode{"min"}
)

func (f FunctionBoostMode) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FunctionBoostMode) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "multiply":
		*f = Multiply
	case "replace":
		*f = Replace
	case "sum":
		*f = Sum
	case "avg":
		*f = Avg
	case "max":
		*f = Max
	case "min":
		*f = Min
	default:
		*f = FunctionBoostMode{string(text)}
	}

	return nil
}

func (f FunctionBoostMode) String() string {
	return f.Name
}
