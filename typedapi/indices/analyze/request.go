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


package analyze

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package analyze
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/analyze/IndicesAnalyzeRequest.ts#L27-L47
type Request struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Attributes []string `json:"attributes,omitempty"`

	CharFilter []types.CharFilter `json:"char_filter,omitempty"`

	Explain *bool `json:"explain,omitempty"`

	Field *types.Field `json:"field,omitempty"`

	Filter []types.TokenFilter `json:"filter,omitempty"`

	Normalizer *string `json:"normalizer,omitempty"`

	Text *types.TextToAnalyze `json:"text,omitempty"`

	Tokenizer *types.Tokenizer `json:"tokenizer,omitempty"`
}

// RequestBuilder is the builder API for the analyze.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Analyze request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Analyzer(analyzer string) *RequestBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *RequestBuilder) Attributes(attributes ...string) *RequestBuilder {
	rb.v.Attributes = attributes
	return rb
}

func (rb *RequestBuilder) CharFilter(char_filter ...types.CharFilter) *RequestBuilder {
	rb.v.CharFilter = char_filter
	return rb
}

func (rb *RequestBuilder) Explain(explain bool) *RequestBuilder {
	rb.v.Explain = &explain
	return rb
}

func (rb *RequestBuilder) Field(field types.Field) *RequestBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *RequestBuilder) Filter(filter ...types.TokenFilter) *RequestBuilder {
	rb.v.Filter = filter
	return rb
}

func (rb *RequestBuilder) Normalizer(normalizer string) *RequestBuilder {
	rb.v.Normalizer = &normalizer
	return rb
}

func (rb *RequestBuilder) Text(text *types.TextToAnalyzeBuilder) *RequestBuilder {
	v := text.Build()
	rb.v.Text = &v
	return rb
}

func (rb *RequestBuilder) Tokenizer(tokenizer *types.TokenizerBuilder) *RequestBuilder {
	v := tokenizer.Build()
	rb.v.Tokenizer = &v
	return rb
}
