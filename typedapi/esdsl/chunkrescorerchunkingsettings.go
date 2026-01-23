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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _chunkRescorerChunkingSettings struct {
	v *types.ChunkRescorerChunkingSettings
}

func NewChunkRescorerChunkingSettings(maxchunksize int) *_chunkRescorerChunkingSettings {

	tmp := &_chunkRescorerChunkingSettings{v: types.NewChunkRescorerChunkingSettings()}

	tmp.MaxChunkSize(maxchunksize)

	return tmp

}

func (s *_chunkRescorerChunkingSettings) MaxChunkSize(maxchunksize int) *_chunkRescorerChunkingSettings {

	s.v.MaxChunkSize = maxchunksize

	return s
}

func (s *_chunkRescorerChunkingSettings) Overlap(overlap int) *_chunkRescorerChunkingSettings {

	s.v.Overlap = &overlap

	return s
}

func (s *_chunkRescorerChunkingSettings) SentenceOverlap(sentenceoverlap int) *_chunkRescorerChunkingSettings {

	s.v.SentenceOverlap = &sentenceoverlap

	return s
}

func (s *_chunkRescorerChunkingSettings) SeparatorGroup(separatorgroup string) *_chunkRescorerChunkingSettings {

	s.v.SeparatorGroup = &separatorgroup

	return s
}

func (s *_chunkRescorerChunkingSettings) Separators(separators ...string) *_chunkRescorerChunkingSettings {

	for _, v := range separators {

		s.v.Separators = append(s.v.Separators, v)

	}
	return s
}

func (s *_chunkRescorerChunkingSettings) Strategy(strategy string) *_chunkRescorerChunkingSettings {

	s.v.Strategy = &strategy

	return s
}

func (s *_chunkRescorerChunkingSettings) ChunkRescorerChunkingSettingsCaster() *types.ChunkRescorerChunkingSettings {
	return s.v
}
