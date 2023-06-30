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
)

// Query type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_types/query_dsl/abstractions.ts#L97-L165
type Query struct {
	Bool              *BoolQuery                        `json:"bool,omitempty"`
	Boosting          *BoostingQuery                    `json:"boosting,omitempty"`
	CombinedFields    *CombinedFieldsQuery              `json:"combined_fields,omitempty"`
	Common            map[string]CommonTermsQuery       `json:"common,omitempty"`
	ConstantScore     *ConstantScoreQuery               `json:"constant_score,omitempty"`
	DisMax            *DisMaxQuery                      `json:"dis_max,omitempty"`
	DistanceFeature   DistanceFeatureQuery              `json:"distance_feature,omitempty"`
	Exists            *ExistsQuery                      `json:"exists,omitempty"`
	FieldMaskingSpan  *SpanFieldMaskingQuery            `json:"field_masking_span,omitempty"`
	FunctionScore     *FunctionScoreQuery               `json:"function_score,omitempty"`
	Fuzzy             map[string]FuzzyQuery             `json:"fuzzy,omitempty"`
	GeoBoundingBox    *GeoBoundingBoxQuery              `json:"geo_bounding_box,omitempty"`
	GeoDistance       *GeoDistanceQuery                 `json:"geo_distance,omitempty"`
	GeoPolygon        *GeoPolygonQuery                  `json:"geo_polygon,omitempty"`
	GeoShape          *GeoShapeQuery                    `json:"geo_shape,omitempty"`
	HasChild          *HasChildQuery                    `json:"has_child,omitempty"`
	HasParent         *HasParentQuery                   `json:"has_parent,omitempty"`
	Ids               *IdsQuery                         `json:"ids,omitempty"`
	Intervals         map[string]IntervalsQuery         `json:"intervals,omitempty"`
	Match             map[string]MatchQuery             `json:"match,omitempty"`
	MatchAll          *MatchAllQuery                    `json:"match_all,omitempty"`
	MatchBoolPrefix   map[string]MatchBoolPrefixQuery   `json:"match_bool_prefix,omitempty"`
	MatchNone         *MatchNoneQuery                   `json:"match_none,omitempty"`
	MatchPhrase       map[string]MatchPhraseQuery       `json:"match_phrase,omitempty"`
	MatchPhrasePrefix map[string]MatchPhrasePrefixQuery `json:"match_phrase_prefix,omitempty"`
	MoreLikeThis      *MoreLikeThisQuery                `json:"more_like_this,omitempty"`
	MultiMatch        *MultiMatchQuery                  `json:"multi_match,omitempty"`
	Nested            *NestedQuery                      `json:"nested,omitempty"`
	ParentId          *ParentIdQuery                    `json:"parent_id,omitempty"`
	Percolate         *PercolateQuery                   `json:"percolate,omitempty"`
	Pinned            *PinnedQuery                      `json:"pinned,omitempty"`
	Prefix            map[string]PrefixQuery            `json:"prefix,omitempty"`
	QueryString       *QueryStringQuery                 `json:"query_string,omitempty"`
	Range             map[string]RangeQuery             `json:"range,omitempty"`
	RankFeature       *RankFeatureQuery                 `json:"rank_feature,omitempty"`
	Regexp            map[string]RegexpQuery            `json:"regexp,omitempty"`
	Script            *ScriptQuery                      `json:"script,omitempty"`
	ScriptScore       *ScriptScoreQuery                 `json:"script_score,omitempty"`
	Shape             *ShapeQuery                       `json:"shape,omitempty"`
	SimpleQueryString *SimpleQueryStringQuery           `json:"simple_query_string,omitempty"`
	SpanContaining    *SpanContainingQuery              `json:"span_containing,omitempty"`
	SpanFirst         *SpanFirstQuery                   `json:"span_first,omitempty"`
	SpanMulti         *SpanMultiTermQuery               `json:"span_multi,omitempty"`
	SpanNear          *SpanNearQuery                    `json:"span_near,omitempty"`
	SpanNot           *SpanNotQuery                     `json:"span_not,omitempty"`
	SpanOr            *SpanOrQuery                      `json:"span_or,omitempty"`
	SpanTerm          map[string]SpanTermQuery          `json:"span_term,omitempty"`
	SpanWithin        *SpanWithinQuery                  `json:"span_within,omitempty"`
	Term              map[string]TermQuery              `json:"term,omitempty"`
	Terms             *TermsQuery                       `json:"terms,omitempty"`
	TermsSet          map[string]TermsSetQuery          `json:"terms_set,omitempty"`
	TextExpansion     *TextExpansionQuery               `json:"text_expansion,omitempty"`
	Type              *TypeQuery                        `json:"type,omitempty"`
	Wildcard          map[string]WildcardQuery          `json:"wildcard,omitempty"`
	Wrapper           *WrapperQuery                     `json:"wrapper,omitempty"`
}

func (s *Query) UnmarshalJSON(data []byte) error {

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

		case "bool":
			if err := dec.Decode(&s.Bool); err != nil {
				return err
			}

		case "boosting":
			if err := dec.Decode(&s.Boosting); err != nil {
				return err
			}

		case "combined_fields":
			if err := dec.Decode(&s.CombinedFields); err != nil {
				return err
			}

		case "common":
			if s.Common == nil {
				s.Common = make(map[string]CommonTermsQuery, 0)
			}
			if err := dec.Decode(&s.Common); err != nil {
				return err
			}

		case "constant_score":
			if err := dec.Decode(&s.ConstantScore); err != nil {
				return err
			}

		case "dis_max":
			if err := dec.Decode(&s.DisMax); err != nil {
				return err
			}

		case "distance_feature":
			if err := dec.Decode(&s.DistanceFeature); err != nil {
				return err
			}

		case "exists":
			if err := dec.Decode(&s.Exists); err != nil {
				return err
			}

		case "field_masking_span":
			if err := dec.Decode(&s.FieldMaskingSpan); err != nil {
				return err
			}

		case "function_score":
			if err := dec.Decode(&s.FunctionScore); err != nil {
				return err
			}

		case "fuzzy":
			if s.Fuzzy == nil {
				s.Fuzzy = make(map[string]FuzzyQuery, 0)
			}
			if err := dec.Decode(&s.Fuzzy); err != nil {
				return err
			}

		case "geo_bounding_box":
			if err := dec.Decode(&s.GeoBoundingBox); err != nil {
				return err
			}

		case "geo_distance":
			if err := dec.Decode(&s.GeoDistance); err != nil {
				return err
			}

		case "geo_polygon":
			if err := dec.Decode(&s.GeoPolygon); err != nil {
				return err
			}

		case "geo_shape":
			if err := dec.Decode(&s.GeoShape); err != nil {
				return err
			}

		case "has_child":
			if err := dec.Decode(&s.HasChild); err != nil {
				return err
			}

		case "has_parent":
			if err := dec.Decode(&s.HasParent); err != nil {
				return err
			}

		case "ids":
			if err := dec.Decode(&s.Ids); err != nil {
				return err
			}

		case "intervals":
			if s.Intervals == nil {
				s.Intervals = make(map[string]IntervalsQuery, 0)
			}
			if err := dec.Decode(&s.Intervals); err != nil {
				return err
			}

		case "match":
			if s.Match == nil {
				s.Match = make(map[string]MatchQuery, 0)
			}
			if err := dec.Decode(&s.Match); err != nil {
				return err
			}

		case "match_all":
			if err := dec.Decode(&s.MatchAll); err != nil {
				return err
			}

		case "match_bool_prefix":
			if s.MatchBoolPrefix == nil {
				s.MatchBoolPrefix = make(map[string]MatchBoolPrefixQuery, 0)
			}
			if err := dec.Decode(&s.MatchBoolPrefix); err != nil {
				return err
			}

		case "match_none":
			if err := dec.Decode(&s.MatchNone); err != nil {
				return err
			}

		case "match_phrase":
			if s.MatchPhrase == nil {
				s.MatchPhrase = make(map[string]MatchPhraseQuery, 0)
			}
			if err := dec.Decode(&s.MatchPhrase); err != nil {
				return err
			}

		case "match_phrase_prefix":
			if s.MatchPhrasePrefix == nil {
				s.MatchPhrasePrefix = make(map[string]MatchPhrasePrefixQuery, 0)
			}
			if err := dec.Decode(&s.MatchPhrasePrefix); err != nil {
				return err
			}

		case "more_like_this":
			if err := dec.Decode(&s.MoreLikeThis); err != nil {
				return err
			}

		case "multi_match":
			if err := dec.Decode(&s.MultiMatch); err != nil {
				return err
			}

		case "nested":
			if err := dec.Decode(&s.Nested); err != nil {
				return err
			}

		case "parent_id":
			if err := dec.Decode(&s.ParentId); err != nil {
				return err
			}

		case "percolate":
			if err := dec.Decode(&s.Percolate); err != nil {
				return err
			}

		case "pinned":
			if err := dec.Decode(&s.Pinned); err != nil {
				return err
			}

		case "prefix":
			if s.Prefix == nil {
				s.Prefix = make(map[string]PrefixQuery, 0)
			}
			if err := dec.Decode(&s.Prefix); err != nil {
				return err
			}

		case "query_string":
			if err := dec.Decode(&s.QueryString); err != nil {
				return err
			}

		case "range":
			if s.Range == nil {
				s.Range = make(map[string]RangeQuery, 0)
			}
			if err := dec.Decode(&s.Range); err != nil {
				return err
			}

		case "rank_feature":
			if err := dec.Decode(&s.RankFeature); err != nil {
				return err
			}

		case "regexp":
			if s.Regexp == nil {
				s.Regexp = make(map[string]RegexpQuery, 0)
			}
			if err := dec.Decode(&s.Regexp); err != nil {
				return err
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		case "script_score":
			if err := dec.Decode(&s.ScriptScore); err != nil {
				return err
			}

		case "shape":
			if err := dec.Decode(&s.Shape); err != nil {
				return err
			}

		case "simple_query_string":
			if err := dec.Decode(&s.SimpleQueryString); err != nil {
				return err
			}

		case "span_containing":
			if err := dec.Decode(&s.SpanContaining); err != nil {
				return err
			}

		case "span_first":
			if err := dec.Decode(&s.SpanFirst); err != nil {
				return err
			}

		case "span_multi":
			if err := dec.Decode(&s.SpanMulti); err != nil {
				return err
			}

		case "span_near":
			if err := dec.Decode(&s.SpanNear); err != nil {
				return err
			}

		case "span_not":
			if err := dec.Decode(&s.SpanNot); err != nil {
				return err
			}

		case "span_or":
			if err := dec.Decode(&s.SpanOr); err != nil {
				return err
			}

		case "span_term":
			if s.SpanTerm == nil {
				s.SpanTerm = make(map[string]SpanTermQuery, 0)
			}
			if err := dec.Decode(&s.SpanTerm); err != nil {
				return err
			}

		case "span_within":
			if err := dec.Decode(&s.SpanWithin); err != nil {
				return err
			}

		case "term":
			if s.Term == nil {
				s.Term = make(map[string]TermQuery, 0)
			}
			if err := dec.Decode(&s.Term); err != nil {
				return err
			}

		case "terms":
			if err := dec.Decode(&s.Terms); err != nil {
				return err
			}

		case "terms_set":
			if s.TermsSet == nil {
				s.TermsSet = make(map[string]TermsSetQuery, 0)
			}
			if err := dec.Decode(&s.TermsSet); err != nil {
				return err
			}

		case "text_expansion":
			if err := dec.Decode(&s.TextExpansion); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "wildcard":
			if s.Wildcard == nil {
				s.Wildcard = make(map[string]WildcardQuery, 0)
			}
			if err := dec.Decode(&s.Wildcard); err != nil {
				return err
			}

		case "wrapper":
			if err := dec.Decode(&s.Wrapper); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewQuery returns a Query.
func NewQuery() *Query {
	r := &Query{
		Common:            make(map[string]CommonTermsQuery, 0),
		Fuzzy:             make(map[string]FuzzyQuery, 0),
		Intervals:         make(map[string]IntervalsQuery, 0),
		Match:             make(map[string]MatchQuery, 0),
		MatchBoolPrefix:   make(map[string]MatchBoolPrefixQuery, 0),
		MatchPhrase:       make(map[string]MatchPhraseQuery, 0),
		MatchPhrasePrefix: make(map[string]MatchPhrasePrefixQuery, 0),
		Prefix:            make(map[string]PrefixQuery, 0),
		Range:             make(map[string]RangeQuery, 0),
		Regexp:            make(map[string]RegexpQuery, 0),
		SpanTerm:          make(map[string]SpanTermQuery, 0),
		Term:              make(map[string]TermQuery, 0),
		TermsSet:          make(map[string]TermsSetQuery, 0),
		Wildcard:          make(map[string]WildcardQuery, 0),
	}

	return r
}
