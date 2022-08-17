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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/normalizemethod"
)

// NormalizeAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L250-L252
type NormalizeAggregation struct {
	// BucketsPath Path to the buckets that contain one set of values to correlate.
	BucketsPath *BucketsPath                     `json:"buckets_path,omitempty"`
	Format      *string                          `json:"format,omitempty"`
	GapPolicy   *gappolicy.GapPolicy             `json:"gap_policy,omitempty"`
	Meta        *Metadata                        `json:"meta,omitempty"`
	Method      *normalizemethod.NormalizeMethod `json:"method,omitempty"`
	Name        *string                          `json:"name,omitempty"`
}

// NormalizeAggregationBuilder holds NormalizeAggregation struct and provides a builder API.
type NormalizeAggregationBuilder struct {
	v *NormalizeAggregation
}

// NewNormalizeAggregation provides a builder for the NormalizeAggregation struct.
func NewNormalizeAggregationBuilder() *NormalizeAggregationBuilder {
	r := NormalizeAggregationBuilder{
		&NormalizeAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the NormalizeAggregation struct
func (rb *NormalizeAggregationBuilder) Build() NormalizeAggregation {
	return *rb.v
}

// BucketsPath Path to the buckets that contain one set of values to correlate.

func (rb *NormalizeAggregationBuilder) BucketsPath(bucketspath *BucketsPathBuilder) *NormalizeAggregationBuilder {
	v := bucketspath.Build()
	rb.v.BucketsPath = &v
	return rb
}

func (rb *NormalizeAggregationBuilder) Format(format string) *NormalizeAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *NormalizeAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *NormalizeAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *NormalizeAggregationBuilder) Meta(meta *MetadataBuilder) *NormalizeAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *NormalizeAggregationBuilder) Method(method normalizemethod.NormalizeMethod) *NormalizeAggregationBuilder {
	rb.v.Method = &method
	return rb
}

func (rb *NormalizeAggregationBuilder) Name(name string) *NormalizeAggregationBuilder {
	rb.v.Name = &name
	return rb
}
