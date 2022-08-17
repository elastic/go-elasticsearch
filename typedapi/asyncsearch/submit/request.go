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


package submit

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package submit
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/async_search/submit/AsyncSearchSubmitRequest.ts#L54-L250
type Request struct {
	Aggregations map[string]types.AggregationContainer `json:"aggregations,omitempty"`

	Collapse *types.FieldCollapse `json:"collapse,omitempty"`

	// DocvalueFields Array of wildcard (*) patterns. The request returns doc values for field
	// names matching these patterns in the hits.fields property of the response.
	DocvalueFields []types.FieldAndFormat `json:"docvalue_fields,omitempty"`

	// Explain If true, returns detailed information about score computation as part of a
	// hit.
	Explain *bool `json:"explain,omitempty"`

	// Fields Array of wildcard (*) patterns. The request returns values for field names
	// matching these patterns in the hits.fields property of the response.
	Fields []types.FieldAndFormat `json:"fields,omitempty"`

	// From Starting document offset. By default, you cannot page through more than
	// 10,000
	// hits using the from and size parameters. To page through more hits, use the
	// search_after parameter.
	From *int `json:"from,omitempty"`

	Highlight *types.Highlight `json:"highlight,omitempty"`

	// IndicesBoost Boosts the _score of documents from specified indices.
	IndicesBoost []map[types.IndexName]float64 `json:"indices_boost,omitempty"`

	// Knn Defines the approximate kNN search to run.
	Knn *types.KnnQuery `json:"knn,omitempty"`

	// MinScore Minimum _score for matching documents. Documents with a lower _score are
	// not included in the search results.
	MinScore *float64 `json:"min_score,omitempty"`

	// Pit Limits the search to a point in time (PIT). If you provide a PIT, you
	// cannot specify an <index> in the request path.
	Pit *types.PointInTimeReference `json:"pit,omitempty"`

	PostFilter *types.QueryContainer `json:"post_filter,omitempty"`

	Profile *bool `json:"profile,omitempty"`

	// Query Defines the search definition using the Query DSL.
	Query *types.QueryContainer `json:"query,omitempty"`

	Rescore []types.Rescore `json:"rescore,omitempty"`

	// RuntimeMappings Defines one or more runtime fields in the search request. These fields take
	// precedence over mapped fields with the same name.
	RuntimeMappings *types.RuntimeFields `json:"runtime_mappings,omitempty"`

	// ScriptFields Retrieve a script evaluation (based on different fields) for each hit.
	ScriptFields map[string]types.ScriptField `json:"script_fields,omitempty"`

	SearchAfter *types.SortResults `json:"search_after,omitempty"`

	// SeqNoPrimaryTerm If true, returns sequence number and primary term of the last modification
	// of each hit. See Optimistic concurrency control.
	SeqNoPrimaryTerm *bool `json:"seq_no_primary_term,omitempty"`

	// Size The number of hits to return. By default, you cannot page through more
	// than 10,000 hits using the from and size parameters. To page through more
	// hits, use the search_after parameter.
	Size *int `json:"size,omitempty"`

	Slice *types.SlicedScroll `json:"slice,omitempty"`

	Sort *types.Sort `json:"sort,omitempty"`

	// Source_ Indicates which source fields are returned for matching documents. These
	// fields are returned in the hits._source property of the search response.
	Source_ *types.SourceConfig `json:"_source,omitempty"`

	// Stats Stats groups to associate with the search. Each group maintains a statistics
	// aggregation for its associated searches. You can retrieve these stats using
	// the indices stats API.
	Stats []string `json:"stats,omitempty"`

	// StoredFields List of stored fields to return as part of a hit. If no fields are specified,
	// no stored fields are included in the response. If this field is specified,
	// the _source
	// parameter defaults to false. You can pass _source: true to return both source
	// fields
	// and stored fields in the search response.
	StoredFields *types.Fields `json:"stored_fields,omitempty"`

	Suggest *types.Suggester `json:"suggest,omitempty"`

	// TerminateAfter Maximum number of documents to collect for each shard. If a query reaches
	// this
	// limit, Elasticsearch terminates the query early. Elasticsearch collects
	// documents
	// before sorting. Defaults to 0, which does not terminate query execution
	// early.
	TerminateAfter *int64 `json:"terminate_after,omitempty"`

	// Timeout Specifies the period of time to wait for a response from each shard. If no
	// response
	// is received before the timeout expires, the request fails and returns an
	// error.
	// Defaults to no timeout.
	Timeout *string `json:"timeout,omitempty"`

	// TrackScores If true, calculate and return document scores, even if the scores are not
	// used for sorting.
	TrackScores *bool `json:"track_scores,omitempty"`

	// TrackTotalHits Number of hits matching the query to count accurately. If true, the exact
	// number of hits is returned at the cost of some performance. If false, the
	// response does not include the total number of hits matching the query.
	// Defaults to 10,000 hits.
	TrackTotalHits *types.TrackHits `json:"track_total_hits,omitempty"`

	// Version If true, returns document version as part of a hit.
	Version *bool `json:"version,omitempty"`
}

// RequestBuilder is the builder API for the submit.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Aggregations: make(map[string]types.AggregationContainer, 0),
			ScriptFields: make(map[string]types.ScriptField, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Submit request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Aggregations(values map[string]*types.AggregationContainerBuilder) *RequestBuilder {
	tmp := make(map[string]types.AggregationContainer, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *RequestBuilder) Collapse(collapse *types.FieldCollapseBuilder) *RequestBuilder {
	v := collapse.Build()
	rb.v.Collapse = &v
	return rb
}

func (rb *RequestBuilder) DocvalueFields(docvalue_fields []types.FieldAndFormatBuilder) *RequestBuilder {
	tmp := make([]types.FieldAndFormat, len(docvalue_fields))
	for _, value := range docvalue_fields {
		tmp = append(tmp, value.Build())
	}
	rb.v.DocvalueFields = tmp
	return rb
}

func (rb *RequestBuilder) Explain(explain bool) *RequestBuilder {
	rb.v.Explain = &explain
	return rb
}

func (rb *RequestBuilder) Fields(fields []types.FieldAndFormatBuilder) *RequestBuilder {
	tmp := make([]types.FieldAndFormat, len(fields))
	for _, value := range fields {
		tmp = append(tmp, value.Build())
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *RequestBuilder) From(from int) *RequestBuilder {
	rb.v.From = &from
	return rb
}

func (rb *RequestBuilder) Highlight(highlight *types.HighlightBuilder) *RequestBuilder {
	v := highlight.Build()
	rb.v.Highlight = &v
	return rb
}

func (rb *RequestBuilder) IndicesBoost(value ...map[types.IndexName]float64) *RequestBuilder {
	rb.v.IndicesBoost = value
	return rb
}

func (rb *RequestBuilder) Knn(knn *types.KnnQueryBuilder) *RequestBuilder {
	v := knn.Build()
	rb.v.Knn = &v
	return rb
}

func (rb *RequestBuilder) MinScore(minscore float64) *RequestBuilder {
	rb.v.MinScore = &minscore
	return rb
}

func (rb *RequestBuilder) Pit(pit *types.PointInTimeReferenceBuilder) *RequestBuilder {
	v := pit.Build()
	rb.v.Pit = &v
	return rb
}

func (rb *RequestBuilder) PostFilter(postfilter *types.QueryContainerBuilder) *RequestBuilder {
	v := postfilter.Build()
	rb.v.PostFilter = &v
	return rb
}

func (rb *RequestBuilder) Profile(profile bool) *RequestBuilder {
	rb.v.Profile = &profile
	return rb
}

func (rb *RequestBuilder) Query(query *types.QueryContainerBuilder) *RequestBuilder {
	v := query.Build()
	rb.v.Query = &v
	return rb
}

func (rb *RequestBuilder) Rescore(arg []types.Rescore) *RequestBuilder {
	rb.v.Rescore = arg
	return rb
}

func (rb *RequestBuilder) RuntimeMappings(runtimemappings *types.RuntimeFieldsBuilder) *RequestBuilder {
	v := runtimemappings.Build()
	rb.v.RuntimeMappings = &v
	return rb
}

func (rb *RequestBuilder) ScriptFields(values map[string]*types.ScriptFieldBuilder) *RequestBuilder {
	tmp := make(map[string]types.ScriptField, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.ScriptFields = tmp
	return rb
}

func (rb *RequestBuilder) SearchAfter(searchafter *types.SortResultsBuilder) *RequestBuilder {
	v := searchafter.Build()
	rb.v.SearchAfter = &v
	return rb
}

func (rb *RequestBuilder) SeqNoPrimaryTerm(seqnoprimaryterm bool) *RequestBuilder {
	rb.v.SeqNoPrimaryTerm = &seqnoprimaryterm
	return rb
}

func (rb *RequestBuilder) Size(size int) *RequestBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *RequestBuilder) Slice(slice *types.SlicedScrollBuilder) *RequestBuilder {
	v := slice.Build()
	rb.v.Slice = &v
	return rb
}

func (rb *RequestBuilder) Sort(sort *types.SortBuilder) *RequestBuilder {
	v := sort.Build()
	rb.v.Sort = &v
	return rb
}

func (rb *RequestBuilder) Source_(source_ *types.SourceConfigBuilder) *RequestBuilder {
	v := source_.Build()
	rb.v.Source_ = &v
	return rb
}

func (rb *RequestBuilder) Stats(stats ...string) *RequestBuilder {
	rb.v.Stats = stats
	return rb
}

func (rb *RequestBuilder) StoredFields(storedfields *types.FieldsBuilder) *RequestBuilder {
	v := storedfields.Build()
	rb.v.StoredFields = &v
	return rb
}

func (rb *RequestBuilder) Suggest(suggest *types.SuggesterBuilder) *RequestBuilder {
	v := suggest.Build()
	rb.v.Suggest = &v
	return rb
}

func (rb *RequestBuilder) TerminateAfter(terminateafter int64) *RequestBuilder {
	rb.v.TerminateAfter = &terminateafter
	return rb
}

func (rb *RequestBuilder) Timeout(timeout string) *RequestBuilder {
	rb.v.Timeout = &timeout
	return rb
}

func (rb *RequestBuilder) TrackScores(trackscores bool) *RequestBuilder {
	rb.v.TrackScores = &trackscores
	return rb
}

func (rb *RequestBuilder) TrackTotalHits(tracktotalhits *types.TrackHitsBuilder) *RequestBuilder {
	v := tracktotalhits.Build()
	rb.v.TrackTotalHits = &v
	return rb
}

func (rb *RequestBuilder) Version(version bool) *RequestBuilder {
	rb.v.Version = &version
	return rb
}
