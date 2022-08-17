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

// ComponentTemplateNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/_types/ComponentTemplate.ts#L31-L36
type ComponentTemplateNode struct {
	Meta_    *Metadata                `json:"_meta,omitempty"`
	Template ComponentTemplateSummary `json:"template"`
	Version  *VersionNumber           `json:"version,omitempty"`
}

// ComponentTemplateNodeBuilder holds ComponentTemplateNode struct and provides a builder API.
type ComponentTemplateNodeBuilder struct {
	v *ComponentTemplateNode
}

// NewComponentTemplateNode provides a builder for the ComponentTemplateNode struct.
func NewComponentTemplateNodeBuilder() *ComponentTemplateNodeBuilder {
	r := ComponentTemplateNodeBuilder{
		&ComponentTemplateNode{},
	}

	return &r
}

// Build finalize the chain and returns the ComponentTemplateNode struct
func (rb *ComponentTemplateNodeBuilder) Build() ComponentTemplateNode {
	return *rb.v
}

func (rb *ComponentTemplateNodeBuilder) Meta_(meta_ *MetadataBuilder) *ComponentTemplateNodeBuilder {
	v := meta_.Build()
	rb.v.Meta_ = &v
	return rb
}

func (rb *ComponentTemplateNodeBuilder) Template(template *ComponentTemplateSummaryBuilder) *ComponentTemplateNodeBuilder {
	v := template.Build()
	rb.v.Template = v
	return rb
}

func (rb *ComponentTemplateNodeBuilder) Version(version VersionNumber) *ComponentTemplateNodeBuilder {
	rb.v.Version = &version
	return rb
}
