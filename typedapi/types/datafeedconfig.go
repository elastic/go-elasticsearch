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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DatafeedConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/Datafeed.ts#L60-L117
type DatafeedConfig struct {
	// Aggregations If set, the datafeed performs aggregation searches. Support for aggregations
	// is limited and should be used only with low cardinality data.
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`
	// ChunkingConfig Datafeeds might be required to search over long time periods, for several
	// months or years. This search is split into time chunks in order to ensure the
	// load on Elasticsearch is managed. Chunking configuration controls how the
	// size of these time chunks are calculated and is an advanced configuration
	// option.
	ChunkingConfig *ChunkingConfig `json:"chunking_config,omitempty"`
	// DatafeedId A numerical character string that uniquely identifies the datafeed. This
	// identifier can contain lowercase alphanumeric characters (a-z and 0-9),
	// hyphens, and underscores. It must start and end with alphanumeric characters.
	// The default value is the job identifier.
	DatafeedId *string `json:"datafeed_id,omitempty"`
	// DelayedDataCheckConfig Specifies whether the datafeed checks for missing data and the size of the
	// window. The datafeed can optionally search over indices that have already
	// been read in an effort to determine whether any data has subsequently been
	// added to the index. If missing data is found, it is a good indication that
	// the `query_delay` option is set too low and the data is being indexed after
	// the datafeed has passed that moment in time. This check runs only on
	// real-time datafeeds.
	DelayedDataCheckConfig *DelayedDataCheckConfig `json:"delayed_data_check_config,omitempty"`
	// Frequency The interval at which scheduled queries are made while the datafeed runs in
	// real time. The default value is either the bucket span for short bucket
	// spans, or, for longer bucket spans, a sensible fraction of the bucket span.
	// For example: `150s`. When `frequency` is shorter than the bucket span,
	// interim results for the last (partial) bucket are written then eventually
	// overwritten by the full bucket results. If the datafeed uses aggregations,
	// this value must be divisible by the interval of the date histogram
	// aggregation.
	Frequency Duration `json:"frequency,omitempty"`
	// Indices An array of index names. Wildcards are supported. If any indices are in
	// remote clusters, the machine learning nodes must have the
	// `remote_cluster_client` role.
	Indices []string `json:"indices,omitempty"`
	// IndicesOptions Specifies index expansion options that are used during search.
	IndicesOptions *IndicesOptions `json:"indices_options,omitempty"`
	JobId          *string         `json:"job_id,omitempty"`
	// MaxEmptySearches If a real-time datafeed has never seen any data (including during any initial
	// training period) then it will automatically stop itself and close its
	// associated job after this many real-time searches that return no documents.
	// In other words, it will stop after `frequency` times `max_empty_searches` of
	// real-time operation. If not set then a datafeed with no end time that sees no
	// data will remain started until it is explicitly stopped.
	MaxEmptySearches *int `json:"max_empty_searches,omitempty"`
	// Query The Elasticsearch query domain-specific language (DSL). This value
	// corresponds to the query object in an Elasticsearch search POST body. All the
	// options that are supported by Elasticsearch can be used, as this object is
	// passed verbatim to Elasticsearch.
	Query *Query `json:"query,omitempty"`
	// QueryDelay The number of seconds behind real time that data is queried. For example, if
	// data from 10:04 a.m. might not be searchable in Elasticsearch until 10:06
	// a.m., set this property to 120 seconds. The default value is randomly
	// selected between `60s` and `120s`. This randomness improves the query
	// performance when there are multiple jobs running on the same node.
	QueryDelay Duration `json:"query_delay,omitempty"`
	// RuntimeMappings Specifies runtime fields for the datafeed search.
	RuntimeMappings RuntimeFields `json:"runtime_mappings,omitempty"`
	// ScriptFields Specifies scripts that evaluate custom expressions and returns script fields
	// to the datafeed. The detector configuration objects in a job can contain
	// functions that use these script fields.
	ScriptFields map[string]ScriptField `json:"script_fields,omitempty"`
	// ScrollSize The size parameter that is used in Elasticsearch searches when the datafeed
	// does not use aggregations. The maximum value is the value of
	// `index.max_result_window`, which is 10,000 by default.
	ScrollSize *int `json:"scroll_size,omitempty"`
}

func (s *DatafeedConfig) UnmarshalJSON(data []byte) error {

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

		case "aggregations", "aggs":
			if s.Aggregations == nil {
				s.Aggregations = make(map[string]Aggregations, 0)
			}
			if err := dec.Decode(&s.Aggregations); err != nil {
				return err
			}

		case "chunking_config":
			if err := dec.Decode(&s.ChunkingConfig); err != nil {
				return err
			}

		case "datafeed_id":
			if err := dec.Decode(&s.DatafeedId); err != nil {
				return err
			}

		case "delayed_data_check_config":
			if err := dec.Decode(&s.DelayedDataCheckConfig); err != nil {
				return err
			}

		case "frequency":
			if err := dec.Decode(&s.Frequency); err != nil {
				return err
			}

		case "indices", "indexes":
			if err := dec.Decode(&s.Indices); err != nil {
				return err
			}

		case "indices_options":
			if err := dec.Decode(&s.IndicesOptions); err != nil {
				return err
			}

		case "job_id":
			if err := dec.Decode(&s.JobId); err != nil {
				return err
			}

		case "max_empty_searches":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxEmptySearches = &value
			case float64:
				f := int(v)
				s.MaxEmptySearches = &f
			}

		case "query":
			if err := dec.Decode(&s.Query); err != nil {
				return err
			}

		case "query_delay":
			if err := dec.Decode(&s.QueryDelay); err != nil {
				return err
			}

		case "runtime_mappings":
			if err := dec.Decode(&s.RuntimeMappings); err != nil {
				return err
			}

		case "script_fields":
			if s.ScriptFields == nil {
				s.ScriptFields = make(map[string]ScriptField, 0)
			}
			if err := dec.Decode(&s.ScriptFields); err != nil {
				return err
			}

		case "scroll_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ScrollSize = &value
			case float64:
				f := int(v)
				s.ScrollSize = &f
			}

		}
	}
	return nil
}

// NewDatafeedConfig returns a DatafeedConfig.
func NewDatafeedConfig() *DatafeedConfig {
	r := &DatafeedConfig{
		Aggregations: make(map[string]Aggregations, 0),
		ScriptFields: make(map[string]ScriptField, 0),
	}

	return r
}
