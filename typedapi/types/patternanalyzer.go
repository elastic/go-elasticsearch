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

// PatternAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/analyzers.ts#L74-L81
type PatternAnalyzer struct {
	Flags     *string        `json:"flags,omitempty"`
	Lowercase *bool          `json:"lowercase,omitempty"`
	Pattern   string         `json:"pattern"`
	Stopwords *StopWords     `json:"stopwords,omitempty"`
	Type      string         `json:"type,omitempty"`
	Version   *VersionString `json:"version,omitempty"`
}

// PatternAnalyzerBuilder holds PatternAnalyzer struct and provides a builder API.
type PatternAnalyzerBuilder struct {
	v *PatternAnalyzer
}

// NewPatternAnalyzer provides a builder for the PatternAnalyzer struct.
func NewPatternAnalyzerBuilder() *PatternAnalyzerBuilder {
	r := PatternAnalyzerBuilder{
		&PatternAnalyzer{},
	}

	r.v.Type = "pattern"

	return &r
}

// Build finalize the chain and returns the PatternAnalyzer struct
func (rb *PatternAnalyzerBuilder) Build() PatternAnalyzer {
	return *rb.v
}

func (rb *PatternAnalyzerBuilder) Flags(flags string) *PatternAnalyzerBuilder {
	rb.v.Flags = &flags
	return rb
}

func (rb *PatternAnalyzerBuilder) Lowercase(lowercase bool) *PatternAnalyzerBuilder {
	rb.v.Lowercase = &lowercase
	return rb
}

func (rb *PatternAnalyzerBuilder) Pattern(pattern string) *PatternAnalyzerBuilder {
	rb.v.Pattern = pattern
	return rb
}

func (rb *PatternAnalyzerBuilder) Stopwords(stopwords *StopWordsBuilder) *PatternAnalyzerBuilder {
	v := stopwords.Build()
	rb.v.Stopwords = &v
	return rb
}

func (rb *PatternAnalyzerBuilder) Version(version VersionString) *PatternAnalyzerBuilder {
	rb.v.Version = &version
	return rb
}
