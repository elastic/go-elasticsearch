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

// LowercaseProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L291-L295
type LowercaseProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
}

// LowercaseProcessorBuilder holds LowercaseProcessor struct and provides a builder API.
type LowercaseProcessorBuilder struct {
	v *LowercaseProcessor
}

// NewLowercaseProcessor provides a builder for the LowercaseProcessor struct.
func NewLowercaseProcessorBuilder() *LowercaseProcessorBuilder {
	r := LowercaseProcessorBuilder{
		&LowercaseProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the LowercaseProcessor struct
func (rb *LowercaseProcessorBuilder) Build() LowercaseProcessor {
	return *rb.v
}

func (rb *LowercaseProcessorBuilder) Field(field Field) *LowercaseProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *LowercaseProcessorBuilder) If_(if_ string) *LowercaseProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *LowercaseProcessorBuilder) IgnoreFailure(ignorefailure bool) *LowercaseProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *LowercaseProcessorBuilder) IgnoreMissing(ignoremissing bool) *LowercaseProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *LowercaseProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *LowercaseProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *LowercaseProcessorBuilder) Tag(tag string) *LowercaseProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *LowercaseProcessorBuilder) TargetField(targetfield Field) *LowercaseProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
