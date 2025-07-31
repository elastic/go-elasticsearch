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

// Role type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/get_role/types.ts#L32-L54
type Role struct {
	Applications      []ApplicationPrivileges                   `json:"applications"`
	Cluster           []clusterprivilege.ClusterPrivilege       `json:"cluster"`
	Description       *string                                   `json:"description,omitempty"`
	Global            map[string]map[string]map[string][]string `json:"global,omitempty"`
	Indices           []IndicesPrivileges                       `json:"indices"`
	Metadata          Metadata                                  `json:"metadata"`
	RemoteCluster     []RemoteClusterPrivileges                 `json:"remote_cluster,omitempty"`
	RemoteIndices     []RemoteIndicesPrivileges                 `json:"remote_indices,omitempty"`
	RoleTemplates     []RoleTemplate                            `json:"role_templates,omitempty"`
	RunAs             []string                                  `json:"run_as,omitempty"`
	TransientMetadata map[string]json.RawMessage                `json:"transient_metadata,omitempty"`
}

func (s *Role) UnmarshalJSON(data []byte) error {

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
				s.Global = make(map[string]map[string]map[string][]string, 0)
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

		case "role_templates":
			if err := dec.Decode(&s.RoleTemplates); err != nil {
				return fmt.Errorf("%s | %w", "RoleTemplates", err)
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

// NewRole returns a Role.
func NewRole() *Role {
	r := &Role{
		Global:            make(map[string]map[string]map[string][]string),
		TransientMetadata: make(map[string]json.RawMessage),
	}

	return r
}
