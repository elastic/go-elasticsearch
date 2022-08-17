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

// TopMetricsValue type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L192-L194
type TopMetricsValue struct {
	Field Field `json:"field"`
}

// TopMetricsValueBuilder holds TopMetricsValue struct and provides a builder API.
type TopMetricsValueBuilder struct {
	v *TopMetricsValue
}

// NewTopMetricsValue provides a builder for the TopMetricsValue struct.
func NewTopMetricsValueBuilder() *TopMetricsValueBuilder {
	r := TopMetricsValueBuilder{
		&TopMetricsValue{},
	}

	return &r
}

// Build finalize the chain and returns the TopMetricsValue struct
func (rb *TopMetricsValueBuilder) Build() TopMetricsValue {
	return *rb.v
}

func (rb *TopMetricsValueBuilder) Field(field Field) *TopMetricsValueBuilder {
	rb.v.Field = field
	return rb
}
