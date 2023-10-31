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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package submit

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package submit
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/async_search/submit/AsyncSearchSubmitRequest.ts#L55-L286
type Request struct {
	Aggregations map[string]types.Aggregations `json:"aggregations,omitempty"`
	Collapse     *types.FieldCollapse          `json:"collapse,omitempty"`
	// DocvalueFields Array of wildcard (*) patterns. The request returns doc values for field
	// names matching these patterns in the hits.fields property of the response.
	DocvalueFields []types.FieldAndFormat `json:"docvalue_fields,omitempty"`
	// Explain If true, returns detailed information about score computation as part of a
	// hit.
	Explain *bool `json:"explain,omitempty"`
	// Ext Configuration of search extensions defined by Elasticsearch plugins.
	Ext map[string]json.RawMessage `json:"ext,omitempty"`
	// Fields Array of wildcard (*) patterns. The request returns values for field names
	// matching these patterns in the hits.fields property of the response.
	Fields []types.FieldAndFormat `json:"fields,omitempty"`
	// From Starting document offset. By default, you cannot page through more than
	// 10,000
	// hits using the from and size parameters. To page through more hits, use the
	// search_after parameter.
	From      *int             `json:"from,omitempty"`
	Highlight *types.Highlight `json:"highlight,omitempty"`
	// IndicesBoost Boosts the _score of documents from specified indices.
	IndicesBoost []map[string]types.Float64 `json:"indices_boost,omitempty"`
	// Knn Defines the approximate kNN search to run.
	Knn []types.KnnQuery `json:"knn,omitempty"`
	// MinScore Minimum _score for matching documents. Documents with a lower _score are
	// not included in the search results.
	MinScore *types.Float64 `json:"min_score,omitempty"`
	// Pit Limits the search to a point in time (PIT). If you provide a PIT, you
	// cannot specify an <index> in the request path.
	Pit        *types.PointInTimeReference `json:"pit,omitempty"`
	PostFilter *types.Query                `json:"post_filter,omitempty"`
	Profile    *bool                       `json:"profile,omitempty"`
	// Query Defines the search definition using the Query DSL.
	Query   *types.Query    `json:"query,omitempty"`
	Rescore []types.Rescore `json:"rescore,omitempty"`
	// RuntimeMappings Defines one or more runtime fields in the search request. These fields take
	// precedence over mapped fields with the same name.
	RuntimeMappings types.RuntimeFields `json:"runtime_mappings,omitempty"`
	// ScriptFields Retrieve a script evaluation (based on different fields) for each hit.
	ScriptFields map[string]types.ScriptField `json:"script_fields,omitempty"`
	SearchAfter  []types.FieldValue           `json:"search_after,omitempty"`
	// SeqNoPrimaryTerm If true, returns sequence number and primary term of the last modification
	// of each hit. See Optimistic concurrency control.
	SeqNoPrimaryTerm *bool `json:"seq_no_primary_term,omitempty"`
	// Size The number of hits to return. By default, you cannot page through more
	// than 10,000 hits using the from and size parameters. To page through more
	// hits, use the search_after parameter.
	Size  *int                     `json:"size,omitempty"`
	Slice *types.SlicedScroll      `json:"slice,omitempty"`
	Sort  []types.SortCombinations `json:"sort,omitempty"`
	// Source_ Indicates which source fields are returned for matching documents. These
	// fields are returned in the hits._source property of the search response.
	Source_ types.SourceConfig `json:"_source,omitempty"`
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
	StoredFields []string         `json:"stored_fields,omitempty"`
	Suggest      *types.Suggester `json:"suggest,omitempty"`
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
	TrackTotalHits types.TrackHits `json:"track_total_hits,omitempty"`
	// Version If true, returns document version as part of a hit.
	Version *bool `json:"version,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Aggregations: make(map[string]types.Aggregations, 0),
		Ext:          make(map[string]json.RawMessage, 0),
		ScriptFields: make(map[string]types.ScriptField, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Submit request: %w", err)
	}

	return &req, nil
}
