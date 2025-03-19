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

// When creating pipeline event batches, how long in milliseconds to wait for
// each event before dispatching an undersized batch to pipeline workers.
func (s *_pipelineSettings) PipelineBatchDelay(pipelinebatchdelay int) *_pipelineSettings {

	s.v.PipelineBatchDelay = pipelinebatchdelay

	return s
}

// The maximum number of events an individual worker thread will collect from
// inputs before attempting to execute its filters and outputs.
func (s *_pipelineSettings) PipelineBatchSize(pipelinebatchsize int) *_pipelineSettings {

	s.v.PipelineBatchSize = pipelinebatchsize

	return s
}

// The number of workers that will, in parallel, execute the filter and output
// stages of the pipeline.
func (s *_pipelineSettings) PipelineWorkers(pipelineworkers int) *_pipelineSettings {

	s.v.PipelineWorkers = pipelineworkers

	return s
}

// The maximum number of written events before forcing a checkpoint when
// persistent queues are enabled (`queue.type: persisted`).
func (s *_pipelineSettings) QueueCheckpointWrites(queuecheckpointwrites int) *_pipelineSettings {

	s.v.QueueCheckpointWrites = queuecheckpointwrites

	return s
}

// The total capacity of the queue (`queue.type: persisted`) in number of bytes.
func (s *_pipelineSettings) QueueMaxBytesNumber(queuemaxbytesnumber int) *_pipelineSettings {

	s.v.QueueMaxBytesNumber = queuemaxbytesnumber

	return s
}

// The total capacity of the queue (`queue.type: persisted`) in terms of units
// of bytes.
func (s *_pipelineSettings) QueueMaxBytesUnits(queuemaxbytesunits string) *_pipelineSettings {

	s.v.QueueMaxBytesUnits = queuemaxbytesunits

	return s
}

// The internal queuing model to use for event buffering.
func (s *_pipelineSettings) QueueType(queuetype string) *_pipelineSettings {

	s.v.QueueType = queuetype

	return s
}

func (s *_pipelineSettings) PipelineSettingsCaster() *types.PipelineSettings {
	return s.v
}
