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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package oidcauthenticate

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package oidcauthenticate
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/security/oidc_authenticate/Request.ts#L22-L61
type Request struct {

	// Nonce Associate a client session with an ID token and mitigate replay attacks.
	// This value needs to be the same as the one that was provided to the
	// `/_security/oidc/prepare` API or the one that was generated by Elasticsearch
	// and included in the response to that call.
	Nonce string `json:"nonce"`
	// Realm The name of the OpenID Connect realm.
	// This property is useful in cases where multiple realms are defined.
	Realm *string `json:"realm,omitempty"`
	// RedirectUri The URL to which the OpenID Connect Provider redirected the User Agent in
	// response to an authentication request after a successful authentication.
	// This URL must be provided as-is (URL encoded), taken from the body of the
	// response or as the value of a location header in the response from the OpenID
	// Connect Provider.
	RedirectUri string `json:"redirect_uri"`
	// State Maintain state between the authentication request and the response.
	// This value needs to be the same as the one that was provided to the
	// `/_security/oidc/prepare` API or the one that was generated by Elasticsearch
	// and included in the response to that call.
	State string `json:"state"`
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
		return nil, fmt.Errorf("could not deserialise json into Oidcauthenticate request: %w", err)
	}

	return &req, nil
}
