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

// IndexTemplateDataStreamConfiguration type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexTemplate.ts#L39-L50
type IndexTemplateDataStreamConfiguration struct {
	// AllowCustomRouting If true, the data stream supports custom routing.
	AllowCustomRouting *bool `json:"allow_custom_routing,omitempty"`
	// Hidden If true, the data stream is hidden.
	Hidden *bool `json:"hidden,omitempty"`
}

// IndexTemplateDataStreamConfigurationBuilder holds IndexTemplateDataStreamConfiguration struct and provides a builder API.
type IndexTemplateDataStreamConfigurationBuilder struct {
	v *IndexTemplateDataStreamConfiguration
}

// NewIndexTemplateDataStreamConfiguration provides a builder for the IndexTemplateDataStreamConfiguration struct.
func NewIndexTemplateDataStreamConfigurationBuilder() *IndexTemplateDataStreamConfigurationBuilder {
	r := IndexTemplateDataStreamConfigurationBuilder{
		&IndexTemplateDataStreamConfiguration{},
	}

	return &r
}

// Build finalize the chain and returns the IndexTemplateDataStreamConfiguration struct
func (rb *IndexTemplateDataStreamConfigurationBuilder) Build() IndexTemplateDataStreamConfiguration {
	return *rb.v
}

// AllowCustomRouting If true, the data stream supports custom routing.

func (rb *IndexTemplateDataStreamConfigurationBuilder) AllowCustomRouting(allowcustomrouting bool) *IndexTemplateDataStreamConfigurationBuilder {
	rb.v.AllowCustomRouting = &allowcustomrouting
	return rb
}

// Hidden If true, the data stream is hidden.

func (rb *IndexTemplateDataStreamConfigurationBuilder) Hidden(hidden bool) *IndexTemplateDataStreamConfigurationBuilder {
	rb.v.Hidden = &hidden
	return rb
}
