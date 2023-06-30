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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Package fieldvaluefactormodifier
package fieldvaluefactormodifier

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/compound.ts#L147-L158
type FieldValueFactorModifier struct {
	Name string
}

var (
	None = FieldValueFactorModifier{"none"}

	Log = FieldValueFactorModifier{"log"}

	Log1p = FieldValueFactorModifier{"log1p"}

	Log2p = FieldValueFactorModifier{"log2p"}

	Ln = FieldValueFactorModifier{"ln"}

	Ln1p = FieldValueFactorModifier{"ln1p"}

	Ln2p = FieldValueFactorModifier{"ln2p"}

	Square = FieldValueFactorModifier{"square"}

	Sqrt = FieldValueFactorModifier{"sqrt"}

	Reciprocal = FieldValueFactorModifier{"reciprocal"}
)

func (f FieldValueFactorModifier) MarshalText() (text []byte, err error) {
	return []byte(f.String()), nil
}

func (f *FieldValueFactorModifier) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "none":
		*f = None
	case "log":
		*f = Log
	case "log1p":
		*f = Log1p
	case "log2p":
		*f = Log2p
	case "ln":
		*f = Ln
	case "ln1p":
		*f = Ln1p
	case "ln2p":
		*f = Ln2p
	case "square":
		*f = Square
	case "sqrt":
		*f = Sqrt
	case "reciprocal":
		*f = Reciprocal
	default:
		*f = FieldValueFactorModifier{string(text)}
	}

	return nil
}

func (f FieldValueFactorModifier) String() string {
	return f.Name
}
