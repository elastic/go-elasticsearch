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

// PropertyBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/Property.ts#L81-L89
type PropertyBase struct {
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
}

// PropertyBaseBuilder holds PropertyBase struct and provides a builder API.
type PropertyBaseBuilder struct {
	v *PropertyBase
}

// NewPropertyBase provides a builder for the PropertyBase struct.
func NewPropertyBaseBuilder() *PropertyBaseBuilder {
	r := PropertyBaseBuilder{
		&PropertyBase{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the PropertyBase struct
func (rb *PropertyBaseBuilder) Build() PropertyBase {
	return *rb.v
}

func (rb *PropertyBaseBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *PropertyBaseBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *PropertyBaseBuilder) Fields(values map[PropertyName]*PropertyBuilder) *PropertyBaseBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *PropertyBaseBuilder) IgnoreAbove(ignoreabove int) *PropertyBaseBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *PropertyBaseBuilder) LocalMetadata(localmetadata *MetadataBuilder) *PropertyBaseBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *PropertyBaseBuilder) Meta(value map[string]string) *PropertyBaseBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *PropertyBaseBuilder) Properties(values map[PropertyName]*PropertyBuilder) *PropertyBaseBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}
