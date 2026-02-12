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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jsonprocessorconflictstrategy"
)

type _jsonProcessor struct {
	v *types.JsonProcessor
}

// Parses a string containing JSON data into a structured object, string, or
// other value.
func NewJsonProcessor() *_jsonProcessor {

	return &_jsonProcessor{v: types.NewJsonProcessor()}

}

func (s *_jsonProcessor) AddToRoot(addtoroot bool) *_jsonProcessor {

	s.v.AddToRoot = &addtoroot

	return s
}

func (s *_jsonProcessor) AddToRootConflictStrategy(addtorootconflictstrategy jsonprocessorconflictstrategy.JsonProcessorConflictStrategy) *_jsonProcessor {

	s.v.AddToRootConflictStrategy = &addtorootconflictstrategy
	return s
}

func (s *_jsonProcessor) AllowDuplicateKeys(allowduplicatekeys bool) *_jsonProcessor {

	s.v.AllowDuplicateKeys = &allowduplicatekeys

	return s
}

func (s *_jsonProcessor) Field(field string) *_jsonProcessor {

	s.v.Field = field

	return s
}

func (s *_jsonProcessor) TargetField(field string) *_jsonProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_jsonProcessor) Description(description string) *_jsonProcessor {

	s.v.Description = &description

	return s
}

func (s *_jsonProcessor) If(if_ types.ScriptVariant) *_jsonProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_jsonProcessor) IgnoreFailure(ignorefailure bool) *_jsonProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_jsonProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_jsonProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_jsonProcessor) Tag(tag string) *_jsonProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_jsonProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Json = s.v

	return container
}

func (s *_jsonProcessor) JsonProcessorCaster() *types.JsonProcessor {
	return s.v
}
