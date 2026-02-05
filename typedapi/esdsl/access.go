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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _access struct {
	v *types.Access
}

func NewAccess() *_access {

	return &_access{v: types.NewAccess()}

}

func (s *_access) Replication(replications ...types.ReplicationAccessVariant) *_access {

	for _, v := range replications {

		s.v.Replication = append(s.v.Replication, *v.ReplicationAccessCaster())

	}
	return s
}

func (s *_access) Search(searches ...types.SearchAccessVariant) *_access {

	for _, v := range searches {

		s.v.Search = append(s.v.Search, *v.SearchAccessCaster())

	}
	return s
}

func (s *_access) AccessCaster() *types.Access {
	return s.v
}
