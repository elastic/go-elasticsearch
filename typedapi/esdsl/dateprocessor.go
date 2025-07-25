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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dateProcessor struct {
	v *types.DateProcessor
}

// Parses dates from fields, and then uses the date or timestamp as the
// timestamp for the document.
func NewDateProcessor() *_dateProcessor {

	return &_dateProcessor{v: types.NewDateProcessor()}

}

func (s *_dateProcessor) Description(description string) *_dateProcessor {

	s.v.Description = &description

	return s
}

func (s *_dateProcessor) Field(field string) *_dateProcessor {

	s.v.Field = field

	return s
}

func (s *_dateProcessor) Formats(formats ...string) *_dateProcessor {

	for _, v := range formats {

		s.v.Formats = append(s.v.Formats, v)

	}
	return s
}

func (s *_dateProcessor) If(if_ types.ScriptVariant) *_dateProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_dateProcessor) IgnoreFailure(ignorefailure bool) *_dateProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_dateProcessor) Locale(locale string) *_dateProcessor {

	s.v.Locale = &locale

	return s
}

func (s *_dateProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dateProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_dateProcessor) OutputFormat(outputformat string) *_dateProcessor {

	s.v.OutputFormat = &outputformat

	return s
}

func (s *_dateProcessor) Tag(tag string) *_dateProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_dateProcessor) TargetField(field string) *_dateProcessor {

	s.v.TargetField = &field

	return s
}

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
