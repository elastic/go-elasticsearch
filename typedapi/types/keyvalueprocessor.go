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

// KeyValueProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L277-L289
type KeyValueProcessor struct {
	ExcludeKeys   []string             `json:"exclude_keys,omitempty"`
	Field         Field                `json:"field"`
	FieldSplit    string               `json:"field_split"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	IncludeKeys   []string             `json:"include_keys,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Prefix        *string              `json:"prefix,omitempty"`
	StripBrackets *bool                `json:"strip_brackets,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
	TrimKey       *string              `json:"trim_key,omitempty"`
	TrimValue     *string              `json:"trim_value,omitempty"`
	ValueSplit    string               `json:"value_split"`
}

// KeyValueProcessorBuilder holds KeyValueProcessor struct and provides a builder API.
type KeyValueProcessorBuilder struct {
	v *KeyValueProcessor
}

// NewKeyValueProcessor provides a builder for the KeyValueProcessor struct.
func NewKeyValueProcessorBuilder() *KeyValueProcessorBuilder {
	r := KeyValueProcessorBuilder{
		&KeyValueProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the KeyValueProcessor struct
func (rb *KeyValueProcessorBuilder) Build() KeyValueProcessor {
	return *rb.v
}

func (rb *KeyValueProcessorBuilder) ExcludeKeys(exclude_keys ...string) *KeyValueProcessorBuilder {
	rb.v.ExcludeKeys = exclude_keys
	return rb
}

func (rb *KeyValueProcessorBuilder) Field(field Field) *KeyValueProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *KeyValueProcessorBuilder) FieldSplit(fieldsplit string) *KeyValueProcessorBuilder {
	rb.v.FieldSplit = fieldsplit
	return rb
}

func (rb *KeyValueProcessorBuilder) If_(if_ string) *KeyValueProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *KeyValueProcessorBuilder) IgnoreFailure(ignorefailure bool) *KeyValueProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *KeyValueProcessorBuilder) IgnoreMissing(ignoremissing bool) *KeyValueProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *KeyValueProcessorBuilder) IncludeKeys(include_keys ...string) *KeyValueProcessorBuilder {
	rb.v.IncludeKeys = include_keys
	return rb
}

func (rb *KeyValueProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *KeyValueProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *KeyValueProcessorBuilder) Prefix(prefix string) *KeyValueProcessorBuilder {
	rb.v.Prefix = &prefix
	return rb
}

func (rb *KeyValueProcessorBuilder) StripBrackets(stripbrackets bool) *KeyValueProcessorBuilder {
	rb.v.StripBrackets = &stripbrackets
	return rb
}

func (rb *KeyValueProcessorBuilder) Tag(tag string) *KeyValueProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *KeyValueProcessorBuilder) TargetField(targetfield Field) *KeyValueProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}

func (rb *KeyValueProcessorBuilder) TrimKey(trimkey string) *KeyValueProcessorBuilder {
	rb.v.TrimKey = &trimkey
	return rb
}

func (rb *KeyValueProcessorBuilder) TrimValue(trimvalue string) *KeyValueProcessorBuilder {
	rb.v.TrimValue = &trimvalue
	return rb
}

func (rb *KeyValueProcessorBuilder) ValueSplit(valuesplit string) *KeyValueProcessorBuilder {
	rb.v.ValueSplit = valuesplit
	return rb
}
