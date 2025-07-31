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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/clusterprivilege"
)

type _roleDescriptor struct {
	v *types.RoleDescriptor
}

func NewRoleDescriptor() *_roleDescriptor {

	return &_roleDescriptor{v: types.NewRoleDescriptor()}

}

func (s *_roleDescriptor) Applications(applications ...types.ApplicationPrivilegesVariant) *_roleDescriptor {

	for _, v := range applications {

		s.v.Applications = append(s.v.Applications, *v.ApplicationPrivilegesCaster())

	}
	return s
}

func (s *_roleDescriptor) Cluster(clusters ...clusterprivilege.ClusterPrivilege) *_roleDescriptor {

	for _, v := range clusters {

		s.v.Cluster = append(s.v.Cluster, v)

	}
	return s
}

func (s *_roleDescriptor) Description(description string) *_roleDescriptor {

	s.v.Description = &description

	return s
}

func (s *_roleDescriptor) Global(globals ...types.GlobalPrivilegeVariant) *_roleDescriptor {

	s.v.Global = make([]types.GlobalPrivilege, len(globals))
	for i, v := range globals {
		s.v.Global[i] = *v.GlobalPrivilegeCaster()
	}

	return s
}

func (s *_roleDescriptor) Indices(indices ...types.IndicesPrivilegesVariant) *_roleDescriptor {

	for _, v := range indices {

		s.v.Indices = append(s.v.Indices, *v.IndicesPrivilegesCaster())

	}
	return s
}

func (s *_roleDescriptor) Metadata(metadata types.MetadataVariant) *_roleDescriptor {

	s.v.Metadata = *metadata.MetadataCaster()

	return s
}

func (s *_roleDescriptor) RemoteCluster(remoteclusters ...types.RemoteClusterPrivilegesVariant) *_roleDescriptor {

	for _, v := range remoteclusters {

		s.v.RemoteCluster = append(s.v.RemoteCluster, *v.RemoteClusterPrivilegesCaster())

	}
	return s
}

func (s *_roleDescriptor) RemoteIndices(remoteindices ...types.RemoteIndicesPrivilegesVariant) *_roleDescriptor {

	for _, v := range remoteindices {

		s.v.RemoteIndices = append(s.v.RemoteIndices, *v.RemoteIndicesPrivilegesCaster())

	}
	return s
}

func (s *_roleDescriptor) Restriction(restriction types.RestrictionVariant) *_roleDescriptor {

	s.v.Restriction = restriction.RestrictionCaster()

	return s
}

func (s *_roleDescriptor) RunAs(runas ...string) *_roleDescriptor {

	for _, v := range runas {

		s.v.RunAs = append(s.v.RunAs, v)

	}
	return s
}

func (s *_roleDescriptor) TransientMetadata(transientmetadata map[string]json.RawMessage) *_roleDescriptor {

	s.v.TransientMetadata = transientmetadata
	return s
}

func (s *_roleDescriptor) AddTransientMetadatum(key string, value json.RawMessage) *_roleDescriptor {

	var tmp map[string]json.RawMessage
	if s.v.TransientMetadata == nil {
		s.v.TransientMetadata = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.TransientMetadata
	}

	tmp[key] = value

	s.v.TransientMetadata = tmp
	return s
}

func (s *_roleDescriptor) RoleDescriptorCaster() *types.RoleDescriptor {
	return s.v
}
