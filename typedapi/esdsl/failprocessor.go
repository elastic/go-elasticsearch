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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _failProcessor struct {
	v *types.FailProcessor
}

// Raises an exception.
// This is useful for when you expect a pipeline to fail and want to relay a
// specific message to the requester.
func NewFailProcessor(message string) *_failProcessor {

	tmp := &_failProcessor{v: types.NewFailProcessor()}

	tmp.Message(message)

	return tmp

}

func (s *_failProcessor) Description(description string) *_failProcessor {

	s.v.Description = &description

	return s
}

func (s *_failProcessor) If(if_ types.ScriptVariant) *_failProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_failProcessor) IgnoreFailure(ignorefailure bool) *_failProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_failProcessor) Message(message string) *_failProcessor {

	s.v.Message = message

	return s
}

func (s *_failProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_failProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_failProcessor) Tag(tag string) *_failProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_failProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Fail = s.v

	return container
}

func (s *_failProcessor) FailProcessorCaster() *types.FailProcessor {
	return s.v
}
