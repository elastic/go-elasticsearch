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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _percolateQuery struct {
	v *types.PercolateQuery
}

// Matches queries stored in an index.
func NewPercolateQuery() *_percolateQuery {

	return &_percolateQuery{v: types.NewPercolateQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_percolateQuery) Boost(boost float32) *_percolateQuery {

	s.v.Boost = &boost

	return s
}

// The source of the document being percolated.
func (s *_percolateQuery) Document(document json.RawMessage) *_percolateQuery {

	s.v.Document = document

	return s
}

// An array of sources of the documents being percolated.
func (s *_percolateQuery) Documents(documents ...json.RawMessage) *_percolateQuery {

	for _, v := range documents {

		s.v.Documents = append(s.v.Documents, v)

	}
	return s
}

// Field that holds the indexed queries. The field must use the `percolator`
// mapping type.
func (s *_percolateQuery) Field(field string) *_percolateQuery {

	s.v.Field = field

	return s
}

// The ID of a stored document to percolate.
func (s *_percolateQuery) Id(id string) *_percolateQuery {

	s.v.Id = &id

	return s
}

// The index of a stored document to percolate.
func (s *_percolateQuery) Index(indexname string) *_percolateQuery {

	s.v.Index = &indexname

	return s
}

// The suffix used for the `_percolator_document_slot` field when multiple
// `percolate` queries are specified.
func (s *_percolateQuery) Name(name string) *_percolateQuery {

	s.v.Name = &name

	return s
}

// Preference used to fetch document to percolate.
func (s *_percolateQuery) Preference(preference string) *_percolateQuery {

	s.v.Preference = &preference

	return s
}

func (s *_percolateQuery) QueryName_(queryname_ string) *_percolateQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Routing used to fetch document to percolate.
func (s *_percolateQuery) Routing(routing string) *_percolateQuery {

	s.v.Routing = &routing

	return s
}

// The expected version of a stored document to percolate.
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
