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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/remoteclusterprivilege"
)

type _remoteClusterPrivileges struct {
	v *types.RemoteClusterPrivileges
}

func NewRemoteClusterPrivileges() *_remoteClusterPrivileges {

	return &_remoteClusterPrivileges{v: types.NewRemoteClusterPrivileges()}

}

// A list of cluster aliases to which the permissions in this entry apply.
func (s *_remoteClusterPrivileges) Clusters(names ...string) *_remoteClusterPrivileges {

	s.v.Clusters = names

	return s
}

// The cluster level privileges that owners of the role have on the remote
// cluster.
func (s *_remoteClusterPrivileges) Privileges(privileges ...remoteclusterprivilege.RemoteClusterPrivilege) *_remoteClusterPrivileges {

	for _, v := range privileges {

		s.v.Privileges = append(s.v.Privileges, v)

	}
	return s
}

func (s *_remoteClusterPrivileges) RemoteClusterPrivilegesCaster() *types.RemoteClusterPrivileges {
	return s.v
}
