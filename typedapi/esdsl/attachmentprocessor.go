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

type _attachmentProcessor struct {
	v *types.AttachmentProcessor
}

// The attachment processor lets Elasticsearch extract file attachments in
// common formats (such as PPT, XLS, and PDF) by using the Apache text
// extraction library Tika.
func NewAttachmentProcessor() *_attachmentProcessor {

	return &_attachmentProcessor{v: types.NewAttachmentProcessor()}

}

// Description of the processor.
// Useful for describing the purpose of the processor or its configuration.
func (s *_attachmentProcessor) Description(description string) *_attachmentProcessor {

	s.v.Description = &description

	return s
}

// The field to get the base64 encoded field from.
func (s *_attachmentProcessor) Field(field string) *_attachmentProcessor {

	s.v.Field = field

	return s
}

// Conditionally execute the processor.
func (s *_attachmentProcessor) If(if_ types.ScriptVariant) *_attachmentProcessor {

	s.v.If = if_.ScriptCaster()

	return s
}

// Ignore failures for the processor.
func (s *_attachmentProcessor) IgnoreFailure(ignorefailure bool) *_attachmentProcessor {

	s.v.IgnoreFailure = &ignorefailure

	return s
}

// If `true` and field does not exist, the processor quietly exits without
// modifying the document.
func (s *_attachmentProcessor) IgnoreMissing(ignoremissing bool) *_attachmentProcessor {

	s.v.IgnoreMissing = &ignoremissing

	return s
}

// The number of chars being used for extraction to prevent huge fields.
// Use `-1` for no limit.
func (s *_attachmentProcessor) IndexedChars(indexedchars int64) *_attachmentProcessor {

	s.v.IndexedChars = &indexedchars

	return s
}

// Field name from which you can overwrite the number of chars being used for
// extraction.
func (s *_attachmentProcessor) IndexedCharsField(field string) *_attachmentProcessor {

	s.v.IndexedCharsField = &field

	return s
}

// Handle failures for the processor.
func (s *_attachmentProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_attachmentProcessor {

	for _, v := range onfailures {

		s.v.OnFailure = append(s.v.OnFailure, *v.ProcessorContainerCaster())

	}
	return s
}

// Array of properties to select to be stored.
// Can be `content`, `title`, `name`, `author`, `keywords`, `date`,
// `content_type`, `content_length`, `language`.
func (s *_attachmentProcessor) Properties(properties ...string) *_attachmentProcessor {

	for _, v := range properties {

		s.v.Properties = append(s.v.Properties, v)

	}
	return s
}

// If true, the binary field will be removed from the document
func (s *_attachmentProcessor) RemoveBinary(removebinary bool) *_attachmentProcessor {

	s.v.RemoveBinary = &removebinary

	return s
}

// Field containing the name of the resource to decode.
// If specified, the processor passes this resource name to the underlying Tika
// library to enable Resource Name Based Detection.
func (s *_attachmentProcessor) ResourceName(resourcename string) *_attachmentProcessor {

	s.v.ResourceName = &resourcename

	return s
}

// Identifier for the processor.
// Useful for debugging and metrics.
func (s *_attachmentProcessor) Tag(tag string) *_attachmentProcessor {

	s.v.Tag = &tag

	return s
}

// The field that will hold the attachment information.
func (s *_attachmentProcessor) TargetField(field string) *_attachmentProcessor {

	s.v.TargetField = &field

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
