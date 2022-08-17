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

// PipelineProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L297-L299
type PipelineProcessor struct {
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	Name          Name                 `json:"name"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// PipelineProcessorBuilder holds PipelineProcessor struct and provides a builder API.
type PipelineProcessorBuilder struct {
	v *PipelineProcessor
}

// NewPipelineProcessor provides a builder for the PipelineProcessor struct.
func NewPipelineProcessorBuilder() *PipelineProcessorBuilder {
	r := PipelineProcessorBuilder{
		&PipelineProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the PipelineProcessor struct
func (rb *PipelineProcessorBuilder) Build() PipelineProcessor {
	return *rb.v
}

func (rb *PipelineProcessorBuilder) If_(if_ string) *PipelineProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *PipelineProcessorBuilder) IgnoreFailure(ignorefailure bool) *PipelineProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *PipelineProcessorBuilder) Name(name Name) *PipelineProcessorBuilder {
	rb.v.Name = name
	return rb
}

func (rb *PipelineProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *PipelineProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *PipelineProcessorBuilder) Tag(tag string) *PipelineProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
