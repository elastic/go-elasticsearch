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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _aggregationRange struct {
	v *types.AggregationRange
}

func NewAggregationRange() *_aggregationRange {

	return &_aggregationRange{v: types.NewAggregationRange()}

}

// Start of the range (inclusive).
func (s *_aggregationRange) From(from types.Float64) *_aggregationRange {

	s.v.From = &from

	return s
}

// Custom key to return the range with.
func (s *_aggregationRange) Key(key string) *_aggregationRange {

	s.v.Key = &key

	return s
}

// End of the range (exclusive).
func (s *_aggregationRange) To(to types.Float64) *_aggregationRange {

	s.v.To = &to

	return s
}

func (s *_aggregationRange) AggregationRangeCaster() *types.AggregationRange {
	return s.v
}
