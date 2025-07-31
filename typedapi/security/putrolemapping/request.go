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

package putrolemapping

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putrolemapping
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/put_role_mapping/SecurityPutRoleMappingRequest.ts#L25-L103
type Request struct {

	// Enabled Mappings that have `enabled` set to `false` are ignored when role mapping is
	// performed.
	Enabled *bool `json:"enabled,omitempty"`
	// Metadata Additional metadata that helps define which roles are assigned to each user.
	// Within the metadata object, keys beginning with `_` are reserved for system
	// usage.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// RoleTemplates A list of Mustache templates that will be evaluated to determine the roles
	// names that should granted to the users that match the role mapping rules.
	// Exactly one of `roles` or `role_templates` must be specified.
	RoleTemplates []types.RoleTemplate `json:"role_templates,omitempty"`
	// Roles A list of role names that are granted to the users that match the role
	// mapping rules.
	// Exactly one of `roles` or `role_templates` must be specified.
	Roles []string `json:"roles,omitempty"`
	// Rules The rules that determine which users should be matched by the mapping.
	// A rule is a logical condition that is expressed by using a JSON DSL.
	Rules *types.RoleMappingRule `json:"rules,omitempty"`
	RunAs []string               `json:"run_as,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putrolemapping request: %w", err)
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

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = &value
			case bool:
				s.Enabled = &v
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "role_templates":
			if err := dec.Decode(&s.RoleTemplates); err != nil {
				return fmt.Errorf("%s | %w", "RoleTemplates", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "rules":
			if err := dec.Decode(&s.Rules); err != nil {
				return fmt.Errorf("%s | %w", "Rules", err)
			}

		case "run_as":
			if err := dec.Decode(&s.RunAs); err != nil {
				return fmt.Errorf("%s | %w", "RunAs", err)
			}

		}
	}
	return nil
}
