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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertype"
)

// HighlightBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/highlighting.ts#L32-L54
type HighlightBase struct {
	BoundaryChars         *string                                      `json:"boundary_chars,omitempty"`
	BoundaryMaxScan       *int                                         `json:"boundary_max_scan,omitempty"`
	BoundaryScanner       *boundaryscanner.BoundaryScanner             `json:"boundary_scanner,omitempty"`
	BoundaryScannerLocale *string                                      `json:"boundary_scanner_locale,omitempty"`
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

// HighlightBaseBuilder holds HighlightBase struct and provides a builder API.
type HighlightBaseBuilder struct {
	v *HighlightBase
}

// NewHighlightBase provides a builder for the HighlightBase struct.
func NewHighlightBaseBuilder() *HighlightBaseBuilder {
	r := HighlightBaseBuilder{
		&HighlightBase{
			Options: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the HighlightBase struct
func (rb *HighlightBaseBuilder) Build() HighlightBase {
	return *rb.v
}

func (rb *HighlightBaseBuilder) BoundaryChars(boundarychars string) *HighlightBaseBuilder {
	rb.v.BoundaryChars = &boundarychars
	return rb
}

func (rb *HighlightBaseBuilder) BoundaryMaxScan(boundarymaxscan int) *HighlightBaseBuilder {
	rb.v.BoundaryMaxScan = &boundarymaxscan
	return rb
}

func (rb *HighlightBaseBuilder) BoundaryScanner(boundaryscanner boundaryscanner.BoundaryScanner) *HighlightBaseBuilder {
	rb.v.BoundaryScanner = &boundaryscanner
	return rb
}

func (rb *HighlightBaseBuilder) BoundaryScannerLocale(boundaryscannerlocale string) *HighlightBaseBuilder {
	rb.v.BoundaryScannerLocale = &boundaryscannerlocale
	return rb
}

func (rb *HighlightBaseBuilder) ForceSource(forcesource bool) *HighlightBaseBuilder {
	rb.v.ForceSource = &forcesource
	return rb
}

func (rb *HighlightBaseBuilder) FragmentSize(fragmentsize int) *HighlightBaseBuilder {
	rb.v.FragmentSize = &fragmentsize
	return rb
}

func (rb *HighlightBaseBuilder) Fragmenter(fragmenter highlighterfragmenter.HighlighterFragmenter) *HighlightBaseBuilder {
	rb.v.Fragmenter = &fragmenter
	return rb
}

func (rb *HighlightBaseBuilder) HighlightFilter(highlightfilter bool) *HighlightBaseBuilder {
	rb.v.HighlightFilter = &highlightfilter
	return rb
}

func (rb *HighlightBaseBuilder) HighlightQuery(highlightquery *QueryContainerBuilder) *HighlightBaseBuilder {
	v := highlightquery.Build()
	rb.v.HighlightQuery = &v
	return rb
}

func (rb *HighlightBaseBuilder) MaxAnalyzedOffset(maxanalyzedoffset int) *HighlightBaseBuilder {
	rb.v.MaxAnalyzedOffset = &maxanalyzedoffset
	return rb
}

func (rb *HighlightBaseBuilder) MaxFragmentLength(maxfragmentlength int) *HighlightBaseBuilder {
	rb.v.MaxFragmentLength = &maxfragmentlength
	return rb
}

func (rb *HighlightBaseBuilder) NoMatchSize(nomatchsize int) *HighlightBaseBuilder {
	rb.v.NoMatchSize = &nomatchsize
	return rb
}

func (rb *HighlightBaseBuilder) NumberOfFragments(numberoffragments int) *HighlightBaseBuilder {
	rb.v.NumberOfFragments = &numberoffragments
	return rb
}

func (rb *HighlightBaseBuilder) Options(value map[string]interface{}) *HighlightBaseBuilder {
	rb.v.Options = value
	return rb
}

func (rb *HighlightBaseBuilder) Order(order highlighterorder.HighlighterOrder) *HighlightBaseBuilder {
	rb.v.Order = &order
	return rb
}

func (rb *HighlightBaseBuilder) PhraseLimit(phraselimit int) *HighlightBaseBuilder {
	rb.v.PhraseLimit = &phraselimit
	return rb
}

func (rb *HighlightBaseBuilder) PostTags(post_tags ...string) *HighlightBaseBuilder {
	rb.v.PostTags = post_tags
	return rb
}

func (rb *HighlightBaseBuilder) PreTags(pre_tags ...string) *HighlightBaseBuilder {
	rb.v.PreTags = pre_tags
	return rb
}

func (rb *HighlightBaseBuilder) RequireFieldMatch(requirefieldmatch bool) *HighlightBaseBuilder {
	rb.v.RequireFieldMatch = &requirefieldmatch
	return rb
}

func (rb *HighlightBaseBuilder) TagsSchema(tagsschema highlightertagsschema.HighlighterTagsSchema) *HighlightBaseBuilder {
	rb.v.TagsSchema = &tagsschema
	return rb
}

func (rb *HighlightBaseBuilder) Type_(type_ highlightertype.HighlighterType) *HighlightBaseBuilder {
	rb.v.Type = &type_
	return rb
}
