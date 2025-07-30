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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/highlightertype"
)

type _highlightField struct {
	v *types.HighlightField
}

func NewHighlightField() *_highlightField {

	return &_highlightField{v: types.NewHighlightField()}

}

func (s *_highlightField) BoundaryChars(boundarychars string) *_highlightField {

	s.v.BoundaryChars = &boundarychars

	return s
}

func (s *_highlightField) BoundaryMaxScan(boundarymaxscan int) *_highlightField {

	s.v.BoundaryMaxScan = &boundarymaxscan

	return s
}

func (s *_highlightField) BoundaryScanner(boundaryscanner boundaryscanner.BoundaryScanner) *_highlightField {

	s.v.BoundaryScanner = &boundaryscanner
	return s
}

func (s *_highlightField) BoundaryScannerLocale(boundaryscannerlocale string) *_highlightField {

	s.v.BoundaryScannerLocale = &boundaryscannerlocale

	return s
}

func (s *_highlightField) ForceSource(forcesource bool) *_highlightField {

	s.v.ForceSource = &forcesource

	return s
}

func (s *_highlightField) FragmentOffset(fragmentoffset int) *_highlightField {

	s.v.FragmentOffset = &fragmentoffset

	return s
}

func (s *_highlightField) FragmentSize(fragmentsize int) *_highlightField {

	s.v.FragmentSize = &fragmentsize

	return s
}

func (s *_highlightField) Fragmenter(fragmenter highlighterfragmenter.HighlighterFragmenter) *_highlightField {

	s.v.Fragmenter = &fragmenter
	return s
}

func (s *_highlightField) HighlightFilter(highlightfilter bool) *_highlightField {

	s.v.HighlightFilter = &highlightfilter

	return s
}

func (s *_highlightField) HighlightQuery(highlightquery types.QueryVariant) *_highlightField {

	s.v.HighlightQuery = highlightquery.QueryCaster()

	return s
}

func (s *_highlightField) MatchedFields(fields ...string) *_highlightField {

	s.v.MatchedFields = fields

	return s
}

func (s *_highlightField) MaxAnalyzedOffset(maxanalyzedoffset int) *_highlightField {

	s.v.MaxAnalyzedOffset = &maxanalyzedoffset

	return s
}

func (s *_highlightField) MaxFragmentLength(maxfragmentlength int) *_highlightField {

	s.v.MaxFragmentLength = &maxfragmentlength

	return s
}

func (s *_highlightField) NoMatchSize(nomatchsize int) *_highlightField {

	s.v.NoMatchSize = &nomatchsize

	return s
}

func (s *_highlightField) NumberOfFragments(numberoffragments int) *_highlightField {

	s.v.NumberOfFragments = &numberoffragments

	return s
}

func (s *_highlightField) Options(options map[string]json.RawMessage) *_highlightField {

	s.v.Options = options
	return s
}

func (s *_highlightField) AddOption(key string, value json.RawMessage) *_highlightField {

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

func (s *_highlightField) Order(order highlighterorder.HighlighterOrder) *_highlightField {

	s.v.Order = &order
	return s
}

func (s *_highlightField) PhraseLimit(phraselimit int) *_highlightField {

	s.v.PhraseLimit = &phraselimit

	return s
}

func (s *_highlightField) PostTags(posttags ...string) *_highlightField {

	for _, v := range posttags {

		s.v.PostTags = append(s.v.PostTags, v)

	}
	return s
}

func (s *_highlightField) PreTags(pretags ...string) *_highlightField {

	for _, v := range pretags {

		s.v.PreTags = append(s.v.PreTags, v)

	}
	return s
}

func (s *_highlightField) RequireFieldMatch(requirefieldmatch bool) *_highlightField {

	s.v.RequireFieldMatch = &requirefieldmatch

	return s
}

func (s *_highlightField) TagsSchema(tagsschema highlightertagsschema.HighlighterTagsSchema) *_highlightField {

	s.v.TagsSchema = &tagsschema
	return s
}

func (s *_highlightField) Type(type_ highlightertype.HighlighterType) *_highlightField {

	s.v.Type = &type_
	return s
}

func (s *_highlightField) HighlightFieldCaster() *types.HighlightField {
	return s.v
}
