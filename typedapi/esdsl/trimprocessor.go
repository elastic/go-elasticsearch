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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _trimProcessor struct {
	v *types.TrimProcessor
}

// Trims whitespace from a field.
// If the field is an array of strings, all members of the array will be
// trimmed.
// This only works on leading and trailing whitespace.
func NewTrimProcessor() *_trimProcessor {

	return &_trimProcessor{v: types.NewTrimProcessor()}

}

func (s *_trimProcessor) Description(description string) *_trimProcessor {

	s.v.Description = &description

	return s
}

func (s *_trimProcessor) Field(field string) *_trimProcessor {

	s.v.Field = field

	return s
}

func (s *_trimProcessor) If(if_ types.ScriptVariant) *_trimProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_trimProcessor) IgnoreFailure(ignorefailure bool) *_trimProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_trimProcessor) IgnoreMissing(ignoremissing bool) *_trimProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_trimProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_trimProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_trimProcessor) Tag(tag string) *_trimProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_trimProcessor) TargetField(field string) *_trimProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_trimProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Trim = s.v

	return container
}

func (s *_trimProcessor) TrimProcessorCaster() *types.TrimProcessor {
	return s.v
}
