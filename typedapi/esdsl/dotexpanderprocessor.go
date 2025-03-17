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

type _dotExpanderProcessor struct {
	v *types.DotExpanderProcessor
}

// Expands a field with dots into an object field.
// This processor allows fields with dots in the name to be accessible by other
// processors in the pipeline.
// Otherwise these fields can’t be accessed by any processor.
func NewDotExpanderProcessor() *_dotExpanderProcessor {

	return &_dotExpanderProcessor{v: types.NewDotExpanderProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_dotExpanderProcessor) Description(description string) *_dotExpanderProcessor {

	s.v.Description = &description

	return s
}

// The field to expand into an object field.
// If set to `*`, all top-level fields will be expanded.
func (s *_dotExpanderProcessor) Field(field string) *_dotExpanderProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_dotExpanderProcessor) If(if_ types.ScriptVariant) *_dotExpanderProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_dotExpanderProcessor) IgnoreFailure(ignorefailure bool) *_dotExpanderProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// Handle failures for the processor.
func (s *_dotExpanderProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dotExpanderProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Controls the behavior when there is already an existing nested object that
// conflicts with the expanded field.
// When `false`, the processor will merge conflicts by combining the old and the
// new values into an array.
// When `true`, the value from the expanded field will overwrite the existing
// value.
func (s *_dotExpanderProcessor) Override(override bool) *_dotExpanderProcessor {

	s.v.Override = &override

	return s
}

// The field that contains the field to expand.
// Only required if the field to expand is part another object field, because
// the `field` option can only understand leaf fields.
func (s *_dotExpanderProcessor) Path(path string) *_dotExpanderProcessor {

	s.v.Path = &path

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_dotExpanderProcessor) Tag(tag string) *_dotExpanderProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_dotExpanderProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.DotExpander = s.v

	return container
}

func (s *_dotExpanderProcessor) DotExpanderProcessorCaster() *types.DotExpanderProcessor {
	return s.v
}
