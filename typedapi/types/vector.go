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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// Vector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/xpack/usage/types.ts#L445-L449
type Vector struct {
	Available               bool `json:"available"`
	DenseVectorDimsAvgCount int  `json:"dense_vector_dims_avg_count"`
	DenseVectorFieldsCount  int  `json:"dense_vector_fields_count"`
	Enabled                 bool `json:"enabled"`
	SparseVectorFieldsCount *int `json:"sparse_vector_fields_count,omitempty"`
}

// NewVector returns a Vector.
func NewVector() *Vector {
	r := &Vector{}

	return r
}
