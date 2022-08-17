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

// DotExpanderProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L194-L197
type DotExpanderProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Path          *string              `json:"path,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// DotExpanderProcessorBuilder holds DotExpanderProcessor struct and provides a builder API.
type DotExpanderProcessorBuilder struct {
	v *DotExpanderProcessor
}

// NewDotExpanderProcessor provides a builder for the DotExpanderProcessor struct.
func NewDotExpanderProcessorBuilder() *DotExpanderProcessorBuilder {
	r := DotExpanderProcessorBuilder{
		&DotExpanderProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the DotExpanderProcessor struct
func (rb *DotExpanderProcessorBuilder) Build() DotExpanderProcessor {
	return *rb.v
}

func (rb *DotExpanderProcessorBuilder) Field(field Field) *DotExpanderProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *DotExpanderProcessorBuilder) If_(if_ string) *DotExpanderProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *DotExpanderProcessorBuilder) IgnoreFailure(ignorefailure bool) *DotExpanderProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *DotExpanderProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *DotExpanderProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *DotExpanderProcessorBuilder) Path(path string) *DotExpanderProcessorBuilder {
	rb.v.Path = &path
	return rb
}

func (rb *DotExpanderProcessorBuilder) Tag(tag string) *DotExpanderProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
