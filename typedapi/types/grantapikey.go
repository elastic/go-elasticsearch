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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// GrantApiKey type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/security/grant_api_key/types.ts#L25-L46
type GrantApiKey struct {
	// Expiration Expiration time for the API key. By default, API keys never expire.
	Expiration *string `json:"expiration,omitempty"`
	// Metadata Arbitrary metadata that you want to associate with the API key.
	// It supports nested data structure.
	// Within the `metadata` object, keys beginning with `_` are reserved for system
	// usage.
	Metadata Metadata `json:"metadata,omitempty"`
	Name     string   `json:"name"`
	// RoleDescriptors The role descriptors for this API key.
	// This parameter is optional.
	// When it is not specified or is an empty array, the API key has a point in
	// time snapshot of permissions of the specified user or access token.
	// If you supply role descriptors, the resultant permissions are an intersection
	// of API keys permissions and the permissions of the user or access token.
	RoleDescriptors []map[string]RoleDescriptor `json:"role_descriptors,omitempty"`
}

func (s *GrantApiKey) UnmarshalJSON(data []byte) error {

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

		case "expiration":
			if err := dec.Decode(&s.Expiration); err != nil {
				return err
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return err
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return err
			}

		case "role_descriptors":

			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			source := bytes.NewReader(rawMsg)
			localDec := json.NewDecoder(source)
			switch rawMsg[0] {
			case '{':
				o := make(map[string]RoleDescriptor, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.RoleDescriptors = append(s.RoleDescriptors, o)
			case '[':
				o := make([]map[string]RoleDescriptor, 0)
				if err := localDec.Decode(&o); err != nil {
					return err
				}
				s.RoleDescriptors = o
			}

		}
	}
	return nil
}

// NewGrantApiKey returns a GrantApiKey.
func NewGrantApiKey() *GrantApiKey {
	r := &GrantApiKey{}

	return r
}
