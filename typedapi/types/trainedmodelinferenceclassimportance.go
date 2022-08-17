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

// TrainedModelInferenceClassImportance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L399-L402
type TrainedModelInferenceClassImportance struct {
	ClassName  string  `json:"class_name"`
	Importance float64 `json:"importance"`
}

// TrainedModelInferenceClassImportanceBuilder holds TrainedModelInferenceClassImportance struct and provides a builder API.
type TrainedModelInferenceClassImportanceBuilder struct {
	v *TrainedModelInferenceClassImportance
}

// NewTrainedModelInferenceClassImportance provides a builder for the TrainedModelInferenceClassImportance struct.
func NewTrainedModelInferenceClassImportanceBuilder() *TrainedModelInferenceClassImportanceBuilder {
	r := TrainedModelInferenceClassImportanceBuilder{
		&TrainedModelInferenceClassImportance{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelInferenceClassImportance struct
func (rb *TrainedModelInferenceClassImportanceBuilder) Build() TrainedModelInferenceClassImportance {
	return *rb.v
}

func (rb *TrainedModelInferenceClassImportanceBuilder) ClassName(classname string) *TrainedModelInferenceClassImportanceBuilder {
	rb.v.ClassName = classname
	return rb
}

func (rb *TrainedModelInferenceClassImportanceBuilder) Importance(importance float64) *TrainedModelInferenceClassImportanceBuilder {
	rb.v.Importance = importance
	return rb
}
