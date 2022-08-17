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

// FiltersAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L166-L171
type FiltersAggregation struct {
	Filters        *BucketsQueryContainer `json:"filters,omitempty"`
	Keyed          *bool                  `json:"keyed,omitempty"`
	Meta           *Metadata              `json:"meta,omitempty"`
	Name           *string                `json:"name,omitempty"`
	OtherBucket    *bool                  `json:"other_bucket,omitempty"`
	OtherBucketKey *string                `json:"other_bucket_key,omitempty"`
}

// FiltersAggregationBuilder holds FiltersAggregation struct and provides a builder API.
type FiltersAggregationBuilder struct {
	v *FiltersAggregation
}

// NewFiltersAggregation provides a builder for the FiltersAggregation struct.
func NewFiltersAggregationBuilder() *FiltersAggregationBuilder {
	r := FiltersAggregationBuilder{
		&FiltersAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the FiltersAggregation struct
func (rb *FiltersAggregationBuilder) Build() FiltersAggregation {
	return *rb.v
}

func (rb *FiltersAggregationBuilder) Filters(filters *BucketsQueryContainerBuilder) *FiltersAggregationBuilder {
	v := filters.Build()
	rb.v.Filters = &v
	return rb
}

func (rb *FiltersAggregationBuilder) Keyed(keyed bool) *FiltersAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

func (rb *FiltersAggregationBuilder) Meta(meta *MetadataBuilder) *FiltersAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *FiltersAggregationBuilder) Name(name string) *FiltersAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *FiltersAggregationBuilder) OtherBucket(otherbucket bool) *FiltersAggregationBuilder {
	rb.v.OtherBucket = &otherbucket
	return rb
}

func (rb *FiltersAggregationBuilder) OtherBucketKey(otherbucketkey string) *FiltersAggregationBuilder {
	rb.v.OtherBucketKey = &otherbucketkey
	return rb
}
