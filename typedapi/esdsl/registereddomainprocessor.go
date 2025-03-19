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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

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

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_registeredDomainProcessor) Description(description string) *_registeredDomainProcessor {

	s.v.Description = &description

	return s
}

// Field containing the source FQDN.
func (s *_registeredDomainProcessor) Field(field string) *_registeredDomainProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_registeredDomainProcessor) If(if_ types.ScriptVariant) *_registeredDomainProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_registeredDomainProcessor) IgnoreFailure(ignorefailure bool) *_registeredDomainProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If true and any required fields are missing, the processor quietly exits
// without modifying the document.
func (s *_registeredDomainProcessor) IgnoreMissing(ignoremissing bool) *_registeredDomainProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_registeredDomainProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_registeredDomainProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_registeredDomainProcessor) Tag(tag string) *_registeredDomainProcessor {

	s.v.Tag = &tag

	return s
}

// Object field containing extracted domain components. If an empty string,
// the processor adds components to the document’s root.
func (s *_registeredDomainProcessor) TargetField(field string) *_registeredDomainProcessor {

	s.v.TargetField = &field

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
