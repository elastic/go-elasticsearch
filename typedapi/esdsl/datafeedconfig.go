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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _datafeedConfig struct {
	v *types.DatafeedConfig
}

func NewDatafeedConfig() *_datafeedConfig {

	return &_datafeedConfig{v: types.NewDatafeedConfig()}

}

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

func (s *_datafeedConfig) ChunkingConfig(chunkingconfig types.ChunkingConfigVariant) *_datafeedConfig {

	s.v.ChunkingConfig = chunkingconfig.ChunkingConfigCaster()

	return s
}

func (s *_datafeedConfig) DatafeedId(id string) *_datafeedConfig {

	s.v.DatafeedId = &id

	return s
}

func (s *_datafeedConfig) DelayedDataCheckConfig(delayeddatacheckconfig types.DelayedDataCheckConfigVariant) *_datafeedConfig {

	s.v.DelayedDataCheckConfig = delayeddatacheckconfig.DelayedDataCheckConfigCaster()

	return s
}

func (s *_datafeedConfig) Frequency(duration types.DurationVariant) *_datafeedConfig {

	s.v.Frequency = *duration.DurationCaster()

	return s
}

func (s *_datafeedConfig) Indices(indices ...string) *_datafeedConfig {

	s.v.Indices = indices

	return s
}

func (s *_datafeedConfig) IndicesOptions(indicesoptions types.IndicesOptionsVariant) *_datafeedConfig {

	s.v.IndicesOptions = indicesoptions.IndicesOptionsCaster()

	return s
}

func (s *_datafeedConfig) JobId(id string) *_datafeedConfig {

	s.v.JobId = &id

	return s
}

func (s *_datafeedConfig) MaxEmptySearches(maxemptysearches int) *_datafeedConfig {

	s.v.MaxEmptySearches = &maxemptysearches

	return s
}

func (s *_datafeedConfig) Query(query types.QueryVariant) *_datafeedConfig {

	s.v.Query = query.QueryCaster()

	return s
}

func (s *_datafeedConfig) QueryDelay(duration types.DurationVariant) *_datafeedConfig {

	s.v.QueryDelay = *duration.DurationCaster()

	return s
}

func (s *_datafeedConfig) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *_datafeedConfig {

	s.v.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return s
}

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

func (s *_datafeedConfig) ScrollSize(scrollsize int) *_datafeedConfig {

	s.v.ScrollSize = &scrollsize

	return s
}

func (s *_datafeedConfig) DatafeedConfigCaster() *types.DatafeedConfig {
	return s.v
}
