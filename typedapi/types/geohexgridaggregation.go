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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// GeohexGridAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/aggregations/bucket.ts#L200-L226
type GeohexGridAggregation struct {
	// Bounds Bounding box used to filter the geo-points in each bucket.
	Bounds *GeoBounds `json:"bounds,omitempty"`
	// Field Field containing indexed geo-point values. Must be explicitly
	// mapped as a `geo_point` field. If the field contains an array
	// `geohex_grid` aggregates all array values.
	Field string                 `json:"field"`
	Meta  map[string]interface{} `json:"meta,omitempty"`
	Name  *string                `json:"name,omitempty"`
	// Precision Integer zoom of the key used to defined cells or buckets
	// in the results. Value should be between 0-15.
	Precision *int `json:"precision,omitempty"`
	// ShardSize Number of buckets returned from each shard.
	ShardSize *int `json:"shard_size,omitempty"`
	// Size Maximum number of buckets to return.
	Size *int `json:"size,omitempty"`
}

// NewGeohexGridAggregation returns a GeohexGridAggregation.
func NewGeohexGridAggregation() *GeohexGridAggregation {
	r := &GeohexGridAggregation{}

	return r
}
