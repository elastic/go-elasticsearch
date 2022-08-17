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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexoptions"
)

// FlattenedProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/complex.ts#L25-L36
type FlattenedProperty struct {
	Boost                    *float64                       `json:"boost,omitempty"`
	DepthLimit               *int                           `json:"depth_limit,omitempty"`
	DocValues                *bool                          `json:"doc_values,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals      *bool                          `json:"eager_global_ordinals,omitempty"`
	Fields                   map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove              *int                           `json:"ignore_above,omitempty"`
	Index                    *bool                          `json:"index,omitempty"`
	IndexOptions             *indexoptions.IndexOptions     `json:"index_options,omitempty"`
	LocalMetadata            *Metadata                      `json:"local_metadata,omitempty"`
	Meta                     map[string]string              `json:"meta,omitempty"`
	NullValue                *string                        `json:"null_value,omitempty"`
	Properties               map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity               *string                        `json:"similarity,omitempty"`
	SplitQueriesOnWhitespace *bool                          `json:"split_queries_on_whitespace,omitempty"`
	Type                     string                         `json:"type,omitempty"`
}

// FlattenedPropertyBuilder holds FlattenedProperty struct and provides a builder API.
type FlattenedPropertyBuilder struct {
	v *FlattenedProperty
}

// NewFlattenedProperty provides a builder for the FlattenedProperty struct.
func NewFlattenedPropertyBuilder() *FlattenedPropertyBuilder {
	r := FlattenedPropertyBuilder{
		&FlattenedProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "flattened"

	return &r
}

// Build finalize the chain and returns the FlattenedProperty struct
func (rb *FlattenedPropertyBuilder) Build() FlattenedProperty {
	return *rb.v
}

func (rb *FlattenedPropertyBuilder) Boost(boost float64) *FlattenedPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *FlattenedPropertyBuilder) DepthLimit(depthlimit int) *FlattenedPropertyBuilder {
	rb.v.DepthLimit = &depthlimit
	return rb
}

func (rb *FlattenedPropertyBuilder) DocValues(docvalues bool) *FlattenedPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *FlattenedPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *FlattenedPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *FlattenedPropertyBuilder) EagerGlobalOrdinals(eagerglobalordinals bool) *FlattenedPropertyBuilder {
	rb.v.EagerGlobalOrdinals = &eagerglobalordinals
	return rb
}

func (rb *FlattenedPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *FlattenedPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *FlattenedPropertyBuilder) IgnoreAbove(ignoreabove int) *FlattenedPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *FlattenedPropertyBuilder) Index(index bool) *FlattenedPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *FlattenedPropertyBuilder) IndexOptions(indexoptions indexoptions.IndexOptions) *FlattenedPropertyBuilder {
	rb.v.IndexOptions = &indexoptions
	return rb
}

func (rb *FlattenedPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *FlattenedPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *FlattenedPropertyBuilder) Meta(value map[string]string) *FlattenedPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *FlattenedPropertyBuilder) NullValue(nullvalue string) *FlattenedPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *FlattenedPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *FlattenedPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *FlattenedPropertyBuilder) Similarity(similarity string) *FlattenedPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *FlattenedPropertyBuilder) SplitQueriesOnWhitespace(splitqueriesonwhitespace bool) *FlattenedPropertyBuilder {
	rb.v.SplitQueriesOnWhitespace = &splitqueriesonwhitespace
	return rb
}
