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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _datafeedConfig struct {
	v *types.DatafeedConfig
}

func NewDatafeedConfig() *_datafeedConfig {

	return &_datafeedConfig{v: types.NewDatafeedConfig()}

}

// If set, the datafeed performs aggregation searches. Support for aggregations
// is limited and should be used only with low cardinality data.
func (s *_datafeedConfig) Aggregations(aggregations map[string]types.Aggregations) *_datafeedConfig {

	s.v.Aggregations = aggregations
	return s
}

func (s *_datafeedConfig) AddAggregation(key string, value types.AggregationsVariant) *_datafeedConfig {

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

// Datafeeds might be required to search over long time periods, for several
// months or years. This search is split into time chunks in order to ensure the
// load on Elasticsearch is managed. Chunking configuration controls how the
// size of these time chunks are calculated and is an advanced configuration
// option.
func (s *_datafeedConfig) ChunkingConfig(chunkingconfig types.ChunkingConfigVariant) *_datafeedConfig {

	s.v.ChunkingConfig = chunkingconfig.ChunkingConfigCaster()

	return s
}

// A numerical character string that uniquely identifies the datafeed. This
// identifier can contain lowercase alphanumeric characters (a-z and 0-9),
// hyphens, and underscores. It must start and end with alphanumeric characters.
// The default value is the job identifier.
func (s *_datafeedConfig) DatafeedId(id string) *_datafeedConfig {

	s.v.DatafeedId = &id

	return s
}

// Specifies whether the datafeed checks for missing data and the size of the
// window. The datafeed can optionally search over indices that have already
// been read in an effort to determine whether any data has subsequently been
// added to the index. If missing data is found, it is a good indication that
// the `query_delay` option is set too low and the data is being indexed after
// the datafeed has passed that moment in time. This check runs only on
// real-time datafeeds.
func (s *_datafeedConfig) DelayedDataCheckConfig(delayeddatacheckconfig types.DelayedDataCheckConfigVariant) *_datafeedConfig {

	s.v.DelayedDataCheckConfig = delayeddatacheckconfig.DelayedDataCheckConfigCaster()

	return s
}

// The interval at which scheduled queries are made while the datafeed runs in
// real time. The default value is either the bucket span for short bucket
// spans, or, for longer bucket spans, a sensible fraction of the bucket span.
// For example: `150s`. When `frequency` is shorter than the bucket span,
// interim results for the last (partial) bucket are written then eventually
// overwritten by the full bucket results. If the datafeed uses aggregations,
// this value must be divisible by the interval of the date histogram
// aggregation.
func (s *_datafeedConfig) Frequency(duration types.DurationVariant) *_datafeedConfig {

	s.v.Frequency = *duration.DurationCaster()

	return s
}

// An array of index names. Wildcards are supported. If any indices are in
// remote clusters, the machine learning nodes must have the
// `remote_cluster_client` role.
func (s *_datafeedConfig) Indices(indices ...string) *_datafeedConfig {

	s.v.Indices = indices

	return s
}

// Specifies index expansion options that are used during search.
func (s *_datafeedConfig) IndicesOptions(indicesoptions types.IndicesOptionsVariant) *_datafeedConfig {

	s.v.IndicesOptions = indicesoptions.IndicesOptionsCaster()

	return s
}

func (s *_datafeedConfig) JobId(id string) *_datafeedConfig {

	s.v.JobId = &id

	return s
}

// If a real-time datafeed has never seen any data (including during any initial
// training period) then it will automatically stop itself and close its
// associated job after this many real-time searches that return no documents.
// In other words, it will stop after `frequency` times `max_empty_searches` of
// real-time operation. If not set then a datafeed with no end time that sees no
// data will remain started until it is explicitly stopped.
func (s *_datafeedConfig) MaxEmptySearches(maxemptysearches int) *_datafeedConfig {

	s.v.MaxEmptySearches = &maxemptysearches

	return s
}

// The Elasticsearch query domain-specific language (DSL). This value
// corresponds to the query object in an Elasticsearch search POST body. All the
// options that are supported by Elasticsearch can be used, as this object is
// passed verbatim to Elasticsearch.
func (s *_datafeedConfig) Query(query types.QueryVariant) *_datafeedConfig {

	s.v.Query = query.QueryCaster()

	return s
}

// The number of seconds behind real time that data is queried. For example, if
// data from 10:04 a.m. might not be searchable in Elasticsearch until 10:06
// a.m., set this property to 120 seconds. The default value is randomly
// selected between `60s` and `120s`. This randomness improves the query
// performance when there are multiple jobs running on the same node.
func (s *_datafeedConfig) QueryDelay(duration types.DurationVariant) *_datafeedConfig {

	s.v.QueryDelay = *duration.DurationCaster()

	return s
}

// Specifies runtime fields for the datafeed search.
func (s *_datafeedConfig) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_datafeedConfig {

	s.v.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return s
}

// Specifies scripts that evaluate custom expressions and returns script fields
// to the datafeed. The detector configuration objects in a job can contain
// functions that use these script fields.
func (s *_datafeedConfig) ScriptFields(scriptfields map[string]types.ScriptField) *_datafeedConfig {

	s.v.ScriptFields = scriptfields
	return s
}

func (s *_datafeedConfig) AddScriptField(key string, value types.ScriptFieldVariant) *_datafeedConfig {

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

// The size parameter that is used in Elasticsearch searches when the datafeed
// does not use aggregations. The maximum value is the value of
// `index.max_result_window`, which is 10,000 by default.
func (s *_datafeedConfig) ScrollSize(scrollsize int) *_datafeedConfig {

	s.v.ScrollSize = &scrollsize

	return s
}

func (s *_datafeedConfig) DatafeedConfigCaster() *types.DatafeedConfig {
	return s.v
}
