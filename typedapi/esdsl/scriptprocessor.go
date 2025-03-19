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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _scriptProcessor struct {
	v *types.ScriptProcessor
}

// Runs an inline or stored script on incoming documents.
// The script runs in the `ingest` context.
func NewScriptProcessor() *_scriptProcessor {

	return &_scriptProcessor{v: types.NewScriptProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_scriptProcessor) Description(description string) *_scriptProcessor {

	s.v.Description = &description

	return s
}

// ID of a stored script.
// If no `source` is specified, this parameter is required.
func (s *_scriptProcessor) Id(id string) *_scriptProcessor {

	s.v.Id = &id

	return s
}

// Conditionally execute the processor.
func (s *_scriptProcessor) If(if_ types.ScriptVariant) *_scriptProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_scriptProcessor) IgnoreFailure(ignorefailure bool) *_scriptProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// Script language.
func (s *_scriptProcessor) Lang(lang string) *_scriptProcessor {

	s.v.Lang = &lang

	return s
}

// Handle failures for the processor.
func (s *_scriptProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_scriptProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Object containing parameters for the script.
func (s *_scriptProcessor) Params(params map[string]json.RawMessage) *_scriptProcessor {

	s.v.Params = params
	return s
}

func (s *_scriptProcessor) AddParam(key string, value json.RawMessage) *_scriptProcessor {

	var tmp map[string]json.RawMessage
	if s.v.Params == nil {
		s.v.Params = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Params
	}

	tmp[key] = value

	s.v.Params = tmp
	return s
}

// Inline script.
// If no `id` is specified, this parameter is required.
func (s *_scriptProcessor) Source(source string) *_scriptProcessor {

	s.v.Source = &source

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_scriptProcessor) Tag(tag string) *_scriptProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_scriptProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Script = s.v

	return container
}

func (s *_scriptProcessor) ScriptProcessorCaster() *types.ScriptProcessor {
	return s.v
}
