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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/valuetype"
)

// WeightedAverageAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L211-L216
type WeightedAverageAggregation struct {
	Format    *string               `json:"format,omitempty"`
	Meta      *Metadata             `json:"meta,omitempty"`
	Name      *string               `json:"name,omitempty"`
	Value     *WeightedAverageValue `json:"value,omitempty"`
	ValueType *valuetype.ValueType  `json:"value_type,omitempty"`
	Weight    *WeightedAverageValue `json:"weight,omitempty"`
}

// WeightedAverageAggregationBuilder holds WeightedAverageAggregation struct and provides a builder API.
type WeightedAverageAggregationBuilder struct {
	v *WeightedAverageAggregation
}

// NewWeightedAverageAggregation provides a builder for the WeightedAverageAggregation struct.
func NewWeightedAverageAggregationBuilder() *WeightedAverageAggregationBuilder {
	r := WeightedAverageAggregationBuilder{
		&WeightedAverageAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the WeightedAverageAggregation struct
func (rb *WeightedAverageAggregationBuilder) Build() WeightedAverageAggregation {
	return *rb.v
}

func (rb *WeightedAverageAggregationBuilder) Format(format string) *WeightedAverageAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *WeightedAverageAggregationBuilder) Meta(meta *MetadataBuilder) *WeightedAverageAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *WeightedAverageAggregationBuilder) Name(name string) *WeightedAverageAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *WeightedAverageAggregationBuilder) Value(value *WeightedAverageValueBuilder) *WeightedAverageAggregationBuilder {
	v := value.Build()
	rb.v.Value = &v
	return rb
}

func (rb *WeightedAverageAggregationBuilder) ValueType(valuetype valuetype.ValueType) *WeightedAverageAggregationBuilder {
	rb.v.ValueType = &valuetype
	return rb
}

func (rb *WeightedAverageAggregationBuilder) Weight(weight *WeightedAverageValueBuilder) *WeightedAverageAggregationBuilder {
	v := weight.Build()
	rb.v.Weight = &v
	return rb
}
