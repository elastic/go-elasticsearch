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

package oidcprepareauthentication

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package oidcprepareauthentication
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/oidc_prepare_authentication/Request.ts#L22-L71
type Request struct {

	// Iss In the case of a third party initiated single sign on, this is the issuer
	// identifier for the OP that the RP is to send the authentication request to.
	// It cannot be specified when *realm* is specified.
	// One of *realm* or *iss* is required.
	Iss *string `json:"iss,omitempty"`
	// LoginHint In the case of a third party initiated single sign on, it is a string value
	// that is included in the authentication request as the *login_hint* parameter.
	// This parameter is not valid when *realm* is specified.
	LoginHint *string `json:"login_hint,omitempty"`
	// Nonce The value used to associate a client session with an ID token and to mitigate
	// replay attacks.
	// If the caller of the API does not provide a value, Elasticsearch will
	// generate one with sufficient entropy and return it in the response.
	Nonce *string `json:"nonce,omitempty"`
	// Realm The name of the OpenID Connect realm in Elasticsearch the configuration of
	// which should be used in order to generate the authentication request.
	// It cannot be specified when *iss* is specified.
	// One of *realm* or *iss* is required.
	Realm *string `json:"realm,omitempty"`
	// State The value used to maintain state between the authentication request and the
	// response, typically used as a Cross-Site Request Forgery mitigation.
	// If the caller of the API does not provide a value, Elasticsearch will
	// generate one with sufficient entropy and return it in the response.
	State *string `json:"state,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Oidcprepareauthentication request: %w", err)
	}

	return &req, nil
}
