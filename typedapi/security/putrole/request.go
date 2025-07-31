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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package putrole

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/clusterprivilege"
)

// Request holds the request body struct for the package putrole
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/put_role/SecurityPutRoleRequest.ts#L32-L111
type Request struct {

	// Applications A list of application privilege entries.
	Applications []types.ApplicationPrivileges `json:"applications,omitempty"`
	// Cluster A list of cluster privileges. These privileges define the cluster-level
	// actions for users with this role.
	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`
	// Description Optional description of the role descriptor
	Description *string `json:"description,omitempty"`
	// Global An object defining global privileges. A global privilege is a form of cluster
	// privilege that is request-aware. Support for global privileges is currently
	// limited to the management of application privileges.
	Global map[string]json.RawMessage `json:"global,omitempty"`
	// Indices A list of indices permissions entries.
	Indices []types.IndicesPrivileges `json:"indices,omitempty"`
	// Metadata Optional metadata. Within the metadata object, keys that begin with an
	// underscore (`_`) are reserved for system use.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// RemoteCluster A list of remote cluster permissions entries.
	RemoteCluster []types.RemoteClusterPrivileges `json:"remote_cluster,omitempty"`
	// RemoteIndices A list of remote indices permissions entries.
	//
	// NOTE: Remote indices are effective for remote clusters configured with the
	// API key based model.
	// They have no effect for remote clusters configured with the certificate based
	// model.
	RemoteIndices []types.RemoteIndicesPrivileges `json:"remote_indices,omitempty"`
	// RunAs A list of users that the owners of this role can impersonate. *Note*: in
	// Serverless, the run-as feature is disabled. For API compatibility, you can
	// still specify an empty `run_as` field, but a non-empty list will be rejected.
	RunAs []string `json:"run_as,omitempty"`
	// TransientMetadata Indicates roles that might be incompatible with the current cluster license,
	// specifically roles with document and field level security. When the cluster
	// license doesn’t allow certain features for a given role, this parameter is
	// updated dynamically to list the incompatible features. If `enabled` is
	// `false`, the role is ignored, but is still listed in the response from the
	// authenticate API.
	TransientMetadata map[string]json.RawMessage `json:"transient_metadata,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		Global:            make(map[string]json.RawMessage, 0),
		TransientMetadata: make(map[string]json.RawMessage, 0),
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

func (s *Request) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "applications":
			if err := dec.Decode(&s.Applications); err != nil {
				return fmt.Errorf("%s | %w", "Applications", err)
			}

		case "cluster":
			if err := dec.Decode(&s.Cluster); err != nil {
				return fmt.Errorf("%s | %w", "Cluster", err)
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "global":
			if s.Global == nil {
				s.Global = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Global); err != nil {
				return fmt.Errorf("%s | %w", "Global", err)
			}

		case "indices":
			if err := dec.Decode(&s.Indices); err != nil {
				return fmt.Errorf("%s | %w", "Indices", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "remote_cluster":
			if err := dec.Decode(&s.RemoteCluster); err != nil {
				return fmt.Errorf("%s | %w", "RemoteCluster", err)
			}

		case "remote_indices":
			if err := dec.Decode(&s.RemoteIndices); err != nil {
				return fmt.Errorf("%s | %w", "RemoteIndices", err)
			}

		case "run_as":
			if err := dec.Decode(&s.RunAs); err != nil {
				return fmt.Errorf("%s | %w", "RunAs", err)
			}

		case "transient_metadata":
			if s.TransientMetadata == nil {
				s.TransientMetadata = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.TransientMetadata); err != nil {
				return fmt.Errorf("%s | %w", "TransientMetadata", err)
			}

		}
	}
	return nil
}
