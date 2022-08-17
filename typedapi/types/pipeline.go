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

// Pipeline type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/logstash/_types/Pipeline.ts#L37-L44
type Pipeline struct {
	Description      string           `json:"description"`
	LastModified     DateTime         `json:"last_modified"`
	Pipeline         string           `json:"pipeline"`
	PipelineMetadata PipelineMetadata `json:"pipeline_metadata"`
	PipelineSettings PipelineSettings `json:"pipeline_settings"`
	Username         string           `json:"username"`
}

// PipelineBuilder holds Pipeline struct and provides a builder API.
type PipelineBuilder struct {
	v *Pipeline
}

// NewPipeline provides a builder for the Pipeline struct.
func NewPipelineBuilder() *PipelineBuilder {
	r := PipelineBuilder{
		&Pipeline{},
	}

	return &r
}

// Build finalize the chain and returns the Pipeline struct
func (rb *PipelineBuilder) Build() Pipeline {
	return *rb.v
}

func (rb *PipelineBuilder) Description(description string) *PipelineBuilder {
	rb.v.Description = description
	return rb
}

func (rb *PipelineBuilder) LastModified(lastmodified *DateTimeBuilder) *PipelineBuilder {
	v := lastmodified.Build()
	rb.v.LastModified = v
	return rb
}

func (rb *PipelineBuilder) Pipeline(pipeline string) *PipelineBuilder {
	rb.v.Pipeline = pipeline
	return rb
}

func (rb *PipelineBuilder) PipelineMetadata(pipelinemetadata *PipelineMetadataBuilder) *PipelineBuilder {
	v := pipelinemetadata.Build()
	rb.v.PipelineMetadata = v
	return rb
}

func (rb *PipelineBuilder) PipelineSettings(pipelinesettings *PipelineSettingsBuilder) *PipelineBuilder {
	v := pipelinesettings.Build()
	rb.v.PipelineSettings = v
	return rb
}

func (rb *PipelineBuilder) Username(username string) *PipelineBuilder {
	rb.v.Username = username
	return rb
}
