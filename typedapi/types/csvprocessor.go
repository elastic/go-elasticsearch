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

// CsvProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L153-L162
type CsvProcessor struct {
	Description   *string              `json:"description,omitempty"`
	EmptyValue    interface{}          `json:"empty_value,omitempty"`
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Quote         *string              `json:"quote,omitempty"`
	Separator     *string              `json:"separator,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetFields  Fields               `json:"target_fields"`
	Trim          bool                 `json:"trim"`
}

// CsvProcessorBuilder holds CsvProcessor struct and provides a builder API.
type CsvProcessorBuilder struct {
	v *CsvProcessor
}

// NewCsvProcessor provides a builder for the CsvProcessor struct.
func NewCsvProcessorBuilder() *CsvProcessorBuilder {
	r := CsvProcessorBuilder{
		&CsvProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the CsvProcessor struct
func (rb *CsvProcessorBuilder) Build() CsvProcessor {
	return *rb.v
}

func (rb *CsvProcessorBuilder) Description(description string) *CsvProcessorBuilder {
	rb.v.Description = &description
	return rb
}

func (rb *CsvProcessorBuilder) EmptyValue(emptyvalue interface{}) *CsvProcessorBuilder {
	rb.v.EmptyValue = emptyvalue
	return rb
}

func (rb *CsvProcessorBuilder) Field(field Field) *CsvProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *CsvProcessorBuilder) If_(if_ string) *CsvProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *CsvProcessorBuilder) IgnoreFailure(ignorefailure bool) *CsvProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *CsvProcessorBuilder) IgnoreMissing(ignoremissing bool) *CsvProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *CsvProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *CsvProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *CsvProcessorBuilder) Quote(quote string) *CsvProcessorBuilder {
	rb.v.Quote = &quote
	return rb
}

func (rb *CsvProcessorBuilder) Separator(separator string) *CsvProcessorBuilder {
	rb.v.Separator = &separator
	return rb
}

func (rb *CsvProcessorBuilder) Tag(tag string) *CsvProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *CsvProcessorBuilder) TargetFields(targetfields *FieldsBuilder) *CsvProcessorBuilder {
	v := targetfields.Build()
	rb.v.TargetFields = v
	return rb
}

func (rb *CsvProcessorBuilder) Trim(trim bool) *CsvProcessorBuilder {
	rb.v.Trim = trim
	return rb
}
