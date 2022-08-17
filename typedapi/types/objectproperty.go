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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// ObjectProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/complex.ts#L45-L48
type ObjectProperty struct {
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Enabled       *bool                          `json:"enabled,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// ObjectPropertyBuilder holds ObjectProperty struct and provides a builder API.
type ObjectPropertyBuilder struct {
	v *ObjectProperty
}

// NewObjectProperty provides a builder for the ObjectProperty struct.
func NewObjectPropertyBuilder() *ObjectPropertyBuilder {
	r := ObjectPropertyBuilder{
		&ObjectProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "object"

	return &r
}

// Build finalize the chain and returns the ObjectProperty struct
func (rb *ObjectPropertyBuilder) Build() ObjectProperty {
	return *rb.v
}

func (rb *ObjectPropertyBuilder) CopyTo(copyto *FieldsBuilder) *ObjectPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *ObjectPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *ObjectPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *ObjectPropertyBuilder) Enabled(enabled bool) *ObjectPropertyBuilder {
	rb.v.Enabled = &enabled
	return rb
}

func (rb *ObjectPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *ObjectPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *ObjectPropertyBuilder) IgnoreAbove(ignoreabove int) *ObjectPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *ObjectPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *ObjectPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *ObjectPropertyBuilder) Meta(value map[string]string) *ObjectPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *ObjectPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *ObjectPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *ObjectPropertyBuilder) Similarity(similarity string) *ObjectPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *ObjectPropertyBuilder) Store(store bool) *ObjectPropertyBuilder {
	rb.v.Store = &store
	return rb
}
