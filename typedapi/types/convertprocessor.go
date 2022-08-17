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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/converttype"
)

// ConvertProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L146-L151
type ConvertProcessor struct {
	Field         Field                   `json:"field"`
	If            *string                 `json:"if,omitempty"`
	IgnoreFailure *bool                   `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                   `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer    `json:"on_failure,omitempty"`
	Tag           *string                 `json:"tag,omitempty"`
	TargetField   Field                   `json:"target_field"`
	Type          converttype.ConvertType `json:"type"`
}

// ConvertProcessorBuilder holds ConvertProcessor struct and provides a builder API.
type ConvertProcessorBuilder struct {
	v *ConvertProcessor
}

// NewConvertProcessor provides a builder for the ConvertProcessor struct.
func NewConvertProcessorBuilder() *ConvertProcessorBuilder {
	r := ConvertProcessorBuilder{
		&ConvertProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the ConvertProcessor struct
func (rb *ConvertProcessorBuilder) Build() ConvertProcessor {
	return *rb.v
}

func (rb *ConvertProcessorBuilder) Field(field Field) *ConvertProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *ConvertProcessorBuilder) If_(if_ string) *ConvertProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *ConvertProcessorBuilder) IgnoreFailure(ignorefailure bool) *ConvertProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *ConvertProcessorBuilder) IgnoreMissing(ignoremissing bool) *ConvertProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *ConvertProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *ConvertProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *ConvertProcessorBuilder) Tag(tag string) *ConvertProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *ConvertProcessorBuilder) TargetField(targetfield Field) *ConvertProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}

func (rb *ConvertProcessorBuilder) Type_(type_ converttype.ConvertType) *ConvertProcessorBuilder {
	rb.v.Type = type_
	return rb
}
