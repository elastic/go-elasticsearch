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

package createapikey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package createapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/create_api_key/SecurityCreateApiKeyRequest.ts#L26-L86
type Request struct {

	// Expiration The expiration time for the API key.
	// By default, API keys never expire.
	Expiration types.Duration `json:"expiration,omitempty"`
	// Metadata Arbitrary metadata that you want to associate with the API key. It supports
	// nested data structure. Within the metadata object, keys beginning with `_`
	// are reserved for system usage.
	Metadata types.Metadata `json:"metadata,omitempty"`
	// Name A name for the API key.
	Name *string `json:"name,omitempty"`
	// RoleDescriptors An array of role descriptors for this API key.
	// When it is not specified or it is an empty array, the API key will have a
	// point in time snapshot of permissions of the authenticated user.
	// If you supply role descriptors, the resultant permissions are an intersection
	// of API keys permissions and the authenticated user's permissions thereby
	// limiting the access scope for API keys.
	// The structure of role descriptor is the same as the request for the create
	// role API.
	// For more details, refer to the create or update roles API.
	//
	// NOTE: Due to the way in which this permission intersection is calculated, it
	// is not possible to create an API key that is a child of another API key,
	// unless the derived key is created without any privileges.
	// In this case, you must explicitly specify a role descriptor with no
	// privileges.
	// The derived API key can be used for authentication; it will not have
	// authority to call Elasticsearch APIs.
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
		return nil, fmt.Errorf("could not deserialise json into Createapikey request: %w", err)
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

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
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
