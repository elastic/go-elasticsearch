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

// StringTermsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L371-L376
type StringTermsAggregate struct {
	Buckets                 BucketsStringTermsBucket `json:"buckets"`
	DocCountErrorUpperBound *int64                   `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    *Metadata                `json:"meta,omitempty"`
	SumOtherDocCount        int64                    `json:"sum_other_doc_count"`
}

// StringTermsAggregateBuilder holds StringTermsAggregate struct and provides a builder API.
type StringTermsAggregateBuilder struct {
	v *StringTermsAggregate
}

// NewStringTermsAggregate provides a builder for the StringTermsAggregate struct.
func NewStringTermsAggregateBuilder() *StringTermsAggregateBuilder {
	r := StringTermsAggregateBuilder{
		&StringTermsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the StringTermsAggregate struct
func (rb *StringTermsAggregateBuilder) Build() StringTermsAggregate {
	return *rb.v
}

func (rb *StringTermsAggregateBuilder) Buckets(buckets *BucketsStringTermsBucketBuilder) *StringTermsAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *StringTermsAggregateBuilder) DocCountErrorUpperBound(doccounterrorupperbound int64) *StringTermsAggregateBuilder {
	rb.v.DocCountErrorUpperBound = &doccounterrorupperbound
	return rb
}

func (rb *StringTermsAggregateBuilder) Meta(meta *MetadataBuilder) *StringTermsAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *StringTermsAggregateBuilder) SumOtherDocCount(sumotherdoccount int64) *StringTermsAggregateBuilder {
	rb.v.SumOtherDocCount = sumotherdoccount
	return rb
}
