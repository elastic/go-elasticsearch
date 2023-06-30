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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterencoder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertype"
)

// Highlight type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/search/_types/highlighting.ts#L57-L60
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
	Options               map[string]json.RawMessage                   `json:"options,omitempty"`
	Order                 *highlighterorder.HighlighterOrder           `json:"order,omitempty"`
	PhraseLimit           *int                                         `json:"phrase_limit,omitempty"`
	PostTags              []string                                     `json:"post_tags,omitempty"`
	PreTags               []string                                     `json:"pre_tags,omitempty"`
	RequireFieldMatch     *bool                                        `json:"require_field_match,omitempty"`
	TagsSchema            *highlightertagsschema.HighlighterTagsSchema `json:"tags_schema,omitempty"`
	Type                  *highlightertype.HighlighterType             `json:"type,omitempty"`
}

func (s *Highlight) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "boundary_chars":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BoundaryChars = &o

		case "boundary_max_scan":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.BoundaryMaxScan = &value
			case float64:
				f := int(v)
				s.BoundaryMaxScan = &f
			}

		case "boundary_scanner":
			if err := dec.Decode(&s.BoundaryScanner); err != nil {
				return err
			}

		case "boundary_scanner_locale":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.BoundaryScannerLocale = &o

		case "encoder":
			if err := dec.Decode(&s.Encoder); err != nil {
				return err
			}

		case "fields":
			if s.Fields == nil {
				s.Fields = make(map[string]HighlightField, 0)
			}
			if err := dec.Decode(&s.Fields); err != nil {
				return err
			}

		case "force_source":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.ForceSource = &value
			case bool:
				s.ForceSource = &v
			}

		case "fragment_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FragmentSize = &value
			case float64:
				f := int(v)
				s.FragmentSize = &f
			}

		case "fragmenter":
			if err := dec.Decode(&s.Fragmenter); err != nil {
				return err
			}

		case "highlight_filter":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.HighlightFilter = &value
			case bool:
				s.HighlightFilter = &v
			}

		case "highlight_query":
			if err := dec.Decode(&s.HighlightQuery); err != nil {
				return err
			}

		case "max_analyzed_offset":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxAnalyzedOffset = &value
			case float64:
				f := int(v)
				s.MaxAnalyzedOffset = &f
			}

		case "max_fragment_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxFragmentLength = &value
			case float64:
				f := int(v)
				s.MaxFragmentLength = &f
			}

		case "no_match_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NoMatchSize = &value
			case float64:
				f := int(v)
				s.NoMatchSize = &f
			}

		case "number_of_fragments":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumberOfFragments = &value
			case float64:
				f := int(v)
				s.NumberOfFragments = &f
			}

		case "options":
			if s.Options == nil {
				s.Options = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Options); err != nil {
				return err
			}

		case "order":
			if err := dec.Decode(&s.Order); err != nil {
				return err
			}

		case "phrase_limit":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PhraseLimit = &value
			case float64:
				f := int(v)
				s.PhraseLimit = &f
			}

		case "post_tags":
			if err := dec.Decode(&s.PostTags); err != nil {
				return err
			}

		case "pre_tags":
			if err := dec.Decode(&s.PreTags); err != nil {
				return err
			}

		case "require_field_match":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.RequireFieldMatch = &value
			case bool:
				s.RequireFieldMatch = &v
			}

		case "tags_schema":
			if err := dec.Decode(&s.TagsSchema); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewHighlight returns a Highlight.
func NewHighlight() *Highlight {
	r := &Highlight{
		Fields:  make(map[string]HighlightField, 0),
		Options: make(map[string]json.RawMessage, 0),
	}

	return r
}
