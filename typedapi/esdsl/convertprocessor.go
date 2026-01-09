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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/converttype"
)

type _convertProcessor struct {
	v *types.ConvertProcessor
}

// Converts a field in the currently ingested document to a different type, such
// as converting a string to an integer.
// If the field value is an array, all members will be converted.
func NewConvertProcessor(type_ converttype.ConvertType) *_convertProcessor {

	tmp := &_convertProcessor{v: types.NewConvertProcessor()}

	tmp.Type(type_)

	return tmp

}

func (s *_convertProcessor) Field(field string) *_convertProcessor {

	s.v.Field = field

	return s
}

func (s *_convertProcessor) IgnoreMissing(ignoremissing bool) *_convertProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_convertProcessor) TargetField(field string) *_convertProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_convertProcessor) Type(type_ converttype.ConvertType) *_convertProcessor {

	s.v.Type = type_
	return s
}

func (s *_convertProcessor) Description(description string) *_convertProcessor {

	s.v.Description = &description

	return s
}

func (s *_convertProcessor) If(if_ types.ScriptVariant) *_convertProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_convertProcessor) IgnoreFailure(ignorefailure bool) *_convertProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_convertProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_convertProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_convertProcessor) Tag(tag string) *_convertProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_convertProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Convert = s.v

	return container
}

func (s *_convertProcessor) ConvertProcessorCaster() *types.ConvertProcessor {
	return s.v
}
