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

// Package fieldvaluefactormodifier
package fieldvaluefactormodifier

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/query_dsl/compound.ts#L331-L374
type FieldValueFactorModifier struct {
	Name string
}

var (

	// None Do not apply any multiplier to the field value.
	None = FieldValueFactorModifier{"none"}

	// Log Take the common logarithm of the field value. Because this function will
	// return a negative value and cause an error if used on values between 0 and 1,
	// it is recommended to use `log1p` instead.
	Log = FieldValueFactorModifier{"log"}

	// Log1p Add 1 to the field value and take the common logarithm.
	Log1p = FieldValueFactorModifier{"log1p"}

	// Log2p Add 2 to the field value and take the common logarithm.
	Log2p = FieldValueFactorModifier{"log2p"}

	// Ln Take the natural logarithm of the field value. Because this function will
	// return a negative value and cause an error if used on values between 0 and 1,
	// it is recommended to use `ln1p` instead.
	Ln = FieldValueFactorModifier{"ln"}

	// Ln1p Add 1 to the field value and take the natural logarithm.
	Ln1p = FieldValueFactorModifier{"ln1p"}

	// Ln2p Add 2 to the field value and take the natural logarithm.
	Ln2p = FieldValueFactorModifier{"ln2p"}

	// Square Square the field value (multiply it by itself).
	Square = FieldValueFactorModifier{"square"}

	// Sqrt Take the square root of the field value.
	Sqrt = FieldValueFactorModifier{"sqrt"}

	// Reciprocal Reciprocate the field value, same as `1/x` where `x` is the field’s value.
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
