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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// PipelineSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/logstash/_types/Pipeline.ts#L28-L36
type PipelineSettings struct {
	PipelineBatchDelay    int    `json:"pipeline.batch.delay"`
	PipelineBatchSize     int    `json:"pipeline.batch.size"`
	PipelineWorkers       int    `json:"pipeline.workers"`
	QueueCheckpointWrites int    `json:"queue.checkpoint.writes"`
	QueueMaxBytesNumber   int    `json:"queue.max_bytes.number"`
	QueueMaxBytesUnits    string `json:"queue.max_bytes.units"`
	QueueType             string `json:"queue.type"`
}

// PipelineSettingsBuilder holds PipelineSettings struct and provides a builder API.
type PipelineSettingsBuilder struct {
	v *PipelineSettings
}

// NewPipelineSettings provides a builder for the PipelineSettings struct.
func NewPipelineSettingsBuilder() *PipelineSettingsBuilder {
	r := PipelineSettingsBuilder{
		&PipelineSettings{},
	}

	return &r
}

// Build finalize the chain and returns the PipelineSettings struct
func (rb *PipelineSettingsBuilder) Build() PipelineSettings {
	return *rb.v
}

func (rb *PipelineSettingsBuilder) PipelineBatchDelay(pipelinebatchdelay int) *PipelineSettingsBuilder {
	rb.v.PipelineBatchDelay = pipelinebatchdelay
	return rb
}

func (rb *PipelineSettingsBuilder) PipelineBatchSize(pipelinebatchsize int) *PipelineSettingsBuilder {
	rb.v.PipelineBatchSize = pipelinebatchsize
	return rb
}

func (rb *PipelineSettingsBuilder) PipelineWorkers(pipelineworkers int) *PipelineSettingsBuilder {
	rb.v.PipelineWorkers = pipelineworkers
	return rb
}

func (rb *PipelineSettingsBuilder) QueueCheckpointWrites(queuecheckpointwrites int) *PipelineSettingsBuilder {
	rb.v.QueueCheckpointWrites = queuecheckpointwrites
	return rb
}

func (rb *PipelineSettingsBuilder) QueueMaxBytesNumber(queuemaxbytesnumber int) *PipelineSettingsBuilder {
	rb.v.QueueMaxBytesNumber = queuemaxbytesnumber
	return rb
}

func (rb *PipelineSettingsBuilder) QueueMaxBytesUnits(queuemaxbytesunits string) *PipelineSettingsBuilder {
	rb.v.QueueMaxBytesUnits = queuemaxbytesunits
	return rb
}

func (rb *PipelineSettingsBuilder) QueueType(queuetype string) *PipelineSettingsBuilder {
	rb.v.QueueType = queuetype
	return rb
}
