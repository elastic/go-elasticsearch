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

package gettoken

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/accesstokengranttype"
)

// Request holds the request body struct for the package gettoken
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/get_token/GetUserAccessTokenRequest.ts#L25-L90
type Request struct {

	// GrantType The type of grant.
	// Supported grant types are: `password`, `_kerberos`, `client_credentials`, and
	// `refresh_token`.
	GrantType *accesstokengranttype.AccessTokenGrantType `json:"grant_type,omitempty"`
	// KerberosTicket The base64 encoded kerberos ticket.
	// If you specify the `_kerberos` grant type, this parameter is required.
	// This parameter is not valid with any other supported grant type.
	KerberosTicket *string `json:"kerberos_ticket,omitempty"`
	// Password The user's password.
	// If you specify the `password` grant type, this parameter is required.
	// This parameter is not valid with any other supported grant type.
	Password *string `json:"password,omitempty"`
	// RefreshToken The string that was returned when you created the token, which enables you to
	// extend its life.
	// If you specify the `refresh_token` grant type, this parameter is required.
	// This parameter is not valid with any other supported grant type.
	RefreshToken *string `json:"refresh_token,omitempty"`
	// Scope The scope of the token.
	// Currently tokens are only issued for a scope of FULL regardless of the value
	// sent with the request.
	Scope *string `json:"scope,omitempty"`
	// Username The username that identifies the user.
	// If you specify the `password` grant type, this parameter is required.
	// This parameter is not valid with any other supported grant type.
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
		return nil, fmt.Errorf("could not deserialise json into Gettoken request: %w", err)
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

		case "grant_type":
			if err := dec.Decode(&s.GrantType); err != nil {
				return fmt.Errorf("%s | %w", "GrantType", err)
			}

		case "kerberos_ticket":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "KerberosTicket", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.KerberosTicket = &o

		case "password":
			if err := dec.Decode(&s.Password); err != nil {
				return fmt.Errorf("%s | %w", "Password", err)
			}

		case "refresh_token":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RefreshToken", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RefreshToken = &o

		case "scope":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Scope", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Scope = &o

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return fmt.Errorf("%s | %w", "Username", err)
			}

		}
	}
	return nil
}
