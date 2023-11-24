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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package putrole

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/clusterprivilege"
)

// Request holds the request body struct for the package putrole
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/security/put_role/SecurityPutRoleRequest.ts#L31-L80
type Request struct {

	// Applications A list of application privilege entries.
	Applications []types.ApplicationPrivileges `json:"applications,omitempty"`
	// Cluster A list of cluster privileges. These privileges define the cluster-level
	// actions for users with this role.
	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`
	// Global An object defining global privileges. A global privilege is a form of cluster
	// privilege that is request-aware. Support for global privileges is currently
	// limited to the management of application privileges.
	Global map[string]json.RawMessage `json:"global,omitempty"`
	// Indices A list of indices permissions entries.
	Indices []types.IndicesPrivileges `json:"indices,omitempty"`
	// Metadata Optional metadata. Within the metadata object, keys that begin with an
	// underscore (`_`) are reserved for system use.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// RunAs A list of users that the owners of this role can impersonate.
	RunAs []string `json:"run_as,omitempty"`
	// TransientMetadata Indicates roles that might be incompatible with the current cluster license,
	// specifically roles with document and field level security. When the cluster
	// license doesn’t allow certain features for a given role, this parameter is
	// updated dynamically to list the incompatible features. If `enabled` is
	// `false`, the role is ignored, but is still listed in the response from the
	// authenticate API.
	TransientMetadata *types.TransientMetadataConfig `json:"transient_metadata,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Global: make(map[string]json.RawMessage, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putrole request: %w", err)
	}

	return &req, nil
}
