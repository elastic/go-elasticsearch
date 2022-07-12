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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// GeoLineAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/aggregations/Aggregate.ts#L735-L739
type GeoLineAggregate struct {
	Geometry GeoLine   `json:"geometry"`
	Meta     *Metadata `json:"meta,omitempty"`
	Type     string    `json:"type"`
}

// GeoLineAggregateBuilder holds GeoLineAggregate struct and provides a builder API.
type GeoLineAggregateBuilder struct {
	v *GeoLineAggregate
}

// NewGeoLineAggregate provides a builder for the GeoLineAggregate struct.
func NewGeoLineAggregateBuilder() *GeoLineAggregateBuilder {
	r := GeoLineAggregateBuilder{
		&GeoLineAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the GeoLineAggregate struct
func (rb *GeoLineAggregateBuilder) Build() GeoLineAggregate {
	return *rb.v
}

func (rb *GeoLineAggregateBuilder) Geometry(geometry *GeoLineBuilder) *GeoLineAggregateBuilder {
	v := geometry.Build()
	rb.v.Geometry = v
	return rb
}

func (rb *GeoLineAggregateBuilder) Meta(meta *MetadataBuilder) *GeoLineAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *GeoLineAggregateBuilder) Type_(type_ string) *GeoLineAggregateBuilder {
	rb.v.Type = type_
	return rb
}
