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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _setSecurityUserProcessor struct {
	v *types.SetSecurityUserProcessor
}

// Sets user-related details (such as `username`, `roles`, `email`, `full_name`,
// `metadata`, `api_key`, `realm` and `authentication_type`) from the current
// authenticated user to the current document by pre-processing the ingest.
func NewSetSecurityUserProcessor() *_setSecurityUserProcessor {

	return &_setSecurityUserProcessor{v: types.NewSetSecurityUserProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_setSecurityUserProcessor) Description(description string) *_setSecurityUserProcessor {

	s.v.Description = &description

	return s
}

// The field to store the user information into.
func (s *_setSecurityUserProcessor) Field(field string) *_setSecurityUserProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_setSecurityUserProcessor) If(if_ types.ScriptVariant) *_setSecurityUserProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_setSecurityUserProcessor) IgnoreFailure(ignorefailure bool) *_setSecurityUserProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// Handle failures for the processor.
func (s *_setSecurityUserProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_setSecurityUserProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Controls what user related properties are added to the field.
func (s *_setSecurityUserProcessor) Properties(properties ...string) *_setSecurityUserProcessor {

	for _, v := range properties {

		s.v.Properties = append(s.v.Properties, v)

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_setSecurityUserProcessor) Tag(tag string) *_setSecurityUserProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_setSecurityUserProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.SetSecurityUser = s.v

	return container
}

func (s *_setSecurityUserProcessor) SetSecurityUserProcessorCaster() *types.SetSecurityUserProcessor {
	return s.v
}
