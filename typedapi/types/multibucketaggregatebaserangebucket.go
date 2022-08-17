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

// MultiBucketAggregateBaseRangeBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L314-L316
type MultiBucketAggregateBaseRangeBucket struct {
	Buckets BucketsRangeBucket `json:"buckets"`
	Meta    *Metadata          `json:"meta,omitempty"`
}

// MultiBucketAggregateBaseRangeBucketBuilder holds MultiBucketAggregateBaseRangeBucket struct and provides a builder API.
type MultiBucketAggregateBaseRangeBucketBuilder struct {
	v *MultiBucketAggregateBaseRangeBucket
}

// NewMultiBucketAggregateBaseRangeBucket provides a builder for the MultiBucketAggregateBaseRangeBucket struct.
func NewMultiBucketAggregateBaseRangeBucketBuilder() *MultiBucketAggregateBaseRangeBucketBuilder {
	r := MultiBucketAggregateBaseRangeBucketBuilder{
		&MultiBucketAggregateBaseRangeBucket{},
	}

	return &r
}

// Build finalize the chain and returns the MultiBucketAggregateBaseRangeBucket struct
func (rb *MultiBucketAggregateBaseRangeBucketBuilder) Build() MultiBucketAggregateBaseRangeBucket {
	return *rb.v
}

func (rb *MultiBucketAggregateBaseRangeBucketBuilder) Buckets(buckets *BucketsRangeBucketBuilder) *MultiBucketAggregateBaseRangeBucketBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *MultiBucketAggregateBaseRangeBucketBuilder) Meta(meta *MetadataBuilder) *MultiBucketAggregateBaseRangeBucketBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
