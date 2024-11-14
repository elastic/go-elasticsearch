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
// https://github.com/elastic/elasticsearch-specification/tree/4fcf747dfafc951e1dcf3077327e3dcee9107db3

// Package converttype
package converttype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4fcf747dfafc951e1dcf3077327e3dcee9107db3/specification/ingest/_types/Processors.ts#L621-L630
type ConvertType struct {
	Name string
}

var (
	Integer = ConvertType{"integer"}

	Long = ConvertType{"long"}

	Double = ConvertType{"double"}

	Float = ConvertType{"float"}

	Boolean = ConvertType{"boolean"}

	Ip = ConvertType{"ip"}

	String = ConvertType{"string"}

	Auto = ConvertType{"auto"}
)

func (c ConvertType) MarshalText() (text []byte, err error) {
	return []byte(c.String()), nil
}

func (c *ConvertType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "integer":
		*c = Integer
	case "long":
		*c = Long
	case "double":
		*c = Double
	case "float":
		*c = Float
	case "boolean":
		*c = Boolean
	case "ip":
		*c = Ip
	case "string":
		*c = String
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
