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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _percolateQuery struct {
	v *types.PercolateQuery
}

// Matches queries stored in an index.
func NewPercolateQuery() *_percolateQuery {

	return &_percolateQuery{v: types.NewPercolateQuery()}

}

func (s *_percolateQuery) Boost(boost float32) *_percolateQuery {

	s.v.Boost = &boost

	return s
}

func (s *_percolateQuery) Document(document json.RawMessage) *_percolateQuery {

	s.v.Document = document

	return s
}

func (s *_percolateQuery) Documents(documents ...json.RawMessage) *_percolateQuery {

	for _, v := range documents {

		s.v.Documents = append(s.v.Documents, v)

	}
	return s
}

func (s *_percolateQuery) Field(field string) *_percolateQuery {

	s.v.Field = field

	return s
}

func (s *_percolateQuery) Id(id string) *_percolateQuery {

	s.v.Id = &id

	return s
}

func (s *_percolateQuery) Index(indexname string) *_percolateQuery {

	s.v.Index = &indexname

	return s
}

func (s *_percolateQuery) Name(name string) *_percolateQuery {

	s.v.Name = &name

	return s
}

func (s *_percolateQuery) Preference(preference string) *_percolateQuery {

	s.v.Preference = &preference

	return s
}

func (s *_percolateQuery) QueryName_(queryname_ string) *_percolateQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_percolateQuery) Routing(routing string) *_percolateQuery {

	s.v.Routing = &routing

	return s
}

func (s *_percolateQuery) Version(versionnumber int64) *_percolateQuery {

	s.v.Version = &versionnumber

	return s
}

func (s *_percolateQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Percolate = s.v

	return container
}

func (s *_percolateQuery) PercolateQueryCaster() *types.PercolateQuery {
	return s.v
}
