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
)

// Authentication type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/delegate_pki/SecurityDelegatePkiResponse.ts#L43-L55
type Authentication struct {
	ApiKey              map[string]string   `json:"api_key,omitempty"`
	AuthenticationRealm AuthenticationRealm `json:"authentication_realm"`
	AuthenticationType  string              `json:"authentication_type"`
	Email               *string             `json:"email,omitempty"`
	Enabled             bool                `json:"enabled"`
	FullName            *string             `json:"full_name,omitempty"`
	LookupRealm         AuthenticationRealm `json:"lookup_realm"`
	Metadata            Metadata            `json:"metadata"`
	Roles               []string            `json:"roles"`
	Token               map[string]string   `json:"token,omitempty"`
	Username            string              `json:"username"`
}

func (s *Authentication) UnmarshalJSON(data []byte) error {

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

		case "api_key":
			if s.ApiKey == nil {
				s.ApiKey = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.ApiKey); err != nil {
				return fmt.Errorf("%s | %w", "ApiKey", err)
			}

		case "authentication_realm":
			if err := dec.Decode(&s.AuthenticationRealm); err != nil {
				return fmt.Errorf("%s | %w", "AuthenticationRealm", err)
			}

		case "authentication_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "AuthenticationType", err)
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
				return fmt.Errorf("%s | %w", "Email", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Email = &o

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "full_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "FullName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.FullName = &o

		case "lookup_realm":
			if err := dec.Decode(&s.LookupRealm); err != nil {
				return fmt.Errorf("%s | %w", "LookupRealm", err)
			}

		case "metadata":
			if err := dec.Decode(&s.Metadata); err != nil {
				return fmt.Errorf("%s | %w", "Metadata", err)
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "token":
			if s.Token == nil {
				s.Token = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Token); err != nil {
				return fmt.Errorf("%s | %w", "Token", err)
			}

		case "username":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Username", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Username = o

		}
	}
	return nil
}

// NewAuthentication returns a Authentication.
func NewAuthentication() *Authentication {
	r := &Authentication{
		ApiKey: make(map[string]string),
		Token:  make(map[string]string),
	}

	return r
}
