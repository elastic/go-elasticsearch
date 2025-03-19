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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _boxplotAggregation struct {
	v *types.BoxplotAggregation
}

// A metrics aggregation that computes a box plot of numeric values extracted
// from the aggregated documents.
func NewBoxplotAggregation() *_boxplotAggregation {

	return &_boxplotAggregation{v: types.NewBoxplotAggregation()}

}

// Limits the maximum number of nodes used by the underlying TDigest algorithm
// to `20 * compression`, enabling control of memory usage and approximation
// error.
func (s *_boxplotAggregation) Compression(compression types.Float64) *_boxplotAggregation {

	s.v.Compression = &compression

	return s
}

// The field on which to run the aggregation.
func (s *_boxplotAggregation) Field(field string) *_boxplotAggregation {

	s.v.Field = &field

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_boxplotAggregation) Missing(missing types.MissingVariant) *_boxplotAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_boxplotAggregation) Script(script types.ScriptVariant) *_boxplotAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_boxplotAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Boxplot = s.v

	return container
}

func (s *_boxplotAggregation) BoxplotAggregationCaster() *types.BoxplotAggregation {
	return s.v
}
