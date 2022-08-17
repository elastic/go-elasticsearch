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

// DateProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L179-L185
type DateProcessor struct {
	Field         Field                `json:"field"`
	Formats       []string             `json:"formats"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	Locale        *string              `json:"locale,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
	Timezone      *string              `json:"timezone,omitempty"`
}

// DateProcessorBuilder holds DateProcessor struct and provides a builder API.
type DateProcessorBuilder struct {
	v *DateProcessor
}

// NewDateProcessor provides a builder for the DateProcessor struct.
func NewDateProcessorBuilder() *DateProcessorBuilder {
	r := DateProcessorBuilder{
		&DateProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the DateProcessor struct
func (rb *DateProcessorBuilder) Build() DateProcessor {
	return *rb.v
}

func (rb *DateProcessorBuilder) Field(field Field) *DateProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *DateProcessorBuilder) Formats(formats ...string) *DateProcessorBuilder {
	rb.v.Formats = formats
	return rb
}

func (rb *DateProcessorBuilder) If_(if_ string) *DateProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *DateProcessorBuilder) IgnoreFailure(ignorefailure bool) *DateProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *DateProcessorBuilder) Locale(locale string) *DateProcessorBuilder {
	rb.v.Locale = &locale
	return rb
}

func (rb *DateProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *DateProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *DateProcessorBuilder) Tag(tag string) *DateProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *DateProcessorBuilder) TargetField(targetfield Field) *DateProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}

func (rb *DateProcessorBuilder) Timezone(timezone string) *DateProcessorBuilder {
	rb.v.Timezone = &timezone
	return rb
}
