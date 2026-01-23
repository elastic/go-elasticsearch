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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
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

func (s *_query) Bool(bool types.BoolQueryVariant) *_query {

	s.v.Bool = bool.BoolQueryCaster()

	return s
}

func (s *_query) Boosting(boosting types.BoostingQueryVariant) *_query {

	s.v.Boosting = boosting.BoostingQueryCaster()

	return s
}

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

func (s *_query) ConstantScore(constantscore types.ConstantScoreQueryVariant) *_query {

	s.v.ConstantScore = constantscore.ConstantScoreQueryCaster()

	return s
}

func (s *_query) DisMax(dismax types.DisMaxQueryVariant) *_query {

	s.v.DisMax = dismax.DisMaxQueryCaster()

	return s
}

func (s *_query) DistanceFeature(distancefeaturequery types.DistanceFeatureQueryVariant) *_query {

	s.v.DistanceFeature = *distancefeaturequery.DistanceFeatureQueryCaster()

	return s
}

func (s *_query) Exists(exists types.ExistsQueryVariant) *_query {

	s.v.Exists = exists.ExistsQueryCaster()

	return s
}

func (s *_query) FunctionScore(functionscore types.FunctionScoreQueryVariant) *_query {

	s.v.FunctionScore = functionscore.FunctionScoreQueryCaster()

	return s
}

// Fuzzy is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Fuzzy(key string, value types.FuzzyQueryVariant) *_query {

	tmp := make(map[string]types.FuzzyQuery)

	tmp[key] = *value.FuzzyQueryCaster()

	s.v.Fuzzy = tmp
	return s
}

func (s *_query) GeoBoundingBox(geoboundingbox types.GeoBoundingBoxQueryVariant) *_query {

	s.v.GeoBoundingBox = geoboundingbox.GeoBoundingBoxQueryCaster()

	return s
}

func (s *_query) GeoDistance(geodistance types.GeoDistanceQueryVariant) *_query {

	s.v.GeoDistance = geodistance.GeoDistanceQueryCaster()

	return s
}

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

func (s *_query) GeoShape(geoshape types.GeoShapeQueryVariant) *_query {

	s.v.GeoShape = geoshape.GeoShapeQueryCaster()

	return s
}

func (s *_query) HasChild(haschild types.HasChildQueryVariant) *_query {

	s.v.HasChild = haschild.HasChildQueryCaster()

	return s
}

func (s *_query) HasParent(hasparent types.HasParentQueryVariant) *_query {

	s.v.HasParent = hasparent.HasParentQueryCaster()

	return s
}

func (s *_query) Ids(ids types.IdsQueryVariant) *_query {

	s.v.Ids = ids.IdsQueryCaster()

	return s
}

// Interval is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Interval(key string, value types.IntervalsQueryVariant) *_query {

	tmp := make(map[string]types.IntervalsQuery)

	tmp[key] = *value.IntervalsQueryCaster()

	s.v.Intervals = tmp
	return s
}

func (s *_query) Knn(knn types.KnnQueryVariant) *_query {

	s.v.Knn = knn.KnnQueryCaster()

	return s
}

// Match is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Match(key string, value types.MatchQueryVariant) *_query {

	tmp := make(map[string]types.MatchQuery)

	tmp[key] = *value.MatchQueryCaster()

	s.v.Match = tmp
	return s
}

func (s *_query) MatchAll(matchall types.MatchAllQueryVariant) *_query {

	s.v.MatchAll = matchall.MatchAllQueryCaster()

	return s
}

// MatchBoolPrefix is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) MatchBoolPrefix(key string, value types.MatchBoolPrefixQueryVariant) *_query {

	tmp := make(map[string]types.MatchBoolPrefixQuery)

	tmp[key] = *value.MatchBoolPrefixQueryCaster()

	s.v.MatchBoolPrefix = tmp
	return s
}

func (s *_query) MatchNone(matchnone types.MatchNoneQueryVariant) *_query {

	s.v.MatchNone = matchnone.MatchNoneQueryCaster()

	return s
}

// MatchPhrase is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) MatchPhrase(key string, value types.MatchPhraseQueryVariant) *_query {

	tmp := make(map[string]types.MatchPhraseQuery)

	tmp[key] = *value.MatchPhraseQueryCaster()

	s.v.MatchPhrase = tmp
	return s
}

// MatchPhrasePrefix is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) MatchPhrasePrefix(key string, value types.MatchPhrasePrefixQueryVariant) *_query {

	tmp := make(map[string]types.MatchPhrasePrefixQuery)

	tmp[key] = *value.MatchPhrasePrefixQueryCaster()

	s.v.MatchPhrasePrefix = tmp
	return s
}

func (s *_query) MoreLikeThis(morelikethis types.MoreLikeThisQueryVariant) *_query {

	s.v.MoreLikeThis = morelikethis.MoreLikeThisQueryCaster()

	return s
}

func (s *_query) MultiMatch(multimatch types.MultiMatchQueryVariant) *_query {

	s.v.MultiMatch = multimatch.MultiMatchQueryCaster()

	return s
}

func (s *_query) Nested(nested types.NestedQueryVariant) *_query {

	s.v.Nested = nested.NestedQueryCaster()

	return s
}

func (s *_query) ParentId(parentid types.ParentIdQueryVariant) *_query {

	s.v.ParentId = parentid.ParentIdQueryCaster()

	return s
}

func (s *_query) Percolate(percolate types.PercolateQueryVariant) *_query {

	s.v.Percolate = percolate.PercolateQueryCaster()

	return s
}

func (s *_query) Pinned(pinned types.PinnedQueryVariant) *_query {

	s.v.Pinned = pinned.PinnedQueryCaster()

	return s
}

// Prefix is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Prefix(key string, value types.PrefixQueryVariant) *_query {

	tmp := make(map[string]types.PrefixQuery)

	tmp[key] = *value.PrefixQueryCaster()

	s.v.Prefix = tmp
	return s
}

func (s *_query) QueryString(querystring types.QueryStringQueryVariant) *_query {

	s.v.QueryString = querystring.QueryStringQueryCaster()

	return s
}

// Range is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Range(key string, value types.RangeQueryVariant) *_query {

	tmp := make(map[string]types.RangeQuery)

	tmp[key] = *value.RangeQueryCaster()

	s.v.Range = tmp
	return s
}

func (s *_query) RankFeature(rankfeature types.RankFeatureQueryVariant) *_query {

	s.v.RankFeature = rankfeature.RankFeatureQueryCaster()

	return s
}

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

func (s *_query) Script(script types.ScriptQueryVariant) *_query {

	s.v.Script = script.ScriptQueryCaster()

	return s
}

func (s *_query) ScriptScore(scriptscore types.ScriptScoreQueryVariant) *_query {

	s.v.ScriptScore = scriptscore.ScriptScoreQueryCaster()

	return s
}

func (s *_query) Semantic(semantic types.SemanticQueryVariant) *_query {

	s.v.Semantic = semantic.SemanticQueryCaster()

	return s
}

func (s *_query) Shape(shape types.ShapeQueryVariant) *_query {

	s.v.Shape = shape.ShapeQueryCaster()

	return s
}

func (s *_query) SimpleQueryString(simplequerystring types.SimpleQueryStringQueryVariant) *_query {

	s.v.SimpleQueryString = simplequerystring.SimpleQueryStringQueryCaster()

	return s
}

func (s *_query) SpanContaining(spancontaining types.SpanContainingQueryVariant) *_query {

	s.v.SpanContaining = spancontaining.SpanContainingQueryCaster()

	return s
}

func (s *_query) SpanFieldMasking(spanfieldmasking types.SpanFieldMaskingQueryVariant) *_query {

	s.v.SpanFieldMasking = spanfieldmasking.SpanFieldMaskingQueryCaster()

	return s
}

func (s *_query) SpanFirst(spanfirst types.SpanFirstQueryVariant) *_query {

	s.v.SpanFirst = spanfirst.SpanFirstQueryCaster()

	return s
}

func (s *_query) SpanMulti(spanmulti types.SpanMultiTermQueryVariant) *_query {

	s.v.SpanMulti = spanmulti.SpanMultiTermQueryCaster()

	return s
}

func (s *_query) SpanNear(spannear types.SpanNearQueryVariant) *_query {

	s.v.SpanNear = spannear.SpanNearQueryCaster()

	return s
}

func (s *_query) SpanNot(spannot types.SpanNotQueryVariant) *_query {

	s.v.SpanNot = spannot.SpanNotQueryCaster()

	return s
}

func (s *_query) SpanOr(spanor types.SpanOrQueryVariant) *_query {

	s.v.SpanOr = spanor.SpanOrQueryCaster()

	return s
}

// SpanTerm is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) SpanTerm(key string, value types.SpanTermQueryVariant) *_query {

	tmp := make(map[string]types.SpanTermQuery)

	tmp[key] = *value.SpanTermQueryCaster()

	s.v.SpanTerm = tmp
	return s
}

func (s *_query) SpanWithin(spanwithin types.SpanWithinQueryVariant) *_query {

	s.v.SpanWithin = spanwithin.SpanWithinQueryCaster()

	return s
}

func (s *_query) SparseVector(sparsevector types.SparseVectorQueryVariant) *_query {

	s.v.SparseVector = sparsevector.SparseVectorQueryCaster()

	return s
}

// Term is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Term(key string, value types.TermQueryVariant) *_query {

	tmp := make(map[string]types.TermQuery)

	tmp[key] = *value.TermQueryCaster()

	s.v.Term = tmp
	return s
}

func (s *_query) Terms(terms types.TermsQueryVariant) *_query {

	s.v.Terms = terms.TermsQueryCaster()

	return s
}

// TermsSet is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) TermsSet(key string, value types.TermsSetQueryVariant) *_query {

	tmp := make(map[string]types.TermsSetQuery)

	tmp[key] = *value.TermsSetQueryCaster()

	s.v.TermsSet = tmp
	return s
}

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

// WeightedToken is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) WeightedToken(key string, value types.WeightedTokensQueryVariant) *_query {

	tmp := make(map[string]types.WeightedTokensQuery)

	tmp[key] = *value.WeightedTokensQueryCaster()

	s.v.WeightedTokens = tmp
	return s
}

// Wildcard is a single key dictionnary.
// It will replace the current value on each call.
func (s *_query) Wildcard(key string, value types.WildcardQueryVariant) *_query {

	tmp := make(map[string]types.WildcardQuery)

	tmp[key] = *value.WildcardQueryCaster()

	s.v.Wildcard = tmp
	return s
}

func (s *_query) Wrapper(wrapper types.WrapperQueryVariant) *_query {

	s.v.Wrapper = wrapper.WrapperQueryCaster()

	return s
}

func (s *_query) QueryCaster() *types.Query {
	return s.v
}
