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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _terminateProcessor struct {
	v *types.TerminateProcessor
}

// Terminates the current ingest pipeline, causing no further processors to be
// run.
// This will normally be executed conditionally, using the `if` option.
func NewTerminateProcessor() *_terminateProcessor {

	return &_terminateProcessor{v: types.NewTerminateProcessor()}

}

func (s *_terminateProcessor) Description(description string) *_terminateProcessor {

	s.v.Description = &description

	return s
}

func (s *_terminateProcessor) If(if_ types.ScriptVariant) *_terminateProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_terminateProcessor) IgnoreFailure(ignorefailure bool) *_terminateProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_terminateProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_terminateProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_terminateProcessor) Tag(tag string) *_terminateProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_terminateProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Terminate = s.v

	return container
}

func (s *_terminateProcessor) TerminateProcessorCaster() *types.TerminateProcessor {
	return s.v
}
