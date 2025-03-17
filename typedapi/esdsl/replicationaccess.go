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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _replicationAccess struct {
	v *types.ReplicationAccess
}

func NewReplicationAccess() *_replicationAccess {

	return &_replicationAccess{v: types.NewReplicationAccess()}

}

// This needs to be set to true if the patterns in the names field should cover
// system indices.
func (s *_replicationAccess) AllowRestrictedIndices(allowrestrictedindices bool) *_replicationAccess {

	s.v.AllowRestrictedIndices = &allowrestrictedindices

	return s
}

// A list of indices (or index name patterns) to which the permissions in this
// entry apply.
func (s *_replicationAccess) Names(names ...string) *_replicationAccess {

	s.v.Names = make([]string, len(names))
	s.v.Names = names

	return s
}

func (s *_replicationAccess) ReplicationAccessCaster() *types.ReplicationAccess {
	return s.v
}
