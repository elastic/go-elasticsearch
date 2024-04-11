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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

// RoleDescriptorRead type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/security/_types/RoleDescriptor.ts#L58-L86
type RoleDescriptorRead struct {
	// Applications A list of application privilege entries
	Applications []ApplicationPrivileges `json:"applications,omitempty"`
	// Cluster A list of cluster privileges. These privileges define the cluster level
	// actions that API keys are able to execute.
	Cluster []string `json:"cluster"`
	// Global An object defining global privileges. A global privilege is a form of cluster
	// privilege that is request-aware. Support for global privileges is currently
	// limited to the management of application privileges.
	Global []GlobalPrivilege `json:"global,omitempty"`
	// Indices A list of indices permissions entries.
	Indices []IndicesPrivileges `json:"indices"`
	// Metadata Optional meta-data. Within the metadata object, keys that begin with `_` are
	// reserved for system usage.
	Metadata Metadata `json:"metadata,omitempty"`
	// RunAs A list of users that the API keys can impersonate.
	RunAs             []string                   `json:"run_as,omitempty"`
	TransientMetadata map[string]json.RawMessage `json:"transient_metadata,omitempty"`
}

func (s *RoleDescriptorRead) UnmarshalJSON(data []byte) error {

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

// NewRoleDescriptorRead returns a RoleDescriptorRead.
func NewRoleDescriptorRead() *RoleDescriptorRead {
	r := &RoleDescriptorRead{
		TransientMetadata: make(map[string]json.RawMessage, 0),
	}

	return r
}
