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

// DoubleTermsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L398-L403
type DoubleTermsAggregate struct {
	Buckets                 BucketsDoubleTermsBucket `json:"buckets"`
	DocCountErrorUpperBound *int64                   `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    *Metadata                `json:"meta,omitempty"`
	SumOtherDocCount        int64                    `json:"sum_other_doc_count"`
}

// DoubleTermsAggregateBuilder holds DoubleTermsAggregate struct and provides a builder API.
type DoubleTermsAggregateBuilder struct {
	v *DoubleTermsAggregate
}

// NewDoubleTermsAggregate provides a builder for the DoubleTermsAggregate struct.
func NewDoubleTermsAggregateBuilder() *DoubleTermsAggregateBuilder {
	r := DoubleTermsAggregateBuilder{
		&DoubleTermsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the DoubleTermsAggregate struct
func (rb *DoubleTermsAggregateBuilder) Build() DoubleTermsAggregate {
	return *rb.v
}

func (rb *DoubleTermsAggregateBuilder) Buckets(buckets *BucketsDoubleTermsBucketBuilder) *DoubleTermsAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *DoubleTermsAggregateBuilder) DocCountErrorUpperBound(doccounterrorupperbound int64) *DoubleTermsAggregateBuilder {
	rb.v.DocCountErrorUpperBound = &doccounterrorupperbound
	return rb
}

func (rb *DoubleTermsAggregateBuilder) Meta(meta *MetadataBuilder) *DoubleTermsAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *DoubleTermsAggregateBuilder) SumOtherDocCount(sumotherdoccount int64) *DoubleTermsAggregateBuilder {
	rb.v.SumOtherDocCount = sumotherdoccount
	return rb
}
