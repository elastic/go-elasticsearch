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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _inferenceChunkingSettings struct {
	v *types.InferenceChunkingSettings
}

func NewInferenceChunkingSettings(service string) *_inferenceChunkingSettings {

	tmp := &_inferenceChunkingSettings{v: types.NewInferenceChunkingSettings()}

	tmp.Service(service)

	return tmp

}

// Chunking configuration object
func (s *_inferenceChunkingSettings) ChunkingSettings(chunkingsettings types.InferenceChunkingSettingsVariant) *_inferenceChunkingSettings {

	s.v.ChunkingSettings = chunkingsettings.InferenceChunkingSettingsCaster()

	return s
}

// The maximum size of a chunk in words.
// This value cannot be higher than `300` or lower than `20` (for `sentence`
// strategy) or `10` (for `word` strategy).
func (s *_inferenceChunkingSettings) MaxChunkSize(maxchunksize int) *_inferenceChunkingSettings {

	s.v.MaxChunkSize = &maxchunksize

	return s
}

// The number of overlapping words for chunks.
// It is applicable only to a `word` chunking strategy.
// This value cannot be higher than half the `max_chunk_size` value.
func (s *_inferenceChunkingSettings) Overlap(overlap int) *_inferenceChunkingSettings {

	s.v.Overlap = &overlap

	return s
}

// The number of overlapping sentences for chunks.
// It is applicable only for a `sentence` chunking strategy.
// It can be either `1` or `0`.
func (s *_inferenceChunkingSettings) SentenceOverlap(sentenceoverlap int) *_inferenceChunkingSettings {

	s.v.SentenceOverlap = &sentenceoverlap

	return s
}

// The service type
func (s *_inferenceChunkingSettings) Service(service string) *_inferenceChunkingSettings {

	s.v.Service = service

	return s
}

// Settings specific to the service
func (s *_inferenceChunkingSettings) ServiceSettings(servicesettings json.RawMessage) *_inferenceChunkingSettings {

	s.v.ServiceSettings = servicesettings

	return s
}

// The chunking strategy: `sentence` or `word`.
func (s *_inferenceChunkingSettings) Strategy(strategy string) *_inferenceChunkingSettings {

	s.v.Strategy = &strategy

	return s
}

// Task settings specific to the service and task type
func (s *_inferenceChunkingSettings) TaskSettings(tasksettings json.RawMessage) *_inferenceChunkingSettings {

	s.v.TaskSettings = tasksettings

	return s
}

func (s *_inferenceChunkingSettings) InferenceChunkingSettingsCaster() *types.InferenceChunkingSettings {
	return s.v
}
