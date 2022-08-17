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

// TemplateMapping type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/TemplateMapping.ts#L27-L34
type TemplateMapping struct {
	Aliases       map[IndexName]Alias    `json:"aliases"`
	IndexPatterns []Name                 `json:"index_patterns"`
	Mappings      TypeMapping            `json:"mappings"`
	Order         int                    `json:"order"`
	Settings      map[string]interface{} `json:"settings"`
	Version       *VersionNumber         `json:"version,omitempty"`
}

// TemplateMappingBuilder holds TemplateMapping struct and provides a builder API.
type TemplateMappingBuilder struct {
	v *TemplateMapping
}

// NewTemplateMapping provides a builder for the TemplateMapping struct.
func NewTemplateMappingBuilder() *TemplateMappingBuilder {
	r := TemplateMappingBuilder{
		&TemplateMapping{
			Aliases:  make(map[IndexName]Alias, 0),
			Settings: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TemplateMapping struct
func (rb *TemplateMappingBuilder) Build() TemplateMapping {
	return *rb.v
}

func (rb *TemplateMappingBuilder) Aliases(values map[IndexName]*AliasBuilder) *TemplateMappingBuilder {
	tmp := make(map[IndexName]Alias, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aliases = tmp
	return rb
}

func (rb *TemplateMappingBuilder) IndexPatterns(index_patterns ...Name) *TemplateMappingBuilder {
	rb.v.IndexPatterns = index_patterns
	return rb
}

func (rb *TemplateMappingBuilder) Mappings(mappings *TypeMappingBuilder) *TemplateMappingBuilder {
	v := mappings.Build()
	rb.v.Mappings = v
	return rb
}

func (rb *TemplateMappingBuilder) Order(order int) *TemplateMappingBuilder {
	rb.v.Order = order
	return rb
}

func (rb *TemplateMappingBuilder) Settings(value map[string]interface{}) *TemplateMappingBuilder {
	rb.v.Settings = value
	return rb
}

func (rb *TemplateMappingBuilder) Version(version VersionNumber) *TemplateMappingBuilder {
	rb.v.Version = &version
	return rb
}
