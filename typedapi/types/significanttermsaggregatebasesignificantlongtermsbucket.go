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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

// SignificantTermsAggregateBaseSignificantLongTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L565-L570
type SignificantTermsAggregateBaseSignificantLongTermsBucket struct {
	BgCount  *int64                            `json:"bg_count,omitempty"`
	Buckets  BucketsSignificantLongTermsBucket `json:"buckets"`
	DocCount *int64                            `json:"doc_count,omitempty"`
	Meta     *Metadata                         `json:"meta,omitempty"`
}

// SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder holds SignificantTermsAggregateBaseSignificantLongTermsBucket struct and provides a builder API.
type SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder struct {
	v *SignificantTermsAggregateBaseSignificantLongTermsBucket
}

// NewSignificantTermsAggregateBaseSignificantLongTermsBucket provides a builder for the SignificantTermsAggregateBaseSignificantLongTermsBucket struct.
func NewSignificantTermsAggregateBaseSignificantLongTermsBucketBuilder() *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder {
	r := SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder{
		&SignificantTermsAggregateBaseSignificantLongTermsBucket{},
	}

	return &r
}

// Build finalize the chain and returns the SignificantTermsAggregateBaseSignificantLongTermsBucket struct
func (rb *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder) Build() SignificantTermsAggregateBaseSignificantLongTermsBucket {
	return *rb.v
}

func (rb *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder) BgCount(bgcount int64) *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder {
	rb.v.BgCount = &bgcount
	return rb
}

func (rb *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder) Buckets(buckets *BucketsSignificantLongTermsBucketBuilder) *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder) DocCount(doccount int64) *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder {
	rb.v.DocCount = &doccount
	return rb
}

func (rb *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder) Meta(meta *MetadataBuilder) *SignificantTermsAggregateBaseSignificantLongTermsBucketBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
