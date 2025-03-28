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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _lowercaseProcessor struct {
	v *types.LowercaseProcessor
}

// Converts a string to its lowercase equivalent.
// If the field is an array of strings, all members of the array will be
// converted.
func NewLowercaseProcessor() *_lowercaseProcessor {

	return &_lowercaseProcessor{v: types.NewLowercaseProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_lowercaseProcessor) Description(description string) *_lowercaseProcessor {

	s.v.Description = &description

	return s
}

// The field to make lowercase.
func (s *_lowercaseProcessor) Field(field string) *_lowercaseProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_lowercaseProcessor) If(if_ types.ScriptVariant) *_lowercaseProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_lowercaseProcessor) IgnoreFailure(ignorefailure bool) *_lowercaseProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist or is `null`, the processor quietly
// exits without modifying the document.
func (s *_lowercaseProcessor) IgnoreMissing(ignoremissing bool) *_lowercaseProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_lowercaseProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_lowercaseProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_lowercaseProcessor) Tag(tag string) *_lowercaseProcessor {

	s.v.Tag = &tag

	return s
}

// The field to assign the converted value to.
// By default, the field is updated in-place.
func (s *_lowercaseProcessor) TargetField(field string) *_lowercaseProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_lowercaseProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Lowercase = s.v

	return container
}

func (s *_lowercaseProcessor) LowercaseProcessorCaster() *types.LowercaseProcessor {
	return s.v
}
