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

// IpRangeProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/range.ts#L46-L48
type IpRangeProperty struct {
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

// IpRangePropertyBuilder holds IpRangeProperty struct and provides a builder API.
type IpRangePropertyBuilder struct {
	v *IpRangeProperty
}

// NewIpRangeProperty provides a builder for the IpRangeProperty struct.
func NewIpRangePropertyBuilder() *IpRangePropertyBuilder {
	r := IpRangePropertyBuilder{
		&IpRangeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "ip_range"

	return &r
}

// Build finalize the chain and returns the IpRangeProperty struct
func (rb *IpRangePropertyBuilder) Build() IpRangeProperty {
	return *rb.v
}

func (rb *IpRangePropertyBuilder) Boost(boost float64) *IpRangePropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *IpRangePropertyBuilder) Coerce(coerce bool) *IpRangePropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *IpRangePropertyBuilder) CopyTo(copyto *FieldsBuilder) *IpRangePropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *IpRangePropertyBuilder) DocValues(docvalues bool) *IpRangePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *IpRangePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *IpRangePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *IpRangePropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *IpRangePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *IpRangePropertyBuilder) IgnoreAbove(ignoreabove int) *IpRangePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *IpRangePropertyBuilder) Index(index bool) *IpRangePropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *IpRangePropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *IpRangePropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *IpRangePropertyBuilder) Meta(value map[string]string) *IpRangePropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *IpRangePropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *IpRangePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *IpRangePropertyBuilder) Similarity(similarity string) *IpRangePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *IpRangePropertyBuilder) Store(store bool) *IpRangePropertyBuilder {
	rb.v.Store = &store
	return rb
}
