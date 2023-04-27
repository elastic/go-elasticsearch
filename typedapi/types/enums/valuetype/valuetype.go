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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

// Package valuetype
package valuetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/aggregations/metric.ts#L198-L209
type ValueType struct {
	Name string
}

var (
	String = ValueType{"string"}

	Long = ValueType{"long"}

	Double = ValueType{"double"}

	Number = ValueType{"number"}

	Date = ValueType{"date"}

	Datenanos = ValueType{"date_nanos"}

	Ip = ValueType{"ip"}

	Numeric = ValueType{"numeric"}

	Geopoint = ValueType{"geo_point"}

	Boolean = ValueType{"boolean"}
)

func (v ValueType) MarshalText() (text []byte, err error) {
	return []byte(v.String()), nil
}

func (v *ValueType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "string":
		*v = String
	case "long":
		*v = Long
	case "double":
		*v = Double
	case "number":
		*v = Number
	case "date":
		*v = Date
	case "date_nanos":
		*v = Datenanos
	case "ip":
		*v = Ip
	case "numeric":
		*v = Numeric
	case "geo_point":
		*v = Geopoint
	case "boolean":
		*v = Boolean
	default:
		*v = ValueType{string(text)}
	}

	return nil
}

func (v ValueType) String() string {
	return v.Name
}
