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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// SortProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L335-L339
type SortProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Order         sortorder.SortOrder  `json:"order"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   Field                `json:"target_field"`
}

// SortProcessorBuilder holds SortProcessor struct and provides a builder API.
type SortProcessorBuilder struct {
	v *SortProcessor
}

// NewSortProcessor provides a builder for the SortProcessor struct.
func NewSortProcessorBuilder() *SortProcessorBuilder {
	r := SortProcessorBuilder{
		&SortProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the SortProcessor struct
func (rb *SortProcessorBuilder) Build() SortProcessor {
	return *rb.v
}

func (rb *SortProcessorBuilder) Field(field Field) *SortProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *SortProcessorBuilder) If_(if_ string) *SortProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *SortProcessorBuilder) IgnoreFailure(ignorefailure bool) *SortProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *SortProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *SortProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *SortProcessorBuilder) Order(order sortorder.SortOrder) *SortProcessorBuilder {
	rb.v.Order = order
	return rb
}

func (rb *SortProcessorBuilder) Tag(tag string) *SortProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *SortProcessorBuilder) TargetField(targetfield Field) *SortProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
