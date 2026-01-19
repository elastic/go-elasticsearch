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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _splitProcessor struct {
	v *types.SplitProcessor
}

// Splits a field into an array using a separator character.
// Only works on string fields.
func NewSplitProcessor(separator string) *_splitProcessor {

	tmp := &_splitProcessor{v: types.NewSplitProcessor()}

	tmp.Separator(separator)

	return tmp

}

func (s *_splitProcessor) Field(field string) *_splitProcessor {

	s.v.Field = field

	return s
}

func (s *_splitProcessor) IgnoreMissing(ignoremissing bool) *_splitProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_splitProcessor) PreserveTrailing(preservetrailing bool) *_splitProcessor {

	s.v.PreserveTrailing = &preservetrailing

	return s
}

func (s *_splitProcessor) Separator(separator string) *_splitProcessor {

	s.v.Separator = separator

	return s
}

func (s *_splitProcessor) TargetField(field string) *_splitProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_splitProcessor) Description(description string) *_splitProcessor {

	s.v.Description = &description

	return s
}

func (s *_splitProcessor) If(if_ types.ScriptVariant) *_splitProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_splitProcessor) IgnoreFailure(ignorefailure bool) *_splitProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_splitProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_splitProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_splitProcessor) Tag(tag string) *_splitProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_splitProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Split = s.v

	return container
}

func (s *_splitProcessor) SplitProcessorCaster() *types.SplitProcessor {
	return s.v
}
