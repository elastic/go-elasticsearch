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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

// Package shapetype
package shapetype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/b1811e10a0722431d79d1c234dd412ff47d8656f/specification/ingest/_types/Processors.ts#L1543-L1546
type ShapeType struct {
	Name string
}

var (
	Geoshape = ShapeType{"geo_shape"}

	Shape = ShapeType{"shape"}
)

func (s ShapeType) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *ShapeType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "geo_shape":
		*s = Geoshape
	case "shape":
		*s = Shape
	default:
		*s = ShapeType{string(text)}
	}

	return nil
}

func (s ShapeType) String() string {
	return s.Name
}
