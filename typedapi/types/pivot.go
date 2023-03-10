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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// Pivot type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/transform/_types/Transform.ts#L54-L68
type Pivot struct {
	// Aggregations Defines how to aggregate the grouped data. The following aggregations are
	// currently supported: average, bucket
	// script, bucket selector, cardinality, filter, geo bounds, geo centroid, geo
	// line, max, median absolute deviation,
	// min, missing, percentiles, rare terms, scripted metric, stats, sum, terms,
	// top metrics, value count, weighted
	// average.
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`
	// GroupBy Defines how to group the data. More than one grouping can be defined per
	// pivot. The following groupings are
	// currently supported: date histogram, geotile grid, histogram, terms.
	GroupBy map[string]PivotGroupByContainer `json:"group_by,omitempty"`
}

// NewPivot returns a Pivot.
func NewPivot() *Pivot {
	r := &Pivot{
		Aggregations: make(map[string]Aggregations, 0),
		GroupBy:      make(map[string]PivotGroupByContainer, 0),
	}

	return r
}
