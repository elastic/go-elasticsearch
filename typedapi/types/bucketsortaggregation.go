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

// BucketSortAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L142-L147
type BucketSortAggregation struct {
	From      *int                 `json:"from,omitempty"`
	GapPolicy *gappolicy.GapPolicy `json:"gap_policy,omitempty"`
	Meta      *Metadata            `json:"meta,omitempty"`
	Name      *string              `json:"name,omitempty"`
	Size      *int                 `json:"size,omitempty"`
	Sort      *Sort                `json:"sort,omitempty"`
}

// BucketSortAggregationBuilder holds BucketSortAggregation struct and provides a builder API.
type BucketSortAggregationBuilder struct {
	v *BucketSortAggregation
}

// NewBucketSortAggregation provides a builder for the BucketSortAggregation struct.
func NewBucketSortAggregationBuilder() *BucketSortAggregationBuilder {
	r := BucketSortAggregationBuilder{
		&BucketSortAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the BucketSortAggregation struct
func (rb *BucketSortAggregationBuilder) Build() BucketSortAggregation {
	return *rb.v
}

func (rb *BucketSortAggregationBuilder) From(from int) *BucketSortAggregationBuilder {
	rb.v.From = &from
	return rb
}

func (rb *BucketSortAggregationBuilder) GapPolicy(gappolicy gappolicy.GapPolicy) *BucketSortAggregationBuilder {
	rb.v.GapPolicy = &gappolicy
	return rb
}

func (rb *BucketSortAggregationBuilder) Meta(meta *MetadataBuilder) *BucketSortAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *BucketSortAggregationBuilder) Name(name string) *BucketSortAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *BucketSortAggregationBuilder) Size(size int) *BucketSortAggregationBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *BucketSortAggregationBuilder) Sort(sort *SortBuilder) *BucketSortAggregationBuilder {
	v := sort.Build()
	rb.v.Sort = &v
	return rb
}
