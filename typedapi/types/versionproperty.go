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

// VersionProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L254-L256
type VersionProperty struct {
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

// VersionPropertyBuilder holds VersionProperty struct and provides a builder API.
type VersionPropertyBuilder struct {
	v *VersionProperty
}

// NewVersionProperty provides a builder for the VersionProperty struct.
func NewVersionPropertyBuilder() *VersionPropertyBuilder {
	r := VersionPropertyBuilder{
		&VersionProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "version"

	return &r
}

// Build finalize the chain and returns the VersionProperty struct
func (rb *VersionPropertyBuilder) Build() VersionProperty {
	return *rb.v
}

func (rb *VersionPropertyBuilder) CopyTo(copyto *FieldsBuilder) *VersionPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *VersionPropertyBuilder) DocValues(docvalues bool) *VersionPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *VersionPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *VersionPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *VersionPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *VersionPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *VersionPropertyBuilder) IgnoreAbove(ignoreabove int) *VersionPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *VersionPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *VersionPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *VersionPropertyBuilder) Meta(value map[string]string) *VersionPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *VersionPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *VersionPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *VersionPropertyBuilder) Similarity(similarity string) *VersionPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *VersionPropertyBuilder) Store(store bool) *VersionPropertyBuilder {
	rb.v.Store = &store
	return rb
}
