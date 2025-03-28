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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _searchRequestBody struct {
	v *types.SearchRequestBody
}

func NewSearchRequestBody() *_searchRequestBody {

	return &_searchRequestBody{v: types.NewSearchRequestBody()}

}

// Defines the aggregations that are run as part of the search request.
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

// Collapses search results the values of the specified field.
func (s *_searchRequestBody) Collapse(collapse types.FieldCollapseVariant) *_searchRequestBody {

	s.v.Collapse = collapse.FieldCollapseCaster()

	return s
}

// An array of wildcard (`*`) field patterns.
// The request returns doc values for field names matching these patterns in the
// `hits.fields` property of the response.
func (s *_searchRequestBody) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *_searchRequestBody {

	for _, v := range docvaluefields {

		s.v.DocvalueFields = append(s.v.DocvalueFields, *v.FieldAndFormatCaster())

	}
	return s
}

// If `true`, the request returns detailed information about score computation
// as part of a hit.
func (s *_searchRequestBody) Explain(explain bool) *_searchRequestBody {

	s.v.Explain = &explain

	return s
}

// Configuration of search extensions defined by Elasticsearch plugins.
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

// An array of wildcard (`*`) field patterns.
// The request returns values for field names matching these patterns in the
// `hits.fields` property of the response.
func (s *_searchRequestBody) Fields(fields ...types.FieldAndFormatVariant) *_searchRequestBody {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, *v.FieldAndFormatCaster())

	}
	return s
}

// The starting document offset, which must be non-negative.
// By default, you cannot page through more than 10,000 hits using the `from`
// and `size` parameters.
// To page through more hits, use the `search_after` parameter.
func (s *_searchRequestBody) From(from int) *_searchRequestBody {

	s.v.From = &from

	return s
}

// Specifies the highlighter to use for retrieving highlighted snippets from one
// or more fields in your search results.
func (s *_searchRequestBody) Highlight(highlight types.HighlightVariant) *_searchRequestBody {

	s.v.Highlight = highlight.HighlightCaster()

	return s
}

// Boost the `_score` of documents from specified indices.
// The boost value is the factor by which scores are multiplied.
// A boost value greater than `1.0` increases the score.
// A boost value between `0` and `1.0` decreases the score.
func (s *_searchRequestBody) IndicesBoost(indicesboost []map[string]types.Float64) *_searchRequestBody {

	s.v.IndicesBoost = indicesboost

	return s
}

// The approximate kNN search to run.
func (s *_searchRequestBody) Knn(knns ...types.KnnSearchVariant) *_searchRequestBody {

	s.v.Knn = make([]types.KnnSearch, len(knns))
	for i, v := range knns {
		s.v.Knn[i] = *v.KnnSearchCaster()
	}

	return s
}

// The minimum `_score` for matching documents.
// Documents with a lower `_score` are not included in the search results.
func (s *_searchRequestBody) MinScore(minscore types.Float64) *_searchRequestBody {

	s.v.MinScore = &minscore

	return s
}

// Limit the search to a point in time (PIT).
// If you provide a PIT, you cannot specify an `<index>` in the request path.
func (s *_searchRequestBody) Pit(pit types.PointInTimeReferenceVariant) *_searchRequestBody {

	s.v.Pit = pit.PointInTimeReferenceCaster()

	return s
}

// Use the `post_filter` parameter to filter search results.
// The search hits are filtered after the aggregations are calculated.
// A post filter has no impact on the aggregation results.
func (s *_searchRequestBody) PostFilter(postfilter types.QueryVariant) *_searchRequestBody {

	s.v.PostFilter = postfilter.QueryCaster()

	return s
}

// Set to `true` to return detailed timing information about the execution of
// individual components in a search request.
// NOTE: This is a debugging tool and adds significant overhead to search
// execution.
func (s *_searchRequestBody) Profile(profile bool) *_searchRequestBody {

	s.v.Profile = &profile

	return s
}

// The search definition using the Query DSL.
func (s *_searchRequestBody) Query(query types.QueryVariant) *_searchRequestBody {

	s.v.Query = query.QueryCaster()

	return s
}

// The Reciprocal Rank Fusion (RRF) to use.
func (s *_searchRequestBody) Rank(rank types.RankContainerVariant) *_searchRequestBody {

	s.v.Rank = rank.RankContainerCaster()

	return s
}

// Can be used to improve precision by reordering just the top (for example 100
// - 500) documents returned by the `query` and `post_filter` phases.
func (s *_searchRequestBody) Rescore(rescores ...types.RescoreVariant) *_searchRequestBody {

	s.v.Rescore = make([]types.Rescore, len(rescores))
	for i, v := range rescores {
		s.v.Rescore[i] = *v.RescoreCaster()
	}

	return s
}

// A retriever is a specification to describe top documents returned from a
// search.
// A retriever replaces other elements of the search API that also return top
// documents such as `query` and `knn`.
func (s *_searchRequestBody) Retriever(retriever types.RetrieverContainerVariant) *_searchRequestBody {

	s.v.Retriever = retriever.RetrieverContainerCaster()

	return s
}

// One or more runtime fields in the search request.
// These fields take precedence over mapped fields with the same name.
func (s *_searchRequestBody) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_searchRequestBody {

	s.v.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return s
}

// Retrieve a script evaluation (based on different fields) for each hit.
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

// Used to retrieve the next page of hits using a set of sort values from the
// previous page.
func (s *_searchRequestBody) SearchAfter(sortresults ...types.FieldValueVariant) *_searchRequestBody {

	for _, v := range sortresults {
		s.v.SearchAfter = append(s.v.SearchAfter, *v.FieldValueCaster())
	}

	return s
}

// If `true`, the request returns sequence number and primary term of the last
// modification of each hit.
func (s *_searchRequestBody) SeqNoPrimaryTerm(seqnoprimaryterm bool) *_searchRequestBody {

	s.v.SeqNoPrimaryTerm = &seqnoprimaryterm

	return s
}

// The number of hits to return, which must not be negative.
// By default, you cannot page through more than 10,000 hits using the `from`
// and `size` parameters.
// To page through more hits, use the `search_after` property.
func (s *_searchRequestBody) Size(size int) *_searchRequestBody {

	s.v.Size = &size

	return s
}

// Split a scrolled search into multiple slices that can be consumed
// independently.
func (s *_searchRequestBody) Slice(slice types.SlicedScrollVariant) *_searchRequestBody {

	s.v.Slice = slice.SlicedScrollCaster()

	return s
}

// A comma-separated list of <field>:<direction> pairs.
func (s *_searchRequestBody) Sort(sorts ...types.SortCombinationsVariant) *_searchRequestBody {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

// The source fields that are returned for matching documents.
// These fields are returned in the `hits._source` property of the search
// response.
// If the `stored_fields` property is specified, the `_source` property defaults
// to `false`.
// Otherwise, it defaults to `true`.
func (s *_searchRequestBody) Source_(sourceconfig types.SourceConfigVariant) *_searchRequestBody {

	s.v.Source_ = *sourceconfig.SourceConfigCaster()

	return s
}

// The stats groups to associate with the search.
// Each group maintains a statistics aggregation for its associated searches.
// You can retrieve these stats using the indices stats API.
func (s *_searchRequestBody) Stats(stats ...string) *_searchRequestBody {

	for _, v := range stats {

		s.v.Stats = append(s.v.Stats, v)

	}
	return s
}

// A comma-separated list of stored fields to return as part of a hit.
// If no fields are specified, no stored fields are included in the response.
// If this field is specified, the `_source` property defaults to `false`.
// You can pass `_source: true` to return both source fields and stored fields
// in the search response.
func (s *_searchRequestBody) StoredFields(fields ...string) *_searchRequestBody {

	s.v.StoredFields = fields

	return s
}

// Defines a suggester that provides similar looking terms based on a provided
// text.
func (s *_searchRequestBody) Suggest(suggest types.SuggesterVariant) *_searchRequestBody {

	s.v.Suggest = suggest.SuggesterCaster()

	return s
}

// The maximum number of documents to collect for each shard.
// If a query reaches this limit, Elasticsearch terminates the query early.
// Elasticsearch collects documents before sorting.
//
// IMPORTANT: Use with caution.
// Elasticsearch applies this property to each shard handling the request.
// When possible, let Elasticsearch perform early termination automatically.
// Avoid specifying this property for requests that target data streams with
// backing indices across multiple data tiers.
//
// If set to `0` (default), the query does not terminate early.
func (s *_searchRequestBody) TerminateAfter(terminateafter int64) *_searchRequestBody {

	s.v.TerminateAfter = &terminateafter

	return s
}

// The period of time to wait for a response from each shard.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// Defaults to no timeout.
func (s *_searchRequestBody) Timeout(timeout string) *_searchRequestBody {

	s.v.Timeout = &timeout

	return s
}

// If `true`, calculate and return document scores, even if the scores are not
// used for sorting.
func (s *_searchRequestBody) TrackScores(trackscores bool) *_searchRequestBody {

	s.v.TrackScores = &trackscores

	return s
}

// Number of hits matching the query to count accurately.
// If `true`, the exact number of hits is returned at the cost of some
// performance.
// If `false`, the  response does not include the total number of hits matching
// the query.
func (s *_searchRequestBody) TrackTotalHits(trackhits types.TrackHitsVariant) *_searchRequestBody {

	s.v.TrackTotalHits = *trackhits.TrackHitsCaster()

	return s
}

// If `true`, the request returns the document version as part of a hit.
func (s *_searchRequestBody) Version(version bool) *_searchRequestBody {

	s.v.Version = &version

	return s
}

func (s *_searchRequestBody) SearchRequestBodyCaster() *types.SearchRequestBody {
	return s.v
}
