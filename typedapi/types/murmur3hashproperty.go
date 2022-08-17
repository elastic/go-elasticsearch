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

// Murmur3HashProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/specialized.ts#L66-L68
type Murmur3HashProperty struct {
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// Murmur3HashPropertyBuilder holds Murmur3HashProperty struct and provides a builder API.
type Murmur3HashPropertyBuilder struct {
	v *Murmur3HashProperty
}

// NewMurmur3HashProperty provides a builder for the Murmur3HashProperty struct.
func NewMurmur3HashPropertyBuilder() *Murmur3HashPropertyBuilder {
	r := Murmur3HashPropertyBuilder{
		&Murmur3HashProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "murmur3"

	return &r
}

// Build finalize the chain and returns the Murmur3HashProperty struct
func (rb *Murmur3HashPropertyBuilder) Build() Murmur3HashProperty {
	return *rb.v
}

func (rb *Murmur3HashPropertyBuilder) CopyTo(copyto *FieldsBuilder) *Murmur3HashPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *Murmur3HashPropertyBuilder) DocValues(docvalues bool) *Murmur3HashPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *Murmur3HashPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *Murmur3HashPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *Murmur3HashPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *Murmur3HashPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *Murmur3HashPropertyBuilder) IgnoreAbove(ignoreabove int) *Murmur3HashPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *Murmur3HashPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *Murmur3HashPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *Murmur3HashPropertyBuilder) Meta(value map[string]string) *Murmur3HashPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *Murmur3HashPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *Murmur3HashPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *Murmur3HashPropertyBuilder) Similarity(similarity string) *Murmur3HashPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *Murmur3HashPropertyBuilder) Store(store bool) *Murmur3HashPropertyBuilder {
	rb.v.Store = &store
	return rb
}
