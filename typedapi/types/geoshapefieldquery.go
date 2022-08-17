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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoshaperelation"
)

// GeoShapeFieldQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/geo.ts#L78-L82
type GeoShapeFieldQuery struct {
	IndexedShape *FieldLookup                       `json:"indexed_shape,omitempty"`
	Relation     *geoshaperelation.GeoShapeRelation `json:"relation,omitempty"`
	Shape        *GeoShape                          `json:"shape,omitempty"`
}

// GeoShapeFieldQueryBuilder holds GeoShapeFieldQuery struct and provides a builder API.
type GeoShapeFieldQueryBuilder struct {
	v *GeoShapeFieldQuery
}

// NewGeoShapeFieldQuery provides a builder for the GeoShapeFieldQuery struct.
func NewGeoShapeFieldQueryBuilder() *GeoShapeFieldQueryBuilder {
	r := GeoShapeFieldQueryBuilder{
		&GeoShapeFieldQuery{},
	}

	return &r
}

// Build finalize the chain and returns the GeoShapeFieldQuery struct
func (rb *GeoShapeFieldQueryBuilder) Build() GeoShapeFieldQuery {
	return *rb.v
}

func (rb *GeoShapeFieldQueryBuilder) IndexedShape(indexedshape *FieldLookupBuilder) *GeoShapeFieldQueryBuilder {
	v := indexedshape.Build()
	rb.v.IndexedShape = &v
	return rb
}

func (rb *GeoShapeFieldQueryBuilder) Relation(relation geoshaperelation.GeoShapeRelation) *GeoShapeFieldQueryBuilder {
	rb.v.Relation = &relation
	return rb
}

func (rb *GeoShapeFieldQueryBuilder) Shape(shape *GeoShapeBuilder) *GeoShapeFieldQueryBuilder {
	v := shape.Build()
	rb.v.Shape = &v
	return rb
}
