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

// StopAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/analyzers.ts#L101-L106
type StopAnalyzer struct {
	Stopwords     *StopWords     `json:"stopwords,omitempty"`
	StopwordsPath *string        `json:"stopwords_path,omitempty"`
	Type          string         `json:"type,omitempty"`
	Version       *VersionString `json:"version,omitempty"`
}

// StopAnalyzerBuilder holds StopAnalyzer struct and provides a builder API.
type StopAnalyzerBuilder struct {
	v *StopAnalyzer
}

// NewStopAnalyzer provides a builder for the StopAnalyzer struct.
func NewStopAnalyzerBuilder() *StopAnalyzerBuilder {
	r := StopAnalyzerBuilder{
		&StopAnalyzer{},
	}

	r.v.Type = "stop"

	return &r
}

// Build finalize the chain and returns the StopAnalyzer struct
func (rb *StopAnalyzerBuilder) Build() StopAnalyzer {
	return *rb.v
}

func (rb *StopAnalyzerBuilder) Stopwords(stopwords *StopWordsBuilder) *StopAnalyzerBuilder {
	v := stopwords.Build()
	rb.v.Stopwords = &v
	return rb
}

func (rb *StopAnalyzerBuilder) StopwordsPath(stopwordspath string) *StopAnalyzerBuilder {
	rb.v.StopwordsPath = &stopwordspath
	return rb
}

func (rb *StopAnalyzerBuilder) Version(version VersionString) *StopAnalyzerBuilder {
	rb.v.Version = &version
	return rb
}
