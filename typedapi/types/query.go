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

// Query type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_types/query_dsl/abstractions.ts#L96-L162
type Query struct {
	Bool              *BoolQuery                        `json:"bool,omitempty"`
	Boosting          *BoostingQuery                    `json:"boosting,omitempty"`
	CombinedFields    *CombinedFieldsQuery              `json:"combined_fields,omitempty"`
	Common            map[string]CommonTermsQuery       `json:"common,omitempty"`
	ConstantScore     *ConstantScoreQuery               `json:"constant_score,omitempty"`
	DisMax            *DisMaxQuery                      `json:"dis_max,omitempty"`
	DistanceFeature   *DistanceFeatureQuery             `json:"distance_feature,omitempty"`
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
	Type              *TypeQuery                        `json:"type,omitempty"`
	Wildcard          map[string]WildcardQuery          `json:"wildcard,omitempty"`
	Wrapper           *WrapperQuery                     `json:"wrapper,omitempty"`
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
