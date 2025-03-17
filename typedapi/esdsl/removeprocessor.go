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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _removeProcessor struct {
	v *types.RemoveProcessor
}

// Removes existing fields.
// If one field doesn’t exist, an exception will be thrown.
func NewRemoveProcessor() *_removeProcessor {

	return &_removeProcessor{v: types.NewRemoveProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_removeProcessor) Description(description string) *_removeProcessor {

	s.v.Description = &description

	return s
}

// Fields to be removed. Supports template snippets.
func (s *_removeProcessor) Field(fields ...string) *_removeProcessor {

	s.v.Field = fields

	return s
}

// Conditionally execute the processor.
func (s *_removeProcessor) If(if_ types.ScriptVariant) *_removeProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_removeProcessor) IgnoreFailure(ignorefailure bool) *_removeProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist or is `null`, the processor quietly
// exits without modifying the document.
func (s *_removeProcessor) IgnoreMissing(ignoremissing bool) *_removeProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Fields to be kept. When set, all fields other than those specified are
// removed.
func (s *_removeProcessor) Keep(fields ...string) *_removeProcessor {

	s.v.Keep = fields

	return s
}

// Handle failures for the processor.
func (s *_removeProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_removeProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_removeProcessor) Tag(tag string) *_removeProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_removeProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Remove = s.v

	return container
}

func (s *_removeProcessor) RemoveProcessorCaster() *types.RemoveProcessor {
	return s.v
}
