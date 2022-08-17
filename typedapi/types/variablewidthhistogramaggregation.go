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

// VariableWidthHistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L426-L431
type VariableWidthHistogramAggregation struct {
	Buckets       *int   `json:"buckets,omitempty"`
	Field         *Field `json:"field,omitempty"`
	InitialBuffer *int   `json:"initial_buffer,omitempty"`
	ShardSize     *int   `json:"shard_size,omitempty"`
}

// VariableWidthHistogramAggregationBuilder holds VariableWidthHistogramAggregation struct and provides a builder API.
type VariableWidthHistogramAggregationBuilder struct {
	v *VariableWidthHistogramAggregation
}

// NewVariableWidthHistogramAggregation provides a builder for the VariableWidthHistogramAggregation struct.
func NewVariableWidthHistogramAggregationBuilder() *VariableWidthHistogramAggregationBuilder {
	r := VariableWidthHistogramAggregationBuilder{
		&VariableWidthHistogramAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the VariableWidthHistogramAggregation struct
func (rb *VariableWidthHistogramAggregationBuilder) Build() VariableWidthHistogramAggregation {
	return *rb.v
}

func (rb *VariableWidthHistogramAggregationBuilder) Buckets(buckets int) *VariableWidthHistogramAggregationBuilder {
	rb.v.Buckets = &buckets
	return rb
}

func (rb *VariableWidthHistogramAggregationBuilder) Field(field Field) *VariableWidthHistogramAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *VariableWidthHistogramAggregationBuilder) InitialBuffer(initialbuffer int) *VariableWidthHistogramAggregationBuilder {
	rb.v.InitialBuffer = &initialbuffer
	return rb
}

func (rb *VariableWidthHistogramAggregationBuilder) ShardSize(shardsize int) *VariableWidthHistogramAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}
