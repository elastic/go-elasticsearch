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

// PredictedValue holds the union for the following types:
//
//	bool
//	float64
//	int
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L410-L410
type PredictedValue interface{}

// PredictedValueBuilder holds PredictedValue struct and provides a builder API.
type PredictedValueBuilder struct {
	v PredictedValue
}

// NewPredictedValue provides a builder for the PredictedValue struct.
func NewPredictedValueBuilder() *PredictedValueBuilder {
	return &PredictedValueBuilder{}
}

// Build finalize the chain and returns the PredictedValue struct
func (u *PredictedValueBuilder) Build() PredictedValue {
	return u.v
}

func (u *PredictedValueBuilder) Bool(bool bool) *PredictedValueBuilder {
	u.v = &bool
	return u
}

func (u *PredictedValueBuilder) Float64(float64 float64) *PredictedValueBuilder {
	u.v = &float64
	return u
}

func (u *PredictedValueBuilder) Int(int int) *PredictedValueBuilder {
	u.v = &int
	return u
}

func (u *PredictedValueBuilder) String(string string) *PredictedValueBuilder {
	u.v = &string
	return u
}
