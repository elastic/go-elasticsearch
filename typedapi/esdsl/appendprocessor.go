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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _appendProcessor struct {
	v *types.AppendProcessor
}

// Appends one or more values to an existing array if the field already exists
// and it is an array.
// Converts a scalar to an array and appends one or more values to it if the
// field exists and it is a scalar.
// Creates an array containing the provided values if the field doesn’t exist.
// Accepts a single value or an array of values.
func NewAppendProcessor() *_appendProcessor {

	return &_appendProcessor{v: types.NewAppendProcessor()}

}

// If `false`, the processor does not append values already present in the
// field.
func (s *_appendProcessor) AllowDuplicates(allowduplicates bool) *_appendProcessor {

	s.v.AllowDuplicates = &allowduplicates

	return s
}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_appendProcessor) Description(description string) *_appendProcessor {

	s.v.Description = &description

	return s
}

// The field to be appended to.
// Supports template snippets.
func (s *_appendProcessor) Field(field string) *_appendProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_appendProcessor) If(if_ types.ScriptVariant) *_appendProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_appendProcessor) IgnoreFailure(ignorefailure bool) *_appendProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// Handle failures for the processor.
func (s *_appendProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_appendProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_appendProcessor) Tag(tag string) *_appendProcessor {

	s.v.Tag = &tag

	return s
}

// The value to be appended. Supports template snippets.
func (s *_appendProcessor) Value(values ...json.RawMessage) *_appendProcessor {

	s.v.Value = make([]json.RawMessage, len(values))
	s.v.Value = values

	return s
}

func (s *_appendProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Append = s.v

	return container
}

func (s *_appendProcessor) AppendProcessorCaster() *types.AppendProcessor {
	return s.v
}
