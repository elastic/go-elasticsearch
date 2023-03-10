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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// MLDatafeed type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/Datafeed.ts#L37-L58
type MLDatafeed struct {
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`
	// Authorization The security privileges that the datafeed uses to run its queries. If Elastic
	// Stack security features were disabled at the time of the most recent update
	// to the datafeed, this property is omitted.
	Authorization          *DatafeedAuthorization  `json:"authorization,omitempty"`
	ChunkingConfig         *ChunkingConfig         `json:"chunking_config,omitempty"`
	DatafeedId             string                  `json:"datafeed_id"`
	DelayedDataCheckConfig DelayedDataCheckConfig  `json:"delayed_data_check_config"`
	Frequency              Duration                `json:"frequency,omitempty"`
	Indexes                []string                `json:"indexes,omitempty"`
	Indices                []string                `json:"indices"`
	IndicesOptions         *IndicesOptions         `json:"indices_options,omitempty"`
	JobId                  string                  `json:"job_id"`
	MaxEmptySearches       *int                    `json:"max_empty_searches,omitempty"`
	Query                  Query                   `json:"query"`
	QueryDelay             Duration                `json:"query_delay,omitempty"`
	RuntimeMappings        map[string]RuntimeField `json:"runtime_mappings,omitempty"`
	ScriptFields           map[string]ScriptField  `json:"script_fields,omitempty"`
	ScrollSize             *int                    `json:"scroll_size,omitempty"`
}

// NewMLDatafeed returns a MLDatafeed.
func NewMLDatafeed() *MLDatafeed {
	r := &MLDatafeed{
		Aggregations: make(map[string]Aggregations, 0),
		ScriptFields: make(map[string]ScriptField, 0),
	}

	return r
}
