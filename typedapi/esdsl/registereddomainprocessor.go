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

type _registeredDomainProcessor struct {
	v *types.RegisteredDomainProcessor
}

// Extracts the registered domain (also known as the effective top-level
// domain or eTLD), sub-domain, and top-level domain from a fully qualified
// domain name (FQDN). Uses the registered domains defined in the Mozilla
// Public Suffix List.
func NewRegisteredDomainProcessor() *_registeredDomainProcessor {

	return &_registeredDomainProcessor{v: types.NewRegisteredDomainProcessor()}

}

func (s *_registeredDomainProcessor) Field(field string) *_registeredDomainProcessor {

	s.v.Field = field

	return s
}

func (s *_registeredDomainProcessor) IgnoreMissing(ignoremissing bool) *_registeredDomainProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_registeredDomainProcessor) TargetField(field string) *_registeredDomainProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_registeredDomainProcessor) Description(description string) *_registeredDomainProcessor {

	s.v.Description = &description

	return s
}

func (s *_registeredDomainProcessor) If(if_ types.ScriptVariant) *_registeredDomainProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_registeredDomainProcessor) IgnoreFailure(ignorefailure bool) *_registeredDomainProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_registeredDomainProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_registeredDomainProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_registeredDomainProcessor) Tag(tag string) *_registeredDomainProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_registeredDomainProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.RegisteredDomain = s.v

	return container
}

func (s *_registeredDomainProcessor) RegisteredDomainProcessorCaster() *types.RegisteredDomainProcessor {
	return s.v
}
