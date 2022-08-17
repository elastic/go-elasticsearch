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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/synonymformat"
)

// SynonymTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L120-L129
type SynonymTokenFilter struct {
	Expand       *bool                        `json:"expand,omitempty"`
	Format       *synonymformat.SynonymFormat `json:"format,omitempty"`
	Lenient      *bool                        `json:"lenient,omitempty"`
	Synonyms     []string                     `json:"synonyms,omitempty"`
	SynonymsPath *string                      `json:"synonyms_path,omitempty"`
	Tokenizer    *string                      `json:"tokenizer,omitempty"`
	Type         string                       `json:"type,omitempty"`
	Updateable   *bool                        `json:"updateable,omitempty"`
	Version      *VersionString               `json:"version,omitempty"`
}

// SynonymTokenFilterBuilder holds SynonymTokenFilter struct and provides a builder API.
type SynonymTokenFilterBuilder struct {
	v *SynonymTokenFilter
}

// NewSynonymTokenFilter provides a builder for the SynonymTokenFilter struct.
func NewSynonymTokenFilterBuilder() *SynonymTokenFilterBuilder {
	r := SynonymTokenFilterBuilder{
		&SynonymTokenFilter{},
	}

	r.v.Type = "synonym"

	return &r
}

// Build finalize the chain and returns the SynonymTokenFilter struct
func (rb *SynonymTokenFilterBuilder) Build() SynonymTokenFilter {
	return *rb.v
}

func (rb *SynonymTokenFilterBuilder) Expand(expand bool) *SynonymTokenFilterBuilder {
	rb.v.Expand = &expand
	return rb
}

func (rb *SynonymTokenFilterBuilder) Format(format synonymformat.SynonymFormat) *SynonymTokenFilterBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *SynonymTokenFilterBuilder) Lenient(lenient bool) *SynonymTokenFilterBuilder {
	rb.v.Lenient = &lenient
	return rb
}

func (rb *SynonymTokenFilterBuilder) Synonyms(synonyms ...string) *SynonymTokenFilterBuilder {
	rb.v.Synonyms = synonyms
	return rb
}

func (rb *SynonymTokenFilterBuilder) SynonymsPath(synonymspath string) *SynonymTokenFilterBuilder {
	rb.v.SynonymsPath = &synonymspath
	return rb
}

func (rb *SynonymTokenFilterBuilder) Tokenizer(tokenizer string) *SynonymTokenFilterBuilder {
	rb.v.Tokenizer = &tokenizer
	return rb
}

func (rb *SynonymTokenFilterBuilder) Updateable(updateable bool) *SynonymTokenFilterBuilder {
	rb.v.Updateable = &updateable
	return rb
}

func (rb *SynonymTokenFilterBuilder) Version(version VersionString) *SynonymTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
