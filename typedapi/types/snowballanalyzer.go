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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/snowballlanguage"
)

// SnowballAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/analyzers.ts#L88-L93
type SnowballAnalyzer struct {
	Language  snowballlanguage.SnowballLanguage `json:"language"`
	Stopwords *StopWords                        `json:"stopwords,omitempty"`
	Type      string                            `json:"type,omitempty"`
	Version   *VersionString                    `json:"version,omitempty"`
}

// SnowballAnalyzerBuilder holds SnowballAnalyzer struct and provides a builder API.
type SnowballAnalyzerBuilder struct {
	v *SnowballAnalyzer
}

// NewSnowballAnalyzer provides a builder for the SnowballAnalyzer struct.
func NewSnowballAnalyzerBuilder() *SnowballAnalyzerBuilder {
	r := SnowballAnalyzerBuilder{
		&SnowballAnalyzer{},
	}

	r.v.Type = "snowball"

	return &r
}

// Build finalize the chain and returns the SnowballAnalyzer struct
func (rb *SnowballAnalyzerBuilder) Build() SnowballAnalyzer {
	return *rb.v
}

func (rb *SnowballAnalyzerBuilder) Language(language snowballlanguage.SnowballLanguage) *SnowballAnalyzerBuilder {
	rb.v.Language = language
	return rb
}

func (rb *SnowballAnalyzerBuilder) Stopwords(stopwords *StopWordsBuilder) *SnowballAnalyzerBuilder {
	v := stopwords.Build()
	rb.v.Stopwords = &v
	return rb
}

func (rb *SnowballAnalyzerBuilder) Version(version VersionString) *SnowballAnalyzerBuilder {
	rb.v.Version = &version
	return rb
}
