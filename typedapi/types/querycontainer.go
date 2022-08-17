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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// QueryContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/abstractions.ts#L96-L162
type QueryContainer struct {
	Bool              *BoolQuery                       `json:"bool,omitempty"`
	Boosting          *BoostingQuery                   `json:"boosting,omitempty"`
	CombinedFields    *CombinedFieldsQuery             `json:"combined_fields,omitempty"`
	Common            map[Field]CommonTermsQuery       `json:"common,omitempty"`
	ConstantScore     *ConstantScoreQuery              `json:"constant_score,omitempty"`
	DisMax            *DisMaxQuery                     `json:"dis_max,omitempty"`
	DistanceFeature   *DistanceFeatureQuery            `json:"distance_feature,omitempty"`
	Exists            *ExistsQuery                     `json:"exists,omitempty"`
	FieldMaskingSpan  *SpanFieldMaskingQuery           `json:"field_masking_span,omitempty"`
	FunctionScore     *FunctionScoreQuery              `json:"function_score,omitempty"`
	Fuzzy             map[Field]FuzzyQuery             `json:"fuzzy,omitempty"`
	GeoBoundingBox    *GeoBoundingBoxQuery             `json:"geo_bounding_box,omitempty"`
	GeoDistance       *GeoDistanceQuery                `json:"geo_distance,omitempty"`
	GeoPolygon        *GeoPolygonQuery                 `json:"geo_polygon,omitempty"`
	GeoShape          *GeoShapeQuery                   `json:"geo_shape,omitempty"`
	HasChild          *HasChildQuery                   `json:"has_child,omitempty"`
	HasParent         *HasParentQuery                  `json:"has_parent,omitempty"`
	Ids               *IdsQuery                        `json:"ids,omitempty"`
	Intervals         map[Field]IntervalsQuery         `json:"intervals,omitempty"`
	Match             map[Field]MatchQuery             `json:"match,omitempty"`
	MatchAll          *MatchAllQuery                   `json:"match_all,omitempty"`
	MatchBoolPrefix   map[Field]MatchBoolPrefixQuery   `json:"match_bool_prefix,omitempty"`
	MatchNone         *MatchNoneQuery                  `json:"match_none,omitempty"`
	MatchPhrase       map[Field]MatchPhraseQuery       `json:"match_phrase,omitempty"`
	MatchPhrasePrefix map[Field]MatchPhrasePrefixQuery `json:"match_phrase_prefix,omitempty"`
	MoreLikeThis      *MoreLikeThisQuery               `json:"more_like_this,omitempty"`
	MultiMatch        *MultiMatchQuery                 `json:"multi_match,omitempty"`
	Nested            *NestedQuery                     `json:"nested,omitempty"`
	ParentId          *ParentIdQuery                   `json:"parent_id,omitempty"`
	Percolate         *PercolateQuery                  `json:"percolate,omitempty"`
	Pinned            *PinnedQuery                     `json:"pinned,omitempty"`
	Prefix            map[Field]PrefixQuery            `json:"prefix,omitempty"`
	QueryString       *QueryStringQuery                `json:"query_string,omitempty"`
	Range             map[Field]RangeQuery             `json:"range,omitempty"`
	RankFeature       *RankFeatureQuery                `json:"rank_feature,omitempty"`
	Regexp            map[Field]RegexpQuery            `json:"regexp,omitempty"`
	Script            *ScriptQuery                     `json:"script,omitempty"`
	ScriptScore       *ScriptScoreQuery                `json:"script_score,omitempty"`
	Shape             *ShapeQuery                      `json:"shape,omitempty"`
	SimpleQueryString *SimpleQueryStringQuery          `json:"simple_query_string,omitempty"`
	SpanContaining    *SpanContainingQuery             `json:"span_containing,omitempty"`
	SpanFirst         *SpanFirstQuery                  `json:"span_first,omitempty"`
	SpanMulti         *SpanMultiTermQuery              `json:"span_multi,omitempty"`
	SpanNear          *SpanNearQuery                   `json:"span_near,omitempty"`
	SpanNot           *SpanNotQuery                    `json:"span_not,omitempty"`
	SpanOr            *SpanOrQuery                     `json:"span_or,omitempty"`
	SpanTerm          map[Field]SpanTermQuery          `json:"span_term,omitempty"`
	SpanWithin        *SpanWithinQuery                 `json:"span_within,omitempty"`
	Term              map[Field]TermQuery              `json:"term,omitempty"`
	Terms             *TermsQuery                      `json:"terms,omitempty"`
	TermsSet          map[Field]TermsSetQuery          `json:"terms_set,omitempty"`
	Type              *TypeQuery                       `json:"type,omitempty"`
	Wildcard          map[Field]WildcardQuery          `json:"wildcard,omitempty"`
	Wrapper           *WrapperQuery                    `json:"wrapper,omitempty"`
}

// QueryContainerBuilder holds QueryContainer struct and provides a builder API.
type QueryContainerBuilder struct {
	v *QueryContainer
}

// NewQueryContainer provides a builder for the QueryContainer struct.
func NewQueryContainerBuilder() *QueryContainerBuilder {
	r := QueryContainerBuilder{
		&QueryContainer{
			Common:            make(map[Field]CommonTermsQuery, 0),
			Fuzzy:             make(map[Field]FuzzyQuery, 0),
			Intervals:         make(map[Field]IntervalsQuery, 0),
			Match:             make(map[Field]MatchQuery, 0),
			MatchBoolPrefix:   make(map[Field]MatchBoolPrefixQuery, 0),
			MatchPhrase:       make(map[Field]MatchPhraseQuery, 0),
			MatchPhrasePrefix: make(map[Field]MatchPhrasePrefixQuery, 0),
			Prefix:            make(map[Field]PrefixQuery, 0),
			Range:             make(map[Field]RangeQuery, 0),
			Regexp:            make(map[Field]RegexpQuery, 0),
			SpanTerm:          make(map[Field]SpanTermQuery, 0),
			Term:              make(map[Field]TermQuery, 0),
			TermsSet:          make(map[Field]TermsSetQuery, 0),
			Wildcard:          make(map[Field]WildcardQuery, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the QueryContainer struct
func (rb *QueryContainerBuilder) Build() QueryContainer {
	return *rb.v
}

func (rb *QueryContainerBuilder) Bool(bool *BoolQueryBuilder) *QueryContainerBuilder {
	v := bool.Build()
	rb.v.Bool = &v
	return rb
}

func (rb *QueryContainerBuilder) Boosting(boosting *BoostingQueryBuilder) *QueryContainerBuilder {
	v := boosting.Build()
	rb.v.Boosting = &v
	return rb
}

func (rb *QueryContainerBuilder) CombinedFields(combinedfields *CombinedFieldsQueryBuilder) *QueryContainerBuilder {
	v := combinedfields.Build()
	rb.v.CombinedFields = &v
	return rb
}

func (rb *QueryContainerBuilder) Common(values map[Field]*CommonTermsQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]CommonTermsQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Common = tmp
	return rb
}

func (rb *QueryContainerBuilder) ConstantScore(constantscore *ConstantScoreQueryBuilder) *QueryContainerBuilder {
	v := constantscore.Build()
	rb.v.ConstantScore = &v
	return rb
}

func (rb *QueryContainerBuilder) DisMax(dismax *DisMaxQueryBuilder) *QueryContainerBuilder {
	v := dismax.Build()
	rb.v.DisMax = &v
	return rb
}

func (rb *QueryContainerBuilder) DistanceFeature(distancefeature *DistanceFeatureQueryBuilder) *QueryContainerBuilder {
	v := distancefeature.Build()
	rb.v.DistanceFeature = &v
	return rb
}

func (rb *QueryContainerBuilder) Exists(exists *ExistsQueryBuilder) *QueryContainerBuilder {
	v := exists.Build()
	rb.v.Exists = &v
	return rb
}

func (rb *QueryContainerBuilder) FieldMaskingSpan(fieldmaskingspan *SpanFieldMaskingQueryBuilder) *QueryContainerBuilder {
	v := fieldmaskingspan.Build()
	rb.v.FieldMaskingSpan = &v
	return rb
}

func (rb *QueryContainerBuilder) FunctionScore(functionscore *FunctionScoreQueryBuilder) *QueryContainerBuilder {
	v := functionscore.Build()
	rb.v.FunctionScore = &v
	return rb
}

func (rb *QueryContainerBuilder) Fuzzy(values map[Field]*FuzzyQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]FuzzyQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fuzzy = tmp
	return rb
}

func (rb *QueryContainerBuilder) GeoBoundingBox(geoboundingbox *GeoBoundingBoxQueryBuilder) *QueryContainerBuilder {
	v := geoboundingbox.Build()
	rb.v.GeoBoundingBox = &v
	return rb
}

func (rb *QueryContainerBuilder) GeoDistance(geodistance *GeoDistanceQueryBuilder) *QueryContainerBuilder {
	v := geodistance.Build()
	rb.v.GeoDistance = &v
	return rb
}

func (rb *QueryContainerBuilder) GeoPolygon(geopolygon *GeoPolygonQueryBuilder) *QueryContainerBuilder {
	v := geopolygon.Build()
	rb.v.GeoPolygon = &v
	return rb
}

func (rb *QueryContainerBuilder) GeoShape(geoshape *GeoShapeQueryBuilder) *QueryContainerBuilder {
	v := geoshape.Build()
	rb.v.GeoShape = &v
	return rb
}

func (rb *QueryContainerBuilder) HasChild(haschild *HasChildQueryBuilder) *QueryContainerBuilder {
	v := haschild.Build()
	rb.v.HasChild = &v
	return rb
}

func (rb *QueryContainerBuilder) HasParent(hasparent *HasParentQueryBuilder) *QueryContainerBuilder {
	v := hasparent.Build()
	rb.v.HasParent = &v
	return rb
}

func (rb *QueryContainerBuilder) Ids(ids *IdsQueryBuilder) *QueryContainerBuilder {
	v := ids.Build()
	rb.v.Ids = &v
	return rb
}

func (rb *QueryContainerBuilder) Intervals(values map[Field]*IntervalsQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]IntervalsQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Intervals = tmp
	return rb
}

func (rb *QueryContainerBuilder) Match(values map[Field]*MatchQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]MatchQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Match = tmp
	return rb
}

func (rb *QueryContainerBuilder) MatchAll(matchall *MatchAllQueryBuilder) *QueryContainerBuilder {
	v := matchall.Build()
	rb.v.MatchAll = &v
	return rb
}

func (rb *QueryContainerBuilder) MatchBoolPrefix(values map[Field]*MatchBoolPrefixQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]MatchBoolPrefixQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.MatchBoolPrefix = tmp
	return rb
}

func (rb *QueryContainerBuilder) MatchNone(matchnone *MatchNoneQueryBuilder) *QueryContainerBuilder {
	v := matchnone.Build()
	rb.v.MatchNone = &v
	return rb
}

func (rb *QueryContainerBuilder) MatchPhrase(values map[Field]*MatchPhraseQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]MatchPhraseQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.MatchPhrase = tmp
	return rb
}

func (rb *QueryContainerBuilder) MatchPhrasePrefix(values map[Field]*MatchPhrasePrefixQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]MatchPhrasePrefixQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.MatchPhrasePrefix = tmp
	return rb
}

func (rb *QueryContainerBuilder) MoreLikeThis(morelikethis *MoreLikeThisQueryBuilder) *QueryContainerBuilder {
	v := morelikethis.Build()
	rb.v.MoreLikeThis = &v
	return rb
}

func (rb *QueryContainerBuilder) MultiMatch(multimatch *MultiMatchQueryBuilder) *QueryContainerBuilder {
	v := multimatch.Build()
	rb.v.MultiMatch = &v
	return rb
}

func (rb *QueryContainerBuilder) Nested(nested *NestedQueryBuilder) *QueryContainerBuilder {
	v := nested.Build()
	rb.v.Nested = &v
	return rb
}

func (rb *QueryContainerBuilder) ParentId(parentid *ParentIdQueryBuilder) *QueryContainerBuilder {
	v := parentid.Build()
	rb.v.ParentId = &v
	return rb
}

func (rb *QueryContainerBuilder) Percolate(percolate *PercolateQueryBuilder) *QueryContainerBuilder {
	v := percolate.Build()
	rb.v.Percolate = &v
	return rb
}

func (rb *QueryContainerBuilder) Pinned(pinned *PinnedQueryBuilder) *QueryContainerBuilder {
	v := pinned.Build()
	rb.v.Pinned = &v
	return rb
}

func (rb *QueryContainerBuilder) Prefix(values map[Field]*PrefixQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]PrefixQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Prefix = tmp
	return rb
}

func (rb *QueryContainerBuilder) QueryString(querystring *QueryStringQueryBuilder) *QueryContainerBuilder {
	v := querystring.Build()
	rb.v.QueryString = &v
	return rb
}

func (rb *QueryContainerBuilder) Range(values map[Field]*RangeQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]RangeQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Range = tmp
	return rb
}

func (rb *QueryContainerBuilder) RankFeature(rankfeature *RankFeatureQueryBuilder) *QueryContainerBuilder {
	v := rankfeature.Build()
	rb.v.RankFeature = &v
	return rb
}

func (rb *QueryContainerBuilder) Regexp(values map[Field]*RegexpQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]RegexpQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Regexp = tmp
	return rb
}

func (rb *QueryContainerBuilder) Script(script *ScriptQueryBuilder) *QueryContainerBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *QueryContainerBuilder) ScriptScore(scriptscore *ScriptScoreQueryBuilder) *QueryContainerBuilder {
	v := scriptscore.Build()
	rb.v.ScriptScore = &v
	return rb
}

func (rb *QueryContainerBuilder) Shape(shape *ShapeQueryBuilder) *QueryContainerBuilder {
	v := shape.Build()
	rb.v.Shape = &v
	return rb
}

func (rb *QueryContainerBuilder) SimpleQueryString(simplequerystring *SimpleQueryStringQueryBuilder) *QueryContainerBuilder {
	v := simplequerystring.Build()
	rb.v.SimpleQueryString = &v
	return rb
}

func (rb *QueryContainerBuilder) SpanContaining(spancontaining *SpanContainingQueryBuilder) *QueryContainerBuilder {
	v := spancontaining.Build()
	rb.v.SpanContaining = &v
	return rb
}

func (rb *QueryContainerBuilder) SpanFirst(spanfirst *SpanFirstQueryBuilder) *QueryContainerBuilder {
	v := spanfirst.Build()
	rb.v.SpanFirst = &v
	return rb
}

func (rb *QueryContainerBuilder) SpanMulti(spanmulti *SpanMultiTermQueryBuilder) *QueryContainerBuilder {
	v := spanmulti.Build()
	rb.v.SpanMulti = &v
	return rb
}

func (rb *QueryContainerBuilder) SpanNear(spannear *SpanNearQueryBuilder) *QueryContainerBuilder {
	v := spannear.Build()
	rb.v.SpanNear = &v
	return rb
}

func (rb *QueryContainerBuilder) SpanNot(spannot *SpanNotQueryBuilder) *QueryContainerBuilder {
	v := spannot.Build()
	rb.v.SpanNot = &v
	return rb
}

func (rb *QueryContainerBuilder) SpanOr(spanor *SpanOrQueryBuilder) *QueryContainerBuilder {
	v := spanor.Build()
	rb.v.SpanOr = &v
	return rb
}

func (rb *QueryContainerBuilder) SpanTerm(values map[Field]*SpanTermQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]SpanTermQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.SpanTerm = tmp
	return rb
}

func (rb *QueryContainerBuilder) SpanWithin(spanwithin *SpanWithinQueryBuilder) *QueryContainerBuilder {
	v := spanwithin.Build()
	rb.v.SpanWithin = &v
	return rb
}

func (rb *QueryContainerBuilder) Term(values map[Field]*TermQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]TermQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Term = tmp
	return rb
}

func (rb *QueryContainerBuilder) Terms(terms *TermsQueryBuilder) *QueryContainerBuilder {
	v := terms.Build()
	rb.v.Terms = &v
	return rb
}

func (rb *QueryContainerBuilder) TermsSet(values map[Field]*TermsSetQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]TermsSetQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.TermsSet = tmp
	return rb
}

func (rb *QueryContainerBuilder) Type_(type_ *TypeQueryBuilder) *QueryContainerBuilder {
	v := type_.Build()
	rb.v.Type = &v
	return rb
}

func (rb *QueryContainerBuilder) Wildcard(values map[Field]*WildcardQueryBuilder) *QueryContainerBuilder {
	tmp := make(map[Field]WildcardQuery, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Wildcard = tmp
	return rb
}

func (rb *QueryContainerBuilder) Wrapper(wrapper *WrapperQueryBuilder) *QueryContainerBuilder {
	v := wrapper.Build()
	rb.v.Wrapper = &v
	return rb
}
