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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// ConstantKeywordProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/mapping/specialized.ts#L44-L47
type ConstantKeywordProperty struct {
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	// Meta Metadata about the field.
	Meta       map[string]string         `json:"meta,omitempty"`
	Properties map[PropertyName]Property `json:"properties,omitempty"`
	Type       string                    `json:"type,omitempty"`
	Value      interface{}               `json:"value,omitempty"`
}

// ConstantKeywordPropertyBuilder holds ConstantKeywordProperty struct and provides a builder API.
type ConstantKeywordPropertyBuilder struct {
	v *ConstantKeywordProperty
}

// NewConstantKeywordProperty provides a builder for the ConstantKeywordProperty struct.
func NewConstantKeywordPropertyBuilder() *ConstantKeywordPropertyBuilder {
	r := ConstantKeywordPropertyBuilder{
		&ConstantKeywordProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "constant_keyword"

	return &r
}

// Build finalize the chain and returns the ConstantKeywordProperty struct
func (rb *ConstantKeywordPropertyBuilder) Build() ConstantKeywordProperty {
	return *rb.v
}

func (rb *ConstantKeywordPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *ConstantKeywordPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *ConstantKeywordPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *ConstantKeywordPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *ConstantKeywordPropertyBuilder) IgnoreAbove(ignoreabove int) *ConstantKeywordPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *ConstantKeywordPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *ConstantKeywordPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

// Meta Metadata about the field.

func (rb *ConstantKeywordPropertyBuilder) Meta(value map[string]string) *ConstantKeywordPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *ConstantKeywordPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *ConstantKeywordPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *ConstantKeywordPropertyBuilder) Value(value interface{}) *ConstantKeywordPropertyBuilder {
	rb.v.Value = value
	return rb
}
