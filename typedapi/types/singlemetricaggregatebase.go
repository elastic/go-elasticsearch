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

// SingleMetricAggregateBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L172-L181
type SingleMetricAggregateBase struct {
	Meta *Metadata `json:"meta,omitempty"`
	// Value The metric value. A missing value generally means that there was no data to
	// aggregate,
	// unless specified otherwise.
	Value         float64 `json:"value,omitempty"`
	ValueAsString *string `json:"value_as_string,omitempty"`
}

// SingleMetricAggregateBaseBuilder holds SingleMetricAggregateBase struct and provides a builder API.
type SingleMetricAggregateBaseBuilder struct {
	v *SingleMetricAggregateBase
}

// NewSingleMetricAggregateBase provides a builder for the SingleMetricAggregateBase struct.
func NewSingleMetricAggregateBaseBuilder() *SingleMetricAggregateBaseBuilder {
	r := SingleMetricAggregateBaseBuilder{
		&SingleMetricAggregateBase{},
	}

	return &r
}

// Build finalize the chain and returns the SingleMetricAggregateBase struct
func (rb *SingleMetricAggregateBaseBuilder) Build() SingleMetricAggregateBase {
	return *rb.v
}

func (rb *SingleMetricAggregateBaseBuilder) Meta(meta *MetadataBuilder) *SingleMetricAggregateBaseBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

// Value The metric value. A missing value generally means that there was no data to
// aggregate,
// unless specified otherwise.

func (rb *SingleMetricAggregateBaseBuilder) Value(value float64) *SingleMetricAggregateBaseBuilder {
	rb.v.Value = value
	return rb
}

func (rb *SingleMetricAggregateBaseBuilder) ValueAsString(valueasstring string) *SingleMetricAggregateBaseBuilder {
	rb.v.ValueAsString = &valueasstring
	return rb
}
