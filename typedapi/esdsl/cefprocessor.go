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

type _cefProcessor struct {
	v *types.CefProcessor
}

// Converts a CEF message into a structured format.
func NewCefProcessor() *_cefProcessor {

	return &_cefProcessor{v: types.NewCefProcessor()}

}

func (s *_cefProcessor) Field(field string) *_cefProcessor {

	s.v.Field = field

	return s
}

func (s *_cefProcessor) IgnoreEmptyValues(ignoreemptyvalues bool) *_cefProcessor {

	s.v.IgnoreEmptyValues = &ignoreemptyvalues

	return s
}

func (s *_cefProcessor) IgnoreMissing(ignoremissing bool) *_cefProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_cefProcessor) TargetField(field string) *_cefProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_cefProcessor) Timezone(timezone string) *_cefProcessor {

	s.v.Timezone = &timezone

	return s
}

func (s *_cefProcessor) Description(description string) *_cefProcessor {

	s.v.Description = &description

	return s
}

func (s *_cefProcessor) If(if_ types.ScriptVariant) *_cefProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_cefProcessor) IgnoreFailure(ignorefailure bool) *_cefProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_cefProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_cefProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_cefProcessor) Tag(tag string) *_cefProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_cefProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Cef = s.v

	return container
}

func (s *_cefProcessor) CefProcessorCaster() *types.CefProcessor {
	return s.v
}
