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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actionstatusoptions"
)

// PipelineSimulation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/simulate/types.ts#L31-L37
type PipelineSimulation struct {
	Doc              *DocumentSimulation                      `json:"doc,omitempty"`
	ProcessorResults []PipelineSimulation                     `json:"processor_results,omitempty"`
	ProcessorType    *string                                  `json:"processor_type,omitempty"`
	Status           *actionstatusoptions.ActionStatusOptions `json:"status,omitempty"`
	Tag              *string                                  `json:"tag,omitempty"`
}

// PipelineSimulationBuilder holds PipelineSimulation struct and provides a builder API.
type PipelineSimulationBuilder struct {
	v *PipelineSimulation
}

// NewPipelineSimulation provides a builder for the PipelineSimulation struct.
func NewPipelineSimulationBuilder() *PipelineSimulationBuilder {
	r := PipelineSimulationBuilder{
		&PipelineSimulation{},
	}

	return &r
}

// Build finalize the chain and returns the PipelineSimulation struct
func (rb *PipelineSimulationBuilder) Build() PipelineSimulation {
	return *rb.v
}

func (rb *PipelineSimulationBuilder) Doc(doc *DocumentSimulationBuilder) *PipelineSimulationBuilder {
	v := doc.Build()
	rb.v.Doc = &v
	return rb
}

func (rb *PipelineSimulationBuilder) ProcessorResults(processor_results []PipelineSimulationBuilder) *PipelineSimulationBuilder {
	tmp := make([]PipelineSimulation, len(processor_results))
	for _, value := range processor_results {
		tmp = append(tmp, value.Build())
	}
	rb.v.ProcessorResults = tmp
	return rb
}

func (rb *PipelineSimulationBuilder) ProcessorType(processortype string) *PipelineSimulationBuilder {
	rb.v.ProcessorType = &processortype
	return rb
}

func (rb *PipelineSimulationBuilder) Status(status actionstatusoptions.ActionStatusOptions) *PipelineSimulationBuilder {
	rb.v.Status = &status
	return rb
}

func (rb *PipelineSimulationBuilder) Tag(tag string) *PipelineSimulationBuilder {
	rb.v.Tag = &tag
	return rb
}
