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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexprivilege"
)

type _indicesPrivileges struct {
	v *types.IndicesPrivileges
}

func NewIndicesPrivileges() *_indicesPrivileges {

	return &_indicesPrivileges{v: types.NewIndicesPrivileges()}

}

// Set to `true` if using wildcard or regular expressions for patterns that
// cover restricted indices. Implicitly, restricted indices have limited
// privileges that can cause pattern tests to fail. If restricted indices are
// explicitly included in the `names` list, Elasticsearch checks privileges
// against these indices regardless of the value set for
// `allow_restricted_indices`.
func (s *_indicesPrivileges) AllowRestrictedIndices(allowrestrictedindices bool) *_indicesPrivileges {

	s.v.AllowRestrictedIndices = &allowrestrictedindices

	return s
}

// The document fields that the owners of the role have read access to.
func (s *_indicesPrivileges) FieldSecurity(fieldsecurity types.FieldSecurityVariant) *_indicesPrivileges {

	s.v.FieldSecurity = fieldsecurity.FieldSecurityCaster()

	return s
}

// A list of indices (or index name patterns) to which the permissions in this
// entry apply.
func (s *_indicesPrivileges) Names(names ...string) *_indicesPrivileges {

	s.v.Names = make([]string, len(names))
	s.v.Names = names

	return s
}

// The index level privileges that owners of the role have on the specified
// indices.
func (s *_indicesPrivileges) Privileges(privileges ...indexprivilege.IndexPrivilege) *_indicesPrivileges {

	for _, v := range privileges {

		s.v.Privileges = append(s.v.Privileges, v)

	}
	return s
}

// A search query that defines the documents the owners of the role have access
// to. A document within the specified indices must match this query for it to
// be accessible by the owners of the role.
func (s *_indicesPrivileges) Query(indicesprivilegesquery types.IndicesPrivilegesQueryVariant) *_indicesPrivileges {

	s.v.Query = *indicesprivilegesquery.IndicesPrivilegesQueryCaster()

	return s
}

func (s *_indicesPrivileges) IndicesPrivilegesCaster() *types.IndicesPrivileges {
	return s.v
}
