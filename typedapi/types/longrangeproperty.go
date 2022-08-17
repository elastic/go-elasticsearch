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

// LongRangeProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/range.ts#L50-L52
type LongRangeProperty struct {
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

// LongRangePropertyBuilder holds LongRangeProperty struct and provides a builder API.
type LongRangePropertyBuilder struct {
	v *LongRangeProperty
}

// NewLongRangeProperty provides a builder for the LongRangeProperty struct.
func NewLongRangePropertyBuilder() *LongRangePropertyBuilder {
	r := LongRangePropertyBuilder{
		&LongRangeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "long_range"

	return &r
}

// Build finalize the chain and returns the LongRangeProperty struct
func (rb *LongRangePropertyBuilder) Build() LongRangeProperty {
	return *rb.v
}

func (rb *LongRangePropertyBuilder) Boost(boost float64) *LongRangePropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *LongRangePropertyBuilder) Coerce(coerce bool) *LongRangePropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *LongRangePropertyBuilder) CopyTo(copyto *FieldsBuilder) *LongRangePropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *LongRangePropertyBuilder) DocValues(docvalues bool) *LongRangePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *LongRangePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *LongRangePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *LongRangePropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *LongRangePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *LongRangePropertyBuilder) IgnoreAbove(ignoreabove int) *LongRangePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *LongRangePropertyBuilder) Index(index bool) *LongRangePropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *LongRangePropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *LongRangePropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *LongRangePropertyBuilder) Meta(value map[string]string) *LongRangePropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *LongRangePropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *LongRangePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *LongRangePropertyBuilder) Similarity(similarity string) *LongRangePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *LongRangePropertyBuilder) Store(store bool) *LongRangePropertyBuilder {
	rb.v.Store = &store
	return rb
}
