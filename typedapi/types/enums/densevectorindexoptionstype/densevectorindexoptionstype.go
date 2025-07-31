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

// Package densevectorindexoptionstype
package densevectorindexoptionstype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/mapping/DenseVectorProperty.ts#L164-L209
type DenseVectorIndexOptionsType struct {
	Name string
}

var (
	Bbqflat = DenseVectorIndexOptionsType{"bbq_flat"}

	Bbqhnsw = DenseVectorIndexOptionsType{"bbq_hnsw"}

	Flat = DenseVectorIndexOptionsType{"flat"}

	Hnsw = DenseVectorIndexOptionsType{"hnsw"}

	Int4flat = DenseVectorIndexOptionsType{"int4_flat"}

	Int4hnsw = DenseVectorIndexOptionsType{"int4_hnsw"}

	Int8flat = DenseVectorIndexOptionsType{"int8_flat"}

	Int8hnsw = DenseVectorIndexOptionsType{"int8_hnsw"}
)

func (d DenseVectorIndexOptionsType) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DenseVectorIndexOptionsType) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "bbq_flat":
		*d = Bbqflat
	case "bbq_hnsw":
		*d = Bbqhnsw
	case "flat":
		*d = Flat
	case "hnsw":
		*d = Hnsw
	case "int4_flat":
		*d = Int4flat
	case "int4_hnsw":
		*d = Int4hnsw
	case "int8_flat":
		*d = Int8flat
	case "int8_hnsw":
		*d = Int8hnsw
	default:
		*d = DenseVectorIndexOptionsType{string(text)}
	}

	return nil
}

func (d DenseVectorIndexOptionsType) String() string {
	return d.Name
}
