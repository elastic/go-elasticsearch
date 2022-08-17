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

// TotalFeatureImportanceClass type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L230-L235
type TotalFeatureImportanceClass struct {
	// ClassName The target class value. Could be a string, boolean, or number.
	ClassName Name `json:"class_name"`
	// Importance A collection of feature importance statistics related to the training data
	// set for this particular feature.
	Importance []TotalFeatureImportanceStatistics `json:"importance"`
}

// TotalFeatureImportanceClassBuilder holds TotalFeatureImportanceClass struct and provides a builder API.
type TotalFeatureImportanceClassBuilder struct {
	v *TotalFeatureImportanceClass
}

// NewTotalFeatureImportanceClass provides a builder for the TotalFeatureImportanceClass struct.
func NewTotalFeatureImportanceClassBuilder() *TotalFeatureImportanceClassBuilder {
	r := TotalFeatureImportanceClassBuilder{
		&TotalFeatureImportanceClass{},
	}

	return &r
}

// Build finalize the chain and returns the TotalFeatureImportanceClass struct
func (rb *TotalFeatureImportanceClassBuilder) Build() TotalFeatureImportanceClass {
	return *rb.v
}

// ClassName The target class value. Could be a string, boolean, or number.

func (rb *TotalFeatureImportanceClassBuilder) ClassName(classname Name) *TotalFeatureImportanceClassBuilder {
	rb.v.ClassName = classname
	return rb
}

// Importance A collection of feature importance statistics related to the training data
// set for this particular feature.

func (rb *TotalFeatureImportanceClassBuilder) Importance(importance []TotalFeatureImportanceStatisticsBuilder) *TotalFeatureImportanceClassBuilder {
	tmp := make([]TotalFeatureImportanceStatistics, len(importance))
	for _, value := range importance {
		tmp = append(tmp, value.Build())
	}
	rb.v.Importance = tmp
	return rb
}
