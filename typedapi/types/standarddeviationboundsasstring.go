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

// StandardDeviationBoundsAsString type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L258-L265
type StandardDeviationBoundsAsString struct {
	Lower           string `json:"lower"`
	LowerPopulation string `json:"lower_population"`
	LowerSampling   string `json:"lower_sampling"`
	Upper           string `json:"upper"`
	UpperPopulation string `json:"upper_population"`
	UpperSampling   string `json:"upper_sampling"`
}

// StandardDeviationBoundsAsStringBuilder holds StandardDeviationBoundsAsString struct and provides a builder API.
type StandardDeviationBoundsAsStringBuilder struct {
	v *StandardDeviationBoundsAsString
}

// NewStandardDeviationBoundsAsString provides a builder for the StandardDeviationBoundsAsString struct.
func NewStandardDeviationBoundsAsStringBuilder() *StandardDeviationBoundsAsStringBuilder {
	r := StandardDeviationBoundsAsStringBuilder{
		&StandardDeviationBoundsAsString{},
	}

	return &r
}

// Build finalize the chain and returns the StandardDeviationBoundsAsString struct
func (rb *StandardDeviationBoundsAsStringBuilder) Build() StandardDeviationBoundsAsString {
	return *rb.v
}

func (rb *StandardDeviationBoundsAsStringBuilder) Lower(lower string) *StandardDeviationBoundsAsStringBuilder {
	rb.v.Lower = lower
	return rb
}

func (rb *StandardDeviationBoundsAsStringBuilder) LowerPopulation(lowerpopulation string) *StandardDeviationBoundsAsStringBuilder {
	rb.v.LowerPopulation = lowerpopulation
	return rb
}

func (rb *StandardDeviationBoundsAsStringBuilder) LowerSampling(lowersampling string) *StandardDeviationBoundsAsStringBuilder {
	rb.v.LowerSampling = lowersampling
	return rb
}

func (rb *StandardDeviationBoundsAsStringBuilder) Upper(upper string) *StandardDeviationBoundsAsStringBuilder {
	rb.v.Upper = upper
	return rb
}

func (rb *StandardDeviationBoundsAsStringBuilder) UpperPopulation(upperpopulation string) *StandardDeviationBoundsAsStringBuilder {
	rb.v.UpperPopulation = upperpopulation
	return rb
}

func (rb *StandardDeviationBoundsAsStringBuilder) UpperSampling(uppersampling string) *StandardDeviationBoundsAsStringBuilder {
	rb.v.UpperSampling = uppersampling
	return rb
}
