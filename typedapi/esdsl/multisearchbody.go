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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _multisearchBody struct {
	v *types.MultisearchBody
}

func NewMultisearchBody() *_multisearchBody {

	return &_multisearchBody{v: types.NewMultisearchBody()}

}

func (s *_multisearchBody) Aggregations(aggregations map[string]types.Aggregations) *_multisearchBody {

	s.v.Aggregations = aggregations
	return s
}

func (s *_multisearchBody) AddAggregation(key string, value types.AggregationsVariant) *_multisearchBody {

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

func (s *_multisearchBody) Collapse(collapse types.FieldCollapseVariant) *_multisearchBody {

	s.v.Collapse = collapse.FieldCollapseCaster()

	return s
}

// Array of wildcard (*) patterns. The request returns doc values for field
// names matching these patterns in the hits.fields property of the response.
func (s *_multisearchBody) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *_multisearchBody {

	for _, v := range docvaluefields {

		s.v.DocvalueFields = append(s.v.DocvalueFields, *v.FieldAndFormatCaster())

	}
	return s
}

// If true, returns detailed information about score computation as part of a
// hit.
func (s *_multisearchBody) Explain(explain bool) *_multisearchBody {

	s.v.Explain = &explain

	return s
}

// Configuration of search extensions defined by Elasticsearch plugins.
func (s *_multisearchBody) Ext(ext map[string]json.RawMessage) *_multisearchBody {

	s.v.Ext = ext
	return s
}

func (s *_multisearchBody) AddExt(key string, value json.RawMessage) *_multisearchBody {

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

// Array of wildcard (*) patterns. The request returns values for field names
// matching these patterns in the hits.fields property of the response.
func (s *_multisearchBody) Fields(fields ...types.FieldAndFormatVariant) *_multisearchBody {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, *v.FieldAndFormatCaster())

	}
	return s
}

// Starting document offset. By default, you cannot page through more than
// 10,000
// hits using the from and size parameters. To page through more hits, use the
// search_after parameter.
func (s *_multisearchBody) From(from int) *_multisearchBody {

	s.v.From = &from

	return s
}

func (s *_multisearchBody) Highlight(highlight types.HighlightVariant) *_multisearchBody {

	s.v.Highlight = highlight.HighlightCaster()

	return s
}

// Boosts the _score of documents from specified indices.
func (s *_multisearchBody) IndicesBoost(indicesboost []map[string]types.Float64) *_multisearchBody {

	s.v.IndicesBoost = indicesboost

	return s
}

// Defines the approximate kNN search to run.
func (s *_multisearchBody) Knn(knns ...types.KnnSearchVariant) *_multisearchBody {

	s.v.Knn = make([]types.KnnSearch, len(knns))
	for i, v := range knns {
		s.v.Knn[i] = *v.KnnSearchCaster()
	}

	return s
}

// Minimum _score for matching documents. Documents with a lower _score are
// not included in the search results.
func (s *_multisearchBody) MinScore(minscore types.Float64) *_multisearchBody {

	s.v.MinScore = &minscore

	return s
}

// Limits the search to a point in time (PIT). If you provide a PIT, you
// cannot specify an <index> in the request path.
func (s *_multisearchBody) Pit(pit types.PointInTimeReferenceVariant) *_multisearchBody {

	s.v.Pit = pit.PointInTimeReferenceCaster()

	return s
}

func (s *_multisearchBody) PostFilter(postfilter types.QueryVariant) *_multisearchBody {

	s.v.PostFilter = postfilter.QueryCaster()

	return s
}

func (s *_multisearchBody) Profile(profile bool) *_multisearchBody {

	s.v.Profile = &profile

	return s
}

// Defines the search definition using the Query DSL.
func (s *_multisearchBody) Query(query types.QueryVariant) *_multisearchBody {

	s.v.Query = query.QueryCaster()

	return s
}

func (s *_multisearchBody) Rescore(rescores ...types.RescoreVariant) *_multisearchBody {

	s.v.Rescore = make([]types.Rescore, len(rescores))
	for i, v := range rescores {
		s.v.Rescore[i] = *v.RescoreCaster()
	}

	return s
}

// Defines one or more runtime fields in the search request. These fields take
// precedence over mapped fields with the same name.
func (s *_multisearchBody) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_multisearchBody {

	s.v.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return s
}

// Retrieve a script evaluation (based on different fields) for each hit.
func (s *_multisearchBody) ScriptFields(scriptfields map[string]types.ScriptField) *_multisearchBody {

	s.v.ScriptFields = scriptfields
	return s
}

func (s *_multisearchBody) AddScriptField(key string, value types.ScriptFieldVariant) *_multisearchBody {

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

func (s *_multisearchBody) SearchAfter(sortresults ...types.FieldValueVariant) *_multisearchBody {

	for _, v := range sortresults {
		s.v.SearchAfter = append(s.v.SearchAfter, *v.FieldValueCaster())
	}

	return s
}

// If true, returns sequence number and primary term of the last modification
// of each hit. See Optimistic concurrency control.
func (s *_multisearchBody) SeqNoPrimaryTerm(seqnoprimaryterm bool) *_multisearchBody {

	s.v.SeqNoPrimaryTerm = &seqnoprimaryterm

	return s
}

// The number of hits to return. By default, you cannot page through more
// than 10,000 hits using the from and size parameters. To page through more
// hits, use the search_after parameter.
func (s *_multisearchBody) Size(size int) *_multisearchBody {

	s.v.Size = &size

	return s
}

func (s *_multisearchBody) Sort(sorts ...types.SortCombinationsVariant) *_multisearchBody {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

// Indicates which source fields are returned for matching documents. These
// fields are returned in the hits._source property of the search response.
func (s *_multisearchBody) Source_(sourceconfig types.SourceConfigVariant) *_multisearchBody {

	s.v.Source_ = *sourceconfig.SourceConfigCaster()

	return s
}

// Stats groups to associate with the search. Each group maintains a statistics
// aggregation for its associated searches. You can retrieve these stats using
// the indices stats API.
func (s *_multisearchBody) Stats(stats ...string) *_multisearchBody {

	for _, v := range stats {

		s.v.Stats = append(s.v.Stats, v)

	}
	return s
}

// List of stored fields to return as part of a hit. If no fields are specified,
// no stored fields are included in the response. If this field is specified,
// the _source
// parameter defaults to false. You can pass _source: true to return both source
// fields
// and stored fields in the search response.
func (s *_multisearchBody) StoredFields(fields ...string) *_multisearchBody {

	s.v.StoredFields = fields

	return s
}

func (s *_multisearchBody) Suggest(suggest types.SuggesterVariant) *_multisearchBody {

	s.v.Suggest = suggest.SuggesterCaster()

	return s
}

// Maximum number of documents to collect for each shard. If a query reaches
// this
// limit, Elasticsearch terminates the query early. Elasticsearch collects
// documents
// before sorting. Defaults to 0, which does not terminate query execution
// early.
func (s *_multisearchBody) TerminateAfter(terminateafter int64) *_multisearchBody {

	s.v.TerminateAfter = &terminateafter

	return s
}

// Specifies the period of time to wait for a response from each shard. If no
// response
// is received before the timeout expires, the request fails and returns an
// error.
// Defaults to no timeout.
func (s *_multisearchBody) Timeout(timeout string) *_multisearchBody {

	s.v.Timeout = &timeout

	return s
}

// If true, calculate and return document scores, even if the scores are not
// used for sorting.
func (s *_multisearchBody) TrackScores(trackscores bool) *_multisearchBody {

	s.v.TrackScores = &trackscores

	return s
}

// Number of hits matching the query to count accurately. If true, the exact
// number of hits is returned at the cost of some performance. If false, the
// response does not include the total number of hits matching the query.
// Defaults to 10,000 hits.
func (s *_multisearchBody) TrackTotalHits(trackhits types.TrackHitsVariant) *_multisearchBody {

	s.v.TrackTotalHits = *trackhits.TrackHitsCaster()

	return s
}

// If true, returns document version as part of a hit.
func (s *_multisearchBody) Version(version bool) *_multisearchBody {

	s.v.Version = &version

	return s
}

func (s *_multisearchBody) MultisearchBodyCaster() *types.MultisearchBody {
	return s.v
}
