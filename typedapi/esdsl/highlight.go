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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterencoder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertype"
)

type _highlight struct {
	v *types.Highlight
}

func NewHighlight() *_highlight {

	return &_highlight{v: types.NewHighlight()}

}

func (s *_highlight) BoundaryChars(boundarychars string) *_highlight {

	s.v.BoundaryChars = &boundarychars

	return s
}

func (s *_highlight) BoundaryMaxScan(boundarymaxscan int) *_highlight {

	s.v.BoundaryMaxScan = &boundarymaxscan

	return s
}

func (s *_highlight) BoundaryScanner(boundaryscanner boundaryscanner.BoundaryScanner) *_highlight {

	s.v.BoundaryScanner = &boundaryscanner
	return s
}

func (s *_highlight) BoundaryScannerLocale(boundaryscannerlocale string) *_highlight {

	s.v.BoundaryScannerLocale = &boundaryscannerlocale

	return s
}

func (s *_highlight) Encoder(encoder highlighterencoder.HighlighterEncoder) *_highlight {

	s.v.Encoder = &encoder
	return s
}

func (s *_highlight) Fields(fields map[string]types.HighlightField) *_highlight {

	s.v.Fields = fields
	return s
}

func (s *_highlight) AddField(key string, value types.HighlightFieldVariant) *_highlight {

	var tmp map[string]types.HighlightField
	if s.v.Fields == nil {
		s.v.Fields = make(map[string]types.HighlightField)
	} else {
		tmp = s.v.Fields
	}

	tmp[key] = *value.HighlightFieldCaster()

	s.v.Fields = tmp
	return s
}

func (s *_highlight) ForceSource(forcesource bool) *_highlight {

	s.v.ForceSource = &forcesource

	return s
}

func (s *_highlight) FragmentSize(fragmentsize int) *_highlight {

	s.v.FragmentSize = &fragmentsize

	return s
}

func (s *_highlight) Fragmenter(fragmenter highlighterfragmenter.HighlighterFragmenter) *_highlight {

	s.v.Fragmenter = &fragmenter
	return s
}

func (s *_highlight) HighlightFilter(highlightfilter bool) *_highlight {

	s.v.HighlightFilter = &highlightfilter

	return s
}

func (s *_highlight) HighlightQuery(highlightquery types.QueryVariant) *_highlight {

	s.v.HighlightQuery = highlightquery.QueryCaster()

	return s
}

func (s *_highlight) MaxAnalyzedOffset(maxanalyzedoffset int) *_highlight {

	s.v.MaxAnalyzedOffset = &maxanalyzedoffset

	return s
}

func (s *_highlight) MaxFragmentLength(maxfragmentlength int) *_highlight {

	s.v.MaxFragmentLength = &maxfragmentlength

	return s
}

func (s *_highlight) NoMatchSize(nomatchsize int) *_highlight {

	s.v.NoMatchSize = &nomatchsize

	return s
}

func (s *_highlight) NumberOfFragments(numberoffragments int) *_highlight {

	s.v.NumberOfFragments = &numberoffragments

	return s
}

func (s *_highlight) Options(options map[string]json.RawMessage) *_highlight {

	s.v.Options = options
	return s
}

func (s *_highlight) AddOption(key string, value json.RawMessage) *_highlight {

	var tmp map[string]json.RawMessage
	if s.v.Options == nil {
		s.v.Options = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Options
	}

	tmp[key] = value

	s.v.Options = tmp
	return s
}

func (s *_highlight) Order(order highlighterorder.HighlighterOrder) *_highlight {

	s.v.Order = &order
	return s
}

func (s *_highlight) PhraseLimit(phraselimit int) *_highlight {

	s.v.PhraseLimit = &phraselimit

	return s
}

func (s *_highlight) PostTags(posttags ...string) *_highlight {

	for _, v := range posttags {

		s.v.PostTags = append(s.v.PostTags, v)

	}
	return s
}

func (s *_highlight) PreTags(pretags ...string) *_highlight {

	for _, v := range pretags {

		s.v.PreTags = append(s.v.PreTags, v)

	}
	return s
}

func (s *_highlight) RequireFieldMatch(requirefieldmatch bool) *_highlight {

	s.v.RequireFieldMatch = &requirefieldmatch

	return s
}

func (s *_highlight) TagsSchema(tagsschema highlightertagsschema.HighlighterTagsSchema) *_highlight {

	s.v.TagsSchema = &tagsschema
	return s
}

func (s *_highlight) Type(type_ highlightertype.HighlighterType) *_highlight {

	s.v.Type = &type_
	return s
}

func (s *_highlight) HighlightCaster() *types.Highlight {
	return s.v
}
