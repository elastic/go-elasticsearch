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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _uppercaseProcessor struct {
	v *types.UppercaseProcessor
}

// Converts a string to its uppercase equivalent. If the field is an array of
// strings, all members of the array will be converted.
func NewUppercaseProcessor() *_uppercaseProcessor {

	return &_uppercaseProcessor{v: types.NewUppercaseProcessor()}

}

func (s *_uppercaseProcessor) Field(field string) *_uppercaseProcessor {

	s.v.Field = field

	return s
}

func (s *_uppercaseProcessor) IgnoreMissing(ignoremissing bool) *_uppercaseProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_uppercaseProcessor) TargetField(field string) *_uppercaseProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_uppercaseProcessor) Description(description string) *_uppercaseProcessor {

	s.v.Description = &description

	return s
}

func (s *_uppercaseProcessor) If(if_ types.ScriptVariant) *_uppercaseProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_uppercaseProcessor) IgnoreFailure(ignorefailure bool) *_uppercaseProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_uppercaseProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_uppercaseProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_uppercaseProcessor) Tag(tag string) *_uppercaseProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_uppercaseProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Uppercase = s.v

	return container
}

func (s *_uppercaseProcessor) UppercaseProcessorCaster() *types.UppercaseProcessor {
	return s.v
}
