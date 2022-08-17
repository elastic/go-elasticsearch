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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

// MinBucketAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L174-L174
type MinBucketAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath         `json:"buckets_path,omitempty"`
	Format      *string              `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta        *Metadata            `json:"meta,omitempty"`
	Name        *string              `json:"name,omitempty"`
}

// MinBucketAggregationBuilder holds MinBucketAggregation struct and provides a builder API.
type MinBucketAggregationBuilder struct {
	v *MinBucketAggregation
}

// NewMinBucketAggregation provides a builder for the MinBucketAggregation struct.
func NewMinBucketAggregationBuilder() *MinBucketAggregationBuilder {
	r := MinBucketAggregationBuilder{
		&MinBucketAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MinBucketAggregation struct
func (rb *MinBucketAggregationBuilder) Build() MinBucketAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *MinBucketAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *MinBucketAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

func (rb *MinBucketAggregationBuilder) Format(format string) *MinBucketAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *MinBucketAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *MinBucketAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *MinBucketAggregationBuilder) Meta(meta *MetadataBuilder) *MinBucketAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *MinBucketAggregationBuilder) Name(name string) *MinBucketAggregationBuilder {
	rb.v.Name = &name
	return rb
}
