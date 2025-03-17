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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _query struct {
	v *types.Query
}

func NewQuery() *_query {
	return &_query{v: types.NewQuery()}
}

// AdditionalQueryProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) AdditionalQueryProperty(key string, value json.RawMessage) *_query {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalQueryProperty = tmp
	return s
}

// matches documents matching boolean combinations of other queries.
func (s *_query) Bool(bool types.BoolQueryVariant) *_query {

	s.v.Bool = bool.BoolQueryCaster()

	return s
}

// Returns documents matching a `positive` query while reducing the relevance
// score of documents that also match a `negative` query.
func (s *_query) Boosting(boosting types.BoostingQueryVariant) *_query {

	s.v.Boosting = boosting.BoostingQueryCaster()

	return s
}

// The `combined_fields` query supports searching multiple text fields as if
// their contents had been indexed into one combined field.
func (s *_query) CombinedFields(combinedfields types.CombinedFieldsQueryVariant) *_query {

	s.v.CombinedFields = combinedfields.CombinedFieldsQueryCaster()

	return s
}

// Common is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Common(key string, value types.CommonTermsQueryVariant) *_query {

	tmp := make(map[string]types.CommonTermsQuery)

	tmp[key] = *value.CommonTermsQueryCaster()

	s.v.Common = tmp
	return s
}

// Wraps a filter query and returns every matching document with a relevance
// score equal to the `boost` parameter value.
func (s *_query) ConstantScore(constantscore types.ConstantScoreQueryVariant) *_query {

	s.v.ConstantScore = constantscore.ConstantScoreQueryCaster()

	return s
}

// Returns documents matching one or more wrapped queries, called query clauses
// or clauses.
// If a returned document matches multiple query clauses, the `dis_max` query
// assigns the document the highest relevance score from any matching clause,
// plus a tie breaking increment for any additional matching subqueries.
func (s *_query) DisMax(dismax types.DisMaxQueryVariant) *_query {

	s.v.DisMax = dismax.DisMaxQueryCaster()

	return s
}

// Boosts the relevance score of documents closer to a provided origin date or
// point.
// For example, you can use this query to give more weight to documents closer
// to a certain date or location.
func (s *_query) DistanceFeature(distancefeaturequery types.DistanceFeatureQueryVariant) *_query {

	s.v.DistanceFeature = *distancefeaturequery.DistanceFeatureQueryCaster()

	return s
}

// Returns documents that contain an indexed value for a field.
func (s *_query) Exists(exists types.ExistsQueryVariant) *_query {

	s.v.Exists = exists.ExistsQueryCaster()

	return s
}

// The `function_score` enables you to modify the score of documents that are
// retrieved by a query.
func (s *_query) FunctionScore(functionscore types.FunctionScoreQueryVariant) *_query {

	s.v.FunctionScore = functionscore.FunctionScoreQueryCaster()

	return s
}

// Returns documents that contain terms similar to the search term, as measured
// by a Levenshtein edit distance.
// Fuzzy is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Fuzzy(key string, value types.FuzzyQueryVariant) *_query {

	tmp := make(map[string]types.FuzzyQuery)

	tmp[key] = *value.FuzzyQueryCaster()

	s.v.Fuzzy = tmp
	return s
}

// Matches geo_point and geo_shape values that intersect a bounding box.
func (s *_query) GeoBoundingBox(geoboundingbox types.GeoBoundingBoxQueryVariant) *_query {

	s.v.GeoBoundingBox = geoboundingbox.GeoBoundingBoxQueryCaster()

	return s
}

// Matches `geo_point` and `geo_shape` values within a given distance of a
// geopoint.
func (s *_query) GeoDistance(geodistance types.GeoDistanceQueryVariant) *_query {

	s.v.GeoDistance = geodistance.GeoDistanceQueryCaster()

	return s
}

// Matches `geo_point` and `geo_shape` values that intersect a grid cell from a
// GeoGrid aggregation.
// GeoGrid is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) GeoGrid(key string, value types.GeoGridQueryVariant) *_query {

	tmp := make(map[string]types.GeoGridQuery)

	tmp[key] = *value.GeoGridQueryCaster()

	s.v.GeoGrid = tmp
	return s
}

func (s *_query) GeoPolygon(geopolygon types.GeoPolygonQueryVariant) *_query {

	s.v.GeoPolygon = geopolygon.GeoPolygonQueryCaster()

	return s
}

// Filter documents indexed using either the `geo_shape` or the `geo_point`
// type.
func (s *_query) GeoShape(geoshape types.GeoShapeQueryVariant) *_query {

	s.v.GeoShape = geoshape.GeoShapeQueryCaster()

	return s
}

// Returns parent documents whose joined child documents match a provided query.
func (s *_query) HasChild(haschild types.HasChildQueryVariant) *_query {

	s.v.HasChild = haschild.HasChildQueryCaster()

	return s
}

// Returns child documents whose joined parent document matches a provided
// query.
func (s *_query) HasParent(hasparent types.HasParentQueryVariant) *_query {

	s.v.HasParent = hasparent.HasParentQueryCaster()

	return s
}

// Returns documents based on their IDs.
// This query uses document IDs stored in the `_id` field.
func (s *_query) Ids(ids types.IdsQueryVariant) *_query {

	s.v.Ids = ids.IdsQueryCaster()

	return s
}

// Returns documents based on the order and proximity of matching terms.
// Interval is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Interval(key string, value types.IntervalsQueryVariant) *_query {

	tmp := make(map[string]types.IntervalsQuery)

	tmp[key] = *value.IntervalsQueryCaster()

	s.v.Intervals = tmp
	return s
}

// Finds the k nearest vectors to a query vector, as measured by a similarity
// metric. knn query finds nearest vectors through approximate search on indexed
// dense_vectors.
func (s *_query) Knn(knn types.KnnQueryVariant) *_query {

	s.v.Knn = knn.KnnQueryCaster()

	return s
}

// Returns documents that match a provided text, number, date or boolean value.
// The provided text is analyzed before matching.
// Match is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Match(key string, value types.MatchQueryVariant) *_query {

	tmp := make(map[string]types.MatchQuery)

	tmp[key] = *value.MatchQueryCaster()

	s.v.Match = tmp
	return s
}

// Matches all documents, giving them all a `_score` of 1.0.
func (s *_query) MatchAll(matchall types.MatchAllQueryVariant) *_query {

	s.v.MatchAll = matchall.MatchAllQueryCaster()

	return s
}

// Analyzes its input and constructs a `bool` query from the terms.
// Each term except the last is used in a `term` query.
// The last term is used in a prefix query.
// MatchBoolPrefix is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) MatchBoolPrefix(key string, value types.MatchBoolPrefixQueryVariant) *_query {

	tmp := make(map[string]types.MatchBoolPrefixQuery)

	tmp[key] = *value.MatchBoolPrefixQueryCaster()

	s.v.MatchBoolPrefix = tmp
	return s
}

// Matches no documents.
func (s *_query) MatchNone(matchnone types.MatchNoneQueryVariant) *_query {

	s.v.MatchNone = matchnone.MatchNoneQueryCaster()

	return s
}

// Analyzes the text and creates a phrase query out of the analyzed text.
// MatchPhrase is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) MatchPhrase(key string, value types.MatchPhraseQueryVariant) *_query {

	tmp := make(map[string]types.MatchPhraseQuery)

	tmp[key] = *value.MatchPhraseQueryCaster()

	s.v.MatchPhrase = tmp
	return s
}

// Returns documents that contain the words of a provided text, in the same
// order as provided.
// The last term of the provided text is treated as a prefix, matching any words
// that begin with that term.
// MatchPhrasePrefix is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) MatchPhrasePrefix(key string, value types.MatchPhrasePrefixQueryVariant) *_query {

	tmp := make(map[string]types.MatchPhrasePrefixQuery)

	tmp[key] = *value.MatchPhrasePrefixQueryCaster()

	s.v.MatchPhrasePrefix = tmp
	return s
}

// Returns documents that are "like" a given set of documents.
func (s *_query) MoreLikeThis(morelikethis types.MoreLikeThisQueryVariant) *_query {

	s.v.MoreLikeThis = morelikethis.MoreLikeThisQueryCaster()

	return s
}

// Enables you to search for a provided text, number, date or boolean value
// across multiple fields.
// The provided text is analyzed before matching.
func (s *_query) MultiMatch(multimatch types.MultiMatchQueryVariant) *_query {

	s.v.MultiMatch = multimatch.MultiMatchQueryCaster()

	return s
}

// Wraps another query to search nested fields.
// If an object matches the search, the nested query returns the root parent
// document.
func (s *_query) Nested(nested types.NestedQueryVariant) *_query {

	s.v.Nested = nested.NestedQueryCaster()

	return s
}

// Returns child documents joined to a specific parent document.
func (s *_query) ParentId(parentid types.ParentIdQueryVariant) *_query {

	s.v.ParentId = parentid.ParentIdQueryCaster()

	return s
}

// Matches queries stored in an index.
func (s *_query) Percolate(percolate types.PercolateQueryVariant) *_query {

	s.v.Percolate = percolate.PercolateQueryCaster()

	return s
}

// Promotes selected documents to rank higher than those matching a given query.
func (s *_query) Pinned(pinned types.PinnedQueryVariant) *_query {

	s.v.Pinned = pinned.PinnedQueryCaster()

	return s
}

// Returns documents that contain a specific prefix in a provided field.
// Prefix is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Prefix(key string, value types.PrefixQueryVariant) *_query {

	tmp := make(map[string]types.PrefixQuery)

	tmp[key] = *value.PrefixQueryCaster()

	s.v.Prefix = tmp
	return s
}

// Returns documents based on a provided query string, using a parser with a
// strict syntax.
func (s *_query) QueryString(querystring types.QueryStringQueryVariant) *_query {

	s.v.QueryString = querystring.QueryStringQueryCaster()

	return s
}

// Returns documents that contain terms within a provided range.
// Range is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Range(key string, value types.RangeQueryVariant) *_query {

	tmp := make(map[string]types.RangeQuery)

	tmp[key] = *value.RangeQueryCaster()

	s.v.Range = tmp
	return s
}

// Boosts the relevance score of documents based on the numeric value of a
// `rank_feature` or `rank_features` field.
func (s *_query) RankFeature(rankfeature types.RankFeatureQueryVariant) *_query {

	s.v.RankFeature = rankfeature.RankFeatureQueryCaster()

	return s
}

// Returns documents that contain terms matching a regular expression.
// Regexp is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Regexp(key string, value types.RegexpQueryVariant) *_query {

	tmp := make(map[string]types.RegexpQuery)

	tmp[key] = *value.RegexpQueryCaster()

	s.v.Regexp = tmp
	return s
}

func (s *_query) Rule(rule types.RuleQueryVariant) *_query {

	s.v.Rule = rule.RuleQueryCaster()

	return s
}

// Filters documents based on a provided script.
// The script query is typically used in a filter context.
func (s *_query) Script(script types.ScriptQueryVariant) *_query {

	s.v.Script = script.ScriptQueryCaster()

	return s
}

// Uses a script to provide a custom score for returned documents.
func (s *_query) ScriptScore(scriptscore types.ScriptScoreQueryVariant) *_query {

	s.v.ScriptScore = scriptscore.ScriptScoreQueryCaster()

	return s
}

// A semantic query to semantic_text field types
func (s *_query) Semantic(semantic types.SemanticQueryVariant) *_query {

	s.v.Semantic = semantic.SemanticQueryCaster()

	return s
}

// Queries documents that contain fields indexed using the `shape` type.
func (s *_query) Shape(shape types.ShapeQueryVariant) *_query {

	s.v.Shape = shape.ShapeQueryCaster()

	return s
}

// Returns documents based on a provided query string, using a parser with a
// limited but fault-tolerant syntax.
func (s *_query) SimpleQueryString(simplequerystring types.SimpleQueryStringQueryVariant) *_query {

	s.v.SimpleQueryString = simplequerystring.SimpleQueryStringQueryCaster()

	return s
}

// Returns matches which enclose another span query.
func (s *_query) SpanContaining(spancontaining types.SpanContainingQueryVariant) *_query {

	s.v.SpanContaining = spancontaining.SpanContainingQueryCaster()

	return s
}

// Wrapper to allow span queries to participate in composite single-field span
// queries by _lying_ about their search field.
func (s *_query) SpanFieldMasking(spanfieldmasking types.SpanFieldMaskingQueryVariant) *_query {

	s.v.SpanFieldMasking = spanfieldmasking.SpanFieldMaskingQueryCaster()

	return s
}

// Matches spans near the beginning of a field.
func (s *_query) SpanFirst(spanfirst types.SpanFirstQueryVariant) *_query {

	s.v.SpanFirst = spanfirst.SpanFirstQueryCaster()

	return s
}

// Allows you to wrap a multi term query (one of `wildcard`, `fuzzy`, `prefix`,
// `range`, or `regexp` query) as a `span` query, so it can be nested.
func (s *_query) SpanMulti(spanmulti types.SpanMultiTermQueryVariant) *_query {

	s.v.SpanMulti = spanmulti.SpanMultiTermQueryCaster()

	return s
}

// Matches spans which are near one another.
// You can specify `slop`, the maximum number of intervening unmatched
// positions, as well as whether matches are required to be in-order.
func (s *_query) SpanNear(spannear types.SpanNearQueryVariant) *_query {

	s.v.SpanNear = spannear.SpanNearQueryCaster()

	return s
}

// Removes matches which overlap with another span query or which are within x
// tokens before (controlled by the parameter `pre`) or y tokens after
// (controlled by the parameter `post`) another span query.
func (s *_query) SpanNot(spannot types.SpanNotQueryVariant) *_query {

	s.v.SpanNot = spannot.SpanNotQueryCaster()

	return s
}

// Matches the union of its span clauses.
func (s *_query) SpanOr(spanor types.SpanOrQueryVariant) *_query {

	s.v.SpanOr = spanor.SpanOrQueryCaster()

	return s
}

// Matches spans containing a term.
// SpanTerm is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) SpanTerm(key string, value types.SpanTermQueryVariant) *_query {

	tmp := make(map[string]types.SpanTermQuery)

	tmp[key] = *value.SpanTermQueryCaster()

	s.v.SpanTerm = tmp
	return s
}

// Returns matches which are enclosed inside another span query.
func (s *_query) SpanWithin(spanwithin types.SpanWithinQueryVariant) *_query {

	s.v.SpanWithin = spanwithin.SpanWithinQueryCaster()

	return s
}

// Using input query vectors or a natural language processing model to convert a
// query into a list of token-weight pairs, queries against a sparse vector
// field.
func (s *_query) SparseVector(sparsevector types.SparseVectorQueryVariant) *_query {

	s.v.SparseVector = sparsevector.SparseVectorQueryCaster()

	return s
}

// Returns documents that contain an exact term in a provided field.
// To return a document, the query term must exactly match the queried field's
// value, including whitespace and capitalization.
// Term is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Term(key string, value types.TermQueryVariant) *_query {

	tmp := make(map[string]types.TermQuery)

	tmp[key] = *value.TermQueryCaster()

	s.v.Term = tmp
	return s
}

// Returns documents that contain one or more exact terms in a provided field.
// To return a document, one or more terms must exactly match a field value,
// including whitespace and capitalization.
func (s *_query) Terms(terms types.TermsQueryVariant) *_query {

	s.v.Terms = terms.TermsQueryCaster()

	return s
}

// Returns documents that contain a minimum number of exact terms in a provided
// field.
// To return a document, a required number of terms must exactly match the field
// values, including whitespace and capitalization.
// TermsSet is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) TermsSet(key string, value types.TermsSetQueryVariant) *_query {

	tmp := make(map[string]types.TermsSetQuery)

	tmp[key] = *value.TermsSetQueryCaster()

	s.v.TermsSet = tmp
	return s
}

// Uses a natural language processing model to convert the query text into a
// list of token-weight pairs which are then used in a query against a sparse
// vector or rank features field.
// TextExpansion is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) TextExpansion(key string, value types.TextExpansionQueryVariant) *_query {

	tmp := make(map[string]types.TextExpansionQuery)

	tmp[key] = *value.TextExpansionQueryCaster()

	s.v.TextExpansion = tmp
	return s
}

func (s *_query) Type(type_ types.TypeQueryVariant) *_query {

	s.v.Type = type_.TypeQueryCaster()

	return s
}

// Supports returning text_expansion query results by sending in precomputed
// tokens with the query.
// WeightedToken is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) WeightedToken(key string, value types.WeightedTokensQueryVariant) *_query {

	tmp := make(map[string]types.WeightedTokensQuery)

	tmp[key] = *value.WeightedTokensQueryCaster()

	s.v.WeightedTokens = tmp
	return s
}

// Returns documents that contain terms matching a wildcard pattern.
// Wildcard is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Wildcard(key string, value types.WildcardQueryVariant) *_query {

	tmp := make(map[string]types.WildcardQuery)

	tmp[key] = *value.WildcardQueryCaster()

	s.v.Wildcard = tmp
	return s
}

// A query that accepts any other query as base64 encoded string.
func (s *_query) Wrapper(wrapper types.WrapperQueryVariant) *_query {

	s.v.Wrapper = wrapper.WrapperQueryCaster()

	return s
}

func (s *_query) QueryCaster() *types.Query {
	return s.v
}
