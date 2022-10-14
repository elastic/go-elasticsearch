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

// IpPrefixAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/bucket.ts#L513-L542
type IpPrefixAggregation struct {
	// AppendPrefixLength Defines whether the prefix length is appended to IP address keys in the
	// response.
	AppendPrefixLength *bool `json:"append_prefix_length,omitempty"`
	// Field The document IP address field to aggregation on. The field mapping type must
	// be `ip`
	Field Field `json:"field"`
	// IsIpv6 Defines whether the prefix applies to IPv6 addresses.
	IsIpv6 *bool `json:"is_ipv6,omitempty"`
	// Keyed Defines whether buckets are returned as a hash rather than an array in the
	// response.
	Keyed *bool     `json:"keyed,omitempty"`
	Meta  *Metadata `json:"meta,omitempty"`
	// MinDocCount Minimum number of documents for buckets to be included in the response.
	MinDocCount *int64  `json:"min_doc_count,omitempty"`
	Name        *string `json:"name,omitempty"`
	// PrefixLength Length of the network prefix. For IPv4 addresses the accepted range is [0,
	// 32].
	// For IPv6 addresses the accepted range is [0, 128].
	PrefixLength int `json:"prefix_length"`
}

// IpPrefixAggregationBuilder holds IpPrefixAggregation struct and provides a builder API.
type IpPrefixAggregationBuilder struct {
	v *IpPrefixAggregation
}

// NewIpPrefixAggregation provides a builder for the IpPrefixAggregation struct.
func NewIpPrefixAggregationBuilder() *IpPrefixAggregationBuilder {
	r := IpPrefixAggregationBuilder{
		&IpPrefixAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the IpPrefixAggregation struct
func (rb *IpPrefixAggregationBuilder) Build() IpPrefixAggregation {
	return *rb.v
}

// AppendPrefixLength Defines whether the prefix length is appended to IP address keys in the
// response.

func (rb *IpPrefixAggregationBuilder) AppendPrefixLength(appendprefixlength bool) *IpPrefixAggregationBuilder {
	rb.v.AppendPrefixLength = &appendprefixlength
	return rb
}

// Field The document IP address field to aggregation on. The field mapping type must
// be `ip`

func (rb *IpPrefixAggregationBuilder) Field(field Field) *IpPrefixAggregationBuilder {
	rb.v.Field = field
	return rb
}

// IsIpv6 Defines whether the prefix applies to IPv6 addresses.

func (rb *IpPrefixAggregationBuilder) IsIpv6(isipv6 bool) *IpPrefixAggregationBuilder {
	rb.v.IsIpv6 = &isipv6
	return rb
}

// Keyed Defines whether buckets are returned as a hash rather than an array in the
// response.

func (rb *IpPrefixAggregationBuilder) Keyed(keyed bool) *IpPrefixAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

func (rb *IpPrefixAggregationBuilder) Meta(meta *MetadataBuilder) *IpPrefixAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

// MinDocCount Minimum number of documents for buckets to be included in the response.

func (rb *IpPrefixAggregationBuilder) MinDocCount(mindoccount int64) *IpPrefixAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

func (rb *IpPrefixAggregationBuilder) Name(name string) *IpPrefixAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// PrefixLength Length of the network prefix. For IPv4 addresses the accepted range is [0,
// 32].
// For IPv6 addresses the accepted range is [0, 128].

func (rb *IpPrefixAggregationBuilder) PrefixLength(prefixlength int) *IpPrefixAggregationBuilder {
	rb.v.PrefixLength = prefixlength
	return rb
}
