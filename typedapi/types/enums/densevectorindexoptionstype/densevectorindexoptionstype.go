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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package densevectorindexoptionstype
package densevectorindexoptionstype

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/_types/mapping/DenseVectorProperty.ts#L181-L235
type DenseVectorIndexOptionsType struct {
	Name string
}

var (

	// Bbqflat This utilizes a brute-force search algorithm in addition to automatically
	// quantizing to binary vectors. Only supports `element_type` of `float`.
	Bbqflat = DenseVectorIndexOptionsType{"bbq_flat"}

	// Bbqhnsw This utilizes the HNSW algorithm in addition to automatic binary quantization
	// for scalable approximate kNN search with `element_type` of `float`.
	//
	// This can reduce the memory footprint by nearly 32x at the cost of some
	// accuracy.
	Bbqhnsw = DenseVectorIndexOptionsType{"bbq_hnsw"}

	// Bbqdisk This utilizes the DiskBBQ algorithm, a version of Inverted Vector File (IVF)
	// that uses BBQ to quantize vectors. Only supports `element_type` of `float`.
	//
	// This not only significantly reduces memory usage, but also allows for
	// indexing and searching of very large datasets that do not fit in memory.
	// Unlike HNSW, this index type loses performance gracefully as the index grows
	// larger than memory.
	Bbqdisk = DenseVectorIndexOptionsType{"bbq_disk"}

	// Flat This utilizes a brute-force search algorithm for exact kNN search. This
	// supports all `element_type` values.
	Flat = DenseVectorIndexOptionsType{"flat"}

	// Hnsw This utilizes the HNSW algorithm for scalable approximate kNN search. This
	// supports all `element_type` values.
	Hnsw = DenseVectorIndexOptionsType{"hnsw"}

	// Int4flat This utilizes a brute-force search algorithm in addition to automatically
	// half-byte scalar quantization. Only supports `element_type` of `float`.
	Int4flat = DenseVectorIndexOptionsType{"int4_flat"}

	// Int4hnsw This utilizes the HNSW algorithm in addition to automatically scalar
	// quantization for scalable approximate kNN search with `element_type` of
	// `float`.
	//
	// This can reduce the memory footprint by 8x at the cost of some accuracy.
	Int4hnsw = DenseVectorIndexOptionsType{"int4_hnsw"}

	// Int8flat This utilizes a brute-force search algorithm in addition to automatically
	// scalar quantization. Only supports `element_type` of `float`.
	Int8flat = DenseVectorIndexOptionsType{"int8_flat"}

	// Int8hnsw The default index type for `float` vectors. This utilizes the HNSW algorithm
	// in addition to automatically scalar quantization for scalable approximate kNN
	// search with `element_type` of `float`.
	//
	// This can reduce the memory footprint by 4x at the cost of some accuracy.
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
	case "bbq_disk":
		*d = Bbqdisk
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
