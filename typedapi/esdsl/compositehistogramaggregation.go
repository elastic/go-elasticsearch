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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type _compositeHistogramAggregation struct {
	v *types.CompositeHistogramAggregation
}

// A histogram aggregation.
func NewCompositeHistogramAggregation(interval types.Float64) *_compositeHistogramAggregation {

	tmp := &_compositeHistogramAggregation{v: types.NewCompositeHistogramAggregation()}

	tmp.Interval(interval)

	return tmp

}

func (s *_compositeHistogramAggregation) Field(field string) *_compositeHistogramAggregation {

	s.v.Field = &field

	return s
}

func (s *_compositeHistogramAggregation) Interval(interval types.Float64) *_compositeHistogramAggregation {

	s.v.Interval = interval

	return s
}

func (s *_compositeHistogramAggregation) MissingBucket(missingbucket bool) *_compositeHistogramAggregation {

	s.v.MissingBucket = &missingbucket

	return s
}

func (s *_compositeHistogramAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_compositeHistogramAggregation {

	s.v.MissingOrder = &missingorder
	return s
}

func (s *_compositeHistogramAggregation) Order(order sortorder.SortOrder) *_compositeHistogramAggregation {

	s.v.Order = &order
	return s
}

func (s *_compositeHistogramAggregation) Script(script types.ScriptVariant) *_compositeHistogramAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_compositeHistogramAggregation) ValueType(valuetype valuetype.ValueType) *_compositeHistogramAggregation {

	s.v.ValueType = &valuetype
	return s
}

func (s *_compositeHistogramAggregation) CompositeAggregationSourceCaster() *types.CompositeAggregationSource {
	container := types.NewCompositeAggregationSource()

	container.Histogram = s.v

	return container
}

func (s *_compositeHistogramAggregation) CompositeHistogramAggregationCaster() *types.CompositeHistogramAggregation {
	return s.v
}
