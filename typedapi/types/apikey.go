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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ApiKey type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/security/_types/ApiKey.ts#L27-L77
type ApiKey struct {
	// Creation Creation time for the API key in milliseconds.
	Creation *int64 `json:"creation,omitempty"`
	// Expiration Expiration time for the API key in milliseconds.
	Expiration *int64 `json:"expiration,omitempty"`
	// Id Id for the API key
	Id string `json:"id"`
	// Invalidated Invalidation status for the API key.
	// If the key has been invalidated, it has a value of `true`. Otherwise, it is
	// `false`.
	Invalidated *bool `json:"invalidated,omitempty"`
	// LimitedBy The owner user’s permissions associated with the API key.
	// It is a point-in-time snapshot captured at creation and subsequent updates.
	// An API key’s effective permissions are an intersection of its assigned
	// privileges and the owner user’s permissions.
	LimitedBy []map[string]RoleDescriptor `json:"limited_by,omitempty"`
	// Metadata Metadata of the API key
	Metadata Metadata `json:"metadata,omitempty"`
	// Name Name of the API key.
	Name string `json:"name"`
	// Realm Realm name of the principal for which this API key was created.
	Realm *string `json:"realm,omitempty"`
	// RoleDescriptors The role descriptors assigned to this API key when it was created or last
	// updated.
	// An empty role descriptor means the API key inherits the owner user’s
	// permissions.
	RoleDescriptors map[string]RoleDescriptor `json:"role_descriptors,omitempty"`
	Sort_           []FieldValue              `json:"_sort,omitempty"`
	// Username Principal for which this API key was created
	Username *string `json:"username,omitempty"`
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

		case "creation":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Creation = &value
			case float64:
				f := int64(v)
				s.Creation = &f
			}

		case "expiration":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.Expiration = &value
			case float64:
				f := int64(v)
				s.Expiration = &f
			}

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return err
			}

		case "invalidated":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Invalidated = &value
			case bool:
				s.Invalidated = &v
			}

		case "limited_by":
			if err := dec.Decode(&s.LimitedBy); err != nil {
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

		case "realm":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Realm = &o

		case "role_descriptors":
			if s.RoleDescriptors == nil {
				s.RoleDescriptors = make(map[string]RoleDescriptor, 0)
			}
			if err := dec.Decode(&s.RoleDescriptors); err != nil {
				return err
			}

		case "_sort":
			if err := dec.Decode(&s.Sort_); err != nil {
				return err
			}

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewApiKey returns a ApiKey.
func NewApiKey() *ApiKey {
	r := &ApiKey{
		RoleDescriptors: make(map[string]RoleDescriptor, 0),
	}

	return r
}
