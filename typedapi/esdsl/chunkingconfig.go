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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/chunkingmode"
)

type _chunkingConfig struct {
	v *types.ChunkingConfig
}

func NewChunkingConfig(mode chunkingmode.ChunkingMode) *_chunkingConfig {

	tmp := &_chunkingConfig{v: types.NewChunkingConfig()}

	tmp.Mode(mode)

	return tmp

}

// If the mode is `auto`, the chunk size is dynamically calculated;
// this is the recommended value when the datafeed does not use aggregations.
// If the mode is `manual`, chunking is applied according to the specified
// `time_span`;
// use this mode when the datafeed uses aggregations. If the mode is `off`, no
// chunking is applied.
func (s *_chunkingConfig) Mode(mode chunkingmode.ChunkingMode) *_chunkingConfig {

	s.v.Mode = mode
	return s
}

// The time span that each search will be querying. This setting is applicable
// only when the `mode` is set to `manual`.
func (s *_chunkingConfig) TimeSpan(duration types.DurationVariant) *_chunkingConfig {

	s.v.TimeSpan = *duration.DurationCaster()

	return s
}

func (s *_chunkingConfig) ChunkingConfigCaster() *types.ChunkingConfig {
	return s.v
}
