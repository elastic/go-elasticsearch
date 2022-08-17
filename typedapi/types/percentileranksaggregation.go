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

// PercentileRanksAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L105-L110
type PercentileRanksAggregation struct {
	Field   *Field     `json:"field,omitempty"`
	Format  *string    `json:"format,omitempty"`
	Hdr     *HdrMethod `json:"hdr,omitempty"`
	Keyed   *bool      `json:"keyed,omitempty"`
	Missing *Missing   `json:"missing,omitempty"`
	Script  *Script    `json:"script,omitempty"`
	Tdigest *TDigest   `json:"tdigest,omitempty"`
	Values  []float64  `json:"values,omitempty"`
}

// PercentileRanksAggregationBuilder holds PercentileRanksAggregation struct and provides a builder API.
type PercentileRanksAggregationBuilder struct {
	v *PercentileRanksAggregation
}

// NewPercentileRanksAggregation provides a builder for the PercentileRanksAggregation struct.
func NewPercentileRanksAggregationBuilder() *PercentileRanksAggregationBuilder {
	r := PercentileRanksAggregationBuilder{
		&PercentileRanksAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the PercentileRanksAggregation struct
func (rb *PercentileRanksAggregationBuilder) Build() PercentileRanksAggregation {
	return *rb.v
}

func (rb *PercentileRanksAggregationBuilder) Field(field Field) *PercentileRanksAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *PercentileRanksAggregationBuilder) Format(format string) *PercentileRanksAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *PercentileRanksAggregationBuilder) Hdr(hdr *HdrMethodBuilder) *PercentileRanksAggregationBuilder {
	v := hdr.Build()
	rb.v.Hdr = &v
	return rb
}

func (rb *PercentileRanksAggregationBuilder) Keyed(keyed bool) *PercentileRanksAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

func (rb *PercentileRanksAggregationBuilder) Missing(missing *MissingBuilder) *PercentileRanksAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *PercentileRanksAggregationBuilder) Script(script *ScriptBuilder) *PercentileRanksAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *PercentileRanksAggregationBuilder) Tdigest(tdigest *TDigestBuilder) *PercentileRanksAggregationBuilder {
	v := tdigest.Build()
	rb.v.Tdigest = &v
	return rb
}

func (rb *PercentileRanksAggregationBuilder) Values(values []float64) *PercentileRanksAggregationBuilder {
	rb.v.Values = values
	return rb
}
