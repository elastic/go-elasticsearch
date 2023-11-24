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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package analyze

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package analyze
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/analyze/IndicesAnalyzeRequest.ts#L27-L92
type Request struct {

	// Analyzer The name of the analyzer that should be applied to the provided `text`.
	// This could be a built-in analyzer, or an analyzer that’s been configured in
	// the index.
	Analyzer *string `json:"analyzer,omitempty"`
	// Attributes Array of token attributes used to filter the output of the `explain`
	// parameter.
	Attributes []string `json:"attributes,omitempty"`
	// CharFilter Array of character filters used to preprocess characters before the
	// tokenizer.
	CharFilter []types.CharFilter `json:"char_filter,omitempty"`
	// Explain If `true`, the response includes token attributes and additional details.
	Explain *bool `json:"explain,omitempty"`
	// Field Field used to derive the analyzer.
	// To use this parameter, you must specify an index.
	// If specified, the `analyzer` parameter overrides this value.
	Field *string `json:"field,omitempty"`
	// Filter Array of token filters used to apply after the tokenizer.
	Filter []types.TokenFilter `json:"filter,omitempty"`
	// Normalizer Normalizer to use to convert text into a single token.
	Normalizer *string `json:"normalizer,omitempty"`
	// Text Text to analyze.
	// If an array of strings is provided, it is analyzed as a multi-value field.
	Text []string `json:"text,omitempty"`
	// Tokenizer Tokenizer to use to convert text into tokens.
	Tokenizer types.Tokenizer `json:"tokenizer,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Analyze request: %w", err)
	}

	return &req, nil
}
