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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _normalizeForStreamProcessor struct {
	v *types.NormalizeForStreamProcessor
}

// Detects whether a document is OpenTelemetry-compliant and, if not, normalizes
// it so that it can be queried by clients that expect either ECS or
// OpenTelemetry Semantic Conventions formats.
func NewNormalizeForStreamProcessor() *_normalizeForStreamProcessor {

	return &_normalizeForStreamProcessor{v: types.NewNormalizeForStreamProcessor()}

}

func (s *_normalizeForStreamProcessor) Description(description string) *_normalizeForStreamProcessor {

	s.v.Description = &description

	return s
}

func (s *_normalizeForStreamProcessor) If(if_ types.ScriptVariant) *_normalizeForStreamProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_normalizeForStreamProcessor) IgnoreFailure(ignorefailure bool) *_normalizeForStreamProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_normalizeForStreamProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_normalizeForStreamProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_normalizeForStreamProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_normalizeForStreamProcessor {

	s.v.OnFailure = onfailurevalues
	return s
}

func (s *_normalizeForStreamProcessor) Tag(tag string) *_normalizeForStreamProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_normalizeForStreamProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.NormalizeForStream = s.v

	return container
}

func (s *_normalizeForStreamProcessor) NormalizeForStreamProcessorCaster() *types.NormalizeForStreamProcessor {
	return s.v
}
