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

// IntegerRangeProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/range.ts#L42-L44
type IntegerRangeProperty struct {
	Boost         *float64                       `json:"boost,omitempty"`
	Coerce        *bool                          `json:"coerce,omitempty"`
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	Index         *bool                          `json:"index,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// IntegerRangePropertyBuilder holds IntegerRangeProperty struct and provides a builder API.
type IntegerRangePropertyBuilder struct {
	v *IntegerRangeProperty
}

// NewIntegerRangeProperty provides a builder for the IntegerRangeProperty struct.
func NewIntegerRangePropertyBuilder() *IntegerRangePropertyBuilder {
	r := IntegerRangePropertyBuilder{
		&IntegerRangeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "integer_range"

	return &r
}

// Build finalize the chain and returns the IntegerRangeProperty struct
func (rb *IntegerRangePropertyBuilder) Build() IntegerRangeProperty {
	return *rb.v
}

func (rb *IntegerRangePropertyBuilder) Boost(boost float64) *IntegerRangePropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *IntegerRangePropertyBuilder) Coerce(coerce bool) *IntegerRangePropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *IntegerRangePropertyBuilder) CopyTo(copyto *FieldsBuilder) *IntegerRangePropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *IntegerRangePropertyBuilder) DocValues(docvalues bool) *IntegerRangePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *IntegerRangePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *IntegerRangePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *IntegerRangePropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *IntegerRangePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *IntegerRangePropertyBuilder) IgnoreAbove(ignoreabove int) *IntegerRangePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *IntegerRangePropertyBuilder) Index(index bool) *IntegerRangePropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *IntegerRangePropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *IntegerRangePropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *IntegerRangePropertyBuilder) Meta(value map[string]string) *IntegerRangePropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *IntegerRangePropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *IntegerRangePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *IntegerRangePropertyBuilder) Similarity(similarity string) *IntegerRangePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *IntegerRangePropertyBuilder) Store(store bool) *IntegerRangePropertyBuilder {
	rb.v.Store = &store
	return rb
}
