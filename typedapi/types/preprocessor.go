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

// Preprocessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/put_trained_model/types.ts#L31-L36
type Preprocessor struct {
	FrequencyEncoding  *FrequencyEncodingPreprocessor  `json:"frequency_encoding,omitempty"`
	OneHotEncoding     *OneHotEncodingPreprocessor     `json:"one_hot_encoding,omitempty"`
	TargetMeanEncoding *TargetMeanEncodingPreprocessor `json:"target_mean_encoding,omitempty"`
}

// PreprocessorBuilder holds Preprocessor struct and provides a builder API.
type PreprocessorBuilder struct {
	v *Preprocessor
}

// NewPreprocessor provides a builder for the Preprocessor struct.
func NewPreprocessorBuilder() *PreprocessorBuilder {
	r := PreprocessorBuilder{
		&Preprocessor{},
	}

	return &r
}

// Build finalize the chain and returns the Preprocessor struct
func (rb *PreprocessorBuilder) Build() Preprocessor {
	return *rb.v
}

func (rb *PreprocessorBuilder) FrequencyEncoding(frequencyencoding *FrequencyEncodingPreprocessorBuilder) *PreprocessorBuilder {
	v := frequencyencoding.Build()
	rb.v.FrequencyEncoding = &v
	return rb
}

func (rb *PreprocessorBuilder) OneHotEncoding(onehotencoding *OneHotEncodingPreprocessorBuilder) *PreprocessorBuilder {
	v := onehotencoding.Build()
	rb.v.OneHotEncoding = &v
	return rb
}

func (rb *PreprocessorBuilder) TargetMeanEncoding(targetmeanencoding *TargetMeanEncodingPreprocessorBuilder) *PreprocessorBuilder {
	v := targetmeanencoding.Build()
	rb.v.TargetMeanEncoding = &v
	return rb
}
