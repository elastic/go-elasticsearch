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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _pipelineSettings struct {
	v *types.PipelineSettings
}

func NewPipelineSettings(pipelinebatchdelay int, pipelinebatchsize int, pipelineworkers int, queuecheckpointwrites int, queuemaxbytesnumber int, queuemaxbytesunits string, queuetype string) *_pipelineSettings {

	tmp := &_pipelineSettings{v: types.NewPipelineSettings()}

	tmp.PipelineBatchDelay(pipelinebatchdelay)

	tmp.PipelineBatchSize(pipelinebatchsize)

	tmp.PipelineWorkers(pipelineworkers)

	tmp.QueueCheckpointWrites(queuecheckpointwrites)

	tmp.QueueMaxBytesNumber(queuemaxbytesnumber)

	tmp.QueueMaxBytesUnits(queuemaxbytesunits)

	tmp.QueueType(queuetype)

	return tmp

}

func (s *_pipelineSettings) PipelineBatchDelay(pipelinebatchdelay int) *_pipelineSettings {

	s.v.PipelineBatchDelay = pipelinebatchdelay

	return s
}

func (s *_pipelineSettings) PipelineBatchSize(pipelinebatchsize int) *_pipelineSettings {

	s.v.PipelineBatchSize = pipelinebatchsize

	return s
}

func (s *_pipelineSettings) PipelineWorkers(pipelineworkers int) *_pipelineSettings {

	s.v.PipelineWorkers = pipelineworkers

	return s
}

func (s *_pipelineSettings) QueueCheckpointWrites(queuecheckpointwrites int) *_pipelineSettings {

	s.v.QueueCheckpointWrites = queuecheckpointwrites

	return s
}

func (s *_pipelineSettings) QueueMaxBytesNumber(queuemaxbytesnumber int) *_pipelineSettings {

	s.v.QueueMaxBytesNumber = queuemaxbytesnumber

	return s
}

func (s *_pipelineSettings) QueueMaxBytesUnits(queuemaxbytesunits string) *_pipelineSettings {

	s.v.QueueMaxBytesUnits = queuemaxbytesunits

	return s
}

func (s *_pipelineSettings) QueueType(queuetype string) *_pipelineSettings {

	s.v.QueueType = queuetype

	return s
}

func (s *_pipelineSettings) PipelineSettingsCaster() *types.PipelineSettings {
	return s.v
}
