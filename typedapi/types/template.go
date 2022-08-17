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

// Template type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/simulate_template/IndicesSimulateTemplateResponse.ts#L33-L37
type Template struct {
	Aliases  map[IndexName]Alias `json:"aliases"`
	Mappings TypeMapping         `json:"mappings"`
	Settings IndexSettings       `json:"settings"`
}

// TemplateBuilder holds Template struct and provides a builder API.
type TemplateBuilder struct {
	v *Template
}

// NewTemplate provides a builder for the Template struct.
func NewTemplateBuilder() *TemplateBuilder {
	r := TemplateBuilder{
		&Template{
			Aliases: make(map[IndexName]Alias, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Template struct
func (rb *TemplateBuilder) Build() Template {
	return *rb.v
}

func (rb *TemplateBuilder) Aliases(values map[IndexName]*AliasBuilder) *TemplateBuilder {
	tmp := make(map[IndexName]Alias, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aliases = tmp
	return rb
}

func (rb *TemplateBuilder) Mappings(mappings *TypeMappingBuilder) *TemplateBuilder {
	v := mappings.Build()
	rb.v.Mappings = v
	return rb
}

func (rb *TemplateBuilder) Settings(settings *IndexSettingsBuilder) *TemplateBuilder {
	v := settings.Build()
	rb.v.Settings = v
	return rb
}
