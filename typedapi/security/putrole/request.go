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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package putrole

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/clusterprivilege"
)

// Request holds the request body struct for the package putrole
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/put_role/SecurityPutRoleRequest.ts#L31-L80
type Request struct {

	// Applications A list of application privilege entries.
	Applications []types.ApplicationPrivileges `json:"applications,omitempty"`

	// Cluster A list of cluster privileges. These privileges define the cluster-level
	// actions for users with this role.
	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`

	// Global An object defining global privileges. A global privilege is a form of cluster
	// privilege that is request-aware. Support for global privileges is currently
	// limited to the management of application privileges.
	Global map[string]interface{} `json:"global,omitempty"`

	// Indices A list of indices permissions entries.
	Indices []types.IndicesPrivileges `json:"indices,omitempty"`

	// Metadata Optional metadata. Within the metadata object, keys that begin with an
	// underscore (`_`) are reserved for system use.
	Metadata *types.Metadata `json:"metadata,omitempty"`

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

// RequestBuilder is the builder API for the putrole.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Global: make(map[string]interface{}, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putrole request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Applications(applications []types.ApplicationPrivilegesBuilder) *RequestBuilder {
	tmp := make([]types.ApplicationPrivileges, len(applications))
	for _, value := range applications {
		tmp = append(tmp, value.Build())
	}
	rb.v.Applications = tmp
	return rb
}

func (rb *RequestBuilder) Cluster(cluster ...clusterprivilege.ClusterPrivilege) *RequestBuilder {
	rb.v.Cluster = cluster
	return rb
}

func (rb *RequestBuilder) Global(value map[string]interface{}) *RequestBuilder {
	rb.v.Global = value
	return rb
}

func (rb *RequestBuilder) Indices(indices []types.IndicesPrivilegesBuilder) *RequestBuilder {
	tmp := make([]types.IndicesPrivileges, len(indices))
	for _, value := range indices {
		tmp = append(tmp, value.Build())
	}
	rb.v.Indices = tmp
	return rb
}

func (rb *RequestBuilder) Metadata(metadata *types.MetadataBuilder) *RequestBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *RequestBuilder) RunAs(run_as ...string) *RequestBuilder {
	rb.v.RunAs = run_as
	return rb
}

func (rb *RequestBuilder) TransientMetadata(transientmetadata *types.TransientMetadataConfigBuilder) *RequestBuilder {
	v := transientmetadata.Build()
	rb.v.TransientMetadata = &v
	return rb
}
