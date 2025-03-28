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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/clusterprivilege"
)

type _roleDescriptor struct {
	v *types.RoleDescriptor
}

func NewRoleDescriptor() *_roleDescriptor {

	return &_roleDescriptor{v: types.NewRoleDescriptor()}

}

// A list of application privilege entries
func (s *_roleDescriptor) Applications(applications ...types.ApplicationPrivilegesVariant) *_roleDescriptor {

	for _, v := range applications {

		s.v.Applications = append(s.v.Applications, *v.ApplicationPrivilegesCaster())

	}
	return s
}

// A list of cluster privileges. These privileges define the cluster level
// actions that API keys are able to execute.
func (s *_roleDescriptor) Cluster(clusters ...clusterprivilege.ClusterPrivilege) *_roleDescriptor {

	for _, v := range clusters {

		s.v.Cluster = append(s.v.Cluster, v)

	}
	return s
}

// Optional description of the role descriptor
func (s *_roleDescriptor) Description(description string) *_roleDescriptor {

	s.v.Description = &description

	return s
}

// An object defining global privileges. A global privilege is a form of cluster
// privilege that is request-aware. Support for global privileges is currently
// limited to the management of application privileges.
func (s *_roleDescriptor) Global(globals ...types.GlobalPrivilegeVariant) *_roleDescriptor {

	s.v.Global = make([]types.GlobalPrivilege, len(globals))
	for i, v := range globals {
		s.v.Global[i] = *v.GlobalPrivilegeCaster()
	}

	return s
}

// A list of indices permissions entries.
func (s *_roleDescriptor) Indices(indices ...types.IndicesPrivilegesVariant) *_roleDescriptor {

	for _, v := range indices {

		s.v.Indices = append(s.v.Indices, *v.IndicesPrivilegesCaster())

	}
	return s
}

// Optional meta-data. Within the metadata object, keys that begin with `_` are
// reserved for system usage.
func (s *_roleDescriptor) Metadata(metadata types.MetadataVariant) *_roleDescriptor {

	s.v.Metadata = *metadata.MetadataCaster()

	return s
}

// A list of cluster permissions for remote clusters.
// NOTE: This is limited a subset of the cluster permissions.
func (s *_roleDescriptor) RemoteCluster(remoteclusters ...types.RemoteClusterPrivilegesVariant) *_roleDescriptor {

	for _, v := range remoteclusters {

		s.v.RemoteCluster = append(s.v.RemoteCluster, *v.RemoteClusterPrivilegesCaster())

	}
	return s
}

// A list of indices permissions for remote clusters.
func (s *_roleDescriptor) RemoteIndices(remoteindices ...types.RemoteIndicesPrivilegesVariant) *_roleDescriptor {

	for _, v := range remoteindices {

		s.v.RemoteIndices = append(s.v.RemoteIndices, *v.RemoteIndicesPrivilegesCaster())

	}
	return s
}

// Restriction for when the role descriptor is allowed to be effective.
func (s *_roleDescriptor) Restriction(restriction types.RestrictionVariant) *_roleDescriptor {

	s.v.Restriction = restriction.RestrictionCaster()

	return s
}

// A list of users that the API keys can impersonate.
// NOTE: In Elastic Cloud Serverless, the run-as feature is disabled.
// For API compatibility, you can still specify an empty `run_as` field, but a
// non-empty list will be rejected.
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
