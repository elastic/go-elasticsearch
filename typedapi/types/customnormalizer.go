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

// CustomNormalizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/normalizers.ts#L30-L34
type CustomNormalizer struct {
	CharFilter []string `json:"char_filter,omitempty"`
	Filter     []string `json:"filter,omitempty"`
	Type       string   `json:"type,omitempty"`
}

// CustomNormalizerBuilder holds CustomNormalizer struct and provides a builder API.
type CustomNormalizerBuilder struct {
	v *CustomNormalizer
}

// NewCustomNormalizer provides a builder for the CustomNormalizer struct.
func NewCustomNormalizerBuilder() *CustomNormalizerBuilder {
	r := CustomNormalizerBuilder{
		&CustomNormalizer{},
	}

	r.v.Type = "custom"

	return &r
}

// Build finalize the chain and returns the CustomNormalizer struct
func (rb *CustomNormalizerBuilder) Build() CustomNormalizer {
	return *rb.v
}

func (rb *CustomNormalizerBuilder) CharFilter(char_filter ...string) *CustomNormalizerBuilder {
	rb.v.CharFilter = char_filter
	return rb
}

func (rb *CustomNormalizerBuilder) Filter(filter ...string) *CustomNormalizerBuilder {
	rb.v.Filter = filter
	return rb
}
