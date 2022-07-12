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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// MultiBucketAggregateBaseHistogramBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/Aggregate.ts#L314-L316
type MultiBucketAggregateBaseHistogramBucket struct {
	Buckets BucketsHistogramBucket `json:"buckets"`
	Meta    *Metadata              `json:"meta,omitempty"`
}

// MultiBucketAggregateBaseHistogramBucketBuilder holds MultiBucketAggregateBaseHistogramBucket struct and provides a builder API.
type MultiBucketAggregateBaseHistogramBucketBuilder struct {
	v *MultiBucketAggregateBaseHistogramBucket
}

// NewMultiBucketAggregateBaseHistogramBucket provides a builder for the MultiBucketAggregateBaseHistogramBucket struct.
func NewMultiBucketAggregateBaseHistogramBucketBuilder() *MultiBucketAggregateBaseHistogramBucketBuilder {
	r := MultiBucketAggregateBaseHistogramBucketBuilder{
		&MultiBucketAggregateBaseHistogramBucket{},
	}

	return &r
}

// Build finalize the chain and returns the MultiBucketAggregateBaseHistogramBucket struct
func (rb *MultiBucketAggregateBaseHistogramBucketBuilder) Build() MultiBucketAggregateBaseHistogramBucket {
	return *rb.v
}

func (rb *MultiBucketAggregateBaseHistogramBucketBuilder) Buckets(buckets *BucketsHistogramBucketBuilder) *MultiBucketAggregateBaseHistogramBucketBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *MultiBucketAggregateBaseHistogramBucketBuilder) Meta(meta *MetadataBuilder) *MultiBucketAggregateBaseHistogramBucketBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
