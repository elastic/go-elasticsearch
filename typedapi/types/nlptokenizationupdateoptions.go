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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tokenizationtruncate"
)

// NlpTokenizationUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L315-L320
type NlpTokenizationUpdateOptions struct {
	// Span Span options to apply
	Span *int `json:"span,omitempty"`
	// Truncate Truncate options to apply
	Truncate *tokenizationtruncate.TokenizationTruncate `json:"truncate,omitempty"`
}

// NlpTokenizationUpdateOptionsBuilder holds NlpTokenizationUpdateOptions struct and provides a builder API.
type NlpTokenizationUpdateOptionsBuilder struct {
	v *NlpTokenizationUpdateOptions
}

// NewNlpTokenizationUpdateOptions provides a builder for the NlpTokenizationUpdateOptions struct.
func NewNlpTokenizationUpdateOptionsBuilder() *NlpTokenizationUpdateOptionsBuilder {
	r := NlpTokenizationUpdateOptionsBuilder{
		&NlpTokenizationUpdateOptions{},
	}

	return &r
}

// Build finalize the chain and returns the NlpTokenizationUpdateOptions struct
func (rb *NlpTokenizationUpdateOptionsBuilder) Build() NlpTokenizationUpdateOptions {
	return *rb.v
}

// Span Span options to apply

func (rb *NlpTokenizationUpdateOptionsBuilder) Span(span int) *NlpTokenizationUpdateOptionsBuilder {
	rb.v.Span = &span
	return rb
}

// Truncate Truncate options to apply

func (rb *NlpTokenizationUpdateOptionsBuilder) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *NlpTokenizationUpdateOptionsBuilder {
	rb.v.Truncate = &truncate
	return rb
}
