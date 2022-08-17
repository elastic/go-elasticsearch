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

// PipelineConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Pipeline.ts#L44-L48
type PipelineConfig struct {
	Description *string              `json:"description,omitempty"`
	Processors  []ProcessorContainer `json:"processors"`
	Version     *VersionNumber       `json:"version,omitempty"`
}

// PipelineConfigBuilder holds PipelineConfig struct and provides a builder API.
type PipelineConfigBuilder struct {
	v *PipelineConfig
}

// NewPipelineConfig provides a builder for the PipelineConfig struct.
func NewPipelineConfigBuilder() *PipelineConfigBuilder {
	r := PipelineConfigBuilder{
		&PipelineConfig{},
	}

	return &r
}

// Build finalize the chain and returns the PipelineConfig struct
func (rb *PipelineConfigBuilder) Build() PipelineConfig {
	return *rb.v
}

func (rb *PipelineConfigBuilder) Description(description string) *PipelineConfigBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *PipelineConfigBuilder) Processors(processors []ProcessorContainerBuilder) *PipelineConfigBuilder {
	tmp := make([]ProcessorContainer, len(processors))
	for _, value := range processors {
		tmp = append(tmp, value.Build())
	}
	rb.v.Processors = tmp
	return rb
}

func (rb *PipelineConfigBuilder) Version(version VersionNumber) *PipelineConfigBuilder {
	rb.v.Version = &version
	return rb
}
