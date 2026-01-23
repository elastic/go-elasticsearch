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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _foreachProcessor struct {
	v *types.ForeachProcessor
}

// Runs an ingest processor on each element of an array or object.
func NewForeachProcessor(processor types.ProcessorContainerVariant) *_foreachProcessor {

	tmp := &_foreachProcessor{v: types.NewForeachProcessor()}

	tmp.Processor(processor)

	return tmp

}

func (s *_foreachProcessor) Field(field string) *_foreachProcessor {

	s.v.Field = field

	return s
}

func (s *_foreachProcessor) IgnoreMissing(ignoremissing bool) *_foreachProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_foreachProcessor) Processor(processor types.ProcessorContainerVariant) *_foreachProcessor {

	s.v.Processor = *processor.ProcessorContainerCaster()

	return s
}

func (s *_foreachProcessor) Description(description string) *_foreachProcessor {

	s.v.Description = &description

	return s
}

func (s *_foreachProcessor) If(if_ types.ScriptVariant) *_foreachProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_foreachProcessor) IgnoreFailure(ignorefailure bool) *_foreachProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_foreachProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_foreachProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_foreachProcessor) Tag(tag string) *_foreachProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_foreachProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Foreach = s.v

	return container
}

func (s *_foreachProcessor) ForeachProcessorCaster() *types.ForeachProcessor {
	return s.v
}
