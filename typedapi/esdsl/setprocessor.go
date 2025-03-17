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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _setProcessor struct {
	v *types.SetProcessor
}

// Adds a field with the specified value.
// If the field already exists, its value will be replaced with the provided
// one.
func NewSetProcessor() *_setProcessor {

	return &_setProcessor{v: types.NewSetProcessor()}

}

// The origin field which will be copied to `field`, cannot set `value`
// simultaneously.
// Supported data types are `boolean`, `number`, `array`, `object`, `string`,
// `date`, etc.
func (s *_setProcessor) CopyFrom(field string) *_setProcessor {

	s.v.CopyFrom = &field

	return s
}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_setProcessor) Description(description string) *_setProcessor {

	s.v.Description = &description

	return s
}

// The field to insert, upsert, or update.
// Supports template snippets.
func (s *_setProcessor) Field(field string) *_setProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_setProcessor) If(if_ types.ScriptVariant) *_setProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// If `true` and `value` is a template snippet that evaluates to `null` or the
// empty string, the processor quietly exits without modifying the document.
func (s *_setProcessor) IgnoreEmptyValue(ignoreemptyvalue bool) *_setProcessor {

	s.v.IgnoreEmptyValue = &ignoreemptyvalue

	return s
}

// Ignore failures for the processor.
func (s *_setProcessor) IgnoreFailure(ignorefailure bool) *_setProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// The media type for encoding `value`.
// Applies only when value is a template snippet.
// Must be one of `application/json`, `text/plain`, or
// `application/x-www-form-urlencoded`.
func (s *_setProcessor) MediaType(mediatype string) *_setProcessor {

	s.v.MediaType = &mediatype

	return s
}

// Handle failures for the processor.
func (s *_setProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_setProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// If `true` processor will update fields with pre-existing non-null-valued
// field.
// When set to `false`, such fields will not be touched.
func (s *_setProcessor) Override(override bool) *_setProcessor {

	s.v.Override = &override

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_setProcessor) Tag(tag string) *_setProcessor {

	s.v.Tag = &tag

	return s
}

// The value to be set for the field.
// Supports template snippets.
// May specify only one of `value` or `copy_from`.
func (s *_setProcessor) Value(value json.RawMessage) *_setProcessor {

	s.v.Value = value

	return s
}

func (s *_setProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Set = s.v

	return container
}

func (s *_setProcessor) SetProcessorCaster() *types.SetProcessor {
	return s.v
}
