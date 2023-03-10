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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertype"

	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// HighlightField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_global/search/_types/highlighting.ts#L88-L92
type HighlightField struct {
	Analyzer              Analyzer                                     `json:"analyzer,omitempty"`
	BoundaryChars         *string                                      `json:"boundary_chars,omitempty"`
	BoundaryMaxScan       *int                                         `json:"boundary_max_scan,omitempty"`
	BoundaryScanner       *boundaryscanner.BoundaryScanner             `json:"boundary_scanner,omitempty"`
	BoundaryScannerLocale *string                                      `json:"boundary_scanner_locale,omitempty"`
	ForceSource           *bool                                        `json:"force_source,omitempty"`
	FragmentOffset        *int                                         `json:"fragment_offset,omitempty"`
	FragmentSize          *int                                         `json:"fragment_size,omitempty"`
	Fragmenter            *highlighterfragmenter.HighlighterFragmenter `json:"fragmenter,omitempty"`
	HighlightFilter       *bool                                        `json:"highlight_filter,omitempty"`
	HighlightQuery        *Query                                       `json:"highlight_query,omitempty"`
	MatchedFields         []string                                     `json:"matched_fields,omitempty"`
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

func (s *HighlightField) UnmarshalJSON(data []byte) error {
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

		case "analyzer":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			kind := make(map[string]string, 0)
			localDec := json.NewDecoder(source)
			localDec.Decode(&kind)
			source.Seek(0, io.SeekStart)

			switch kind["type"] {

			case "custom":
				o := NewCustomAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "fingerprint":
				o := NewFingerprintAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "keyword":
				o := NewKeywordAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "language":
				o := NewLanguageAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "nori":
				o := NewNoriAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "pattern":
				o := NewPatternAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "simple":
				o := NewSimpleAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "standard":
				o := NewStandardAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "stop":
				o := NewStopAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "whitespace":
				o := NewWhitespaceAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "icu_analyzer":
				o := NewIcuAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "kuromoji":
				o := NewKuromojiAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "snowball":
				o := NewSnowballAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			case "dutch":
				o := NewDutchAnalyzer()
				if err := localDec.Decode(o); err != nil {
					return err
				}
				s.Analyzer = *o
			default:
				if err := dec.Decode(&s.Analyzer); err != nil {
					return err
				}
			}

		case "boundary_chars":
			if err := dec.Decode(&s.BoundaryChars); err != nil {
				return err
			}

		case "boundary_max_scan":
			if err := dec.Decode(&s.BoundaryMaxScan); err != nil {
				return err
			}

		case "boundary_scanner":
			if err := dec.Decode(&s.BoundaryScanner); err != nil {
				return err
			}

		case "boundary_scanner_locale":
			if err := dec.Decode(&s.BoundaryScannerLocale); err != nil {
				return err
			}

		case "force_source":
			if err := dec.Decode(&s.ForceSource); err != nil {
				return err
			}

		case "fragment_offset":
			if err := dec.Decode(&s.FragmentOffset); err != nil {
				return err
			}

		case "fragment_size":
			if err := dec.Decode(&s.FragmentSize); err != nil {
				return err
			}

		case "fragmenter":
			if err := dec.Decode(&s.Fragmenter); err != nil {
				return err
			}

		case "highlight_filter":
			if err := dec.Decode(&s.HighlightFilter); err != nil {
				return err
			}

		case "highlight_query":
			if err := dec.Decode(&s.HighlightQuery); err != nil {
				return err
			}

		case "matched_fields":
			if err := dec.Decode(&s.MatchedFields); err != nil {
				return err
			}

		case "max_analyzed_offset":
			if err := dec.Decode(&s.MaxAnalyzedOffset); err != nil {
				return err
			}

		case "max_fragment_length":
			if err := dec.Decode(&s.MaxFragmentLength); err != nil {
				return err
			}

		case "no_match_size":
			if err := dec.Decode(&s.NoMatchSize); err != nil {
				return err
			}

		case "number_of_fragments":
			if err := dec.Decode(&s.NumberOfFragments); err != nil {
				return err
			}

		case "options":
			if err := dec.Decode(&s.Options); err != nil {
				return err
			}

		case "order":
			if err := dec.Decode(&s.Order); err != nil {
				return err
			}

		case "phrase_limit":
			if err := dec.Decode(&s.PhraseLimit); err != nil {
				return err
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
			if err := dec.Decode(&s.RequireFieldMatch); err != nil {
				return err
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

// NewHighlightField returns a HighlightField.
func NewHighlightField() *HighlightField {
	r := &HighlightField{
		Options: make(map[string]json.RawMessage, 0),
	}

	return r
}
