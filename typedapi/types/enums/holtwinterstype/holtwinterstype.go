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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Package holtwinterstype
package holtwinterstype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/aggregations/pipeline.ts#L309-L312
type HoltWintersType struct {
	Name string
}

var (
	Additive = HoltWintersType{"add"}

	Multiplicative = HoltWintersType{"mult"}
)

func (h HoltWintersType) MarshalText() (text []byte, err error) {
	return []byte(h.String()), nil
}

func (h *HoltWintersType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "add":
		*h = Additive
	case "mult":
		*h = Multiplicative
	default:
		*h = HoltWintersType{string(text)}
	}

	return nil
}

func (h HoltWintersType) String() string {
	return h.Name
}
