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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _urlDecodeProcessor struct {
	v *types.UrlDecodeProcessor
}

// URL-decodes a string.
// If the field is an array of strings, all members of the array will be
// decoded.
func NewUrlDecodeProcessor() *_urlDecodeProcessor {

	return &_urlDecodeProcessor{v: types.NewUrlDecodeProcessor()}

}

func (s *_urlDecodeProcessor) Field(field string) *_urlDecodeProcessor {

	s.v.Field = field

	return s
}

func (s *_urlDecodeProcessor) IgnoreMissing(ignoremissing bool) *_urlDecodeProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_urlDecodeProcessor) TargetField(field string) *_urlDecodeProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_urlDecodeProcessor) Description(description string) *_urlDecodeProcessor {

	s.v.Description = &description

	return s
}

func (s *_urlDecodeProcessor) If(if_ types.ScriptVariant) *_urlDecodeProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_urlDecodeProcessor) IgnoreFailure(ignorefailure bool) *_urlDecodeProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_urlDecodeProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_urlDecodeProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_urlDecodeProcessor) Tag(tag string) *_urlDecodeProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_urlDecodeProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Urldecode = s.v

	return container
}

func (s *_urlDecodeProcessor) UrlDecodeProcessorCaster() *types.UrlDecodeProcessor {
	return s.v
}
