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

// AdjacencyMatrixAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L552-L554
type AdjacencyMatrixAggregate struct {
	Buckets BucketsAdjacencyMatrixBucket `json:"buckets"`
	Meta    *Metadata                    `json:"meta,omitempty"`
}

// AdjacencyMatrixAggregateBuilder holds AdjacencyMatrixAggregate struct and provides a builder API.
type AdjacencyMatrixAggregateBuilder struct {
	v *AdjacencyMatrixAggregate
}

// NewAdjacencyMatrixAggregate provides a builder for the AdjacencyMatrixAggregate struct.
func NewAdjacencyMatrixAggregateBuilder() *AdjacencyMatrixAggregateBuilder {
	r := AdjacencyMatrixAggregateBuilder{
		&AdjacencyMatrixAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the AdjacencyMatrixAggregate struct
func (rb *AdjacencyMatrixAggregateBuilder) Build() AdjacencyMatrixAggregate {
	return *rb.v
}

func (rb *AdjacencyMatrixAggregateBuilder) Buckets(buckets *BucketsAdjacencyMatrixBucketBuilder) *AdjacencyMatrixAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *AdjacencyMatrixAggregateBuilder) Meta(meta *MetadataBuilder) *AdjacencyMatrixAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
