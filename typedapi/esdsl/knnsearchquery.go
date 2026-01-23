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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _knnSearchQuery struct {
	v *types.KnnSearchQuery
}

func NewKnnSearchQuery(k int, numcandidates int) *_knnSearchQuery {

	tmp := &_knnSearchQuery{v: types.NewKnnSearchQuery()}

	tmp.K(k)

	tmp.NumCandidates(numcandidates)

	return tmp

}

func (s *_knnSearchQuery) Field(field string) *_knnSearchQuery {

	s.v.Field = field

	return s
}

func (s *_knnSearchQuery) K(k int) *_knnSearchQuery {

	s.v.K = k

	return s
}

func (s *_knnSearchQuery) NumCandidates(numcandidates int) *_knnSearchQuery {

	s.v.NumCandidates = numcandidates

	return s
}

func (s *_knnSearchQuery) QueryVector(queryvectors ...float32) *_knnSearchQuery {

	s.v.QueryVector = queryvectors

	return s
}

func (s *_knnSearchQuery) KnnSearchQueryCaster() *types.KnnSearchQuery {
	return s.v
}
