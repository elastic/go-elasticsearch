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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _joinProcessor struct {
	v *types.JoinProcessor
}

// Joins each element of an array into a single string using a separator
// character between each element.
// Throws an error when the field is not an array.
func NewJoinProcessor(separator string) *_joinProcessor {

	tmp := &_joinProcessor{v: types.NewJoinProcessor()}

	tmp.Separator(separator)

	return tmp

}

func (s *_joinProcessor) Field(field string) *_joinProcessor {

	s.v.Field = field

	return s
}

func (s *_joinProcessor) Separator(separator string) *_joinProcessor {

	s.v.Separator = separator

	return s
}

func (s *_joinProcessor) TargetField(field string) *_joinProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_joinProcessor) Description(description string) *_joinProcessor {

	s.v.Description = &description

	return s
}

func (s *_joinProcessor) If(if_ types.ScriptVariant) *_joinProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_joinProcessor) IgnoreFailure(ignorefailure bool) *_joinProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_joinProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_joinProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_joinProcessor) Tag(tag string) *_joinProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_joinProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Join = s.v

	return container
}

func (s *_joinProcessor) JoinProcessorCaster() *types.JoinProcessor {
	return s.v
}
