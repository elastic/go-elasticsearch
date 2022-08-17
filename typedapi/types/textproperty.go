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

// TextProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L236-L252
type TextProperty struct {
	Analyzer                 *string                            `json:"analyzer,omitempty"`
	Boost                    *float64                           `json:"boost,omitempty"`
	CopyTo                   *Fields                            `json:"copy_to,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping     `json:"dynamic,omitempty"`
	EagerGlobalOrdinals      *bool                              `json:"eager_global_ordinals,omitempty"`
	Fielddata                *bool                              `json:"fielddata,omitempty"`
	FielddataFrequencyFilter *FielddataFrequencyFilter          `json:"fielddata_frequency_filter,omitempty"`
	Fields                   map[PropertyName]Property          `json:"fields,omitempty"`
	IgnoreAbove              *int                               `json:"ignore_above,omitempty"`
	Index                    *bool                              `json:"index,omitempty"`
	IndexOptions             *indexoptions.IndexOptions         `json:"index_options,omitempty"`
	IndexPhrases             *bool                              `json:"index_phrases,omitempty"`
	IndexPrefixes            *TextIndexPrefixes                 `json:"index_prefixes,omitempty"`
	LocalMetadata            *Metadata                          `json:"local_metadata,omitempty"`
	Meta                     map[string]string                  `json:"meta,omitempty"`
	Norms                    *bool                              `json:"norms,omitempty"`
	PositionIncrementGap     *int                               `json:"position_increment_gap,omitempty"`
	Properties               map[PropertyName]Property          `json:"properties,omitempty"`
	SearchAnalyzer           *string                            `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer      *string                            `json:"search_quote_analyzer,omitempty"`
	Similarity               *string                            `json:"similarity,omitempty"`
	Store                    *bool                              `json:"store,omitempty"`
	TermVector               *termvectoroption.TermVectorOption `json:"term_vector,omitempty"`
	Type                     string                             `json:"type,omitempty"`
}

// TextPropertyBuilder holds TextProperty struct and provides a builder API.
type TextPropertyBuilder struct {
	v *TextProperty
}

// NewTextProperty provides a builder for the TextProperty struct.
func NewTextPropertyBuilder() *TextPropertyBuilder {
	r := TextPropertyBuilder{
		&TextProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "text"

	return &r
}

// Build finalize the chain and returns the TextProperty struct
func (rb *TextPropertyBuilder) Build() TextProperty {
	return *rb.v
}

func (rb *TextPropertyBuilder) Analyzer(analyzer string) *TextPropertyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *TextPropertyBuilder) Boost(boost float64) *TextPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *TextPropertyBuilder) CopyTo(copyto *FieldsBuilder) *TextPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *TextPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *TextPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *TextPropertyBuilder) EagerGlobalOrdinals(eagerglobalordinals bool) *TextPropertyBuilder {
	rb.v.EagerGlobalOrdinals = &eagerglobalordinals
	return rb
}

func (rb *TextPropertyBuilder) Fielddata(fielddata bool) *TextPropertyBuilder {
	rb.v.Fielddata = &fielddata
	return rb
}

func (rb *TextPropertyBuilder) FielddataFrequencyFilter(fielddatafrequencyfilter *FielddataFrequencyFilterBuilder) *TextPropertyBuilder {
	v := fielddatafrequencyfilter.Build()
	rb.v.FielddataFrequencyFilter = &v
	return rb
}

func (rb *TextPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *TextPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *TextPropertyBuilder) IgnoreAbove(ignoreabove int) *TextPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *TextPropertyBuilder) Index(index bool) *TextPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *TextPropertyBuilder) IndexOptions(indexoptions indexoptions.IndexOptions) *TextPropertyBuilder {
	rb.v.IndexOptions = &indexoptions
	return rb
}

func (rb *TextPropertyBuilder) IndexPhrases(indexphrases bool) *TextPropertyBuilder {
	rb.v.IndexPhrases = &indexphrases
	return rb
}

func (rb *TextPropertyBuilder) IndexPrefixes(indexprefixes *TextIndexPrefixesBuilder) *TextPropertyBuilder {
	v := indexprefixes.Build()
	rb.v.IndexPrefixes = &v
	return rb
}

func (rb *TextPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *TextPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *TextPropertyBuilder) Meta(value map[string]string) *TextPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *TextPropertyBuilder) Norms(norms bool) *TextPropertyBuilder {
	rb.v.Norms = &norms
	return rb
}

func (rb *TextPropertyBuilder) PositionIncrementGap(positionincrementgap int) *TextPropertyBuilder {
	rb.v.PositionIncrementGap = &positionincrementgap
	return rb
}

func (rb *TextPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *TextPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *TextPropertyBuilder) SearchAnalyzer(searchanalyzer string) *TextPropertyBuilder {
	rb.v.SearchAnalyzer = &searchanalyzer
	return rb
}

func (rb *TextPropertyBuilder) SearchQuoteAnalyzer(searchquoteanalyzer string) *TextPropertyBuilder {
	rb.v.SearchQuoteAnalyzer = &searchquoteanalyzer
	return rb
}

func (rb *TextPropertyBuilder) Similarity(similarity string) *TextPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *TextPropertyBuilder) Store(store bool) *TextPropertyBuilder {
	rb.v.Store = &store
	return rb
}

func (rb *TextPropertyBuilder) TermVector(termvector termvectoroption.TermVectorOption) *TextPropertyBuilder {
	rb.v.TermVector = &termvector
	return rb
}
