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

// WildcardProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L258-L262
type WildcardProperty struct {
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	NullValue     *string                        `json:"null_value,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// WildcardPropertyBuilder holds WildcardProperty struct and provides a builder API.
type WildcardPropertyBuilder struct {
	v *WildcardProperty
}

// NewWildcardProperty provides a builder for the WildcardProperty struct.
func NewWildcardPropertyBuilder() *WildcardPropertyBuilder {
	r := WildcardPropertyBuilder{
		&WildcardProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "wildcard"

	return &r
}

// Build finalize the chain and returns the WildcardProperty struct
func (rb *WildcardPropertyBuilder) Build() WildcardProperty {
	return *rb.v
}

func (rb *WildcardPropertyBuilder) CopyTo(copyto *FieldsBuilder) *WildcardPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *WildcardPropertyBuilder) DocValues(docvalues bool) *WildcardPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *WildcardPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *WildcardPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *WildcardPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *WildcardPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *WildcardPropertyBuilder) IgnoreAbove(ignoreabove int) *WildcardPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *WildcardPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *WildcardPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *WildcardPropertyBuilder) Meta(value map[string]string) *WildcardPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *WildcardPropertyBuilder) NullValue(nullvalue string) *WildcardPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *WildcardPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *WildcardPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *WildcardPropertyBuilder) Similarity(similarity string) *WildcardPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *WildcardPropertyBuilder) Store(store bool) *WildcardPropertyBuilder {
	rb.v.Store = &store
	return rb
}
