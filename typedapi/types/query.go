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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// Query type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/query_dsl/abstractions.ts#L98-L391
type Query struct {
	// Bool matches documents matching boolean combinations of other queries.
	Bool *BoolQuery `json:"bool,omitempty"`
	// Boosting Returns documents matching a `positive` query while reducing the relevance
	// score of documents that also match a `negative` query.
	Boosting *BoostingQuery `json:"boosting,omitempty"`
	// CombinedFields The `combined_fields` query supports searching multiple text fields as if
	// their contents had been indexed into one combined field.
	CombinedFields *CombinedFieldsQuery        `json:"combined_fields,omitempty"`
	Common         map[string]CommonTermsQuery `json:"common,omitempty"`
	// ConstantScore Wraps a filter query and returns every matching document with a relevance
	// score equal to the `boost` parameter value.
	ConstantScore *ConstantScoreQuery `json:"constant_score,omitempty"`
	// DisMax Returns documents matching one or more wrapped queries, called query clauses
	// or clauses.
	// If a returned document matches multiple query clauses, the `dis_max` query
	// assigns the document the highest relevance score from any matching clause,
	// plus a tie breaking increment for any additional matching subqueries.
	DisMax *DisMaxQuery `json:"dis_max,omitempty"`
	// DistanceFeature Boosts the relevance score of documents closer to a provided origin date or
	// point.
	// For example, you can use this query to give more weight to documents closer
	// to a certain date or location.
	DistanceFeature DistanceFeatureQuery `json:"distance_feature,omitempty"`
	// Exists Returns documents that contain an indexed value for a field.
	Exists *ExistsQuery `json:"exists,omitempty"`
	// FieldMaskingSpan Wrapper to allow span queries to participate in composite single-field span
	// queries by _lying_ about their search field.
	FieldMaskingSpan *SpanFieldMaskingQuery `json:"field_masking_span,omitempty"`
	// FunctionScore The `function_score` enables you to modify the score of documents that are
	// retrieved by a query.
	FunctionScore *FunctionScoreQuery `json:"function_score,omitempty"`
	// Fuzzy Returns documents that contain terms similar to the search term, as measured
	// by a Levenshtein edit distance.
	Fuzzy map[string]FuzzyQuery `json:"fuzzy,omitempty"`
	// GeoBoundingBox Matches geo_point and geo_shape values that intersect a bounding box.
	GeoBoundingBox *GeoBoundingBoxQuery `json:"geo_bounding_box,omitempty"`
	// GeoDistance Matches `geo_point` and `geo_shape` values within a given distance of a
	// geopoint.
	GeoDistance *GeoDistanceQuery `json:"geo_distance,omitempty"`
	GeoPolygon  *GeoPolygonQuery  `json:"geo_polygon,omitempty"`
	// GeoShape Filter documents indexed using either the `geo_shape` or the `geo_point`
	// type.
	GeoShape *GeoShapeQuery `json:"geo_shape,omitempty"`
	// HasChild Returns parent documents whose joined child documents match a provided query.
	HasChild *HasChildQuery `json:"has_child,omitempty"`
	// HasParent Returns child documents whose joined parent document matches a provided
	// query.
	HasParent *HasParentQuery `json:"has_parent,omitempty"`
	// Ids Returns documents based on their IDs.
	// This query uses document IDs stored in the `_id` field.
	Ids *IdsQuery `json:"ids,omitempty"`
	// Intervals Returns documents based on the order and proximity of matching terms.
	Intervals map[string]IntervalsQuery `json:"intervals,omitempty"`
	// Match Returns documents that match a provided text, number, date or boolean value.
	// The provided text is analyzed before matching.
	Match map[string]MatchQuery `json:"match,omitempty"`
	// MatchAll Matches all documents, giving them all a `_score` of 1.0.
	MatchAll *MatchAllQuery `json:"match_all,omitempty"`
	// MatchBoolPrefix Analyzes its input and constructs a `bool` query from the terms.
	// Each term except the last is used in a `term` query.
	// The last term is used in a prefix query.
	MatchBoolPrefix map[string]MatchBoolPrefixQuery `json:"match_bool_prefix,omitempty"`
	// MatchNone Matches no documents.
	MatchNone *MatchNoneQuery `json:"match_none,omitempty"`
	// MatchPhrase Analyzes the text and creates a phrase query out of the analyzed text.
	MatchPhrase map[string]MatchPhraseQuery `json:"match_phrase,omitempty"`
	// MatchPhrasePrefix Returns documents that contain the words of a provided text, in the same
	// order as provided.
	// The last term of the provided text is treated as a prefix, matching any words
	// that begin with that term.
	MatchPhrasePrefix map[string]MatchPhrasePrefixQuery `json:"match_phrase_prefix,omitempty"`
	// MoreLikeThis Returns documents that are "like" a given set of documents.
	MoreLikeThis *MoreLikeThisQuery `json:"more_like_this,omitempty"`
	// MultiMatch Enables you to search for a provided text, number, date or boolean value
	// across multiple fields.
	// The provided text is analyzed before matching.
	MultiMatch *MultiMatchQuery `json:"multi_match,omitempty"`
	// Nested Wraps another query to search nested fields.
	// If an object matches the search, the nested query returns the root parent
	// document.
	Nested *NestedQuery `json:"nested,omitempty"`
	// ParentId Returns child documents joined to a specific parent document.
	ParentId *ParentIdQuery `json:"parent_id,omitempty"`
	// Percolate Matches queries stored in an index.
	Percolate *PercolateQuery `json:"percolate,omitempty"`
	// Pinned Promotes selected documents to rank higher than those matching a given query.
	Pinned *PinnedQuery `json:"pinned,omitempty"`
	// Prefix Returns documents that contain a specific prefix in a provided field.
	Prefix map[string]PrefixQuery `json:"prefix,omitempty"`
	// QueryString Returns documents based on a provided query string, using a parser with a
	// strict syntax.
	QueryString *QueryStringQuery `json:"query_string,omitempty"`
	// Range Returns documents that contain terms within a provided range.
	Range map[string]RangeQuery `json:"range,omitempty"`
	// RankFeature Boosts the relevance score of documents based on the numeric value of a
	// `rank_feature` or `rank_features` field.
	RankFeature *RankFeatureQuery `json:"rank_feature,omitempty"`
	// Regexp Returns documents that contain terms matching a regular expression.
	Regexp    map[string]RegexpQuery `json:"regexp,omitempty"`
	RuleQuery *RuleQuery             `json:"rule_query,omitempty"`
	// Script Filters documents based on a provided script.
	// The script query is typically used in a filter context.
	Script *ScriptQuery `json:"script,omitempty"`
	// ScriptScore Uses a script to provide a custom score for returned documents.
	ScriptScore *ScriptScoreQuery `json:"script_score,omitempty"`
	// Shape Queries documents that contain fields indexed using the `shape` type.
	Shape *ShapeQuery `json:"shape,omitempty"`
	// SimpleQueryString Returns documents based on a provided query string, using a parser with a
	// limited but fault-tolerant syntax.
	SimpleQueryString *SimpleQueryStringQuery `json:"simple_query_string,omitempty"`
	// SpanContaining Returns matches which enclose another span query.
	SpanContaining *SpanContainingQuery `json:"span_containing,omitempty"`
	// SpanFirst Matches spans near the beginning of a field.
	SpanFirst *SpanFirstQuery `json:"span_first,omitempty"`
	// SpanMulti Allows you to wrap a multi term query (one of `wildcard`, `fuzzy`, `prefix`,
	// `range`, or `regexp` query) as a `span` query, so it can be nested.
	SpanMulti *SpanMultiTermQuery `json:"span_multi,omitempty"`
	// SpanNear Matches spans which are near one another.
	// You can specify `slop`, the maximum number of intervening unmatched
	// positions, as well as whether matches are required to be in-order.
	SpanNear *SpanNearQuery `json:"span_near,omitempty"`
	// SpanNot Removes matches which overlap with another span query or which are within x
	// tokens before (controlled by the parameter `pre`) or y tokens after
	// (controlled by the parameter `post`) another span query.
	SpanNot *SpanNotQuery `json:"span_not,omitempty"`
	// SpanOr Matches the union of its span clauses.
	SpanOr *SpanOrQuery `json:"span_or,omitempty"`
	// SpanTerm Matches spans containing a term.
	SpanTerm map[string]SpanTermQuery `json:"span_term,omitempty"`
	// SpanWithin Returns matches which are enclosed inside another span query.
	SpanWithin *SpanWithinQuery `json:"span_within,omitempty"`
	// Term Returns documents that contain an exact term in a provided field.
	// To return a document, the query term must exactly match the queried field's
	// value, including whitespace and capitalization.
	Term map[string]TermQuery `json:"term,omitempty"`
	// Terms Returns documents that contain one or more exact terms in a provided field.
	// To return a document, one or more terms must exactly match a field value,
	// including whitespace and capitalization.
	Terms *TermsQuery `json:"terms,omitempty"`
	// TermsSet Returns documents that contain a minimum number of exact terms in a provided
	// field.
	// To return a document, a required number of terms must exactly match the field
	// values, including whitespace and capitalization.
	TermsSet map[string]TermsSetQuery `json:"terms_set,omitempty"`
	// TextExpansion Uses a natural language processing model to convert the query text into a
	// list of token-weight pairs which are then used in a query against a sparse
	// vector or rank features field.
	TextExpansion map[string]TextExpansionQuery `json:"text_expansion,omitempty"`
	Type          *TypeQuery                    `json:"type,omitempty"`
	// Wildcard Returns documents that contain terms matching a wildcard pattern.
	Wildcard map[string]WildcardQuery `json:"wildcard,omitempty"`
	// Wrapper A query that accepts any other query as base64 encoded string.
	Wrapper *WrapperQuery `json:"wrapper,omitempty"`
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

		case "rule_query":
			if err := dec.Decode(&s.RuleQuery); err != nil {
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
			if s.TextExpansion == nil {
				s.TextExpansion = make(map[string]TextExpansionQuery, 0)
			}
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
		TextExpansion:     make(map[string]TextExpansionQuery, 0),
		Wildcard:          make(map[string]WildcardQuery, 0),
	}

	return r
}
