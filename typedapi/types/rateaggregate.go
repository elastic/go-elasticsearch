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

// RateAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L692-L696
type RateAggregate struct {
	Meta          *Metadata `json:"meta,omitempty"`
	Value         float64   `json:"value"`
	ValueAsString *string   `json:"value_as_string,omitempty"`
}

// RateAggregateBuilder holds RateAggregate struct and provides a builder API.
type RateAggregateBuilder struct {
	v *RateAggregate
}

// NewRateAggregate provides a builder for the RateAggregate struct.
func NewRateAggregateBuilder() *RateAggregateBuilder {
	r := RateAggregateBuilder{
		&RateAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the RateAggregate struct
func (rb *RateAggregateBuilder) Build() RateAggregate {
	return *rb.v
}

func (rb *RateAggregateBuilder) Meta(meta *MetadataBuilder) *RateAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *RateAggregateBuilder) Value(value float64) *RateAggregateBuilder {
	rb.v.Value = value
	return rb
}

func (rb *RateAggregateBuilder) ValueAsString(valueasstring string) *RateAggregateBuilder {
	rb.v.ValueAsString = &valueasstring
	return rb
}
