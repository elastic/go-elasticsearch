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

// GeoHashGridAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L181-L187
type GeoHashGridAggregation struct {
	Bounds    *GeoBounds        `json:"bounds,omitempty"`
	Field     *Field            `json:"field,omitempty"`
	Meta      *Metadata         `json:"meta,omitempty"`
	Name      *string           `json:"name,omitempty"`
	Precision *GeoHashPrecision `json:"precision,omitempty"`
	ShardSize *int              `json:"shard_size,omitempty"`
	Size      *int              `json:"size,omitempty"`
}

// GeoHashGridAggregationBuilder holds GeoHashGridAggregation struct and provides a builder API.
type GeoHashGridAggregationBuilder struct {
	v *GeoHashGridAggregation
}

// NewGeoHashGridAggregation provides a builder for the GeoHashGridAggregation struct.
func NewGeoHashGridAggregationBuilder() *GeoHashGridAggregationBuilder {
	r := GeoHashGridAggregationBuilder{
		&GeoHashGridAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the GeoHashGridAggregation struct
func (rb *GeoHashGridAggregationBuilder) Build() GeoHashGridAggregation {
	return *rb.v
}

func (rb *GeoHashGridAggregationBuilder) Bounds(bounds *GeoBoundsBuilder) *GeoHashGridAggregationBuilder {
	v := bounds.Build()
	rb.v.Bounds = &v
	return rb
}

func (rb *GeoHashGridAggregationBuilder) Field(field Field) *GeoHashGridAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *GeoHashGridAggregationBuilder) Meta(meta *MetadataBuilder) *GeoHashGridAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *GeoHashGridAggregationBuilder) Name(name string) *GeoHashGridAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *GeoHashGridAggregationBuilder) Precision(precision *GeoHashPrecisionBuilder) *GeoHashGridAggregationBuilder {
	v := precision.Build()
	rb.v.Precision = &v
	return rb
}

func (rb *GeoHashGridAggregationBuilder) ShardSize(shardsize int) *GeoHashGridAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

func (rb *GeoHashGridAggregationBuilder) Size(size int) *GeoHashGridAggregationBuilder {
	rb.v.Size = &size
	return rb
}
