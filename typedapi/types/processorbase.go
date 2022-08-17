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

// ProcessorBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L69-L74
type ProcessorBase struct {
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// ProcessorBaseBuilder holds ProcessorBase struct and provides a builder API.
type ProcessorBaseBuilder struct {
	v *ProcessorBase
}

// NewProcessorBase provides a builder for the ProcessorBase struct.
func NewProcessorBaseBuilder() *ProcessorBaseBuilder {
	r := ProcessorBaseBuilder{
		&ProcessorBase{},
	}

	return &r
}

// Build finalize the chain and returns the ProcessorBase struct
func (rb *ProcessorBaseBuilder) Build() ProcessorBase {
	return *rb.v
}

func (rb *ProcessorBaseBuilder) If_(if_ string) *ProcessorBaseBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *ProcessorBaseBuilder) IgnoreFailure(ignorefailure bool) *ProcessorBaseBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *ProcessorBaseBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *ProcessorBaseBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *ProcessorBaseBuilder) Tag(tag string) *ProcessorBaseBuilder {
	rb.v.Tag = &tag
	return rb
}
