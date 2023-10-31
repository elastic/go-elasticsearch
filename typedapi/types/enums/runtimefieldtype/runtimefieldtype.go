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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Package runtimefieldtype
package runtimefieldtype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/mapping/RuntimeFields.ts#L56-L65
type RuntimeFieldType struct {
	Name string
}

var (
	Boolean = RuntimeFieldType{"boolean"}

	Date = RuntimeFieldType{"date"}

	Double = RuntimeFieldType{"double"}

	Geopoint = RuntimeFieldType{"geo_point"}

	Ip = RuntimeFieldType{"ip"}

	Keyword = RuntimeFieldType{"keyword"}

	Long = RuntimeFieldType{"long"}

	Lookup = RuntimeFieldType{"lookup"}
)

func (r RuntimeFieldType) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RuntimeFieldType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "boolean":
		*r = Boolean
	case "date":
		*r = Date
	case "double":
		*r = Double
	case "geo_point":
		*r = Geopoint
	case "ip":
		*r = Ip
	case "keyword":
		*r = Keyword
	case "long":
		*r = Long
	case "lookup":
		*r = Lookup
	default:
		*r = RuntimeFieldType{string(text)}
	}

	return nil
}

func (r RuntimeFieldType) String() string {
	return r.Name
}
