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

// ComponentTemplateSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/_types/ComponentTemplate.ts#L38-L45
type ComponentTemplateSummary struct {
	Aliases  map[string]AliasDefinition  `json:"aliases,omitempty"`
	Mappings *TypeMapping                `json:"mappings,omitempty"`
	Meta_    *Metadata                   `json:"_meta,omitempty"`
	Settings map[IndexName]IndexSettings `json:"settings"`
	Version  *VersionNumber              `json:"version,omitempty"`
}

// ComponentTemplateSummaryBuilder holds ComponentTemplateSummary struct and provides a builder API.
type ComponentTemplateSummaryBuilder struct {
	v *ComponentTemplateSummary
}

// NewComponentTemplateSummary provides a builder for the ComponentTemplateSummary struct.
func NewComponentTemplateSummaryBuilder() *ComponentTemplateSummaryBuilder {
	r := ComponentTemplateSummaryBuilder{
		&ComponentTemplateSummary{
			Aliases:  make(map[string]AliasDefinition, 0),
			Settings: make(map[IndexName]IndexSettings, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ComponentTemplateSummary struct
func (rb *ComponentTemplateSummaryBuilder) Build() ComponentTemplateSummary {
	return *rb.v
}

func (rb *ComponentTemplateSummaryBuilder) Aliases(values map[string]*AliasDefinitionBuilder) *ComponentTemplateSummaryBuilder {
	tmp := make(map[string]AliasDefinition, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aliases = tmp
	return rb
}

func (rb *ComponentTemplateSummaryBuilder) Mappings(mappings *TypeMappingBuilder) *ComponentTemplateSummaryBuilder {
	v := mappings.Build()
	rb.v.Mappings = &v
	return rb
}

func (rb *ComponentTemplateSummaryBuilder) Meta_(meta_ *MetadataBuilder) *ComponentTemplateSummaryBuilder {
	v := meta_.Build()
	rb.v.Meta_ = &v
	return rb
}

func (rb *ComponentTemplateSummaryBuilder) Settings(values map[IndexName]*IndexSettingsBuilder) *ComponentTemplateSummaryBuilder {
	tmp := make(map[IndexName]IndexSettings, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Settings = tmp
	return rb
}

func (rb *ComponentTemplateSummaryBuilder) Version(version VersionNumber) *ComponentTemplateSummaryBuilder {
	rb.v.Version = &version
	return rb
}
