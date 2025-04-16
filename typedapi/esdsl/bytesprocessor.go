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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bytesProcessor struct {
	v *types.BytesProcessor
}

// Converts a human readable byte value (for example `1kb`) to its value in
// bytes (for example `1024`).
// If the field is an array of strings, all members of the array will be
// converted.
// Supported human readable units are "b", "kb", "mb", "gb", "tb", "pb" case
// insensitive.
// An error will occur if the field is not a supported format or resultant value
// exceeds 2^63.
func NewBytesProcessor() *_bytesProcessor {

	return &_bytesProcessor{v: types.NewBytesProcessor()}

}

func (s *_bytesProcessor) Description(description string) *_bytesProcessor {

	s.v.Description = &description

	return s
}

func (s *_bytesProcessor) Field(field string) *_bytesProcessor {

	s.v.Field = field

	return s
}

func (s *_bytesProcessor) If(if_ types.ScriptVariant) *_bytesProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_bytesProcessor) IgnoreFailure(ignorefailure bool) *_bytesProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_bytesProcessor) IgnoreMissing(ignoremissing bool) *_bytesProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_bytesProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_bytesProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_bytesProcessor) Tag(tag string) *_bytesProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_bytesProcessor) TargetField(field string) *_bytesProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_bytesProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Bytes = s.v

	return container
}

func (s *_bytesProcessor) BytesProcessorCaster() *types.BytesProcessor {
	return s.v
}
