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

// ElisionTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L186-L191
type ElisionTokenFilter struct {
	Articles     []string       `json:"articles,omitempty"`
	ArticlesCase *bool          `json:"articles_case,omitempty"`
	ArticlesPath *string        `json:"articles_path,omitempty"`
	Type         string         `json:"type,omitempty"`
	Version      *VersionString `json:"version,omitempty"`
}

// ElisionTokenFilterBuilder holds ElisionTokenFilter struct and provides a builder API.
type ElisionTokenFilterBuilder struct {
	v *ElisionTokenFilter
}

// NewElisionTokenFilter provides a builder for the ElisionTokenFilter struct.
func NewElisionTokenFilterBuilder() *ElisionTokenFilterBuilder {
	r := ElisionTokenFilterBuilder{
		&ElisionTokenFilter{},
	}

	r.v.Type = "elision"

	return &r
}

// Build finalize the chain and returns the ElisionTokenFilter struct
func (rb *ElisionTokenFilterBuilder) Build() ElisionTokenFilter {
	return *rb.v
}

func (rb *ElisionTokenFilterBuilder) Articles(articles ...string) *ElisionTokenFilterBuilder {
	rb.v.Articles = articles
	return rb
}

func (rb *ElisionTokenFilterBuilder) ArticlesCase(articlescase bool) *ElisionTokenFilterBuilder {
	rb.v.ArticlesCase = &articlescase
	return rb
}

func (rb *ElisionTokenFilterBuilder) ArticlesPath(articlespath string) *ElisionTokenFilterBuilder {
	rb.v.ArticlesPath = &articlespath
	return rb
}

func (rb *ElisionTokenFilterBuilder) Version(version VersionString) *ElisionTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
