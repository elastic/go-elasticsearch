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
// https://github.com/elastic/elasticsearch-specification/tree/f1932ce6b46a53a8342db522b1a7883bcc9e0996

// Package densevectorelementtype
package densevectorelementtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/f1932ce6b46a53a8342db522b1a7883bcc9e0996/specification/_types/mapping/DenseVectorProperty.ts#L64-L80
type DenseVectorElementType struct {
	Name string
}

var (
	Bit = DenseVectorElementType{"bit"}

	Byte = DenseVectorElementType{"byte"}

	Float = DenseVectorElementType{"float"}
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
	default:
		*d = DenseVectorElementType{string(text)}
	}

	return nil
}

func (d DenseVectorElementType) String() string {
	return d.Name
}
