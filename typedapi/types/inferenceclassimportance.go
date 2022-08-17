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

// InferenceClassImportance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/Aggregate.ts#L635-L638
type InferenceClassImportance struct {
	ClassName  string  `json:"class_name"`
	Importance float64 `json:"importance"`
}

// InferenceClassImportanceBuilder holds InferenceClassImportance struct and provides a builder API.
type InferenceClassImportanceBuilder struct {
	v *InferenceClassImportance
}

// NewInferenceClassImportance provides a builder for the InferenceClassImportance struct.
func NewInferenceClassImportanceBuilder() *InferenceClassImportanceBuilder {
	r := InferenceClassImportanceBuilder{
		&InferenceClassImportance{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceClassImportance struct
func (rb *InferenceClassImportanceBuilder) Build() InferenceClassImportance {
	return *rb.v
}

func (rb *InferenceClassImportanceBuilder) ClassName(classname string) *InferenceClassImportanceBuilder {
	rb.v.ClassName = classname
	return rb
}

func (rb *InferenceClassImportanceBuilder) Importance(importance float64) *InferenceClassImportanceBuilder {
	rb.v.Importance = importance
	return rb
}
