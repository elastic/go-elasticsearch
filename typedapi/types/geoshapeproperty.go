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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoorientation"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geostrategy"
)

// GeoShapeProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/geo.ts#L37-L50
type GeoShapeProperty struct {
	Coerce          *bool                          `json:"coerce,omitempty"`
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	Orientation     *geoorientation.GeoOrientation `json:"orientation,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Strategy        *geostrategy.GeoStrategy       `json:"strategy,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// GeoShapePropertyBuilder holds GeoShapeProperty struct and provides a builder API.
type GeoShapePropertyBuilder struct {
	v *GeoShapeProperty
}

// NewGeoShapeProperty provides a builder for the GeoShapeProperty struct.
func NewGeoShapePropertyBuilder() *GeoShapePropertyBuilder {
	r := GeoShapePropertyBuilder{
		&GeoShapeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "geo_shape"

	return &r
}

// Build finalize the chain and returns the GeoShapeProperty struct
func (rb *GeoShapePropertyBuilder) Build() GeoShapeProperty {
	return *rb.v
}

func (rb *GeoShapePropertyBuilder) Coerce(coerce bool) *GeoShapePropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *GeoShapePropertyBuilder) CopyTo(copyto *FieldsBuilder) *GeoShapePropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *GeoShapePropertyBuilder) DocValues(docvalues bool) *GeoShapePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *GeoShapePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *GeoShapePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *GeoShapePropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *GeoShapePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *GeoShapePropertyBuilder) IgnoreAbove(ignoreabove int) *GeoShapePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *GeoShapePropertyBuilder) IgnoreMalformed(ignoremalformed bool) *GeoShapePropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

func (rb *GeoShapePropertyBuilder) IgnoreZValue(ignorezvalue bool) *GeoShapePropertyBuilder {
	rb.v.IgnoreZValue = &ignorezvalue
	return rb
}

func (rb *GeoShapePropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *GeoShapePropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

func (rb *GeoShapePropertyBuilder) Meta(value map[string]string) *GeoShapePropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *GeoShapePropertyBuilder) Orientation(orientation geoorientation.GeoOrientation) *GeoShapePropertyBuilder {
	rb.v.Orientation = &orientation
	return rb
}

func (rb *GeoShapePropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *GeoShapePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *GeoShapePropertyBuilder) Similarity(similarity string) *GeoShapePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *GeoShapePropertyBuilder) Store(store bool) *GeoShapePropertyBuilder {
	rb.v.Store = &store
	return rb
}

func (rb *GeoShapePropertyBuilder) Strategy(strategy geostrategy.GeoStrategy) *GeoShapePropertyBuilder {
	rb.v.Strategy = &strategy
	return rb
}
