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


package putdatafeed

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putdatafeed
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/put_datafeed/MlPutDatafeedRequest.ts#L37-L171
type Request struct {

	// Aggregations If set, the datafeed performs aggregation searches.
	// Support for aggregations is limited and should be used only with low
	// cardinality data.
	Aggregations map[string]types.Aggregations `json:"aggregations,omitempty"`
	// ChunkingConfig Datafeeds might be required to search over long time periods, for several
	// months or years.
	// This search is split into time chunks in order to ensure the load on
	// Elasticsearch is managed.
	// Chunking configuration controls how the size of these time chunks are
	// calculated;
	// it is an advanced configuration option.
	ChunkingConfig *types.ChunkingConfig `json:"chunking_config,omitempty"`
	// DelayedDataCheckConfig Specifies whether the datafeed checks for missing data and the size of the
	// window.
	// The datafeed can optionally search over indices that have already been read
	// in an effort to determine whether
	// any data has subsequently been added to the index. If missing data is found,
	// it is a good indication that the
	// `query_delay` is set too low and the data is being indexed after the datafeed
	// has passed that moment in time.
	// This check runs only on real-time datafeeds.
	DelayedDataCheckConfig *types.DelayedDataCheckConfig `json:"delayed_data_check_config,omitempty"`
	// Frequency The interval at which scheduled queries are made while the datafeed runs in
	// real time.
	// The default value is either the bucket span for short bucket spans, or, for
	// longer bucket spans, a sensible
	// fraction of the bucket span. When `frequency` is shorter than the bucket
	// span, interim results for the last
	// (partial) bucket are written then eventually overwritten by the full bucket
	// results. If the datafeed uses
	// aggregations, this value must be divisible by the interval of the date
	// histogram aggregation.
	Frequency *types.Duration     `json:"frequency,omitempty"`
	Headers   map[string][]string `json:"headers,omitempty"`
	// Indices An array of index names. Wildcards are supported. If any of the indices are
	// in remote clusters, the machine
	// learning nodes must have the `remote_cluster_client` role.
	Indices []string `json:"indices,omitempty"`
	// IndicesOptions Specifies index expansion options that are used during search
	IndicesOptions *types.IndicesOptions `json:"indices_options,omitempty"`
	// JobId Identifier for the anomaly detection job.
	JobId *string `json:"job_id,omitempty"`
	// MaxEmptySearches If a real-time datafeed has never seen any data (including during any initial
	// training period), it automatically
	// stops and closes the associated job after this many real-time searches return
	// no documents. In other words,
	// it stops after `frequency` times `max_empty_searches` of real-time operation.
	// If not set, a datafeed with no
	// end time that sees no data remains started until it is explicitly stopped. By
	// default, it is not set.
	MaxEmptySearches *int `json:"max_empty_searches,omitempty"`
	// Query The Elasticsearch query domain-specific language (DSL). This value
	// corresponds to the query object in an
	// Elasticsearch search POST body. All the options that are supported by
	// Elasticsearch can be used, as this
	// object is passed verbatim to Elasticsearch.
	Query *types.Query `json:"query,omitempty"`
	// QueryDelay The number of seconds behind real time that data is queried. For example, if
	// data from 10:04 a.m. might
	// not be searchable in Elasticsearch until 10:06 a.m., set this property to 120
	// seconds. The default
	// value is randomly selected between `60s` and `120s`. This randomness improves
	// the query performance
	// when there are multiple jobs running on the same node.
	QueryDelay *types.Duration `json:"query_delay,omitempty"`
	// RuntimeMappings Specifies runtime fields for the datafeed search.
	RuntimeMappings map[string]types.RuntimeField `json:"runtime_mappings,omitempty"`
	// ScriptFields Specifies scripts that evaluate custom expressions and returns script fields
	// to the datafeed.
	// The detector configuration objects in a job can contain functions that use
	// these script fields.
	ScriptFields map[string]types.ScriptField `json:"script_fields,omitempty"`
	// ScrollSize The size parameter that is used in Elasticsearch searches when the datafeed
	// does not use aggregations.
	// The maximum value is the value of `index.max_result_window`, which is 10,000
	// by default.
	ScrollSize *int `json:"scroll_size,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Aggregations: make(map[string]types.Aggregations, 0),
		ScriptFields: make(map[string]types.ScriptField, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putdatafeed request: %w", err)
	}

	return &req, nil
}
