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

// Definition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/put_trained_model/types.ts#L24-L29
type Definition struct {
	// Preprocessors Collection of preprocessors
	Preprocessors []Preprocessor `json:"preprocessors,omitempty"`
	// TrainedModel The definition of the trained model.
	TrainedModel TrainedModel `json:"trained_model"`
}

// DefinitionBuilder holds Definition struct and provides a builder API.
type DefinitionBuilder struct {
	v *Definition
}

// NewDefinition provides a builder for the Definition struct.
func NewDefinitionBuilder() *DefinitionBuilder {
	r := DefinitionBuilder{
		&Definition{},
	}

	return &r
}

// Build finalize the chain and returns the Definition struct
func (rb *DefinitionBuilder) Build() Definition {
	return *rb.v
}

// Preprocessors Collection of preprocessors

func (rb *DefinitionBuilder) Preprocessors(preprocessors []PreprocessorBuilder) *DefinitionBuilder {
	tmp := make([]Preprocessor, len(preprocessors))
	for _, value := range preprocessors {
		tmp = append(tmp, value.Build())
	}
	rb.v.Preprocessors = tmp
	return rb
}

// TrainedModel The definition of the trained model.

func (rb *DefinitionBuilder) TrainedModel(trainedmodel *TrainedModelBuilder) *DefinitionBuilder {
	v := trainedmodel.Build()
	rb.v.TrainedModel = v
	return rb
}
