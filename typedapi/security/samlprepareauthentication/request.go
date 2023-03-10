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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package samlprepareauthentication

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package samlprepareauthentication
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/security/saml_prepare_authentication/Request.ts#L22-L46
type Request struct {

	// Acs The Assertion Consumer Service URL that matches the one of the SAML realms in
	// Elasticsearch.
	// The realm is used to generate the authentication request. You must specify
	// either this parameter or the realm parameter.
	Acs *string `json:"acs,omitempty"`
	// Realm The name of the SAML realm in Elasticsearch for which the configuration is
	// used to generate the authentication request.
	// You must specify either this parameter or the acs parameter.
	Realm *string `json:"realm,omitempty"`
	// RelayState A string that will be included in the redirect URL that this API returns as
	// the RelayState query parameter.
	// If the Authentication Request is signed, this value is used as part of the
	// signature computation.
	RelayState *string `json:"relay_state,omitempty"`
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
		return nil, fmt.Errorf("could not deserialise json into Samlprepareauthentication request: %w", err)
	}

	return &req, nil
}
