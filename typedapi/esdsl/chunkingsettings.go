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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _chunkingSettings struct {
	v *types.ChunkingSettings
}

func NewChunkingSettings(maxchunksize int, strategy string) *_chunkingSettings {

	tmp := &_chunkingSettings{v: types.NewChunkingSettings()}

	tmp.MaxChunkSize(maxchunksize)

	tmp.Strategy(strategy)

	return tmp

}

func (s *_chunkingSettings) MaxChunkSize(maxchunksize int) *_chunkingSettings {

	s.v.MaxChunkSize = maxchunksize

	return s
}

func (s *_chunkingSettings) Overlap(overlap int) *_chunkingSettings {

	s.v.Overlap = &overlap

	return s
}

func (s *_chunkingSettings) SentenceOverlap(sentenceoverlap int) *_chunkingSettings {

	s.v.SentenceOverlap = &sentenceoverlap

	return s
}

func (s *_chunkingSettings) Strategy(strategy string) *_chunkingSettings {

	s.v.Strategy = strategy

	return s
}

func (s *_chunkingSettings) ChunkingSettingsCaster() *types.ChunkingSettings {
	return s.v
}
