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

type _dissectProcessor struct {
	v *types.DissectProcessor
}

// Extracts structured fields out of a single text field by matching the text
// field against a delimiter-based pattern.
func NewDissectProcessor(pattern string) *_dissectProcessor {

	tmp := &_dissectProcessor{v: types.NewDissectProcessor()}

	tmp.Pattern(pattern)

	return tmp

}

// The character(s) that separate the appended fields.
func (s *_dissectProcessor) AppendSeparator(appendseparator string) *_dissectProcessor {

	s.v.AppendSeparator = &appendseparator

	return s
}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_dissectProcessor) Description(description string) *_dissectProcessor {

	s.v.Description = &description

	return s
}

// The field to dissect.
func (s *_dissectProcessor) Field(field string) *_dissectProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_dissectProcessor) If(if_ types.ScriptVariant) *_dissectProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_dissectProcessor) IgnoreFailure(ignorefailure bool) *_dissectProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist or is `null`, the processor quietly
// exits without modifying the document.
func (s *_dissectProcessor) IgnoreMissing(ignoremissing bool) *_dissectProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_dissectProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dissectProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// The pattern to apply to the field.
func (s *_dissectProcessor) Pattern(pattern string) *_dissectProcessor {

	s.v.Pattern = pattern

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_dissectProcessor) Tag(tag string) *_dissectProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_dissectProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Dissect = s.v

	return container
}

func (s *_dissectProcessor) DissectProcessorCaster() *types.DissectProcessor {
	return s.v
}
