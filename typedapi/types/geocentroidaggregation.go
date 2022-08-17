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

// GeoCentroidAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L76-L79
type GeoCentroidAggregation struct {
	Count    *int64       `json:"count,omitempty"`
	Field    *Field       `json:"field,omitempty"`
	Location *GeoLocation `json:"location,omitempty"`
	Missing  *Missing     `json:"missing,omitempty"`
	Script   *Script      `json:"script,omitempty"`
}

// GeoCentroidAggregationBuilder holds GeoCentroidAggregation struct and provides a builder API.
type GeoCentroidAggregationBuilder struct {
	v *GeoCentroidAggregation
}

// NewGeoCentroidAggregation provides a builder for the GeoCentroidAggregation struct.
func NewGeoCentroidAggregationBuilder() *GeoCentroidAggregationBuilder {
	r := GeoCentroidAggregationBuilder{
		&GeoCentroidAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the GeoCentroidAggregation struct
func (rb *GeoCentroidAggregationBuilder) Build() GeoCentroidAggregation {
	return *rb.v
}

func (rb *GeoCentroidAggregationBuilder) Count(count int64) *GeoCentroidAggregationBuilder {
	rb.v.Count = &count
	return rb
}

func (rb *GeoCentroidAggregationBuilder) Field(field Field) *GeoCentroidAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *GeoCentroidAggregationBuilder) Location(location *GeoLocationBuilder) *GeoCentroidAggregationBuilder {
	v := location.Build()
	rb.v.Location = &v
	return rb
}

func (rb *GeoCentroidAggregationBuilder) Missing(missing *MissingBuilder) *GeoCentroidAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *GeoCentroidAggregationBuilder) Script(script *ScriptBuilder) *GeoCentroidAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
