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

// TopMetrics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L680-L684
type TopMetrics struct {
	Metrics map[string]FieldValue `json:"metrics"`
	Sort    []FieldValue          `json:"sort"`
}

// TopMetricsBuilder holds TopMetrics struct and provides a builder API.
type TopMetricsBuilder struct {
	v *TopMetrics
}

// NewTopMetrics provides a builder for the TopMetrics struct.
func NewTopMetricsBuilder() *TopMetricsBuilder {
	r := TopMetricsBuilder{
		&TopMetrics{
			Metrics: make(map[string]FieldValue, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TopMetrics struct
func (rb *TopMetricsBuilder) Build() TopMetrics {
	return *rb.v
}

func (rb *TopMetricsBuilder) Metrics(value map[string]FieldValue) *TopMetricsBuilder {
	rb.v.Metrics = value
	return rb
}

func (rb *TopMetricsBuilder) Sort(sort ...FieldValue) *TopMetricsBuilder {
	rb.v.Sort = sort
	return rb
}
