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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// ExtendedBoundsFieldDateMath type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/aggregations/bucket.ts#L230-L233
type ExtendedBoundsFieldDateMath struct {
	Max FieldDateMath `json:"max"`
	Min FieldDateMath `json:"min"`
}

// ExtendedBoundsFieldDateMathBuilder holds ExtendedBoundsFieldDateMath struct and provides a builder API.
type ExtendedBoundsFieldDateMathBuilder struct {
	v *ExtendedBoundsFieldDateMath
}

// NewExtendedBoundsFieldDateMath provides a builder for the ExtendedBoundsFieldDateMath struct.
func NewExtendedBoundsFieldDateMathBuilder() *ExtendedBoundsFieldDateMathBuilder {
	r := ExtendedBoundsFieldDateMathBuilder{
		&ExtendedBoundsFieldDateMath{},
	}

	return &r
}

// Build finalize the chain and returns the ExtendedBoundsFieldDateMath struct
func (rb *ExtendedBoundsFieldDateMathBuilder) Build() ExtendedBoundsFieldDateMath {
	return *rb.v
}

func (rb *ExtendedBoundsFieldDateMathBuilder) Max(max *FieldDateMathBuilder) *ExtendedBoundsFieldDateMathBuilder {
	v := max.Build()
	rb.v.Max = v
	return rb
}

func (rb *ExtendedBoundsFieldDateMathBuilder) Min(min *FieldDateMathBuilder) *ExtendedBoundsFieldDateMathBuilder {
	v := min.Build()
	rb.v.Min = v
	return rb
}
