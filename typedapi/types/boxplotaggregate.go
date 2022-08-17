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

// BoxPlotAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L657-L673
type BoxPlotAggregate struct {
	Lower         float64   `json:"lower"`
	LowerAsString *string   `json:"lower_as_string,omitempty"`
	Max           float64   `json:"max"`
	MaxAsString   *string   `json:"max_as_string,omitempty"`
	Meta          *Metadata `json:"meta,omitempty"`
	Min           float64   `json:"min"`
	MinAsString   *string   `json:"min_as_string,omitempty"`
	Q1            float64   `json:"q1"`
	Q1AsString    *string   `json:"q1_as_string,omitempty"`
	Q2            float64   `json:"q2"`
	Q2AsString    *string   `json:"q2_as_string,omitempty"`
	Q3            float64   `json:"q3"`
	Q3AsString    *string   `json:"q3_as_string,omitempty"`
	Upper         float64   `json:"upper"`
	UpperAsString *string   `json:"upper_as_string,omitempty"`
}

// BoxPlotAggregateBuilder holds BoxPlotAggregate struct and provides a builder API.
type BoxPlotAggregateBuilder struct {
	v *BoxPlotAggregate
}

// NewBoxPlotAggregate provides a builder for the BoxPlotAggregate struct.
func NewBoxPlotAggregateBuilder() *BoxPlotAggregateBuilder {
	r := BoxPlotAggregateBuilder{
		&BoxPlotAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the BoxPlotAggregate struct
func (rb *BoxPlotAggregateBuilder) Build() BoxPlotAggregate {
	return *rb.v
}

func (rb *BoxPlotAggregateBuilder) Lower(lower float64) *BoxPlotAggregateBuilder {
	rb.v.Lower = lower
	return rb
}

func (rb *BoxPlotAggregateBuilder) LowerAsString(lowerasstring string) *BoxPlotAggregateBuilder {
	rb.v.LowerAsString = &lowerasstring
	return rb
}

func (rb *BoxPlotAggregateBuilder) Max(max float64) *BoxPlotAggregateBuilder {
	rb.v.Max = max
	return rb
}

func (rb *BoxPlotAggregateBuilder) MaxAsString(maxasstring string) *BoxPlotAggregateBuilder {
	rb.v.MaxAsString = &maxasstring
	return rb
}

func (rb *BoxPlotAggregateBuilder) Meta(meta *MetadataBuilder) *BoxPlotAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *BoxPlotAggregateBuilder) Min(min float64) *BoxPlotAggregateBuilder {
	rb.v.Min = min
	return rb
}

func (rb *BoxPlotAggregateBuilder) MinAsString(minasstring string) *BoxPlotAggregateBuilder {
	rb.v.MinAsString = &minasstring
	return rb
}

func (rb *BoxPlotAggregateBuilder) Q1(q1 float64) *BoxPlotAggregateBuilder {
	rb.v.Q1 = q1
	return rb
}

func (rb *BoxPlotAggregateBuilder) Q1AsString(q1asstring string) *BoxPlotAggregateBuilder {
	rb.v.Q1AsString = &q1asstring
	return rb
}

func (rb *BoxPlotAggregateBuilder) Q2(q2 float64) *BoxPlotAggregateBuilder {
	rb.v.Q2 = q2
	return rb
}

func (rb *BoxPlotAggregateBuilder) Q2AsString(q2asstring string) *BoxPlotAggregateBuilder {
	rb.v.Q2AsString = &q2asstring
	return rb
}

func (rb *BoxPlotAggregateBuilder) Q3(q3 float64) *BoxPlotAggregateBuilder {
	rb.v.Q3 = q3
	return rb
}

func (rb *BoxPlotAggregateBuilder) Q3AsString(q3asstring string) *BoxPlotAggregateBuilder {
	rb.v.Q3AsString = &q3asstring
	return rb
}

func (rb *BoxPlotAggregateBuilder) Upper(upper float64) *BoxPlotAggregateBuilder {
	rb.v.Upper = upper
	return rb
}

func (rb *BoxPlotAggregateBuilder) UpperAsString(upperasstring string) *BoxPlotAggregateBuilder {
	rb.v.UpperAsString = &upperasstring
	return rb
}
