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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _spanFieldMaskingQuery struct {
	v *types.SpanFieldMaskingQuery
}

// Wrapper to allow span queries to participate in composite single-field span
// queries by _lying_ about their search field.
func NewSpanFieldMaskingQuery(query types.SpanQueryVariant) *_spanFieldMaskingQuery {

	tmp := &_spanFieldMaskingQuery{v: types.NewSpanFieldMaskingQuery()}

	tmp.Query(query)

	return tmp

}

func (s *_spanFieldMaskingQuery) Field(field string) *_spanFieldMaskingQuery {

	s.v.Field = field

	return s
}

func (s *_spanFieldMaskingQuery) Query(query types.SpanQueryVariant) *_spanFieldMaskingQuery {

	s.v.Query = *query.SpanQueryCaster()

	return s
}

func (s *_spanFieldMaskingQuery) Boost(boost float32) *_spanFieldMaskingQuery {

	s.v.Boost = &boost

	return s
}

func (s *_spanFieldMaskingQuery) QueryName_(queryname_ string) *_spanFieldMaskingQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_spanFieldMaskingQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.SpanFieldMasking = s.v

	return container
}

func (s *_spanFieldMaskingQuery) SpanQueryCaster() *types.SpanQuery {
	container := types.NewSpanQuery()

	container.SpanFieldMasking = s.v

	return container
}

func (s *_spanFieldMaskingQuery) SpanFieldMaskingQueryCaster() *types.SpanFieldMaskingQuery {
	return s.v
}
