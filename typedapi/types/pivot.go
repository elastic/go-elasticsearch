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

// Pivot type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/_types/Transform.ts#L54-L68
type Pivot struct {
	// Aggregations Defines how to aggregate the grouped data. The following aggregations are
	// currently supported: average, bucket
	// script, bucket selector, cardinality, filter, geo bounds, geo centroid, geo
	// line, max, median absolute deviation,
	// min, missing, percentiles, rare terms, scripted metric, stats, sum, terms,
	// top metrics, value count, weighted
	// average.
	Aggregations map[string]AggregationContainer `json:"aggregations,omitempty"`
	// GroupBy Defines how to group the data. More than one grouping can be defined per
	// pivot. The following groupings are
	// currently supported: date histogram, geotile grid, histogram, terms.
	GroupBy map[string]PivotGroupByContainer `json:"group_by,omitempty"`
}

// PivotBuilder holds Pivot struct and provides a builder API.
type PivotBuilder struct {
	v *Pivot
}

// NewPivot provides a builder for the Pivot struct.
func NewPivotBuilder() *PivotBuilder {
	r := PivotBuilder{
		&Pivot{
			Aggregations: make(map[string]AggregationContainer, 0),
			GroupBy:      make(map[string]PivotGroupByContainer, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Pivot struct
func (rb *PivotBuilder) Build() Pivot {
	return *rb.v
}

// Aggregations Defines how to aggregate the grouped data. The following aggregations are
// currently supported: average, bucket
// script, bucket selector, cardinality, filter, geo bounds, geo centroid, geo
// line, max, median absolute deviation,
// min, missing, percentiles, rare terms, scripted metric, stats, sum, terms,
// top metrics, value count, weighted
// average.

func (rb *PivotBuilder) Aggregations(values map[string]*AggregationContainerBuilder) *PivotBuilder {
	tmp := make(map[string]AggregationContainer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

// GroupBy Defines how to group the data. More than one grouping can be defined per
// pivot. The following groupings are
// currently supported: date histogram, geotile grid, histogram, terms.

func (rb *PivotBuilder) GroupBy(values map[string]*PivotGroupByContainerBuilder) *PivotBuilder {
	tmp := make(map[string]PivotGroupByContainer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.GroupBy = tmp
	return rb
}
