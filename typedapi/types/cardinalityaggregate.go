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

// CardinalityAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L127-L130
type CardinalityAggregate struct {
	Meta  *Metadata `json:"meta,omitempty"`
	Value int64     `json:"value"`
}

// CardinalityAggregateBuilder holds CardinalityAggregate struct and provides a builder API.
type CardinalityAggregateBuilder struct {
	v *CardinalityAggregate
}

// NewCardinalityAggregate provides a builder for the CardinalityAggregate struct.
func NewCardinalityAggregateBuilder() *CardinalityAggregateBuilder {
	r := CardinalityAggregateBuilder{
		&CardinalityAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the CardinalityAggregate struct
func (rb *CardinalityAggregateBuilder) Build() CardinalityAggregate {
	return *rb.v
}

func (rb *CardinalityAggregateBuilder) Meta(meta *MetadataBuilder) *CardinalityAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *CardinalityAggregateBuilder) Value(value int64) *CardinalityAggregateBuilder {
	rb.v.Value = value
	return rb
}
