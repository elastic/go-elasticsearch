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

// RangeAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/aggregations/bucket.ts#L289-L296
type RangeAggregation struct {
	Field   *string                `json:"field,omitempty"`
	Format  *string                `json:"format,omitempty"`
	Keyed   *bool                  `json:"keyed,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
	Missing *int                   `json:"missing,omitempty"`
	Name    *string                `json:"name,omitempty"`
	Ranges  []AggregationRange     `json:"ranges,omitempty"`
	Script  *Script                `json:"script,omitempty"`
}

// NewRangeAggregation returns a RangeAggregation.
func NewRangeAggregation() *RangeAggregation {
	r := &RangeAggregation{}

	return r
}
