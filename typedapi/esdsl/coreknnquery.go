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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _coreKnnQuery struct {
	v *types.CoreKnnQuery
}

func NewCoreKnnQuery(k int, numcandidates int) *_coreKnnQuery {

	tmp := &_coreKnnQuery{v: types.NewCoreKnnQuery()}

	tmp.K(k)

	tmp.NumCandidates(numcandidates)

	return tmp

}

func (s *_coreKnnQuery) Field(field string) *_coreKnnQuery {

	s.v.Field = field

	return s
}

func (s *_coreKnnQuery) K(k int) *_coreKnnQuery {

	s.v.K = k

	return s
}

func (s *_coreKnnQuery) NumCandidates(numcandidates int) *_coreKnnQuery {

	s.v.NumCandidates = numcandidates

	return s
}

func (s *_coreKnnQuery) QueryVector(queryvectors ...float32) *_coreKnnQuery {

	s.v.QueryVector = queryvectors

	return s
}

func (s *_coreKnnQuery) CoreKnnQueryCaster() *types.CoreKnnQuery {
	return s.v
}
