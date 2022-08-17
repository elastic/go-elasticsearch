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

// RangeAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L511-L512
type RangeAggregate struct {
	Buckets BucketsRangeBucket `json:"buckets"`
	Meta    *Metadata          `json:"meta,omitempty"`
}

// RangeAggregateBuilder holds RangeAggregate struct and provides a builder API.
type RangeAggregateBuilder struct {
	v *RangeAggregate
}

// NewRangeAggregate provides a builder for the RangeAggregate struct.
func NewRangeAggregateBuilder() *RangeAggregateBuilder {
	r := RangeAggregateBuilder{
		&RangeAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the RangeAggregate struct
func (rb *RangeAggregateBuilder) Build() RangeAggregate {
	return *rb.v
}

func (rb *RangeAggregateBuilder) Buckets(buckets *BucketsRangeBucketBuilder) *RangeAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *RangeAggregateBuilder) Meta(meta *MetadataBuilder) *RangeAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
