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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _sparseVectorQuery struct {
	v *types.SparseVectorQuery
}

func NewSparseVectorQuery() *_sparseVectorQuery {
	return &_sparseVectorQuery{v: types.NewSparseVectorQuery()}
}

// AdditionalSparseVectorQueryProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_sparseVectorQuery) AdditionalSparseVectorQueryProperty(key string, value json.RawMessage) *_sparseVectorQuery {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalSparseVectorQueryProperty = tmp
	return s
}

func (s *_sparseVectorQuery) Boost(boost float32) *_sparseVectorQuery {

	s.v.Boost = &boost

	return s
}

func (s *_sparseVectorQuery) Field(field string) *_sparseVectorQuery {

	s.v.Field = field

	return s
}

func (s *_sparseVectorQuery) InferenceId(id string) *_sparseVectorQuery {

	s.v.InferenceId = &id

	return s
}

func (s *_sparseVectorQuery) Prune(prune bool) *_sparseVectorQuery {

	s.v.Prune = &prune

	return s
}

func (s *_sparseVectorQuery) PruningConfig(pruningconfig types.TokenPruningConfigVariant) *_sparseVectorQuery {

	s.v.PruningConfig = pruningconfig.TokenPruningConfigCaster()

	return s
}

func (s *_sparseVectorQuery) Query(query string) *_sparseVectorQuery {

	s.v.Query = &query

	return s
}

func (s *_sparseVectorQuery) QueryName_(queryname_ string) *_sparseVectorQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_sparseVectorQuery) QueryVector(queryvector map[string]float32) *_sparseVectorQuery {

	s.v.QueryVector = queryvector
	return s
}

func (s *_sparseVectorQuery) AddQueryVector(key string, value float32) *_sparseVectorQuery {

	var tmp map[string]float32
	if s.v.QueryVector == nil {
		s.v.QueryVector = make(map[string]float32)
	} else {
		tmp = s.v.QueryVector
	}

	tmp[key] = value

	s.v.QueryVector = tmp
	return s
}

func (s *_sparseVectorQuery) SparseVectorQueryCaster() *types.SparseVectorQuery {
	return s.v
}
