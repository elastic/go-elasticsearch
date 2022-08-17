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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterencoder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertype"
)

// Highlight type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/highlighting.ts#L56-L59
type Highlight struct {
	BoundaryChars         *string                                      `json:"boundary_chars,omitempty"`
	BoundaryMaxScan       *int                                         `json:"boundary_max_scan,omitempty"`
	BoundaryScanner       *boundaryscanner.BoundaryScanner             `json:"boundary_scanner,omitempty"`
	BoundaryScannerLocale *string                                      `json:"boundary_scanner_locale,omitempty"`
	Encoder               *highlighterencoder.HighlighterEncoder       `json:"encoder,omitempty"`
	Fields                map[Field]HighlightField                     `json:"fields"`
	ForceSource           *bool                                        `json:"force_source,omitempty"`
	FragmentSize          *int                                         `json:"fragment_size,omitempty"`
	Fragmenter            *highlighterfragmenter.HighlighterFragmenter `json:"fragmenter,omitempty"`
	HighlightFilter       *bool                                        `json:"highlight_filter,omitempty"`
	HighlightQuery        *QueryContainer                              `json:"highlight_query,omitempty"`
	MaxAnalyzedOffset     *int                                         `json:"max_analyzed_offset,omitempty"`
	MaxFragmentLength     *int                                         `json:"max_fragment_length,omitempty"`
	NoMatchSize           *int                                         `json:"no_match_size,omitempty"`
	NumberOfFragments     *int                                         `json:"number_of_fragments,omitempty"`
	Options               map[string]interface{}                       `json:"options,omitempty"`
	Order                 *highlighterorder.HighlighterOrder           `json:"order,omitempty"`
	PhraseLimit           *int                                         `json:"phrase_limit,omitempty"`
	PostTags              []string                                     `json:"post_tags,omitempty"`
	PreTags               []string                                     `json:"pre_tags,omitempty"`
	RequireFieldMatch     *bool                                        `json:"require_field_match,omitempty"`
	TagsSchema            *highlightertagsschema.HighlighterTagsSchema `json:"tags_schema,omitempty"`
	Type                  *highlightertype.HighlighterType             `json:"type,omitempty"`
}

// HighlightBuilder holds Highlight struct and provides a builder API.
type HighlightBuilder struct {
	v *Highlight
}

// NewHighlight provides a builder for the Highlight struct.
func NewHighlightBuilder() *HighlightBuilder {
	r := HighlightBuilder{
		&Highlight{
			Fields:  make(map[Field]HighlightField, 0),
			Options: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Highlight struct
func (rb *HighlightBuilder) Build() Highlight {
	return *rb.v
}

func (rb *HighlightBuilder) BoundaryChars(boundarychars string) *HighlightBuilder {
	rb.v.BoundaryChars = &boundarychars
	return rb
}

func (rb *HighlightBuilder) BoundaryMaxScan(boundarymaxscan int) *HighlightBuilder {
	rb.v.BoundaryMaxScan = &boundarymaxscan
	return rb
}

func (rb *HighlightBuilder) BoundaryScanner(boundaryscanner boundaryscanner.BoundaryScanner) *HighlightBuilder {
	rb.v.BoundaryScanner = &boundaryscanner
	return rb
}

func (rb *HighlightBuilder) BoundaryScannerLocale(boundaryscannerlocale string) *HighlightBuilder {
	rb.v.BoundaryScannerLocale = &boundaryscannerlocale
	return rb
}

func (rb *HighlightBuilder) Encoder(encoder highlighterencoder.HighlighterEncoder) *HighlightBuilder {
	rb.v.Encoder = &encoder
	return rb
}

func (rb *HighlightBuilder) Fields(values map[Field]*HighlightFieldBuilder) *HighlightBuilder {
	tmp := make(map[Field]HighlightField, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *HighlightBuilder) ForceSource(forcesource bool) *HighlightBuilder {
	rb.v.ForceSource = &forcesource
	return rb
}

func (rb *HighlightBuilder) FragmentSize(fragmentsize int) *HighlightBuilder {
	rb.v.FragmentSize = &fragmentsize
	return rb
}

func (rb *HighlightBuilder) Fragmenter(fragmenter highlighterfragmenter.HighlighterFragmenter) *HighlightBuilder {
	rb.v.Fragmenter = &fragmenter
	return rb
}

func (rb *HighlightBuilder) HighlightFilter(highlightfilter bool) *HighlightBuilder {
	rb.v.HighlightFilter = &highlightfilter
	return rb
}

func (rb *HighlightBuilder) HighlightQuery(highlightquery *QueryContainerBuilder) *HighlightBuilder {
	v := highlightquery.Build()
	rb.v.HighlightQuery = &v
	return rb
}

func (rb *HighlightBuilder) MaxAnalyzedOffset(maxanalyzedoffset int) *HighlightBuilder {
	rb.v.MaxAnalyzedOffset = &maxanalyzedoffset
	return rb
}

func (rb *HighlightBuilder) MaxFragmentLength(maxfragmentlength int) *HighlightBuilder {
	rb.v.MaxFragmentLength = &maxfragmentlength
	return rb
}

func (rb *HighlightBuilder) NoMatchSize(nomatchsize int) *HighlightBuilder {
	rb.v.NoMatchSize = &nomatchsize
	return rb
}

func (rb *HighlightBuilder) NumberOfFragments(numberoffragments int) *HighlightBuilder {
	rb.v.NumberOfFragments = &numberoffragments
	return rb
}

func (rb *HighlightBuilder) Options(value map[string]interface{}) *HighlightBuilder {
	rb.v.Options = value
	return rb
}

func (rb *HighlightBuilder) Order(order highlighterorder.HighlighterOrder) *HighlightBuilder {
	rb.v.Order = &order
	return rb
}

func (rb *HighlightBuilder) PhraseLimit(phraselimit int) *HighlightBuilder {
	rb.v.PhraseLimit = &phraselimit
	return rb
}

func (rb *HighlightBuilder) PostTags(post_tags ...string) *HighlightBuilder {
	rb.v.PostTags = post_tags
	return rb
}

func (rb *HighlightBuilder) PreTags(pre_tags ...string) *HighlightBuilder {
	rb.v.PreTags = pre_tags
	return rb
}

func (rb *HighlightBuilder) RequireFieldMatch(requirefieldmatch bool) *HighlightBuilder {
	rb.v.RequireFieldMatch = &requirefieldmatch
	return rb
}

func (rb *HighlightBuilder) TagsSchema(tagsschema highlightertagsschema.HighlighterTagsSchema) *HighlightBuilder {
	rb.v.TagsSchema = &tagsschema
	return rb
}

func (rb *HighlightBuilder) Type_(type_ highlightertype.HighlighterType) *HighlightBuilder {
	rb.v.Type = &type_
	return rb
}
