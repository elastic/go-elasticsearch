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

// KeywordMarkerTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L230-L236
type KeywordMarkerTokenFilter struct {
	IgnoreCase      *bool          `json:"ignore_case,omitempty"`
	Keywords        []string       `json:"keywords,omitempty"`
	KeywordsPath    *string        `json:"keywords_path,omitempty"`
	KeywordsPattern *string        `json:"keywords_pattern,omitempty"`
	Type            string         `json:"type,omitempty"`
	Version         *VersionString `json:"version,omitempty"`
}

// KeywordMarkerTokenFilterBuilder holds KeywordMarkerTokenFilter struct and provides a builder API.
type KeywordMarkerTokenFilterBuilder struct {
	v *KeywordMarkerTokenFilter
}

// NewKeywordMarkerTokenFilter provides a builder for the KeywordMarkerTokenFilter struct.
func NewKeywordMarkerTokenFilterBuilder() *KeywordMarkerTokenFilterBuilder {
	r := KeywordMarkerTokenFilterBuilder{
		&KeywordMarkerTokenFilter{},
	}

	r.v.Type = "keyword_marker"

	return &r
}

// Build finalize the chain and returns the KeywordMarkerTokenFilter struct
func (rb *KeywordMarkerTokenFilterBuilder) Build() KeywordMarkerTokenFilter {
	return *rb.v
}

func (rb *KeywordMarkerTokenFilterBuilder) IgnoreCase(ignorecase bool) *KeywordMarkerTokenFilterBuilder {
	rb.v.IgnoreCase = &ignorecase
	return rb
}

func (rb *KeywordMarkerTokenFilterBuilder) Keywords(keywords ...string) *KeywordMarkerTokenFilterBuilder {
	rb.v.Keywords = keywords
	return rb
}

func (rb *KeywordMarkerTokenFilterBuilder) KeywordsPath(keywordspath string) *KeywordMarkerTokenFilterBuilder {
	rb.v.KeywordsPath = &keywordspath
	return rb
}

func (rb *KeywordMarkerTokenFilterBuilder) KeywordsPattern(keywordspattern string) *KeywordMarkerTokenFilterBuilder {
	rb.v.KeywordsPattern = &keywordspattern
	return rb
}

func (rb *KeywordMarkerTokenFilterBuilder) Version(version VersionString) *KeywordMarkerTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
