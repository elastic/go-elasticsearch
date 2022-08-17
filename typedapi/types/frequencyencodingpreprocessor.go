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

// FrequencyEncodingPreprocessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/put_trained_model/types.ts#L38-L42
type FrequencyEncodingPreprocessor struct {
	FeatureName  string             `json:"feature_name"`
	Field        string             `json:"field"`
	FrequencyMap map[string]float64 `json:"frequency_map"`
}

// FrequencyEncodingPreprocessorBuilder holds FrequencyEncodingPreprocessor struct and provides a builder API.
type FrequencyEncodingPreprocessorBuilder struct {
	v *FrequencyEncodingPreprocessor
}

// NewFrequencyEncodingPreprocessor provides a builder for the FrequencyEncodingPreprocessor struct.
func NewFrequencyEncodingPreprocessorBuilder() *FrequencyEncodingPreprocessorBuilder {
	r := FrequencyEncodingPreprocessorBuilder{
		&FrequencyEncodingPreprocessor{
			FrequencyMap: make(map[string]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the FrequencyEncodingPreprocessor struct
func (rb *FrequencyEncodingPreprocessorBuilder) Build() FrequencyEncodingPreprocessor {
	return *rb.v
}

func (rb *FrequencyEncodingPreprocessorBuilder) FeatureName(featurename string) *FrequencyEncodingPreprocessorBuilder {
	rb.v.FeatureName = featurename
	return rb
}

func (rb *FrequencyEncodingPreprocessorBuilder) Field(field string) *FrequencyEncodingPreprocessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *FrequencyEncodingPreprocessorBuilder) FrequencyMap(value map[string]float64) *FrequencyEncodingPreprocessorBuilder {
	rb.v.FrequencyMap = value
	return rb
}
