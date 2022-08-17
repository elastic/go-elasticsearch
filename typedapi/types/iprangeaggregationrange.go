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

// IpRangeAggregationRange type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L251-L255
type IpRangeAggregationRange struct {
	From string  `json:"from,omitempty"`
	Mask *string `json:"mask,omitempty"`
	To   string  `json:"to,omitempty"`
}

// IpRangeAggregationRangeBuilder holds IpRangeAggregationRange struct and provides a builder API.
type IpRangeAggregationRangeBuilder struct {
	v *IpRangeAggregationRange
}

// NewIpRangeAggregationRange provides a builder for the IpRangeAggregationRange struct.
func NewIpRangeAggregationRangeBuilder() *IpRangeAggregationRangeBuilder {
	r := IpRangeAggregationRangeBuilder{
		&IpRangeAggregationRange{},
	}

	return &r
}

// Build finalize the chain and returns the IpRangeAggregationRange struct
func (rb *IpRangeAggregationRangeBuilder) Build() IpRangeAggregationRange {
	return *rb.v
}

func (rb *IpRangeAggregationRangeBuilder) From(from string) *IpRangeAggregationRangeBuilder {
	rb.v.From = from
	return rb
}

func (rb *IpRangeAggregationRangeBuilder) Mask(mask string) *IpRangeAggregationRangeBuilder {
	rb.v.Mask = &mask
	return rb
}

func (rb *IpRangeAggregationRangeBuilder) To(to string) *IpRangeAggregationRangeBuilder {
	rb.v.To = to
	return rb
}
