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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/chunkingmode"
)

// ChunkingConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/ml/_types/Datafeed.ts#L164-L177
type ChunkingConfig struct {
	// Mode If the mode is `auto`, the chunk size is dynamically calculated;
	// this is the recommended value when the datafeed does not use aggregations.
	// If the mode is `manual`, chunking is applied according to the specified
	// `time_span`;
	// use this mode when the datafeed uses aggregations. If the mode is `off`, no
	// chunking is applied.
	Mode chunkingmode.ChunkingMode `json:"mode"`
	// TimeSpan The time span that each search will be querying. This setting is applicable
	// only when the `mode` is set to `manual`.
	TimeSpan *Time `json:"time_span,omitempty"`
}

// ChunkingConfigBuilder holds ChunkingConfig struct and provides a builder API.
type ChunkingConfigBuilder struct {
	v *ChunkingConfig
}

// NewChunkingConfig provides a builder for the ChunkingConfig struct.
func NewChunkingConfigBuilder() *ChunkingConfigBuilder {
	r := ChunkingConfigBuilder{
		&ChunkingConfig{},
	}

	return &r
}

// Build finalize the chain and returns the ChunkingConfig struct
func (rb *ChunkingConfigBuilder) Build() ChunkingConfig {
	return *rb.v
}

// Mode If the mode is `auto`, the chunk size is dynamically calculated;
// this is the recommended value when the datafeed does not use aggregations.
// If the mode is `manual`, chunking is applied according to the specified
// `time_span`;
// use this mode when the datafeed uses aggregations. If the mode is `off`, no
// chunking is applied.

func (rb *ChunkingConfigBuilder) Mode(mode chunkingmode.ChunkingMode) *ChunkingConfigBuilder {
	rb.v.Mode = mode
	return rb
}

// TimeSpan The time span that each search will be querying. This setting is applicable
// only when the `mode` is set to `manual`.

func (rb *ChunkingConfigBuilder) TimeSpan(timespan *TimeBuilder) *ChunkingConfigBuilder {
	v := timespan.Build()
	rb.v.TimeSpan = &v
	return rb
}
