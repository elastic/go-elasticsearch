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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _renameProcessor struct {
	v *types.RenameProcessor
}

// Renames an existing field.
// If the field doesn’t exist or the new name is already used, an exception will
// be thrown.
func NewRenameProcessor() *_renameProcessor {

	return &_renameProcessor{v: types.NewRenameProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_renameProcessor) Description(description string) *_renameProcessor {

	s.v.Description = &description

	return s
}

// The field to be renamed.
// Supports template snippets.
func (s *_renameProcessor) Field(field string) *_renameProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_renameProcessor) If(if_ types.ScriptVariant) *_renameProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_renameProcessor) IgnoreFailure(ignorefailure bool) *_renameProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and `field` does not exist, the processor quietly exits without
// modifying the document.
func (s *_renameProcessor) IgnoreMissing(ignoremissing bool) *_renameProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// Handle failures for the processor.
func (s *_renameProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_renameProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_renameProcessor) Tag(tag string) *_renameProcessor {

	s.v.Tag = &tag

	return s
}

// The new name of the field.
// Supports template snippets.
func (s *_renameProcessor) TargetField(field string) *_renameProcessor {

	s.v.TargetField = field

	return s
}

func (s *_renameProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Rename = s.v

	return container
}

func (s *_renameProcessor) RenameProcessorCaster() *types.RenameProcessor {
	return s.v
}
