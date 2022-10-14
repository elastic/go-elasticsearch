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

// MultiBucketAggregateBaseLongTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/aggregations/Aggregate.ts#L318-L320
type MultiBucketAggregateBaseLongTermsBucket struct {
	Buckets BucketsLongTermsBucket `json:"buckets"`
	Meta    *Metadata              `json:"meta,omitempty"`
}

// MultiBucketAggregateBaseLongTermsBucketBuilder holds MultiBucketAggregateBaseLongTermsBucket struct and provides a builder API.
type MultiBucketAggregateBaseLongTermsBucketBuilder struct {
	v *MultiBucketAggregateBaseLongTermsBucket
}

// NewMultiBucketAggregateBaseLongTermsBucket provides a builder for the MultiBucketAggregateBaseLongTermsBucket struct.
func NewMultiBucketAggregateBaseLongTermsBucketBuilder() *MultiBucketAggregateBaseLongTermsBucketBuilder {
	r := MultiBucketAggregateBaseLongTermsBucketBuilder{
		&MultiBucketAggregateBaseLongTermsBucket{},
	}

	return &r
}

// Build finalize the chain and returns the MultiBucketAggregateBaseLongTermsBucket struct
func (rb *MultiBucketAggregateBaseLongTermsBucketBuilder) Build() MultiBucketAggregateBaseLongTermsBucket {
	return *rb.v
}

func (rb *MultiBucketAggregateBaseLongTermsBucketBuilder) Buckets(buckets *BucketsLongTermsBucketBuilder) *MultiBucketAggregateBaseLongTermsBucketBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *MultiBucketAggregateBaseLongTermsBucketBuilder) Meta(meta *MetadataBuilder) *MultiBucketAggregateBaseLongTermsBucketBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
