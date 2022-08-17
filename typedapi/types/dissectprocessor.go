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

// DissectProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L187-L192
type DissectProcessor struct {
	AppendSeparator string               `json:"append_separator"`
	Field           Field                `json:"field"`
	If              *string              `json:"if,omitempty"`
	IgnoreFailure   *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing   bool                 `json:"ignore_missing"`
	OnFailure       []ProcessorContainer `json:"on_failure,omitempty"`
	Pattern         string               `json:"pattern"`
	Tag             *string              `json:"tag,omitempty"`
}

// DissectProcessorBuilder holds DissectProcessor struct and provides a builder API.
type DissectProcessorBuilder struct {
	v *DissectProcessor
}

// NewDissectProcessor provides a builder for the DissectProcessor struct.
func NewDissectProcessorBuilder() *DissectProcessorBuilder {
	r := DissectProcessorBuilder{
		&DissectProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the DissectProcessor struct
func (rb *DissectProcessorBuilder) Build() DissectProcessor {
	return *rb.v
}

func (rb *DissectProcessorBuilder) AppendSeparator(appendseparator string) *DissectProcessorBuilder {
	rb.v.AppendSeparator = appendseparator
	return rb
}

func (rb *DissectProcessorBuilder) Field(field Field) *DissectProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *DissectProcessorBuilder) If_(if_ string) *DissectProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *DissectProcessorBuilder) IgnoreFailure(ignorefailure bool) *DissectProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *DissectProcessorBuilder) IgnoreMissing(ignoremissing bool) *DissectProcessorBuilder {
	rb.v.IgnoreMissing = ignoremissing
	return rb
}

func (rb *DissectProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *DissectProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *DissectProcessorBuilder) Pattern(pattern string) *DissectProcessorBuilder {
	rb.v.Pattern = pattern
	return rb
}

func (rb *DissectProcessorBuilder) Tag(tag string) *DissectProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
