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

// BytesProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L122-L126
type BytesProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
}

// BytesProcessorBuilder holds BytesProcessor struct and provides a builder API.
type BytesProcessorBuilder struct {
	v *BytesProcessor
}

// NewBytesProcessor provides a builder for the BytesProcessor struct.
func NewBytesProcessorBuilder() *BytesProcessorBuilder {
	r := BytesProcessorBuilder{
		&BytesProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the BytesProcessor struct
func (rb *BytesProcessorBuilder) Build() BytesProcessor {
	return *rb.v
}

func (rb *BytesProcessorBuilder) Field(field Field) *BytesProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *BytesProcessorBuilder) If_(if_ string) *BytesProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *BytesProcessorBuilder) IgnoreFailure(ignorefailure bool) *BytesProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *BytesProcessorBuilder) IgnoreMissing(ignoremissing bool) *BytesProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *BytesProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *BytesProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *BytesProcessorBuilder) Tag(tag string) *BytesProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *BytesProcessorBuilder) TargetField(targetfield Field) *BytesProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
