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

// LimitTokenCountTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L248-L252
type LimitTokenCountTokenFilter struct {
	ConsumeAllTokens *bool          `json:"consume_all_tokens,omitempty"`
	MaxTokenCount    *int           `json:"max_token_count,omitempty"`
	Type             string         `json:"type,omitempty"`
	Version          *VersionString `json:"version,omitempty"`
}

// LimitTokenCountTokenFilterBuilder holds LimitTokenCountTokenFilter struct and provides a builder API.
type LimitTokenCountTokenFilterBuilder struct {
	v *LimitTokenCountTokenFilter
}

// NewLimitTokenCountTokenFilter provides a builder for the LimitTokenCountTokenFilter struct.
func NewLimitTokenCountTokenFilterBuilder() *LimitTokenCountTokenFilterBuilder {
	r := LimitTokenCountTokenFilterBuilder{
		&LimitTokenCountTokenFilter{},
	}

	r.v.Type = "limit"

	return &r
}

// Build finalize the chain and returns the LimitTokenCountTokenFilter struct
func (rb *LimitTokenCountTokenFilterBuilder) Build() LimitTokenCountTokenFilter {
	return *rb.v
}

func (rb *LimitTokenCountTokenFilterBuilder) ConsumeAllTokens(consumealltokens bool) *LimitTokenCountTokenFilterBuilder {
	rb.v.ConsumeAllTokens = &consumealltokens
	return rb
}

func (rb *LimitTokenCountTokenFilterBuilder) MaxTokenCount(maxtokencount int) *LimitTokenCountTokenFilterBuilder {
	rb.v.MaxTokenCount = &maxtokencount
	return rb
}

func (rb *LimitTokenCountTokenFilterBuilder) Version(version VersionString) *LimitTokenCountTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
