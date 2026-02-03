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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _grokProcessor struct {
	v *types.GrokProcessor
}

// Extracts structured fields out of a single text field within a document.
// You choose which field to extract matched fields from, as well as the grok
// pattern you expect will match.
// A grok pattern is like a regular expression that supports aliased expressions
// that can be reused.
func NewGrokProcessor() *_grokProcessor {

	return &_grokProcessor{v: types.NewGrokProcessor()}

}

func (s *_grokProcessor) EcsCompatibility(ecscompatibility string) *_grokProcessor {

	s.v.EcsCompatibility = &ecscompatibility

	return s
}

func (s *_grokProcessor) Field(field string) *_grokProcessor {

	s.v.Field = field

	return s
}

func (s *_grokProcessor) IgnoreMissing(ignoremissing bool) *_grokProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_grokProcessor) PatternDefinitions(patterndefinitions map[string]string) *_grokProcessor {

	s.v.PatternDefinitions = patterndefinitions
	return s
}

func (s *_grokProcessor) AddPatternDefinition(key string, value string) *_grokProcessor {

	var tmp map[string]string
	if s.v.PatternDefinitions == nil {
		s.v.PatternDefinitions = make(map[string]string)
	} else {
		tmp = s.v.PatternDefinitions
	}

	tmp[key] = value

	s.v.PatternDefinitions = tmp
	return s
}

func (s *_grokProcessor) Patterns(patterns ...string) *_grokProcessor {

	for _, v := range patterns {

		s.v.Patterns = append(s.v.Patterns, v)

	}
	return s
}

func (s *_grokProcessor) TraceMatch(tracematch bool) *_grokProcessor {

	s.v.TraceMatch = &tracematch

	return s
}

func (s *_grokProcessor) Description(description string) *_grokProcessor {

	s.v.Description = &description

	return s
}

func (s *_grokProcessor) If(if_ types.ScriptVariant) *_grokProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_grokProcessor) IgnoreFailure(ignorefailure bool) *_grokProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_grokProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_grokProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_grokProcessor) Tag(tag string) *_grokProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_grokProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Grok = s.v

	return container
}

func (s *_grokProcessor) GrokProcessorCaster() *types.GrokProcessor {
	return s.v
}
