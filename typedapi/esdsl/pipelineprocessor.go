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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pipelineProcessor struct {
	v *types.PipelineProcessor
}

// Executes another pipeline.
func NewPipelineProcessor() *_pipelineProcessor {

	return &_pipelineProcessor{v: types.NewPipelineProcessor()}

}

func (s *_pipelineProcessor) Description(description string) *_pipelineProcessor {

	s.v.Description = &description

	return s
}

func (s *_pipelineProcessor) If(if_ types.ScriptVariant) *_pipelineProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_pipelineProcessor) IgnoreFailure(ignorefailure bool) *_pipelineProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_pipelineProcessor) IgnoreMissingPipeline(ignoremissingpipeline bool) *_pipelineProcessor {

	s.v.IgnoreMissingPipeline = &ignoremissingpipeline

	return s
}

func (s *_pipelineProcessor) Name(name string) *_pipelineProcessor {

	s.v.Name = name

	return s
}

func (s *_pipelineProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_pipelineProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_pipelineProcessor) Tag(tag string) *_pipelineProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_pipelineProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Pipeline = s.v

	return container
}

func (s *_pipelineProcessor) PipelineProcessorCaster() *types.PipelineProcessor {
	return s.v
}
