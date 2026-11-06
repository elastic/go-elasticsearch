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

type _recoverFailureDocumentProcessor struct {
	v *types.RecoverFailureDocumentProcessor
}

// Recovers documents that have been stored in a data stream's failure store by
// restoring them to their original format.
func NewRecoverFailureDocumentProcessor() *_recoverFailureDocumentProcessor {

	return &_recoverFailureDocumentProcessor{v: types.NewRecoverFailureDocumentProcessor()}

}

func (s *_recoverFailureDocumentProcessor) Description(description string) *_recoverFailureDocumentProcessor {

	s.v.Description = &description

	return s
}

func (s *_recoverFailureDocumentProcessor) If(if_ types.ScriptVariant) *_recoverFailureDocumentProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_recoverFailureDocumentProcessor) IgnoreFailure(ignorefailure bool) *_recoverFailureDocumentProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_recoverFailureDocumentProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_recoverFailureDocumentProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_recoverFailureDocumentProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_recoverFailureDocumentProcessor {

	s.v.OnFailure = onfailurevalues
	return s
}

func (s *_recoverFailureDocumentProcessor) Tag(tag string) *_recoverFailureDocumentProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_recoverFailureDocumentProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.RecoverFailureDocument = s.v

	return container
}

func (s *_recoverFailureDocumentProcessor) RecoverFailureDocumentProcessorCaster() *types.RecoverFailureDocumentProcessor {
	return s.v
}
