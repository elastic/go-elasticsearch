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

type _dateProcessor struct {
	v *types.DateProcessor
}

// Parses dates from fields, and then uses the date or timestamp as the
// timestamp for the document.
func NewDateProcessor() *_dateProcessor {

	return &_dateProcessor{v: types.NewDateProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_dateProcessor) Description(description string) *_dateProcessor {

	s.v.Description = &description

	return s
}

// The field to get the date from.
func (s *_dateProcessor) Field(field string) *_dateProcessor {

	s.v.Field = field

	return s
}

// An array of the expected date formats.
// Can be a java time pattern or one of the following formats: ISO8601, UNIX,
// UNIX_MS, or TAI64N.
func (s *_dateProcessor) Formats(formats ...string) *_dateProcessor {

	for _, v := range formats {

		s.v.Formats = append(s.v.Formats, v)

	}
	return s
}

// Conditionally execute the processor.
func (s *_dateProcessor) If(if_ types.ScriptVariant) *_dateProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_dateProcessor) IgnoreFailure(ignorefailure bool) *_dateProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// The locale to use when parsing the date, relevant when parsing month names or
// week days.
// Supports template snippets.
func (s *_dateProcessor) Locale(locale string) *_dateProcessor {

	s.v.Locale = &locale

	return s
}

// Handle failures for the processor.
func (s *_dateProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dateProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// The format to use when writing the date to target_field. Must be a valid
// java time pattern.
func (s *_dateProcessor) OutputFormat(outputformat string) *_dateProcessor {

	s.v.OutputFormat = &outputformat

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_dateProcessor) Tag(tag string) *_dateProcessor {

	s.v.Tag = &tag

	return s
}

// The field that will hold the parsed date.
func (s *_dateProcessor) TargetField(field string) *_dateProcessor {

	s.v.TargetField = &field

	return s
}

// The timezone to use when parsing the date.
// Supports template snippets.
func (s *_dateProcessor) Timezone(timezone string) *_dateProcessor {

	s.v.Timezone = &timezone

	return s
}

func (s *_dateProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Date = s.v

	return container
}

func (s *_dateProcessor) DateProcessorCaster() *types.DateProcessor {
	return s.v
}
