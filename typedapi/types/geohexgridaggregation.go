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

// GeohexGridAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L197-L223
type GeohexGridAggregation struct {
	// Bounds Bounding box used to filter the geo-points in each bucket.
	Bounds *GeoBounds `json:"bounds,omitempty"`
	// Field Field containing indexed geo-point values. Must be explicitly
	// mapped as a `geo_point` field. If the field contains an array
	// `geohex_grid` aggregates all array values.
	Field Field     `json:"field"`
	Meta  *Metadata `json:"meta,omitempty"`
	Name  *string   `json:"name,omitempty"`
	// Precision Integer zoom of the key used to defined cells or buckets
	// in the results. Value should be between 0-15.
	Precision *int `json:"precision,omitempty"`
	// ShardSize Number of buckets returned from each shard.
	ShardSize *int `json:"shard_size,omitempty"`
	// Size Maximum number of buckets to return.
	Size *int `json:"size,omitempty"`
}

// GeohexGridAggregationBuilder holds GeohexGridAggregation struct and provides a builder API.
type GeohexGridAggregationBuilder struct {
	v *GeohexGridAggregation
}

// NewGeohexGridAggregation provides a builder for the GeohexGridAggregation struct.
func NewGeohexGridAggregationBuilder() *GeohexGridAggregationBuilder {
	r := GeohexGridAggregationBuilder{
		&GeohexGridAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the GeohexGridAggregation struct
func (rb *GeohexGridAggregationBuilder) Build() GeohexGridAggregation {
	return *rb.v
}

// Bounds Bounding box used to filter the geo-points in each bucket.

func (rb *GeohexGridAggregationBuilder) Bounds(bounds *GeoBoundsBuilder) *GeohexGridAggregationBuilder {
	v := bounds.Build()
	rb.v.Bounds = &v
	return rb
}

// Field Field containing indexed geo-point values. Must be explicitly
// mapped as a `geo_point` field. If the field contains an array
// `geohex_grid` aggregates all array values.

func (rb *GeohexGridAggregationBuilder) Field(field Field) *GeohexGridAggregationBuilder {
	rb.v.Field = field
	return rb
}

func (rb *GeohexGridAggregationBuilder) Meta(meta *MetadataBuilder) *GeohexGridAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *GeohexGridAggregationBuilder) Name(name string) *GeohexGridAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Precision Integer zoom of the key used to defined cells or buckets
// in the results. Value should be between 0-15.

func (rb *GeohexGridAggregationBuilder) Precision(precision int) *GeohexGridAggregationBuilder {
	rb.v.Precision = &precision
	return rb
}

// ShardSize Number of buckets returned from each shard.

func (rb *GeohexGridAggregationBuilder) ShardSize(shardsize int) *GeohexGridAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

// Size Maximum number of buckets to return.

func (rb *GeohexGridAggregationBuilder) Size(size int) *GeohexGridAggregationBuilder {
	rb.v.Size = &size
	return rb
}
