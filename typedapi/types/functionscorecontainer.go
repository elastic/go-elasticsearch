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

// FunctionScoreContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L107-L127
type FunctionScoreContainer struct {
	Exp              *DecayFunction                 `json:"exp,omitempty"`
	FieldValueFactor *FieldValueFactorScoreFunction `json:"field_value_factor,omitempty"`
	Filter           *QueryContainer                `json:"filter,omitempty"`
	Gauss            *DecayFunction                 `json:"gauss,omitempty"`
	Linear           *DecayFunction                 `json:"linear,omitempty"`
	RandomScore      *RandomScoreFunction           `json:"random_score,omitempty"`
	ScriptScore      *ScriptScoreFunction           `json:"script_score,omitempty"`
	Weight           *float64                       `json:"weight,omitempty"`
}

// FunctionScoreContainerBuilder holds FunctionScoreContainer struct and provides a builder API.
type FunctionScoreContainerBuilder struct {
	v *FunctionScoreContainer
}

// NewFunctionScoreContainer provides a builder for the FunctionScoreContainer struct.
func NewFunctionScoreContainerBuilder() *FunctionScoreContainerBuilder {
	r := FunctionScoreContainerBuilder{
		&FunctionScoreContainer{},
	}

	return &r
}

// Build finalize the chain and returns the FunctionScoreContainer struct
func (rb *FunctionScoreContainerBuilder) Build() FunctionScoreContainer {
	return *rb.v
}

func (rb *FunctionScoreContainerBuilder) Exp(exp *DecayFunctionBuilder) *FunctionScoreContainerBuilder {
	v := exp.Build()
	rb.v.Exp = &v
	return rb
}

func (rb *FunctionScoreContainerBuilder) FieldValueFactor(fieldvaluefactor *FieldValueFactorScoreFunctionBuilder) *FunctionScoreContainerBuilder {
	v := fieldvaluefactor.Build()
	rb.v.FieldValueFactor = &v
	return rb
}

func (rb *FunctionScoreContainerBuilder) Filter(filter *QueryContainerBuilder) *FunctionScoreContainerBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *FunctionScoreContainerBuilder) Gauss(gauss *DecayFunctionBuilder) *FunctionScoreContainerBuilder {
	v := gauss.Build()
	rb.v.Gauss = &v
	return rb
}

func (rb *FunctionScoreContainerBuilder) Linear(linear *DecayFunctionBuilder) *FunctionScoreContainerBuilder {
	v := linear.Build()
	rb.v.Linear = &v
	return rb
}

func (rb *FunctionScoreContainerBuilder) RandomScore(randomscore *RandomScoreFunctionBuilder) *FunctionScoreContainerBuilder {
	v := randomscore.Build()
	rb.v.RandomScore = &v
	return rb
}

func (rb *FunctionScoreContainerBuilder) ScriptScore(scriptscore *ScriptScoreFunctionBuilder) *FunctionScoreContainerBuilder {
	v := scriptscore.Build()
	rb.v.ScriptScore = &v
	return rb
}

func (rb *FunctionScoreContainerBuilder) Weight(weight float64) *FunctionScoreContainerBuilder {
	rb.v.Weight = &weight
	return rb
}
