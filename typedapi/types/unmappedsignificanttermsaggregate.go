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

// UnmappedSignificantTermsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/aggregations/Aggregate.ts#L594-L600
type UnmappedSignificantTermsAggregate struct {
	BgCount  *int64      `json:"bg_count,omitempty"`
	Buckets  BucketsVoid `json:"buckets"`
	DocCount *int64      `json:"doc_count,omitempty"`
	Meta     *Metadata   `json:"meta,omitempty"`
}

// UnmappedSignificantTermsAggregateBuilder holds UnmappedSignificantTermsAggregate struct and provides a builder API.
type UnmappedSignificantTermsAggregateBuilder struct {
	v *UnmappedSignificantTermsAggregate
}

// NewUnmappedSignificantTermsAggregate provides a builder for the UnmappedSignificantTermsAggregate struct.
func NewUnmappedSignificantTermsAggregateBuilder() *UnmappedSignificantTermsAggregateBuilder {
	r := UnmappedSignificantTermsAggregateBuilder{
		&UnmappedSignificantTermsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the UnmappedSignificantTermsAggregate struct
func (rb *UnmappedSignificantTermsAggregateBuilder) Build() UnmappedSignificantTermsAggregate {
	return *rb.v
}

func (rb *UnmappedSignificantTermsAggregateBuilder) BgCount(bgcount int64) *UnmappedSignificantTermsAggregateBuilder {
	rb.v.BgCount = &bgcount
	return rb
}

func (rb *UnmappedSignificantTermsAggregateBuilder) Buckets(buckets *BucketsVoidBuilder) *UnmappedSignificantTermsAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *UnmappedSignificantTermsAggregateBuilder) DocCount(doccount int64) *UnmappedSignificantTermsAggregateBuilder {
	rb.v.DocCount = &doccount
	return rb
}

func (rb *UnmappedSignificantTermsAggregateBuilder) Meta(meta *MetadataBuilder) *UnmappedSignificantTermsAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
