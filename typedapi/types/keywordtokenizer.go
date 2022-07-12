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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// KeywordTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/analysis/tokenizers.ts#L61-L64
type KeywordTokenizer struct {
	BufferSize int            `json:"buffer_size"`
	Type       string         `json:"type,omitempty"`
	Version    *VersionString `json:"version,omitempty"`
}

// KeywordTokenizerBuilder holds KeywordTokenizer struct and provides a builder API.
type KeywordTokenizerBuilder struct {
	v *KeywordTokenizer
}

// NewKeywordTokenizer provides a builder for the KeywordTokenizer struct.
func NewKeywordTokenizerBuilder() *KeywordTokenizerBuilder {
	r := KeywordTokenizerBuilder{
		&KeywordTokenizer{},
	}

	r.v.Type = "keyword"

	return &r
}

// Build finalize the chain and returns the KeywordTokenizer struct
func (rb *KeywordTokenizerBuilder) Build() KeywordTokenizer {
	return *rb.v
}

func (rb *KeywordTokenizerBuilder) BufferSize(buffersize int) *KeywordTokenizerBuilder {
	rb.v.BufferSize = buffersize
	return rb
}

func (rb *KeywordTokenizerBuilder) Version(version VersionString) *KeywordTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
