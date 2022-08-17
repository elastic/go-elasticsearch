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

// PointProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/geo.ts#L62-L67
type PointProperty struct {
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	NullValue       *string                        `json:"null_value,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// PointPropertyBuilder holds PointProperty struct and provides a builder API.
type PointPropertyBuilder struct {
	v *PointProperty
}

// NewPointProperty provides a builder for the PointProperty struct.
func NewPointPropertyBuilder() *PointPropertyBuilder {
	r := PointPropertyBuilder{
		&PointProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "point"

	return &r
}

// Build finalize the chain and returns the PointProperty struct
func (rb *PointPropertyBuilder) Build() PointProperty {
	return *rb.v
}

func (rb *PointPropertyBuilder) CopyTo(copyto *FieldsBuilder) *PointPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *PointPropertyBuilder) DocValues(docvalues bool) *PointPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *PointPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *PointPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *PointPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *PointPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *PointPropertyBuilder) IgnoreAbove(ignoreabove int) *PointPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *PointPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *PointPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *PointPropertyBuilder) IgnoreZValue(ignorezvalue bool) *PointPropertyBuilder {
	rb.v.IgnoreZValue = &ignorezvalue
	return rb
}

func (rb *PointPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *PointPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *PointPropertyBuilder) Meta(value map[string]string) *PointPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *PointPropertyBuilder) NullValue(nullvalue string) *PointPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *PointPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *PointPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *PointPropertyBuilder) Similarity(similarity string) *PointPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *PointPropertyBuilder) Store(store bool) *PointPropertyBuilder {
	rb.v.Store = &store
	return rb
}
