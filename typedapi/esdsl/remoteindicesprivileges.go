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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/indexprivilege"
)

type _remoteIndicesPrivileges struct {
	v *types.RemoteIndicesPrivileges
}

func NewRemoteIndicesPrivileges() *_remoteIndicesPrivileges {

	return &_remoteIndicesPrivileges{v: types.NewRemoteIndicesPrivileges()}

}

func (s *_remoteIndicesPrivileges) AllowRestrictedIndices(allowrestrictedindices bool) *_remoteIndicesPrivileges {

	s.v.AllowRestrictedIndices = &allowrestrictedindices

	return s
}

func (s *_remoteIndicesPrivileges) Clusters(names ...string) *_remoteIndicesPrivileges {

	s.v.Clusters = names

	return s
}

func (s *_remoteIndicesPrivileges) FieldSecurity(fieldsecurity types.FieldSecurityVariant) *_remoteIndicesPrivileges {

	s.v.FieldSecurity = fieldsecurity.FieldSecurityCaster()

	return s
}

func (s *_remoteIndicesPrivileges) Names(names ...string) *_remoteIndicesPrivileges {

	s.v.Names = make([]string, len(names))
	s.v.Names = names

	return s
}

func (s *_remoteIndicesPrivileges) Privileges(privileges ...indexprivilege.IndexPrivilege) *_remoteIndicesPrivileges {

	for _, v := range privileges {

		s.v.Privileges = append(s.v.Privileges, v)

	}
	return s
}

func (s *_remoteIndicesPrivileges) Query(indicesprivilegesquery types.IndicesPrivilegesQueryVariant) *_remoteIndicesPrivileges {

	s.v.Query = *indicesprivilegesquery.IndicesPrivilegesQueryCaster()

	return s
}

func (s *_remoteIndicesPrivileges) RemoteIndicesPrivilegesCaster() *types.RemoteIndicesPrivileges {
	return s.v
}
