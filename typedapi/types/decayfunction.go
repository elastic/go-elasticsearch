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

// DecayFunction holds the union for the following types:
//
//	DateDecayFunction
//	GeoDecayFunction
//	NumericDecayFunction
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L100-L105
type DecayFunction interface{}

// DecayFunctionBuilder holds DecayFunction struct and provides a builder API.
type DecayFunctionBuilder struct {
	v DecayFunction
}

// NewDecayFunction provides a builder for the DecayFunction struct.
func NewDecayFunctionBuilder() *DecayFunctionBuilder {
	return &DecayFunctionBuilder{}
}

// Build finalize the chain and returns the DecayFunction struct
func (u *DecayFunctionBuilder) Build() DecayFunction {
	return u.v
}

func (u *DecayFunctionBuilder) DateDecayFunction(datedecayfunction *DateDecayFunctionBuilder) *DecayFunctionBuilder {
	v := datedecayfunction.Build()
	u.v = &v
	return u
}

func (u *DecayFunctionBuilder) GeoDecayFunction(geodecayfunction *GeoDecayFunctionBuilder) *DecayFunctionBuilder {
	v := geodecayfunction.Build()
	u.v = &v
	return u
}

func (u *DecayFunctionBuilder) NumericDecayFunction(numericdecayfunction *NumericDecayFunctionBuilder) *DecayFunctionBuilder {
	v := numericdecayfunction.Build()
	u.v = &v
	return u
}
