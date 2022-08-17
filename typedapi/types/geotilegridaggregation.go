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

// GeoTileGridAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L189-L195
type GeoTileGridAggregation struct {
	Bounds    *GeoBounds        `json:"bounds,omitempty"`
	Field     *Field            `json:"field,omitempty"`
	Meta      *Metadata         `json:"meta,omitempty"`
	Name      *string           `json:"name,omitempty"`
	Precision *GeoTilePrecision `json:"precision,omitempty"`
	ShardSize *int              `json:"shard_size,omitempty"`
	Size      *int              `json:"size,omitempty"`
}

// GeoTileGridAggregationBuilder holds GeoTileGridAggregation struct and provides a builder API.
type GeoTileGridAggregationBuilder struct {
	v *GeoTileGridAggregation
}

// NewGeoTileGridAggregation provides a builder for the GeoTileGridAggregation struct.
func NewGeoTileGridAggregationBuilder() *GeoTileGridAggregationBuilder {
	r := GeoTileGridAggregationBuilder{
		&GeoTileGridAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the GeoTileGridAggregation struct
func (rb *GeoTileGridAggregationBuilder) Build() GeoTileGridAggregation {
	return *rb.v
}

func (rb *GeoTileGridAggregationBuilder) Bounds(bounds *GeoBoundsBuilder) *GeoTileGridAggregationBuilder {
	v := bounds.Build()
	rb.v.Bounds = &v
	return rb
}

func (rb *GeoTileGridAggregationBuilder) Field(field Field) *GeoTileGridAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *GeoTileGridAggregationBuilder) Meta(meta *MetadataBuilder) *GeoTileGridAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *GeoTileGridAggregationBuilder) Name(name string) *GeoTileGridAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *GeoTileGridAggregationBuilder) Precision(precision GeoTilePrecision) *GeoTileGridAggregationBuilder {
	rb.v.Precision = &precision
	return rb
}

func (rb *GeoTileGridAggregationBuilder) ShardSize(shardsize int) *GeoTileGridAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

func (rb *GeoTileGridAggregationBuilder) Size(size int) *GeoTileGridAggregationBuilder {
	rb.v.Size = &size
	return rb
}
