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

// RenameProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L306-L310
type RenameProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   Field                `json:"target_field"`
}

// RenameProcessorBuilder holds RenameProcessor struct and provides a builder API.
type RenameProcessorBuilder struct {
	v *RenameProcessor
}

// NewRenameProcessor provides a builder for the RenameProcessor struct.
func NewRenameProcessorBuilder() *RenameProcessorBuilder {
	r := RenameProcessorBuilder{
		&RenameProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the RenameProcessor struct
func (rb *RenameProcessorBuilder) Build() RenameProcessor {
	return *rb.v
}

func (rb *RenameProcessorBuilder) Field(field Field) *RenameProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *RenameProcessorBuilder) If_(if_ string) *RenameProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *RenameProcessorBuilder) IgnoreFailure(ignorefailure bool) *RenameProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *RenameProcessorBuilder) IgnoreMissing(ignoremissing bool) *RenameProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *RenameProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *RenameProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *RenameProcessorBuilder) Tag(tag string) *RenameProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *RenameProcessorBuilder) TargetField(targetfield Field) *RenameProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
