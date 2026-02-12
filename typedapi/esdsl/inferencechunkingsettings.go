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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inferenceChunkingSettings struct {
	v *types.InferenceChunkingSettings
}

func NewInferenceChunkingSettings() *_inferenceChunkingSettings {

	return &_inferenceChunkingSettings{v: types.NewInferenceChunkingSettings()}

}

func (s *_inferenceChunkingSettings) MaxChunkSize(maxchunksize int) *_inferenceChunkingSettings {

	s.v.MaxChunkSize = &maxchunksize

	return s
}

func (s *_inferenceChunkingSettings) Overlap(overlap int) *_inferenceChunkingSettings {

	s.v.Overlap = &overlap

	return s
}

func (s *_inferenceChunkingSettings) SentenceOverlap(sentenceoverlap int) *_inferenceChunkingSettings {

	s.v.SentenceOverlap = &sentenceoverlap

	return s
}

func (s *_inferenceChunkingSettings) SeparatorGroup(separatorgroup string) *_inferenceChunkingSettings {

	s.v.SeparatorGroup = &separatorgroup

	return s
}

func (s *_inferenceChunkingSettings) Separators(separators ...string) *_inferenceChunkingSettings {

	for _, v := range separators {

		s.v.Separators = append(s.v.Separators, v)

	}
	return s
}

func (s *_inferenceChunkingSettings) Strategy(strategy string) *_inferenceChunkingSettings {

	s.v.Strategy = &strategy

	return s
}

func (s *_inferenceChunkingSettings) InferenceChunkingSettingsCaster() *types.InferenceChunkingSettings {
	return s.v
}
