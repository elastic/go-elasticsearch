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

// SimpleAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/analyzers.ts#L83-L86
type SimpleAnalyzer struct {
	Type    string         `json:"type,omitempty"`
	Version *VersionString `json:"version,omitempty"`
}

// SimpleAnalyzerBuilder holds SimpleAnalyzer struct and provides a builder API.
type SimpleAnalyzerBuilder struct {
	v *SimpleAnalyzer
}

// NewSimpleAnalyzer provides a builder for the SimpleAnalyzer struct.
func NewSimpleAnalyzerBuilder() *SimpleAnalyzerBuilder {
	r := SimpleAnalyzerBuilder{
		&SimpleAnalyzer{},
	}

	r.v.Type = "simple"

	return &r
}

// Build finalize the chain and returns the SimpleAnalyzer struct
func (rb *SimpleAnalyzerBuilder) Build() SimpleAnalyzer {
	return *rb.v
}

func (rb *SimpleAnalyzerBuilder) Version(version VersionString) *SimpleAnalyzerBuilder {
	rb.v.Version = &version
	return rb
}
