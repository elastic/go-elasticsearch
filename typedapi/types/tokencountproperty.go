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

// TokenCountProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/specialized.ts#L70-L77
type TokenCountProperty struct {
	Analyzer                 *string                        `json:"analyzer,omitempty"`
	Boost                    *float64                       `json:"boost,omitempty"`
	CopyTo                   *Fields                        `json:"copy_to,omitempty"`
	DocValues                *bool                          `json:"doc_values,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EnablePositionIncrements *bool                          `json:"enable_position_increments,omitempty"`
	Fields                   map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove              *int                           `json:"ignore_above,omitempty"`
	Index                    *bool                          `json:"index,omitempty"`
	LocalMetadata            *Metadata                      `json:"local_metadata,omitempty"`
	Meta                     map[string]string              `json:"meta,omitempty"`
	NullValue                *float64                       `json:"null_value,omitempty"`
	Properties               map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity               *string                        `json:"similarity,omitempty"`
	Store                    *bool                          `json:"store,omitempty"`
	Type                     string                         `json:"type,omitempty"`
}

// TokenCountPropertyBuilder holds TokenCountProperty struct and provides a builder API.
type TokenCountPropertyBuilder struct {
	v *TokenCountProperty
}

// NewTokenCountProperty provides a builder for the TokenCountProperty struct.
func NewTokenCountPropertyBuilder() *TokenCountPropertyBuilder {
	r := TokenCountPropertyBuilder{
		&TokenCountProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "token_count"

	return &r
}

// Build finalize the chain and returns the TokenCountProperty struct
func (rb *TokenCountPropertyBuilder) Build() TokenCountProperty {
	return *rb.v
}

func (rb *TokenCountPropertyBuilder) Analyzer(analyzer string) *TokenCountPropertyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *TokenCountPropertyBuilder) Boost(boost float64) *TokenCountPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *TokenCountPropertyBuilder) CopyTo(copyto *FieldsBuilder) *TokenCountPropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *TokenCountPropertyBuilder) DocValues(docvalues bool) *TokenCountPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *TokenCountPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *TokenCountPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *TokenCountPropertyBuilder) EnablePositionIncrements(enablepositionincrements bool) *TokenCountPropertyBuilder {
	rb.v.EnablePositionIncrements = &enablepositionincrements
	return rb
}

func (rb *TokenCountPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *TokenCountPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *TokenCountPropertyBuilder) IgnoreAbove(ignoreabove int) *TokenCountPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *TokenCountPropertyBuilder) Index(index bool) *TokenCountPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *TokenCountPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *TokenCountPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *TokenCountPropertyBuilder) Meta(value map[string]string) *TokenCountPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *TokenCountPropertyBuilder) NullValue(nullvalue float64) *TokenCountPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

func (rb *TokenCountPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *TokenCountPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *TokenCountPropertyBuilder) Similarity(similarity string) *TokenCountPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *TokenCountPropertyBuilder) Store(store bool) *TokenCountPropertyBuilder {
	rb.v.Store = &store
	return rb
}
