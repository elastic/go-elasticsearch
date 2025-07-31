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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package rankvectorelementtype
package rankvectorelementtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/core.ts#L393-L397
type RankVectorElementType struct {
	Name string
}

var (
	Byte = RankVectorElementType{"byte"}

	Float = RankVectorElementType{"float"}

	Bit = RankVectorElementType{"bit"}
)

func (r RankVectorElementType) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RankVectorElementType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "byte":
		*r = Byte
	case "float":
		*r = Float
	case "bit":
		*r = Bit
	default:
		*r = RankVectorElementType{string(text)}
	}

	return nil
}

func (r RankVectorElementType) String() string {
	return r.Name
}
