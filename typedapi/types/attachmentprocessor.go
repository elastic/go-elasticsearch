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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// AttachmentProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L95-L103
type AttachmentProcessor struct {
	Field             Field                `json:"field"`
	If                *string              `json:"if,omitempty"`
	IgnoreFailure     *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing     *bool                `json:"ignore_missing,omitempty"`
	IndexedChars      *int64               `json:"indexed_chars,omitempty"`
	IndexedCharsField *Field               `json:"indexed_chars_field,omitempty"`
	OnFailure         []ProcessorContainer `json:"on_failure,omitempty"`
	Properties        []string             `json:"properties,omitempty"`
	ResourceName      *string              `json:"resource_name,omitempty"`
	Tag               *string              `json:"tag,omitempty"`
	TargetField       *Field               `json:"target_field,omitempty"`
}

// AttachmentProcessorBuilder holds AttachmentProcessor struct and provides a builder API.
type AttachmentProcessorBuilder struct {
	v *AttachmentProcessor
}

// NewAttachmentProcessor provides a builder for the AttachmentProcessor struct.
func NewAttachmentProcessorBuilder() *AttachmentProcessorBuilder {
	r := AttachmentProcessorBuilder{
		&AttachmentProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the AttachmentProcessor struct
func (rb *AttachmentProcessorBuilder) Build() AttachmentProcessor {
	return *rb.v
}

func (rb *AttachmentProcessorBuilder) Field(field Field) *AttachmentProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *AttachmentProcessorBuilder) If_(if_ string) *AttachmentProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *AttachmentProcessorBuilder) IgnoreFailure(ignorefailure bool) *AttachmentProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *AttachmentProcessorBuilder) IgnoreMissing(ignoremissing bool) *AttachmentProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *AttachmentProcessorBuilder) IndexedChars(indexedchars int64) *AttachmentProcessorBuilder {
	rb.v.IndexedChars = &indexedchars
	return rb
}

func (rb *AttachmentProcessorBuilder) IndexedCharsField(indexedcharsfield Field) *AttachmentProcessorBuilder {
	rb.v.IndexedCharsField = &indexedcharsfield
	return rb
}

func (rb *AttachmentProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *AttachmentProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *AttachmentProcessorBuilder) Properties(properties ...string) *AttachmentProcessorBuilder {
	rb.v.Properties = properties
	return rb
}

func (rb *AttachmentProcessorBuilder) ResourceName(resourcename string) *AttachmentProcessorBuilder {
	rb.v.ResourceName = &resourcename
	return rb
}

func (rb *AttachmentProcessorBuilder) Tag(tag string) *AttachmentProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *AttachmentProcessorBuilder) TargetField(targetfield Field) *AttachmentProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
