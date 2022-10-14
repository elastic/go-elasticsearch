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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// VariableWidthHistogramAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/aggregations/Aggregate.ts#L353-L355
type VariableWidthHistogramAggregate struct {
	Buckets BucketsVariableWidthHistogramBucket `json:"buckets"`
	Meta    *Metadata                           `json:"meta,omitempty"`
}

// VariableWidthHistogramAggregateBuilder holds VariableWidthHistogramAggregate struct and provides a builder API.
type VariableWidthHistogramAggregateBuilder struct {
	v *VariableWidthHistogramAggregate
}

// NewVariableWidthHistogramAggregate provides a builder for the VariableWidthHistogramAggregate struct.
func NewVariableWidthHistogramAggregateBuilder() *VariableWidthHistogramAggregateBuilder {
	r := VariableWidthHistogramAggregateBuilder{
		&VariableWidthHistogramAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the VariableWidthHistogramAggregate struct
func (rb *VariableWidthHistogramAggregateBuilder) Build() VariableWidthHistogramAggregate {
	return *rb.v
}

func (rb *VariableWidthHistogramAggregateBuilder) Buckets(buckets *BucketsVariableWidthHistogramBucketBuilder) *VariableWidthHistogramAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *VariableWidthHistogramAggregateBuilder) Meta(meta *MetadataBuilder) *VariableWidthHistogramAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
