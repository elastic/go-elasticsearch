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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ttesttype"
)

type _tTestAggregation struct {
	v *types.TTestAggregation
}

// A metrics aggregation that performs a statistical hypothesis test in which
// the test statistic follows a Student’s t-distribution under the null
// hypothesis on numeric values extracted from the aggregated documents.
func NewTTestAggregation() *_tTestAggregation {

	return &_tTestAggregation{v: types.NewTTestAggregation()}

}

// Test population A.
func (s *_tTestAggregation) A(a types.TestPopulationVariant) *_tTestAggregation {

	s.v.A = a.TestPopulationCaster()

	return s
}

// Test population B.
func (s *_tTestAggregation) B(b types.TestPopulationVariant) *_tTestAggregation {

	s.v.B = b.TestPopulationCaster()

	return s
}

// The type of test.
func (s *_tTestAggregation) Type(type_ ttesttype.TTestType) *_tTestAggregation {

	s.v.Type = &type_
	return s
}

func (s *_tTestAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.TTest = s.v

	return container
}

func (s *_tTestAggregation) TTestAggregationCaster() *types.TTestAggregation {
	return s.v
}
