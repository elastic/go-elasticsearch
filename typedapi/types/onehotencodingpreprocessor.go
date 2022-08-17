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

// OneHotEncodingPreprocessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/put_trained_model/types.ts#L44-L47
type OneHotEncodingPreprocessor struct {
	Field  string            `json:"field"`
	HotMap map[string]string `json:"hot_map"`
}

// OneHotEncodingPreprocessorBuilder holds OneHotEncodingPreprocessor struct and provides a builder API.
type OneHotEncodingPreprocessorBuilder struct {
	v *OneHotEncodingPreprocessor
}

// NewOneHotEncodingPreprocessor provides a builder for the OneHotEncodingPreprocessor struct.
func NewOneHotEncodingPreprocessorBuilder() *OneHotEncodingPreprocessorBuilder {
	r := OneHotEncodingPreprocessorBuilder{
		&OneHotEncodingPreprocessor{
			HotMap: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the OneHotEncodingPreprocessor struct
func (rb *OneHotEncodingPreprocessorBuilder) Build() OneHotEncodingPreprocessor {
	return *rb.v
}

func (rb *OneHotEncodingPreprocessorBuilder) Field(field string) *OneHotEncodingPreprocessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *OneHotEncodingPreprocessorBuilder) HotMap(value map[string]string) *OneHotEncodingPreprocessorBuilder {
	rb.v.HotMap = value
	return rb
}
