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

// DenseVectorProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/complex.ts#L50-L56
type DenseVectorProperty struct {
	Dims          int                            `json:"dims"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	Index         *bool                          `json:"index,omitempty"`
	IndexOptions  *DenseVectorIndexOptions       `json:"index_options,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// DenseVectorPropertyBuilder holds DenseVectorProperty struct and provides a builder API.
type DenseVectorPropertyBuilder struct {
	v *DenseVectorProperty
}

// NewDenseVectorProperty provides a builder for the DenseVectorProperty struct.
func NewDenseVectorPropertyBuilder() *DenseVectorPropertyBuilder {
	r := DenseVectorPropertyBuilder{
		&DenseVectorProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "dense_vector"

	return &r
}

// Build finalize the chain and returns the DenseVectorProperty struct
func (rb *DenseVectorPropertyBuilder) Build() DenseVectorProperty {
	return *rb.v
}

func (rb *DenseVectorPropertyBuilder) Dims(dims int) *DenseVectorPropertyBuilder {
	rb.v.Dims = dims
	return rb
}

func (rb *DenseVectorPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *DenseVectorPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *DenseVectorPropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *DenseVectorPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *DenseVectorPropertyBuilder) IgnoreAbove(ignoreabove int) *DenseVectorPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *DenseVectorPropertyBuilder) Index(index bool) *DenseVectorPropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *DenseVectorPropertyBuilder) IndexOptions(indexoptions *DenseVectorIndexOptionsBuilder) *DenseVectorPropertyBuilder {
	v := indexoptions.Build()
	rb.v.IndexOptions = &v
	return rb
}

func (rb *DenseVectorPropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *DenseVectorPropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *DenseVectorPropertyBuilder) Meta(value map[string]string) *DenseVectorPropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *DenseVectorPropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *DenseVectorPropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *DenseVectorPropertyBuilder) Similarity(similarity string) *DenseVectorPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}
