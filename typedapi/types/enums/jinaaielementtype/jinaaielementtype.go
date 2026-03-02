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
// https://github.com/elastic/elasticsearch-specification/tree/224e96968e3ab27c2d1d33f015783b44ed183c1f

// Package jinaaielementtype
package jinaaielementtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/224e96968e3ab27c2d1d33f015783b44ed183c1f/specification/inference/_types/CommonTypes.ts#L1490-L1494
type JinaAIElementType struct {
	Name string
}

var (
	Binary = JinaAIElementType{"binary"}

	Bit = JinaAIElementType{"bit"}

	Float = JinaAIElementType{"float"}
)

func (j JinaAIElementType) MarshalText() (text []byte, err error) {
	return []byte(j.String()), nil
}

func (j *JinaAIElementType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "binary":
		*j = Binary
	case "bit":
		*j = Bit
	case "float":
		*j = Float
	default:
		*j = JinaAIElementType{string(text)}
	}

	return nil
}

func (j JinaAIElementType) String() string {
	return j.Name
}
