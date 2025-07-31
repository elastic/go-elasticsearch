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

package updateapikey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package updateapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/update_api_key/Request.ts#L26-L91
type Request struct {

	// Expiration The expiration time for the API key.
	// By default, API keys never expire.
	// This property can be omitted to leave the expiration unchanged.
	Expiration types.Duration `json:"expiration,omitempty"`
	// Metadata Arbitrary metadata that you want to associate with the API key.
	// It supports a nested data structure.
	// Within the metadata object, keys beginning with `_` are reserved for system
	// usage.
	// When specified, this value fully replaces the metadata previously associated
	// with the API key.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// RoleDescriptors The role descriptors to assign to this API key.
	// The API key's effective permissions are an intersection of its assigned
	// privileges and the point in time snapshot of permissions of the owner user.
	// You can assign new privileges by specifying them in this parameter.
	// To remove assigned privileges, you can supply an empty `role_descriptors`
	// parameter, that is to say, an empty object `{}`.
	// If an API key has no assigned privileges, it inherits the owner user's full
	// permissions.
	// The snapshot of the owner's permissions is always updated, whether you supply
	// the `role_descriptors` parameter or not.
	// The structure of a role descriptor is the same as the request for the create
	// API keys API.
	RoleDescriptors map[string]types.RoleDescriptor `json:"role_descriptors,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		RoleDescriptors: make(map[string]types.RoleDescriptor, 0),
	}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Updateapikey request: %w", err)
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

		case "expiration":
			if err := dec.Decode(&s.Expiration); err != nil {
				return fmt.Errorf("%s | %w", "Expiration", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "role_descriptors":
			if s.RoleDescriptors == nil {
				s.RoleDescriptors = make(map[string]types.RoleDescriptor, 0)
			}
			if err := dec.Decode(&s.RoleDescriptors); err != nil {
				return fmt.Errorf("%s | %w", "RoleDescriptors", err)
			}

		}
	}
	return nil
}
