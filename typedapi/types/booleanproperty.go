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

// BooleanProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/core.ts#L53-L59
type BooleanProperty struct {
	Boost         *float64                       `json:"boost,omitempty"`
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fielddata     *NumericFielddata              `json:"fielddata,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	Index         *bool                          `json:"index,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	NullValue     *bool                          `json:"null_value,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// BooleanPropertyBuilder holds BooleanProperty struct and provides a builder API.
type BooleanPropertyBuilder struct {
	v *BooleanProperty
}

// NewBooleanProperty provides a builder for the BooleanProperty struct.
func NewBooleanPropertyBuilder() *BooleanPropertyBuilder {
	r := BooleanPropertyBuilder{
		&BooleanProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "boolean"

	return &r
}

// Build finalize the chain and returns the BooleanProperty struct
func (rb *BooleanPropertyBuilder) Build() BooleanProperty {
	return *rb.v
}

func (rb *BooleanPropertyBuilder) Boost(boost float64) *BooleanPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *BooleanPropertyBuilder) CopyTo(copyto *FieldsBuilder) *BooleanPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *BooleanPropertyBuilder) DocValues(docvalues bool) *BooleanPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *BooleanPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *BooleanPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *BooleanPropertyBuilder) Fielddata(fielddata *NumericFielddataBuilder) *BooleanPropertyBuilder {
	v := fielddata.Build()
	rb.v.Fielddata = &v
	return rb
}

func (rb *BooleanPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *BooleanPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *BooleanPropertyBuilder) IgnoreAbove(ignoreabove int) *BooleanPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *BooleanPropertyBuilder) Index(index bool) *BooleanPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *BooleanPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *BooleanPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *BooleanPropertyBuilder) Meta(value map[string]string) *BooleanPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *BooleanPropertyBuilder) NullValue(nullvalue bool) *BooleanPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *BooleanPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *BooleanPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *BooleanPropertyBuilder) Similarity(similarity string) *BooleanPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *BooleanPropertyBuilder) Store(store bool) *BooleanPropertyBuilder {
	rb.v.Store = &store
	return rb
}
