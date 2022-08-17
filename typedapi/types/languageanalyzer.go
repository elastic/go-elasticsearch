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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/language"
)

// LanguageAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/analyzers.ts#L52-L59
type LanguageAnalyzer struct {
	Language      language.Language `json:"language"`
	StemExclusion []string          `json:"stem_exclusion"`
	Stopwords     *StopWords        `json:"stopwords,omitempty"`
	StopwordsPath *string           `json:"stopwords_path,omitempty"`
	Type          string            `json:"type,omitempty"`
	Version       *VersionString    `json:"version,omitempty"`
}

// LanguageAnalyzerBuilder holds LanguageAnalyzer struct and provides a builder API.
type LanguageAnalyzerBuilder struct {
	v *LanguageAnalyzer
}

// NewLanguageAnalyzer provides a builder for the LanguageAnalyzer struct.
func NewLanguageAnalyzerBuilder() *LanguageAnalyzerBuilder {
	r := LanguageAnalyzerBuilder{
		&LanguageAnalyzer{},
	}

	r.v.Type = "language"

	return &r
}

// Build finalize the chain and returns the LanguageAnalyzer struct
func (rb *LanguageAnalyzerBuilder) Build() LanguageAnalyzer {
	return *rb.v
}

func (rb *LanguageAnalyzerBuilder) Language(language language.Language) *LanguageAnalyzerBuilder {
	rb.v.Language = language
	return rb
}

func (rb *LanguageAnalyzerBuilder) StemExclusion(stem_exclusion ...string) *LanguageAnalyzerBuilder {
	rb.v.StemExclusion = stem_exclusion
	return rb
}

func (rb *LanguageAnalyzerBuilder) Stopwords(stopwords *StopWordsBuilder) *LanguageAnalyzerBuilder {
	v := stopwords.Build()
	rb.v.Stopwords = &v
	return rb
}

func (rb *LanguageAnalyzerBuilder) StopwordsPath(stopwordspath string) *LanguageAnalyzerBuilder {
	rb.v.StopwordsPath = &stopwordspath
	return rb
}

func (rb *LanguageAnalyzerBuilder) Version(version VersionString) *LanguageAnalyzerBuilder {
	rb.v.Version = &version
	return rb
}
