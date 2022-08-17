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

// SplitProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L341-L347
type SplitProcessor struct {
	Field            Field                `json:"field"`
	If               *string              `json:"if,omitempty"`
	IgnoreFailure    *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing    *bool                `json:"ignore_missing,omitempty"`
	OnFailure        []ProcessorContainer `json:"on_failure,omitempty"`
	PreserveTrailing *bool                `json:"preserve_trailing,omitempty"`
	Separator        string               `json:"separator"`
	Tag              *string              `json:"tag,omitempty"`
	TargetField      *Field               `json:"target_field,omitempty"`
}

// SplitProcessorBuilder holds SplitProcessor struct and provides a builder API.
type SplitProcessorBuilder struct {
	v *SplitProcessor
}

// NewSplitProcessor provides a builder for the SplitProcessor struct.
func NewSplitProcessorBuilder() *SplitProcessorBuilder {
	r := SplitProcessorBuilder{
		&SplitProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the SplitProcessor struct
func (rb *SplitProcessorBuilder) Build() SplitProcessor {
	return *rb.v
}

func (rb *SplitProcessorBuilder) Field(field Field) *SplitProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *SplitProcessorBuilder) If_(if_ string) *SplitProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *SplitProcessorBuilder) IgnoreFailure(ignorefailure bool) *SplitProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *SplitProcessorBuilder) IgnoreMissing(ignoremissing bool) *SplitProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *SplitProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *SplitProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *SplitProcessorBuilder) PreserveTrailing(preservetrailing bool) *SplitProcessorBuilder {
	rb.v.PreserveTrailing = &preservetrailing
	return rb
}

func (rb *SplitProcessorBuilder) Separator(separator string) *SplitProcessorBuilder {
	rb.v.Separator = separator
	return rb
}

func (rb *SplitProcessorBuilder) Tag(tag string) *SplitProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *SplitProcessorBuilder) TargetField(targetfield Field) *SplitProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
