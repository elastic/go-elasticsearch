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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexprivilege"
)

type _remoteIndicesPrivileges struct {
	v *types.RemoteIndicesPrivileges
}

func NewRemoteIndicesPrivileges() *_remoteIndicesPrivileges {

	return &_remoteIndicesPrivileges{v: types.NewRemoteIndicesPrivileges()}

}

// Set to `true` if using wildcard or regular expressions for patterns that
// cover restricted indices. Implicitly, restricted indices have limited
// privileges that can cause pattern tests to fail. If restricted indices are
// explicitly included in the `names` list, Elasticsearch checks privileges
// against these indices regardless of the value set for
// `allow_restricted_indices`.
func (s *_remoteIndicesPrivileges) AllowRestrictedIndices(allowrestrictedindices bool) *_remoteIndicesPrivileges {

	s.v.AllowRestrictedIndices = &allowrestrictedindices

	return s
}

// A list of cluster aliases to which the permissions in this entry apply.
func (s *_remoteIndicesPrivileges) Clusters(names ...string) *_remoteIndicesPrivileges {

	s.v.Clusters = names

	return s
}

// The document fields that the owners of the role have read access to.
func (s *_remoteIndicesPrivileges) FieldSecurity(fieldsecurity types.FieldSecurityVariant) *_remoteIndicesPrivileges {

	s.v.FieldSecurity = fieldsecurity.FieldSecurityCaster()

	return s
}

// A list of indices (or index name patterns) to which the permissions in this
// entry apply.
func (s *_remoteIndicesPrivileges) Names(names ...string) *_remoteIndicesPrivileges {

	s.v.Names = make([]string, len(names))
	s.v.Names = names

	return s
}

// The index level privileges that owners of the role have on the specified
// indices.
func (s *_remoteIndicesPrivileges) Privileges(privileges ...indexprivilege.IndexPrivilege) *_remoteIndicesPrivileges {

	for _, v := range privileges {

		s.v.Privileges = append(s.v.Privileges, v)

	}
	return s
}

// A search query that defines the documents the owners of the role have access
// to. A document within the specified indices must match this query for it to
// be accessible by the owners of the role.
func (s *_remoteIndicesPrivileges) Query(indicesprivilegesquery types.IndicesPrivilegesQueryVariant) *_remoteIndicesPrivileges {

	s.v.Query = *indicesprivilegesquery.IndicesPrivilegesQueryCaster()

	return s
}

func (s *_remoteIndicesPrivileges) RemoteIndicesPrivilegesCaster() *types.RemoteIndicesPrivileges {
	return s.v
}
