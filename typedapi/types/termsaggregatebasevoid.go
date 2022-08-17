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

// TermsAggregateBaseVoid type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L364-L369
type TermsAggregateBaseVoid struct {
	Buckets                 BucketsVoid `json:"buckets"`
	DocCountErrorUpperBound *int64      `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    *Metadata   `json:"meta,omitempty"`
	SumOtherDocCount        int64       `json:"sum_other_doc_count"`
}

// TermsAggregateBaseVoidBuilder holds TermsAggregateBaseVoid struct and provides a builder API.
type TermsAggregateBaseVoidBuilder struct {
	v *TermsAggregateBaseVoid
}

// NewTermsAggregateBaseVoid provides a builder for the TermsAggregateBaseVoid struct.
func NewTermsAggregateBaseVoidBuilder() *TermsAggregateBaseVoidBuilder {
	r := TermsAggregateBaseVoidBuilder{
		&TermsAggregateBaseVoid{},
	}

	return &r
}

// Build finalize the chain and returns the TermsAggregateBaseVoid struct
func (rb *TermsAggregateBaseVoidBuilder) Build() TermsAggregateBaseVoid {
	return *rb.v
}

func (rb *TermsAggregateBaseVoidBuilder) Buckets(buckets *BucketsVoidBuilder) *TermsAggregateBaseVoidBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *TermsAggregateBaseVoidBuilder) DocCountErrorUpperBound(doccounterrorupperbound int64) *TermsAggregateBaseVoidBuilder {
	rb.v.DocCountErrorUpperBound = &doccounterrorupperbound
	return rb
}

func (rb *TermsAggregateBaseVoidBuilder) Meta(meta *MetadataBuilder) *TermsAggregateBaseVoidBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *TermsAggregateBaseVoidBuilder) SumOtherDocCount(sumotherdoccount int64) *TermsAggregateBaseVoidBuilder {
	rb.v.SumOtherDocCount = sumotherdoccount
	return rb
}
