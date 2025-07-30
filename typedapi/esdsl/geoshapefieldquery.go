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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/geoshaperelation"
)

type _geoShapeFieldQuery struct {
	v *types.GeoShapeFieldQuery
}

func NewGeoShapeFieldQuery() *_geoShapeFieldQuery {

	return &_geoShapeFieldQuery{v: types.NewGeoShapeFieldQuery()}

}

func (s *_geoShapeFieldQuery) IndexedShape(indexedshape types.FieldLookupVariant) *_geoShapeFieldQuery {

	s.v.IndexedShape = indexedshape.FieldLookupCaster()

	return s
}

func (s *_geoShapeFieldQuery) Relation(relation geoshaperelation.GeoShapeRelation) *_geoShapeFieldQuery {

	s.v.Relation = &relation
	return s
}

func (s *_geoShapeFieldQuery) Shape(geoshape json.RawMessage) *_geoShapeFieldQuery {

	s.v.Shape = geoshape

	return s
}

func (s *_geoShapeFieldQuery) GeoShapeFieldQueryCaster() *types.GeoShapeFieldQuery {
	return s.v
}
