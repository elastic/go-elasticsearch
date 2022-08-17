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

// IndexSettingsAnalysis type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L310-L316
type IndexSettingsAnalysis struct {
	Analyzer   map[string]Analyzer    `json:"analyzer,omitempty"`
	CharFilter map[string]CharFilter  `json:"char_filter,omitempty"`
	Filter     map[string]TokenFilter `json:"filter,omitempty"`
	Normalizer map[string]Normalizer  `json:"normalizer,omitempty"`
	Tokenizer  map[string]Tokenizer   `json:"tokenizer,omitempty"`
}

// IndexSettingsAnalysisBuilder holds IndexSettingsAnalysis struct and provides a builder API.
type IndexSettingsAnalysisBuilder struct {
	v *IndexSettingsAnalysis
}

// NewIndexSettingsAnalysis provides a builder for the IndexSettingsAnalysis struct.
func NewIndexSettingsAnalysisBuilder() *IndexSettingsAnalysisBuilder {
	r := IndexSettingsAnalysisBuilder{
		&IndexSettingsAnalysis{
			Analyzer:   make(map[string]Analyzer, 0),
			CharFilter: make(map[string]CharFilter, 0),
			Filter:     make(map[string]TokenFilter, 0),
			Normalizer: make(map[string]Normalizer, 0),
			Tokenizer:  make(map[string]Tokenizer, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexSettingsAnalysis struct
func (rb *IndexSettingsAnalysisBuilder) Build() IndexSettingsAnalysis {
	return *rb.v
}

func (rb *IndexSettingsAnalysisBuilder) Analyzer(values map[string]*AnalyzerBuilder) *IndexSettingsAnalysisBuilder {
	tmp := make(map[string]Analyzer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Analyzer = tmp
	return rb
}

func (rb *IndexSettingsAnalysisBuilder) CharFilter(values map[string]*CharFilterBuilder) *IndexSettingsAnalysisBuilder {
	tmp := make(map[string]CharFilter, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.CharFilter = tmp
	return rb
}

func (rb *IndexSettingsAnalysisBuilder) Filter(values map[string]*TokenFilterBuilder) *IndexSettingsAnalysisBuilder {
	tmp := make(map[string]TokenFilter, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Filter = tmp
	return rb
}

func (rb *IndexSettingsAnalysisBuilder) Normalizer(values map[string]*NormalizerBuilder) *IndexSettingsAnalysisBuilder {
	tmp := make(map[string]Normalizer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Normalizer = tmp
	return rb
}

func (rb *IndexSettingsAnalysisBuilder) Tokenizer(values map[string]*TokenizerBuilder) *IndexSettingsAnalysisBuilder {
	tmp := make(map[string]Tokenizer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Tokenizer = tmp
	return rb
}
