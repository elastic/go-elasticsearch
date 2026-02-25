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

type _redactProcessor struct {
	v *types.RedactProcessor
}

// The Redact processor uses the Grok rules engine to obscure text in the input
// document matching the given Grok patterns. The processor can be used to
// obscure Personal Identifying Information (PII) by configuring it to detect
// known patterns such as email or IP addresses. Text that matches a Grok
// pattern is replaced with a configurable string such as `<EMAIL>` where an
// email address is matched or simply replace all matches with the text
// `<REDACTED>` if preferred.
func NewRedactProcessor() *_redactProcessor {

	return &_redactProcessor{v: types.NewRedactProcessor()}

}

func (s *_redactProcessor) Field(field string) *_redactProcessor {

	s.v.Field = field

	return s
}

func (s *_redactProcessor) IgnoreMissing(ignoremissing bool) *_redactProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_redactProcessor) PatternDefinitions(patterndefinitions map[string]string) *_redactProcessor {

	s.v.PatternDefinitions = patterndefinitions
	return s
}

func (s *_redactProcessor) AddPatternDefinition(key string, value string) *_redactProcessor {

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

func (s *_redactProcessor) Patterns(patterns ...string) *_redactProcessor {

	for _, v := range patterns {

		s.v.Patterns = append(s.v.Patterns, v)

	}
	return s
}

func (s *_redactProcessor) Prefix(prefix string) *_redactProcessor {

	s.v.Prefix = &prefix

	return s
}

func (s *_redactProcessor) SkipIfUnlicensed(skipifunlicensed bool) *_redactProcessor {

	s.v.SkipIfUnlicensed = &skipifunlicensed

	return s
}

func (s *_redactProcessor) Suffix(suffix string) *_redactProcessor {

	s.v.Suffix = &suffix

	return s
}

func (s *_redactProcessor) TraceRedact(traceredact bool) *_redactProcessor {

	s.v.TraceRedact = &traceredact

	return s
}

func (s *_redactProcessor) Description(description string) *_redactProcessor {

	s.v.Description = &description

	return s
}

func (s *_redactProcessor) If(if_ types.ScriptVariant) *_redactProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_redactProcessor) IgnoreFailure(ignorefailure bool) *_redactProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_redactProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_redactProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_redactProcessor) Tag(tag string) *_redactProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_redactProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Redact = s.v

	return container
}

func (s *_redactProcessor) RedactProcessorCaster() *types.RedactProcessor {
	return s.v
}
