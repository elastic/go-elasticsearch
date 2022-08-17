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

// GeoPointProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/geo.ts#L23-L28
type GeoPointProperty struct {
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	NullValue       *GeoLocation                   `json:"null_value,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// GeoPointPropertyBuilder holds GeoPointProperty struct and provides a builder API.
type GeoPointPropertyBuilder struct {
	v *GeoPointProperty
}

// NewGeoPointProperty provides a builder for the GeoPointProperty struct.
func NewGeoPointPropertyBuilder() *GeoPointPropertyBuilder {
	r := GeoPointPropertyBuilder{
		&GeoPointProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "geo_point"

	return &r
}

// Build finalize the chain and returns the GeoPointProperty struct
func (rb *GeoPointPropertyBuilder) Build() GeoPointProperty {
	return *rb.v
}

func (rb *GeoPointPropertyBuilder) CopyTo(copyto *FieldsBuilder) *GeoPointPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *GeoPointPropertyBuilder) DocValues(docvalues bool) *GeoPointPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *GeoPointPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *GeoPointPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *GeoPointPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *GeoPointPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *GeoPointPropertyBuilder) IgnoreAbove(ignoreabove int) *GeoPointPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *GeoPointPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *GeoPointPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *GeoPointPropertyBuilder) IgnoreZValue(ignorezvalue bool) *GeoPointPropertyBuilder {
	rb.v.IgnoreZValue = &ignorezvalue
	return rb
}

func (rb *GeoPointPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *GeoPointPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *GeoPointPropertyBuilder) Meta(value map[string]string) *GeoPointPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *GeoPointPropertyBuilder) NullValue(nullvalue *GeoLocationBuilder) *GeoPointPropertyBuilder {
	v := nullvalue.Build()
	rb.v.NullValue = &v
	return rb
}

func (rb *GeoPointPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *GeoPointPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *GeoPointPropertyBuilder) Similarity(similarity string) *GeoPointPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *GeoPointPropertyBuilder) Store(store bool) *GeoPointPropertyBuilder {
	rb.v.Store = &store
	return rb
}
