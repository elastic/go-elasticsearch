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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

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
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_global/search/_types/highlighting.ts#L153-L156
type Highlight struct {
	// BoundaryChars A string that contains each boundary character.
	BoundaryChars *string `json:"boundary_chars,omitempty"`
	// BoundaryMaxScan How far to scan for boundary characters.
	BoundaryMaxScan *int `json:"boundary_max_scan,omitempty"`
	// BoundaryScanner Specifies how to break the highlighted fragments: chars, sentence, or word.
	// Only valid for the unified and fvh highlighters.
	// Defaults to `sentence` for the `unified` highlighter. Defaults to `chars` for
	// the `fvh` highlighter.
	BoundaryScanner *boundaryscanner.BoundaryScanner `json:"boundary_scanner,omitempty"`
	// BoundaryScannerLocale Controls which locale is used to search for sentence and word boundaries.
	// This parameter takes a form of a language tag, for example: `"en-US"`,
	// `"fr-FR"`, `"ja-JP"`.
	BoundaryScannerLocale *string                                `json:"boundary_scanner_locale,omitempty"`
	Encoder               *highlighterencoder.HighlighterEncoder `json:"encoder,omitempty"`
	Fields                map[string]HighlightField              `json:"fields"`
	ForceSource           *bool                                  `json:"force_source,omitempty"`
	// FragmentSize The size of the highlighted fragment in characters.
	FragmentSize *int `json:"fragment_size,omitempty"`
	// Fragmenter Specifies how text should be broken up in highlight snippets: `simple` or
	// `span`.
	// Only valid for the `plain` highlighter.
	Fragmenter      *highlighterfragmenter.HighlighterFragmenter `json:"fragmenter,omitempty"`
	HighlightFilter *bool                                        `json:"highlight_filter,omitempty"`
	// HighlightQuery Highlight matches for a query other than the search query.
	// This is especially useful if you use a rescore query because those are not
	// taken into account by highlighting by default.
	HighlightQuery *Query `json:"highlight_query,omitempty"`
	// MaxAnalyzedOffset If set to a non-negative value, highlighting stops at this defined maximum
	// limit.
	// The rest of the text is not processed, thus not highlighted and no error is
	// returned
	// The `max_analyzed_offset` query setting does not override the
	// `index.highlight.max_analyzed_offset` setting, which prevails when it’s set
	// to lower value than the query setting.
	MaxAnalyzedOffset *int `json:"max_analyzed_offset,omitempty"`
	MaxFragmentLength *int `json:"max_fragment_length,omitempty"`
	// NoMatchSize The amount of text you want to return from the beginning of the field if
	// there are no matching fragments to highlight.
	NoMatchSize *int `json:"no_match_size,omitempty"`
	// NumberOfFragments The maximum number of fragments to return.
	// If the number of fragments is set to `0`, no fragments are returned.
	// Instead, the entire field contents are highlighted and returned.
	// This can be handy when you need to highlight short texts such as a title or
	// address, but fragmentation is not required.
	// If `number_of_fragments` is `0`, `fragment_size` is ignored.
	NumberOfFragments *int                       `json:"number_of_fragments,omitempty"`
	Options           map[string]json.RawMessage `json:"options,omitempty"`
	// Order Sorts highlighted fragments by score when set to `score`.
	// By default, fragments will be output in the order they appear in the field
	// (order: `none`).
	// Setting this option to `score` will output the most relevant fragments first.
	// Each highlighter applies its own logic to compute relevancy scores.
	Order *highlighterorder.HighlighterOrder `json:"order,omitempty"`
	// PhraseLimit Controls the number of matching phrases in a document that are considered.
	// Prevents the `fvh` highlighter from analyzing too many phrases and consuming
	// too much memory.
	// When using `matched_fields`, `phrase_limit` phrases per matched field are
	// considered. Raising the limit increases query time and consumes more memory.
	// Only supported by the `fvh` highlighter.
	PhraseLimit *int `json:"phrase_limit,omitempty"`
	// PostTags Use in conjunction with `pre_tags` to define the HTML tags to use for the
	// highlighted text.
	// By default, highlighted text is wrapped in `<em>` and `</em>` tags.
	PostTags []string `json:"post_tags,omitempty"`
	// PreTags Use in conjunction with `post_tags` to define the HTML tags to use for the
	// highlighted text.
	// By default, highlighted text is wrapped in `<em>` and `</em>` tags.
	PreTags []string `json:"pre_tags,omitempty"`
	// RequireFieldMatch By default, only fields that contains a query match are highlighted.
	// Set to `false` to highlight all fields.
	RequireFieldMatch *bool `json:"require_field_match,omitempty"`
	// TagsSchema Set to `styled` to use the built-in tag schema.
	TagsSchema *highlightertagsschema.HighlighterTagsSchema `json:"tags_schema,omitempty"`
	Type       *highlightertype.HighlighterType             `json:"type,omitempty"`
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
