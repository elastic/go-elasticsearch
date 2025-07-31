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

type _spanTermQuery struct {
	k string
	v *types.SpanTermQuery
}

// Matches spans containing a term.
func NewSpanTermQuery(field string, value types.FieldValueVariant) *_spanTermQuery {
	tmp := &_spanTermQuery{
		k: field,
		v: types.NewSpanTermQuery(),
	}

	tmp.Value(value)
	return tmp
}

func (s *_spanTermQuery) Boost(boost float32) *_spanTermQuery {

	s.v.Boost = &boost

	return s
}

func (s *_spanTermQuery) QueryName_(queryname_ string) *_spanTermQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_spanTermQuery) Value(fieldvalue types.FieldValueVariant) *_spanTermQuery {

	s.v.Value = *fieldvalue.FieldValueCaster()

	return s
}

func (s *_spanTermQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.SpanTerm = map[string]types.SpanTermQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_spanTermQuery) SpanQueryCaster() *types.SpanQuery {
	container := types.NewSpanQuery()
	container.SpanTerm = map[string]types.SpanTermQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleSpanTermQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleSpanTermQuery() *_spanTermQuery {
	return &_spanTermQuery{
		k: "",
		v: types.NewSpanTermQuery(),
	}
}

func (s *_spanTermQuery) SpanTermQueryCaster() *types.SpanTermQuery {
	return s.v.SpanTermQueryCaster()
}
