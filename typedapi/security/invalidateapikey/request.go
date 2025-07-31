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

package invalidateapikey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Request holds the request body struct for the package invalidateapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/invalidate_api_key/SecurityInvalidateApiKeyRequest.ts#L23-L82
type Request struct {
	Id *string `json:"id,omitempty"`
	// Ids A list of API key ids.
	// This parameter cannot be used with any of `name`, `realm_name`, or
	// `username`.
	Ids []string `json:"ids,omitempty"`
	// Name An API key name.
	// This parameter cannot be used with any of `ids`, `realm_name` or `username`.
	Name *string `json:"name,omitempty"`
	// Owner Query API keys owned by the currently authenticated user.
	// The `realm_name` or `username` parameters cannot be specified when this
	// parameter is set to `true` as they are assumed to be the currently
	// authenticated ones.
	//
	// NOTE: At least one of `ids`, `name`, `username`, and `realm_name` must be
	// specified if `owner` is `false`.
	Owner *bool `json:"owner,omitempty"`
	// RealmName The name of an authentication realm.
	// This parameter cannot be used with either `ids` or `name`, or when `owner`
	// flag is set to `true`.
	RealmName *string `json:"realm_name,omitempty"`
	// Username The username of a user.
	// This parameter cannot be used with either `ids` or `name` or when `owner`
	// flag is set to `true`.
	Username *string `json:"username,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Invalidateapikey request: %w", err)
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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "ids":
			if err := dec.Decode(&s.Ids); err != nil {
				return fmt.Errorf("%s | %w", "Ids", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "owner":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Owner", err)
				}
				s.Owner = &value
			case bool:
				s.Owner = &v
			}

		case "realm_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RealmName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RealmName = &o

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return fmt.Errorf("%s | %w", "Username", err)
			}

		}
	}
	return nil
}
