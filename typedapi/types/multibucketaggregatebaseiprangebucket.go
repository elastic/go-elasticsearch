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

// MultiBucketAggregateBaseIpRangeBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/aggregations/Aggregate.ts#L318-L320
type MultiBucketAggregateBaseIpRangeBucket struct {
	Buckets BucketsIpRangeBucket `json:"buckets"`
	Meta    *Metadata            `json:"meta,omitempty"`
}

// MultiBucketAggregateBaseIpRangeBucketBuilder holds MultiBucketAggregateBaseIpRangeBucket struct and provides a builder API.
type MultiBucketAggregateBaseIpRangeBucketBuilder struct {
	v *MultiBucketAggregateBaseIpRangeBucket
}

// NewMultiBucketAggregateBaseIpRangeBucket provides a builder for the MultiBucketAggregateBaseIpRangeBucket struct.
func NewMultiBucketAggregateBaseIpRangeBucketBuilder() *MultiBucketAggregateBaseIpRangeBucketBuilder {
	r := MultiBucketAggregateBaseIpRangeBucketBuilder{
		&MultiBucketAggregateBaseIpRangeBucket{},
	}

	return &r
}

// Build finalize the chain and returns the MultiBucketAggregateBaseIpRangeBucket struct
func (rb *MultiBucketAggregateBaseIpRangeBucketBuilder) Build() MultiBucketAggregateBaseIpRangeBucket {
	return *rb.v
}

func (rb *MultiBucketAggregateBaseIpRangeBucketBuilder) Buckets(buckets *BucketsIpRangeBucketBuilder) *MultiBucketAggregateBaseIpRangeBucketBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *MultiBucketAggregateBaseIpRangeBucketBuilder) Meta(meta *MetadataBuilder) *MultiBucketAggregateBaseIpRangeBucketBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
