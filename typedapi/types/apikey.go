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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/apikeytype"
)

// ApiKey type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/_types/ApiKey.ts#L27-L113
type ApiKey struct {
	// Access The access granted to cross-cluster API keys.
	// The access is composed of permissions for cross cluster search and cross
	// cluster replication.
	// At least one of them must be specified.
	// When specified, the new access assignment fully replaces the previously
	// assigned access.
	Access *Access `json:"access,omitempty"`
	// Creation Creation time for the API key in milliseconds.
	Creation int64 `json:"creation"`
	// Expiration Expiration time for the API key in milliseconds.
	Expiration *int64 `json:"expiration,omitempty"`
	// Id Id for the API key
	Id string `json:"id"`
	// Invalidated Invalidation status for the API key.
	// If the key has been invalidated, it has a value of `true`. Otherwise, it is
	// `false`.
	Invalidated bool `json:"invalidated"`
	// Invalidation If the key has been invalidated, invalidation time in milliseconds.
	Invalidation *int64 `json:"invalidation,omitempty"`
	// LimitedBy The owner user’s permissions associated with the API key.
	// It is a point-in-time snapshot captured at creation and subsequent updates.
	// An API key’s effective permissions are an intersection of its assigned
	// privileges and the owner user’s permissions.
	LimitedBy []map[string]RoleDescriptor `json:"limited_by,omitempty"`
	// Metadata Metadata of the API key
	Metadata Metadata `json:"metadata"`
	// Name Name of the API key.
	Name string `json:"name"`
	// ProfileUid The profile uid for the API key owner principal, if requested and if it
	// exists
	ProfileUid *string `json:"profile_uid,omitempty"`
	// Realm Realm name of the principal for which this API key was created.
	Realm string `json:"realm"`
	// RealmType Realm type of the principal for which this API key was created
	RealmType *string `json:"realm_type,omitempty"`
	// RoleDescriptors The role descriptors assigned to this API key when it was created or last
	// updated.
	// An empty role descriptor means the API key inherits the owner user’s
	// permissions.
	RoleDescriptors map[string]RoleDescriptor `json:"role_descriptors,omitempty"`
	// Sort_ Sorting values when using the `sort` parameter with the
	// `security.query_api_keys` API.
	Sort_ []FieldValue `json:"_sort,omitempty"`
	// Type The type of the API key (e.g. `rest` or `cross_cluster`).
	Type apikeytype.ApiKeyType `json:"type"`
	// Username Principal for which this API key was created
	Username string `json:"username"`
}

func (s *ApiKey) UnmarshalJSON(data []byte) error {

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

		case "access":
			if err := dec.Decode(&s.Access); err != nil {
				return fmt.Errorf("%s | %w", "Access", err)
			}

		case "creation":
			if err := dec.Decode(&s.Creation); err != nil {
				return fmt.Errorf("%s | %w", "Creation", err)
			}

		case "expiration":
			if err := dec.Decode(&s.Expiration); err != nil {
				return fmt.Errorf("%s | %w", "Expiration", err)
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "invalidated":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Invalidated", err)
				}
				s.Invalidated = value
			case bool:
				s.Invalidated = v
			}

		case "invalidation":
			if err := dec.Decode(&s.Invalidation); err != nil {
				return fmt.Errorf("%s | %w", "Invalidation", err)
			}

		case "limited_by":
			if err := dec.Decode(&s.LimitedBy); err != nil {
				return fmt.Errorf("%s | %w", "LimitedBy", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "profile_uid":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ProfileUid", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ProfileUid = &o

		case "realm":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Realm", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Realm = o

		case "realm_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RealmType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RealmType = &o

		case "role_descriptors":
			if s.RoleDescriptors == nil {
				s.RoleDescriptors = make(map[string]RoleDescriptor, 0)
			}
			if err := dec.Decode(&s.RoleDescriptors); err != nil {
				return fmt.Errorf("%s | %w", "RoleDescriptors", err)
			}

		case "_sort":
			if err := dec.Decode(&s.Sort_); err != nil {
				return fmt.Errorf("%s | %w", "Sort_", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return fmt.Errorf("%s | %w", "Username", err)
			}

		}
	}
	return nil
}

// NewApiKey returns a ApiKey.
func NewApiKey() *ApiKey {
	r := &ApiKey{
		RoleDescriptors: make(map[string]RoleDescriptor),
	}

	return r
}
