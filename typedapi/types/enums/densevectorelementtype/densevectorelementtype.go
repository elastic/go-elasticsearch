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

// Package densevectorelementtype
package densevectorelementtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/mapping/DenseVectorProperty.ts#L64-L85
type DenseVectorElementType struct {
	Name string
}

var (

	// Bit Indexes a single bit per dimension. Useful for very high-dimensional vectors
	// or models that specifically support bit vectors.
	//
	// NOTE: when using `bit`, the number of dimensions must be a multiple of `8`
	// and must represent the number of bits.
	Bit = DenseVectorElementType{"bit"}

	// Byte Indexes a 1-byte integer value per dimension.
	Byte = DenseVectorElementType{"byte"}

	// Float Indexes a 4-byte floating-point value per dimension.
	Float = DenseVectorElementType{"float"}

	// Bfloat16 Indexes a 2-byte floating-point value per dimension.
	Bfloat16 = DenseVectorElementType{"bfloat16"}
)

func (d DenseVectorElementType) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DenseVectorElementType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "bit":
		*d = Bit
	case "byte":
		*d = Byte
	case "float":
		*d = Float
	case "bfloat16":
		*d = Bfloat16
	default:
		*d = DenseVectorElementType{string(text)}
	}

	return nil
}

func (d DenseVectorElementType) String() string {
	return d.Name
}
