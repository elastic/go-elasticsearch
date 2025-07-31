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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/clusterprivilege"
)

// QueryRole type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/query_role/types.ts#L103-L109
type QueryRole struct {
	// Applications A list of application privilege entries
	Applications []ApplicationPrivileges `json:"applications,omitempty"`
	// Cluster A list of cluster privileges. These privileges define the cluster level
	// actions that API keys are able to execute.
	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`
	// Description Optional description of the role descriptor
	Description *string `json:"description,omitempty"`
	// Global An object defining global privileges. A global privilege is a form of cluster
	// privilege that is request-aware. Support for global privileges is currently
	// limited to the management of application privileges.
	Global []GlobalPrivilege `json:"global,omitempty"`
	// Indices A list of indices permissions entries.
	Indices []IndicesPrivileges `json:"indices,omitempty"`
	// Metadata Optional meta-data. Within the metadata object, keys that begin with `_` are
	// reserved for system usage.
	Metadata Metadata `json:"metadata,omitempty"`
	// Name Name of the role.
	Name string `json:"name"`
	// RemoteCluster A list of cluster permissions for remote clusters.
	// NOTE: This is limited a subset of the cluster permissions.
	RemoteCluster []RemoteClusterPrivileges `json:"remote_cluster,omitempty"`
	// RemoteIndices A list of indices permissions for remote clusters.
	RemoteIndices []RemoteIndicesPrivileges `json:"remote_indices,omitempty"`
	// Restriction Restriction for when the role descriptor is allowed to be effective.
	Restriction *Restriction `json:"restriction,omitempty"`
	// RunAs A list of users that the API keys can impersonate.
	// NOTE: In Elastic Cloud Serverless, the run-as feature is disabled.
	// For API compatibility, you can still specify an empty `run_as` field, but a
	// non-empty list will be rejected.
	RunAs             []string                   `json:"run_as,omitempty"`
	Sort_             []FieldValue               `json:"_sort,omitempty"`
	TransientMetadata map[string]json.RawMessage `json:"transient_metadata,omitempty"`
}

func (s *QueryRole) UnmarshalJSON(data []byte) error {

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
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewGlobalPrivilege()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Global", err)
				}

				s.Global = append(s.Global, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Global); err != nil {
					return fmt.Errorf("%s | %w", "Global", err)
				}
			}

		case "indices", "index":
			if err := dec.Decode(&s.Indices); err != nil {
				return fmt.Errorf("%s | %w", "Indices", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Name = o

		case "remote_cluster":
			if err := dec.Decode(&s.RemoteCluster); err != nil {
				return fmt.Errorf("%s | %w", "RemoteCluster", err)
			}

		case "remote_indices":
			if err := dec.Decode(&s.RemoteIndices); err != nil {
				return fmt.Errorf("%s | %w", "RemoteIndices", err)
			}

		case "restriction":
			if err := dec.Decode(&s.Restriction); err != nil {
				return fmt.Errorf("%s | %w", "Restriction", err)
			}

		case "run_as":
			if err := dec.Decode(&s.RunAs); err != nil {
				return fmt.Errorf("%s | %w", "RunAs", err)
			}

		case "_sort":
			if err := dec.Decode(&s.Sort_); err != nil {
				return fmt.Errorf("%s | %w", "Sort_", err)
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

// NewQueryRole returns a QueryRole.
func NewQueryRole() *QueryRole {
	r := &QueryRole{
		TransientMetadata: make(map[string]json.RawMessage),
	}

	return r
}
