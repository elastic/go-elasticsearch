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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package grantapikey

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/apikeygranttype"
)

// Request holds the request body struct for the package grantapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/security/grant_api_key/SecurityGrantApiKeyRequest.ts#L28-L125
type Request struct {
	// AccessToken The user's access token. If you specify the `access_token` grant type, this
	// parameter is required. It is not valid with other grant types.
	AccessToken *string `json:"access_token,omitempty"`
	// ApiKey The API key.
	ApiKey types.GrantApiKey `json:"api_key"`
	// ClientAuthentication The client authentication for a JWT that requires it, which is what is
	// normally sent in the `ES-Client-Authentication` request header. It is valid
	// only with the `access_token` grant type, when the access token is a JWT.
	ClientAuthentication *types.GrantApiKeyClientAuthentication `json:"client_authentication,omitempty"`
	// GrantType The type of grant. Supported grant types are: `access_token`, `password`, and
	// `_user_managed_service_account`.
	GrantType apikeygranttype.ApiKeyGrantType `json:"grant_type"`
	// Password The user's password. If you specify the `password` grant type, this parameter
	// is required. It is not valid with other grant types.
	Password *string `json:"password,omitempty"`
	// RunAs The name of the user to be impersonated. It is not valid with the
	// `_user_managed_service_account` grant type.
	RunAs *string `json:"run_as,omitempty"`
	// ServiceAccountToken The service account token of a user-managed service account, as returned by
	// the create service account token API. If you specify the
	// `_user_managed_service_account` grant type, this parameter is required. It is
	// not valid with other grant types.
	ServiceAccountToken *string `json:"service_account_token,omitempty"`
	// Username The user name that identifies the user. If you specify the `password` grant
	// type, this parameter is required. It is not valid with other grant types.
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
		return nil, fmt.Errorf("could not deserialise json into Grantapikey request: %w", err)
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

		case "access_token":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "AccessToken", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AccessToken = &o

		case "api_key":
			if err := dec.Decode(&s.ApiKey); err != nil {
				return fmt.Errorf("%s | %w", "ApiKey", err)
			}

		case "client_authentication":
			if err := dec.Decode(&s.ClientAuthentication); err != nil {
				return fmt.Errorf("%s | %w", "ClientAuthentication", err)
			}

		case "grant_type":
			if err := dec.Decode(&s.GrantType); err != nil {
				return fmt.Errorf("%s | %w", "GrantType", err)
			}

		case "password":
			if err := dec.Decode(&s.Password); err != nil {
				return fmt.Errorf("%s | %w", "Password", err)
			}

		case "run_as":
			if err := dec.Decode(&s.RunAs); err != nil {
				return fmt.Errorf("%s | %w", "RunAs", err)
			}

		case "service_account_token":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ServiceAccountToken", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ServiceAccountToken = &o

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return fmt.Errorf("%s | %w", "Username", err)
			}

		}
	}
	return nil
}
