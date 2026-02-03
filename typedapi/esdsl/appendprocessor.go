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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
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

func (s *_appendProcessor) AllowDuplicates(allowduplicates bool) *_appendProcessor {

	s.v.AllowDuplicates = &allowduplicates

	return s
}

func (s *_appendProcessor) CopyFrom(field string) *_appendProcessor {

	s.v.CopyFrom = &field

	return s
}

func (s *_appendProcessor) Field(field string) *_appendProcessor {

	s.v.Field = field

	return s
}

func (s *_appendProcessor) Value(values ...json.RawMessage) *_appendProcessor {

	s.v.Value = make([]json.RawMessage, len(values))
	s.v.Value = values

	return s
}

func (s *_appendProcessor) Description(description string) *_appendProcessor {

	s.v.Description = &description

	return s
}

func (s *_appendProcessor) If(if_ types.ScriptVariant) *_appendProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_appendProcessor) IgnoreFailure(ignorefailure bool) *_appendProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_appendProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_appendProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_appendProcessor) Tag(tag string) *_appendProcessor {

	s.v.Tag = &tag

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
