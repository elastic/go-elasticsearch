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

// TotalFeatureImportance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L221-L228
type TotalFeatureImportance struct {
	// Classes If the trained model is a classification model, feature importance statistics
	// are gathered per target class value.
	Classes []TotalFeatureImportanceClass `json:"classes"`
	// FeatureName The feature for which this importance was calculated.
	FeatureName Name `json:"feature_name"`
	// Importance A collection of feature importance statistics related to the training data
	// set for this particular feature.
	Importance []TotalFeatureImportanceStatistics `json:"importance"`
}

// TotalFeatureImportanceBuilder holds TotalFeatureImportance struct and provides a builder API.
type TotalFeatureImportanceBuilder struct {
	v *TotalFeatureImportance
}

// NewTotalFeatureImportance provides a builder for the TotalFeatureImportance struct.
func NewTotalFeatureImportanceBuilder() *TotalFeatureImportanceBuilder {
	r := TotalFeatureImportanceBuilder{
		&TotalFeatureImportance{},
	}

	return &r
}

// Build finalize the chain and returns the TotalFeatureImportance struct
func (rb *TotalFeatureImportanceBuilder) Build() TotalFeatureImportance {
	return *rb.v
}

// Classes If the trained model is a classification model, feature importance statistics
// are gathered per target class value.

func (rb *TotalFeatureImportanceBuilder) Classes(classes []TotalFeatureImportanceClassBuilder) *TotalFeatureImportanceBuilder {
	tmp := make([]TotalFeatureImportanceClass, len(classes))
	for _, value := range classes {
		tmp = append(tmp, value.Build())
	}
	rb.v.Classes = tmp
	return rb
}

// FeatureName The feature for which this importance was calculated.

func (rb *TotalFeatureImportanceBuilder) FeatureName(featurename Name) *TotalFeatureImportanceBuilder {
	rb.v.FeatureName = featurename
	return rb
}

// Importance A collection of feature importance statistics related to the training data
// set for this particular feature.

func (rb *TotalFeatureImportanceBuilder) Importance(importance []TotalFeatureImportanceStatisticsBuilder) *TotalFeatureImportanceBuilder {
	tmp := make([]TotalFeatureImportanceStatistics, len(importance))
	for _, value := range importance {
		tmp = append(tmp, value.Build())
	}
	rb.v.Importance = tmp
	return rb
}
