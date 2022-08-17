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

// GeoCentroidAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L295-L299
type GeoCentroidAggregate struct {
	Count    int64        `json:"count"`
	Location *GeoLocation `json:"location,omitempty"`
	Meta     *Metadata    `json:"meta,omitempty"`
}

// GeoCentroidAggregateBuilder holds GeoCentroidAggregate struct and provides a builder API.
type GeoCentroidAggregateBuilder struct {
	v *GeoCentroidAggregate
}

// NewGeoCentroidAggregate provides a builder for the GeoCentroidAggregate struct.
func NewGeoCentroidAggregateBuilder() *GeoCentroidAggregateBuilder {
	r := GeoCentroidAggregateBuilder{
		&GeoCentroidAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the GeoCentroidAggregate struct
func (rb *GeoCentroidAggregateBuilder) Build() GeoCentroidAggregate {
	return *rb.v
}

func (rb *GeoCentroidAggregateBuilder) Count(count int64) *GeoCentroidAggregateBuilder {
	rb.v.Count = count
	return rb
}

func (rb *GeoCentroidAggregateBuilder) Location(location *GeoLocationBuilder) *GeoCentroidAggregateBuilder {
	v := location.Build()
	rb.v.Location = &v
	return rb
}

func (rb *GeoCentroidAggregateBuilder) Meta(meta *MetadataBuilder) *GeoCentroidAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
