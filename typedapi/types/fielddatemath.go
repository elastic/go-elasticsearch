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

// FieldDateMath holds the union for the following types:
//
//	DateMath
//	float64
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L137-L144
type FieldDateMath interface{}

// FieldDateMathBuilder holds FieldDateMath struct and provides a builder API.
type FieldDateMathBuilder struct {
	v FieldDateMath
}

// NewFieldDateMath provides a builder for the FieldDateMath struct.
func NewFieldDateMathBuilder() *FieldDateMathBuilder {
	return &FieldDateMathBuilder{}
}

// Build finalize the chain and returns the FieldDateMath struct
func (u *FieldDateMathBuilder) Build() FieldDateMath {
	return u.v
}

func (u *FieldDateMathBuilder) DateMath(datemath DateMath) *FieldDateMathBuilder {
	u.v = &datemath
	return u
}

func (u *FieldDateMathBuilder) Float64(float64 float64) *FieldDateMathBuilder {
	u.v = &float64
	return u
}
