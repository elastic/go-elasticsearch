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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type _indicesPrivileges struct {
	v *types.IndicesPrivileges
}

func NewIndicesPrivileges() *_indicesPrivileges {

	return &_indicesPrivileges{v: types.NewIndicesPrivileges()}

}

func (s *_indicesPrivileges) AllowRestrictedIndices(allowrestrictedindices bool) *_indicesPrivileges {

	s.v.AllowRestrictedIndices = &allowrestrictedindices

	return s
}

func (s *_indicesPrivileges) FieldSecurity(fieldsecurity types.FieldSecurityVariant) *_indicesPrivileges {

	s.v.FieldSecurity = fieldsecurity.FieldSecurityCaster()

	return s
}

func (s *_indicesPrivileges) Names(names ...string) *_indicesPrivileges {

	s.v.Names = make([]string, len(names))
	s.v.Names = names

	return s
}

func (s *_indicesPrivileges) Privileges(privileges ...indexprivilege.IndexPrivilege) *_indicesPrivileges {

	for _, v := range privileges {

		s.v.Privileges = append(s.v.Privileges, v)

	}
	return s
}

func (s *_indicesPrivileges) Query(indicesprivilegesquery types.IndicesPrivilegesQueryVariant) *_indicesPrivileges {

	s.v.Query = *indicesprivilegesquery.IndicesPrivilegesQueryCaster()

	return s
}

func (s *_indicesPrivileges) IndicesPrivilegesCaster() *types.IndicesPrivileges {
	return s.v
}
