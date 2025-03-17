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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/densevectorindexoptionstype"
)

type _denseVectorIndexOptions struct {
	v *types.DenseVectorIndexOptions
}

func NewDenseVectorIndexOptions(type_ densevectorindexoptionstype.DenseVectorIndexOptionsType) *_denseVectorIndexOptions {

	tmp := &_denseVectorIndexOptions{v: types.NewDenseVectorIndexOptions()}

	tmp.Type(type_)

	return tmp

}

// The confidence interval to use when quantizing the vectors. Can be any value
// between and including `0.90` and
// `1.0` or exactly `0`. When the value is `0`, this indicates that dynamic
// quantiles should be calculated for
// optimized quantization. When between `0.90` and `1.0`, this value restricts
// the values used when calculating
// the quantization thresholds.
//
// For example, a value of `0.95` will only use the middle `95%` of the values
// when calculating the quantization
// thresholds (e.g. the highest and lowest `2.5%` of values will be ignored).
//
// Defaults to `1/(dims + 1)` for `int8` quantized vectors and `0` for `int4`
// for dynamic quantile calculation.
//
// Only applicable to `int8_hnsw`, `int4_hnsw`, `int8_flat`, and `int4_flat`
// index types.
func (s *_denseVectorIndexOptions) ConfidenceInterval(confidenceinterval float32) *_denseVectorIndexOptions {

	s.v.ConfidenceInterval = &confidenceinterval

	return s
}

// The number of candidates to track while assembling the list of nearest
// neighbors for each new node.
//
// Only applicable to `hnsw`, `int8_hnsw`, and `int4_hnsw` index types.
func (s *_denseVectorIndexOptions) EfConstruction(efconstruction int) *_denseVectorIndexOptions {

	s.v.EfConstruction = &efconstruction

	return s
}

// The number of neighbors each node will be connected to in the HNSW graph.
//
// Only applicable to `hnsw`, `int8_hnsw`, and `int4_hnsw` index types.
func (s *_denseVectorIndexOptions) M(m int) *_denseVectorIndexOptions {

	s.v.M = &m

	return s
}

// The type of kNN algorithm to use.
func (s *_denseVectorIndexOptions) Type(type_ densevectorindexoptionstype.DenseVectorIndexOptionsType) *_denseVectorIndexOptions {

	s.v.Type = type_
	return s
}

func (s *_denseVectorIndexOptions) DenseVectorIndexOptionsCaster() *types.DenseVectorIndexOptions {
	return s.v
}
