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


// Package converttype
package converttype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ingest/_types/Processors.ts#L137-L145
type ConvertType struct {
	Name string
}

var (
	Integer = ConvertType{"integer"}

	Long = ConvertType{"long"}

	Float = ConvertType{"float"}

	Double = ConvertType{"double"}

	String = ConvertType{"string"}

	Boolean = ConvertType{"boolean"}

	Auto = ConvertType{"auto"}
)

func (c ConvertType) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ConvertType) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "integer":
		*c = Integer
	case "long":
		*c = Long
	case "float":
		*c = Float
	case "double":
		*c = Double
	case "string":
		*c = String
	case "boolean":
		*c = Boolean
	case "auto":
		*c = Auto
	default:
		*c = ConvertType{string(text)}
	}

	return nil
}

func (c ConvertType) String() string {
	return c.Name
}
