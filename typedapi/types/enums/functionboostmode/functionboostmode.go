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
// https://github.com/elastic/elasticsearch-specification/tree/f6a370d0fba975752c644fc730f7c45610e28f36

// Package functionboostmode
package functionboostmode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/f6a370d0fba975752c644fc730f7c45610e28f36/specification/_types/query_dsl/compound.ts#L295-L321
type FunctionBoostMode struct {
	Name string
}

var (
	Multiply = FunctionBoostMode{"multiply"}

	Replace = FunctionBoostMode{"replace"}

	Sum = FunctionBoostMode{"sum"}

	Avg = FunctionBoostMode{"avg"}

	Max = FunctionBoostMode{"max"}

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
