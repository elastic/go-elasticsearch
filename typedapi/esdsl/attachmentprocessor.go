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

type _attachmentProcessor struct {
	v *types.AttachmentProcessor
}

// The attachment processor lets Elasticsearch extract file attachments in
// common formats (such as PPT, XLS, and PDF) by using the Apache text
// extraction library Tika.
func NewAttachmentProcessor() *_attachmentProcessor {

	return &_attachmentProcessor{v: types.NewAttachmentProcessor()}

}

func (s *_attachmentProcessor) Field(field string) *_attachmentProcessor {

	s.v.Field = field

	return s
}

func (s *_attachmentProcessor) IgnoreMissing(ignoremissing bool) *_attachmentProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

func (s *_attachmentProcessor) IndexedChars(indexedchars int64) *_attachmentProcessor {

	s.v.IndexedChars = &indexedchars

	return s
}

func (s *_attachmentProcessor) IndexedCharsField(field string) *_attachmentProcessor {

	s.v.IndexedCharsField = &field

	return s
}

func (s *_attachmentProcessor) Properties(properties ...string) *_attachmentProcessor {

	for _, v := range properties {

		s.v.Properties = append(s.v.Properties, v)

	}
	return s
}

func (s *_attachmentProcessor) RemoveBinary(removebinary bool) *_attachmentProcessor {

	s.v.RemoveBinary = &removebinary

	return s
}

func (s *_attachmentProcessor) ResourceName(resourcename string) *_attachmentProcessor {

	s.v.ResourceName = &resourcename

	return s
}

func (s *_attachmentProcessor) TargetField(field string) *_attachmentProcessor {

	s.v.TargetField = &field

	return s
}

func (s *_attachmentProcessor) Description(description string) *_attachmentProcessor {

	s.v.Description = &description

	return s
}

func (s *_attachmentProcessor) If(if_ types.ScriptVariant) *_attachmentProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

func (s *_attachmentProcessor) IgnoreFailure(ignorefailure bool) *_attachmentProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

func (s *_attachmentProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_attachmentProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

func (s *_attachmentProcessor) Tag(tag string) *_attachmentProcessor {

	s.v.Tag = &tag

	return s
}

func (s *_attachmentProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	container := types.NewProcessorContainer()

	container.Attachment = s.v

	return container
}

func (s *_attachmentProcessor) AttachmentProcessorCaster() *types.AttachmentProcessor {
	return s.v
}
