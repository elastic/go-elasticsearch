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
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoshaperelation"
)

// ShapeFieldQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/_types/query_dsl/specialized.ts#L183-L187
type ShapeFieldQuery struct {
	IndexedShape *FieldLookup                       `json:"indexed_shape,omitempty"`
	Relation     *geoshaperelation.GeoShapeRelation `json:"relation,omitempty"`
	Shape        *GeoShape                          `json:"shape,omitempty"`
}

// ShapeFieldQueryBuilder holds ShapeFieldQuery struct and provides a builder API.
type ShapeFieldQueryBuilder struct {
	v *ShapeFieldQuery
}

// NewShapeFieldQuery provides a builder for the ShapeFieldQuery struct.
func NewShapeFieldQueryBuilder() *ShapeFieldQueryBuilder {
	r := ShapeFieldQueryBuilder{
		&ShapeFieldQuery{},
	}

	return &r
}

// Build finalize the chain and returns the ShapeFieldQuery struct
func (rb *ShapeFieldQueryBuilder) Build() ShapeFieldQuery {
	return *rb.v
}

func (rb *ShapeFieldQueryBuilder) IndexedShape(indexedshape *FieldLookupBuilder) *ShapeFieldQueryBuilder {
	v := indexedshape.Build()
	rb.v.IndexedShape = &v
	return rb
}

func (rb *ShapeFieldQueryBuilder) Relation(relation geoshaperelation.GeoShapeRelation) *ShapeFieldQueryBuilder {
	rb.v.Relation = &relation
	return rb
}

func (rb *ShapeFieldQueryBuilder) Shape(shape *GeoShapeBuilder) *ShapeFieldQueryBuilder {
	v := shape.Build()
	rb.v.Shape = &v
	return rb
}
