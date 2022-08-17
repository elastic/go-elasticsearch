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

// AnalyzeDetail type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/analyze/types.ts#L24-L30
type AnalyzeDetail struct {
	Analyzer       *AnalyzerDetail    `json:"analyzer,omitempty"`
	Charfilters    []CharFilterDetail `json:"charfilters,omitempty"`
	CustomAnalyzer bool               `json:"custom_analyzer"`
	Tokenfilters   []TokenDetail      `json:"tokenfilters,omitempty"`
	Tokenizer      *TokenDetail       `json:"tokenizer,omitempty"`
}

// AnalyzeDetailBuilder holds AnalyzeDetail struct and provides a builder API.
type AnalyzeDetailBuilder struct {
	v *AnalyzeDetail
}

// NewAnalyzeDetail provides a builder for the AnalyzeDetail struct.
func NewAnalyzeDetailBuilder() *AnalyzeDetailBuilder {
	r := AnalyzeDetailBuilder{
		&AnalyzeDetail{},
	}

	return &r
}

// Build finalize the chain and returns the AnalyzeDetail struct
func (rb *AnalyzeDetailBuilder) Build() AnalyzeDetail {
	return *rb.v
}

func (rb *AnalyzeDetailBuilder) Analyzer(analyzer *AnalyzerDetailBuilder) *AnalyzeDetailBuilder {
	v := analyzer.Build()
	rb.v.Analyzer = &v
	return rb
}

func (rb *AnalyzeDetailBuilder) Charfilters(charfilters []CharFilterDetailBuilder) *AnalyzeDetailBuilder {
	tmp := make([]CharFilterDetail, len(charfilters))
	for _, value := range charfilters {
		tmp = append(tmp, value.Build())
	}
	rb.v.Charfilters = tmp
	return rb
}

func (rb *AnalyzeDetailBuilder) CustomAnalyzer(customanalyzer bool) *AnalyzeDetailBuilder {
	rb.v.CustomAnalyzer = customanalyzer
	return rb
}

func (rb *AnalyzeDetailBuilder) Tokenfilters(tokenfilters []TokenDetailBuilder) *AnalyzeDetailBuilder {
	tmp := make([]TokenDetail, len(tokenfilters))
	for _, value := range tokenfilters {
		tmp = append(tmp, value.Build())
	}
	rb.v.Tokenfilters = tmp
	return rb
}

func (rb *AnalyzeDetailBuilder) Tokenizer(tokenizer *TokenDetailBuilder) *AnalyzeDetailBuilder {
	v := tokenizer.Build()
	rb.v.Tokenizer = &v
	return rb
}
