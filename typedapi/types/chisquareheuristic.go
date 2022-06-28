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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// ChiSquareHeuristic type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/bucket.ts#L315-L318
type ChiSquareHeuristic struct {
	BackgroundIsSuperset bool `json:"background_is_superset"`
	IncludeNegatives     bool `json:"include_negatives"`
}

// ChiSquareHeuristicBuilder holds ChiSquareHeuristic struct and provides a builder API.
type ChiSquareHeuristicBuilder struct {
	v *ChiSquareHeuristic
}

// NewChiSquareHeuristic provides a builder for the ChiSquareHeuristic struct.
func NewChiSquareHeuristicBuilder() *ChiSquareHeuristicBuilder {
	r := ChiSquareHeuristicBuilder{
		&ChiSquareHeuristic{},
	}

	return &r
}

// Build finalize the chain and returns the ChiSquareHeuristic struct
func (rb *ChiSquareHeuristicBuilder) Build() ChiSquareHeuristic {
	return *rb.v
}

func (rb *ChiSquareHeuristicBuilder) BackgroundIsSuperset(backgroundissuperset bool) *ChiSquareHeuristicBuilder {
	rb.v.BackgroundIsSuperset = backgroundissuperset
	return rb
}

func (rb *ChiSquareHeuristicBuilder) IncludeNegatives(includenegatives bool) *ChiSquareHeuristicBuilder {
	rb.v.IncludeNegatives = includenegatives
	return rb
}
