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

// AutoDateHistogramAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L343-L347
type AutoDateHistogramAggregate struct {
	Buckets  BucketsDateHistogramBucket `json:"buckets"`
	Interval DurationLarge              `json:"interval"`
	Meta     *Metadata                  `json:"meta,omitempty"`
}

// AutoDateHistogramAggregateBuilder holds AutoDateHistogramAggregate struct and provides a builder API.
type AutoDateHistogramAggregateBuilder struct {
	v *AutoDateHistogramAggregate
}

// NewAutoDateHistogramAggregate provides a builder for the AutoDateHistogramAggregate struct.
func NewAutoDateHistogramAggregateBuilder() *AutoDateHistogramAggregateBuilder {
	r := AutoDateHistogramAggregateBuilder{
		&AutoDateHistogramAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the AutoDateHistogramAggregate struct
func (rb *AutoDateHistogramAggregateBuilder) Build() AutoDateHistogramAggregate {
	return *rb.v
}

func (rb *AutoDateHistogramAggregateBuilder) Buckets(buckets *BucketsDateHistogramBucketBuilder) *AutoDateHistogramAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *AutoDateHistogramAggregateBuilder) Interval(interval DurationLarge) *AutoDateHistogramAggregateBuilder {
	rb.v.Interval = interval
	return rb
}

func (rb *AutoDateHistogramAggregateBuilder) Meta(meta *MetadataBuilder) *AutoDateHistogramAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
