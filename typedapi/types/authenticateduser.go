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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// AuthenticatedUser type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/security/get_token/types.ts#L40-L45
type AuthenticatedUser struct {
	AuthenticationProvider *AuthenticationProvider `json:"authentication_provider,omitempty"`
	AuthenticationRealm    UserRealm               `json:"authentication_realm"`
	AuthenticationType     string                  `json:"authentication_type"`
	Email                  string                  `json:"email,omitempty"`
	Enabled                bool                    `json:"enabled"`
	FullName               string                  `json:"full_name,omitempty"`
	LookupRealm            UserRealm               `json:"lookup_realm"`
	Metadata               Metadata                `json:"metadata"`
	ProfileUid             *string                 `json:"profile_uid,omitempty"`
	Roles                  []string                `json:"roles"`
	Username               string                  `json:"username"`
}

func (s *AuthenticatedUser) UnmarshalJSON(data []byte) error {

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

		case "authentication_provider":
			if err := dec.Decode(&s.AuthenticationProvider); err != nil {
				return err
			}

		case "authentication_realm":
			if err := dec.Decode(&s.AuthenticationRealm); err != nil {
				return err
			}

		case "authentication_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AuthenticationType = o

		case "email":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Email = o

		case "enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "full_name":
			if err := dec.Decode(&s.FullName); err != nil {
				return err
			}

		case "lookup_realm":
			if err := dec.Decode(&s.LookupRealm); err != nil {
				return err
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return err
			}

		case "profile_uid":
			if err := dec.Decode(&s.ProfileUid); err != nil {
				return err
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
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

// NewAuthenticatedUser returns a AuthenticatedUser.
func NewAuthenticatedUser() *AuthenticatedUser {
	r := &AuthenticatedUser{}

	return r
}
