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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


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
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_global/search/_types/highlighting.ts#L57-L60
type Highlight struct {
	BoundaryChars         *string                                      `json:"boundary_chars,omitempty"`
	BoundaryMaxScan       *int                                         `json:"boundary_max_scan,omitempty"`
	BoundaryScanner       *boundaryscanner.BoundaryScanner             `json:"boundary_scanner,omitempty"`
	BoundaryScannerLocale *string                                      `json:"boundary_scanner_locale,omitempty"`
	Encoder               *highlighterencoder.HighlighterEncoder       `json:"encoder,omitempty"`
	Fields                map[string]HighlightField                    `json:"fields"`
	ForceSource           *bool                                        `json:"force_source,omitempty"`
	FragmentSize          *int                                         `json:"fragment_size,omitempty"`
	Fragmenter            *highlighterfragmenter.HighlighterFragmenter `json:"fragmenter,omitempty"`
	HighlightFilter       *bool                                        `json:"highlight_filter,omitempty"`
	HighlightQuery        *Query                                       `json:"highlight_query,omitempty"`
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

// NewHighlight returns a Highlight.
func NewHighlight() *Highlight {
	r := &Highlight{
		Fields:  make(map[string]HighlightField, 0),
		Options: make(map[string]interface{}, 0),
	}

	return r
}
