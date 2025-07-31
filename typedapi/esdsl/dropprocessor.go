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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dropProcessor struct {
	v *types.DropProcessor
}

// Drops the document without raising any errors.
// This is useful to prevent the document from getting indexed based on some
// condition.
func NewDropProcessor() *_dropProcessor {

	return &_dropProcessor{v: types.NewDropProcessor()}

}

func (s *_dropProcessor) Description(description string) *_dropProcessor {

	s.v.Description = &description

	return s
}

func (s *_dropProcessor) If(if_ types.ScriptVariant) *_dropProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_dropProcessor) IgnoreFailure(ignorefailure bool) *_dropProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_dropProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dropProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_dropProcessor) Tag(tag string) *_dropProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_dropProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Drop = s.v

	return container
}

func (s *_dropProcessor) DropProcessorCaster() *types.DropProcessor {
	return s.v
}
