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

// TotalFeatureImportanceStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L237-L244
type TotalFeatureImportanceStatistics struct {
	// Max The maximum importance value across all the training data for this feature.
	Max int `json:"max"`
	// MeanMagnitude The average magnitude of this feature across all the training data. This
	// value is the average of the absolute values of the importance for this
	// feature.
	MeanMagnitude float64 `json:"mean_magnitude"`
	// Min The minimum importance value across all the training data for this feature.
	Min int `json:"min"`
}

// TotalFeatureImportanceStatisticsBuilder holds TotalFeatureImportanceStatistics struct and provides a builder API.
type TotalFeatureImportanceStatisticsBuilder struct {
	v *TotalFeatureImportanceStatistics
}

// NewTotalFeatureImportanceStatistics provides a builder for the TotalFeatureImportanceStatistics struct.
func NewTotalFeatureImportanceStatisticsBuilder() *TotalFeatureImportanceStatisticsBuilder {
	r := TotalFeatureImportanceStatisticsBuilder{
		&TotalFeatureImportanceStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the TotalFeatureImportanceStatistics struct
func (rb *TotalFeatureImportanceStatisticsBuilder) Build() TotalFeatureImportanceStatistics {
	return *rb.v
}

// Max The maximum importance value across all the training data for this feature.

func (rb *TotalFeatureImportanceStatisticsBuilder) Max(max int) *TotalFeatureImportanceStatisticsBuilder {
	rb.v.Max = max
	return rb
}

// MeanMagnitude The average magnitude of this feature across all the training data. This
// value is the average of the absolute values of the importance for this
// feature.

func (rb *TotalFeatureImportanceStatisticsBuilder) MeanMagnitude(meanmagnitude float64) *TotalFeatureImportanceStatisticsBuilder {
	rb.v.MeanMagnitude = meanmagnitude
	return rb
}

// Min The minimum importance value across all the training data for this feature.

func (rb *TotalFeatureImportanceStatisticsBuilder) Min(min int) *TotalFeatureImportanceStatisticsBuilder {
	rb.v.Min = min
	return rb
}
