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

// Missing holds the union for the following types:
//
//	bool
//	float64
//	int
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/AggregationContainer.ts#L209-L209
type Missing interface{}

// MissingBuilder holds Missing struct and provides a builder API.
type MissingBuilder struct {
	v Missing
}

// NewMissing provides a builder for the Missing struct.
func NewMissingBuilder() *MissingBuilder {
	return &MissingBuilder{}
}

// Build finalize the chain and returns the Missing struct
func (u *MissingBuilder) Build() Missing {
	return u.v
}

func (u *MissingBuilder) Bool(bool bool) *MissingBuilder {
	u.v = &bool
	return u
}

func (u *MissingBuilder) Float64(float64 float64) *MissingBuilder {
	u.v = &float64
	return u
}

func (u *MissingBuilder) Int(int int) *MissingBuilder {
	u.v = &int
	return u
}

func (u *MissingBuilder) String(string string) *MissingBuilder {
	u.v = &string
	return u
}
