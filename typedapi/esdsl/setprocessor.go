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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
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

func (s *_setProcessor) CopyFrom(field string) *_setProcessor {

	s.v.CopyFrom = &field

	return s
}

func (s *_setProcessor) Description(description string) *_setProcessor {

	s.v.Description = &description

	return s
}

func (s *_setProcessor) Field(field string) *_setProcessor {

	s.v.Field = field

	return s
}

func (s *_setProcessor) If(if_ types.ScriptVariant) *_setProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_setProcessor) IgnoreEmptyValue(ignoreemptyvalue bool) *_setProcessor {

	s.v.IgnoreEmptyValue = &ignoreemptyvalue

	return s
}

func (s *_setProcessor) IgnoreFailure(ignorefailure bool) *_setProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_setProcessor) MediaType(mediatype string) *_setProcessor {

	s.v.MediaType = &mediatype

	return s
}

func (s *_setProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_setProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_setProcessor) Override(override bool) *_setProcessor {

	s.v.Override = &override

	return s
}

func (s *_setProcessor) Tag(tag string) *_setProcessor {

	s.v.Tag = &tag

	return s
}

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
