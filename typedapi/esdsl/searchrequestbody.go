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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _searchRequestBody struct {
	v *types.SearchRequestBody
}

func NewSearchRequestBody() *_searchRequestBody {

	return &_searchRequestBody{v: types.NewSearchRequestBody()}

}

func (s *_searchRequestBody) Aggregations(aggregations map[string]types.Aggregations) *_searchRequestBody {

	s.v.Aggregations = aggregations
	return s
}

func (s *_searchRequestBody) AddAggregation(key string, value types.AggregationsVariant) *_searchRequestBody {

	var tmp map[string]types.Aggregations
	if s.v.Aggregations == nil {
		s.v.Aggregations = make(map[string]types.Aggregations)
	} else {
		tmp = s.v.Aggregations
	}

	tmp[key] = *value.AggregationsCaster()

	s.v.Aggregations = tmp
	return s
}

func (s *_searchRequestBody) Collapse(collapse types.FieldCollapseVariant) *_searchRequestBody {

	s.v.Collapse = collapse.FieldCollapseCaster()

	return s
}

func (s *_searchRequestBody) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *_searchRequestBody {

	for _, v := range docvaluefields {

		s.v.DocvalueFields = append(s.v.DocvalueFields, *v.FieldAndFormatCaster())

	}
	return s
}

func (s *_searchRequestBody) Explain(explain bool) *_searchRequestBody {

	s.v.Explain = &explain

	return s
}

func (s *_searchRequestBody) Ext(ext map[string]json.RawMessage) *_searchRequestBody {

	s.v.Ext = ext
	return s
}

func (s *_searchRequestBody) AddExt(key string, value json.RawMessage) *_searchRequestBody {

	var tmp map[string]json.RawMessage
	if s.v.Ext == nil {
		s.v.Ext = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Ext
	}

	tmp[key] = value

	s.v.Ext = tmp
	return s
}

func (s *_searchRequestBody) Fields(fields ...types.FieldAndFormatVariant) *_searchRequestBody {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, *v.FieldAndFormatCaster())

	}
	return s
}

func (s *_searchRequestBody) From(from int) *_searchRequestBody {

	s.v.From = &from

	return s
}

func (s *_searchRequestBody) Highlight(highlight types.HighlightVariant) *_searchRequestBody {

	s.v.Highlight = highlight.HighlightCaster()

	return s
}

func (s *_searchRequestBody) IndicesBoost(indicesboost []map[string]types.Float64) *_searchRequestBody {

	s.v.IndicesBoost = indicesboost

	return s
}

func (s *_searchRequestBody) Knn(knns ...types.KnnSearchVariant) *_searchRequestBody {

	s.v.Knn = make([]types.KnnSearch, len(knns))
	for i, v := range knns {
		s.v.Knn[i] = *v.KnnSearchCaster()
	}

	return s
}

func (s *_searchRequestBody) MinScore(minscore types.Float64) *_searchRequestBody {

	s.v.MinScore = &minscore

	return s
}

func (s *_searchRequestBody) Pit(pit types.PointInTimeReferenceVariant) *_searchRequestBody {

	s.v.Pit = pit.PointInTimeReferenceCaster()

	return s
}

func (s *_searchRequestBody) PostFilter(postfilter types.QueryVariant) *_searchRequestBody {

	s.v.PostFilter = postfilter.QueryCaster()

	return s
}

func (s *_searchRequestBody) Profile(profile bool) *_searchRequestBody {

	s.v.Profile = &profile

	return s
}

func (s *_searchRequestBody) Query(query types.QueryVariant) *_searchRequestBody {

	s.v.Query = query.QueryCaster()

	return s
}

func (s *_searchRequestBody) Rank(rank types.RankContainerVariant) *_searchRequestBody {

	s.v.Rank = rank.RankContainerCaster()

	return s
}

func (s *_searchRequestBody) Rescore(rescores ...types.RescoreVariant) *_searchRequestBody {

	s.v.Rescore = make([]types.Rescore, len(rescores))
	for i, v := range rescores {
		s.v.Rescore[i] = *v.RescoreCaster()
	}

	return s
}

func (s *_searchRequestBody) Retriever(retriever types.RetrieverContainerVariant) *_searchRequestBody {

	s.v.Retriever = retriever.RetrieverContainerCaster()

	return s
}

func (s *_searchRequestBody) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_searchRequestBody {

	s.v.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return s
}

func (s *_searchRequestBody) ScriptFields(scriptfields map[string]types.ScriptField) *_searchRequestBody {

	s.v.ScriptFields = scriptfields
	return s
}

func (s *_searchRequestBody) AddScriptField(key string, value types.ScriptFieldVariant) *_searchRequestBody {

	var tmp map[string]types.ScriptField
	if s.v.ScriptFields == nil {
		s.v.ScriptFields = make(map[string]types.ScriptField)
	} else {
		tmp = s.v.ScriptFields
	}

	tmp[key] = *value.ScriptFieldCaster()

	s.v.ScriptFields = tmp
	return s
}

func (s *_searchRequestBody) SearchAfter(sortresults ...types.FieldValueVariant) *_searchRequestBody {

	for _, v := range sortresults {
		s.v.SearchAfter = append(s.v.SearchAfter, *v.FieldValueCaster())
	}

	return s
}

func (s *_searchRequestBody) SeqNoPrimaryTerm(seqnoprimaryterm bool) *_searchRequestBody {

	s.v.SeqNoPrimaryTerm = &seqnoprimaryterm

	return s
}

func (s *_searchRequestBody) Size(size int) *_searchRequestBody {

	s.v.Size = &size

	return s
}

func (s *_searchRequestBody) Slice(slice types.SlicedScrollVariant) *_searchRequestBody {

	s.v.Slice = slice.SlicedScrollCaster()

	return s
}

func (s *_searchRequestBody) Sort(sorts ...types.SortCombinationsVariant) *_searchRequestBody {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

func (s *_searchRequestBody) Source_(sourceconfig types.SourceConfigVariant) *_searchRequestBody {

	s.v.Source_ = *sourceconfig.SourceConfigCaster()

	return s
}

func (s *_searchRequestBody) Stats(stats ...string) *_searchRequestBody {

	for _, v := range stats {

		s.v.Stats = append(s.v.Stats, v)

	}
	return s
}

func (s *_searchRequestBody) StoredFields(fields ...string) *_searchRequestBody {

	s.v.StoredFields = fields

	return s
}

func (s *_searchRequestBody) Suggest(suggest types.SuggesterVariant) *_searchRequestBody {

	s.v.Suggest = suggest.SuggesterCaster()

	return s
}

func (s *_searchRequestBody) TerminateAfter(terminateafter int64) *_searchRequestBody {

	s.v.TerminateAfter = &terminateafter

	return s
}

func (s *_searchRequestBody) Timeout(timeout string) *_searchRequestBody {

	s.v.Timeout = &timeout

	return s
}

func (s *_searchRequestBody) TrackScores(trackscores bool) *_searchRequestBody {

	s.v.TrackScores = &trackscores

	return s
}

func (s *_searchRequestBody) TrackTotalHits(trackhits types.TrackHitsVariant) *_searchRequestBody {

	s.v.TrackTotalHits = *trackhits.TrackHitsCaster()

	return s
}

func (s *_searchRequestBody) Version(version bool) *_searchRequestBody {

	s.v.Version = &version

	return s
}

func (s *_searchRequestBody) SearchRequestBodyCaster() *types.SearchRequestBody {
	return s.v
}
