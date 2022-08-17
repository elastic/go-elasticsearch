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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termvectoroption"
)

// SearchAsYouTypeProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L179-L189
type SearchAsYouTypeProperty struct {
	Analyzer            *string                            `json:"analyzer,omitempty"`
	CopyTo              *Fields                            `json:"copy_to,omitempty"`
	Dynamic             *dynamicmapping.DynamicMapping     `json:"dynamic,omitempty"`
	Fields              map[PropertyName]Property          `json:"fields,omitempty"`
	IgnoreAbove         *int                               `json:"ignore_above,omitempty"`
	Index               *bool                              `json:"index,omitempty"`
	IndexOptions        *indexoptions.IndexOptions         `json:"index_options,omitempty"`
	LocalMetadata       *Metadata                          `json:"local_metadata,omitempty"`
	MaxShingleSize      *int                               `json:"max_shingle_size,omitempty"`
	Meta                map[string]string                  `json:"meta,omitempty"`
	Norms               *bool                              `json:"norms,omitempty"`
	Properties          map[PropertyName]Property          `json:"properties,omitempty"`
	SearchAnalyzer      *string                            `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer *string                            `json:"search_quote_analyzer,omitempty"`
	Similarity          *string                            `json:"similarity,omitempty"`
	Store               *bool                              `json:"store,omitempty"`
	TermVector          *termvectoroption.TermVectorOption `json:"term_vector,omitempty"`
	Type                string                             `json:"type,omitempty"`
}

// SearchAsYouTypePropertyBuilder holds SearchAsYouTypeProperty struct and provides a builder API.
type SearchAsYouTypePropertyBuilder struct {
	v *SearchAsYouTypeProperty
}

// NewSearchAsYouTypeProperty provides a builder for the SearchAsYouTypeProperty struct.
func NewSearchAsYouTypePropertyBuilder() *SearchAsYouTypePropertyBuilder {
	r := SearchAsYouTypePropertyBuilder{
		&SearchAsYouTypeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "search_as_you_type"

	return &r
}

// Build finalize the chain and returns the SearchAsYouTypeProperty struct
func (rb *SearchAsYouTypePropertyBuilder) Build() SearchAsYouTypeProperty {
	return *rb.v
}

func (rb *SearchAsYouTypePropertyBuilder) Analyzer(analyzer string) *SearchAsYouTypePropertyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) CopyTo(copyto *FieldsBuilder) *SearchAsYouTypePropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *SearchAsYouTypePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *SearchAsYouTypePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) IgnoreAbove(ignoreabove int) *SearchAsYouTypePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) Index(index bool) *SearchAsYouTypePropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) IndexOptions(indexoptions indexoptions.IndexOptions) *SearchAsYouTypePropertyBuilder {
	rb.v.IndexOptions = &indexoptions
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *SearchAsYouTypePropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) MaxShingleSize(maxshinglesize int) *SearchAsYouTypePropertyBuilder {
	rb.v.MaxShingleSize = &maxshinglesize
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) Meta(value map[string]string) *SearchAsYouTypePropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) Norms(norms bool) *SearchAsYouTypePropertyBuilder {
	rb.v.Norms = &norms
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *SearchAsYouTypePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) SearchAnalyzer(searchanalyzer string) *SearchAsYouTypePropertyBuilder {
	rb.v.SearchAnalyzer = &searchanalyzer
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) SearchQuoteAnalyzer(searchquoteanalyzer string) *SearchAsYouTypePropertyBuilder {
	rb.v.SearchQuoteAnalyzer = &searchquoteanalyzer
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) Similarity(similarity string) *SearchAsYouTypePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) Store(store bool) *SearchAsYouTypePropertyBuilder {
	rb.v.Store = &store
	return rb
}

func (rb *SearchAsYouTypePropertyBuilder) TermVector(termvector termvectoroption.TermVectorOption) *SearchAsYouTypePropertyBuilder {
	rb.v.TermVector = &termvector
	return rb
}
