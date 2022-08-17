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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/useragentproperty"
)

// UserAgentProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L114-L120
type UserAgentProcessor struct {
	Field         Field                                 `json:"field"`
	If            *string                               `json:"if,omitempty"`
	IgnoreFailure *bool                                 `json:"ignore_failure,omitempty"`
	IgnoreMissing bool                                  `json:"ignore_missing"`
	OnFailure     []ProcessorContainer                  `json:"on_failure,omitempty"`
	Options       []useragentproperty.UserAgentProperty `json:"options"`
	RegexFile     string                                `json:"regex_file"`
	Tag           *string                               `json:"tag,omitempty"`
	TargetField   Field                                 `json:"target_field"`
}

// UserAgentProcessorBuilder holds UserAgentProcessor struct and provides a builder API.
type UserAgentProcessorBuilder struct {
	v *UserAgentProcessor
}

// NewUserAgentProcessor provides a builder for the UserAgentProcessor struct.
func NewUserAgentProcessorBuilder() *UserAgentProcessorBuilder {
	r := UserAgentProcessorBuilder{
		&UserAgentProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the UserAgentProcessor struct
func (rb *UserAgentProcessorBuilder) Build() UserAgentProcessor {
	return *rb.v
}

func (rb *UserAgentProcessorBuilder) Field(field Field) *UserAgentProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *UserAgentProcessorBuilder) If_(if_ string) *UserAgentProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *UserAgentProcessorBuilder) IgnoreFailure(ignorefailure bool) *UserAgentProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *UserAgentProcessorBuilder) IgnoreMissing(ignoremissing bool) *UserAgentProcessorBuilder {
	rb.v.IgnoreMissing = ignoremissing
	return rb
}

func (rb *UserAgentProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *UserAgentProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *UserAgentProcessorBuilder) Options(options ...useragentproperty.UserAgentProperty) *UserAgentProcessorBuilder {
	rb.v.Options = options
	return rb
}

func (rb *UserAgentProcessorBuilder) RegexFile(regexfile string) *UserAgentProcessorBuilder {
	rb.v.RegexFile = regexfile
	return rb
}

func (rb *UserAgentProcessorBuilder) Tag(tag string) *UserAgentProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *UserAgentProcessorBuilder) TargetField(targetfield Field) *UserAgentProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
