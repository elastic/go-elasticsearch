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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package samlprepareauthentication

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package samlprepareauthentication
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/saml_prepare_authentication/Request.ts#L22-L46
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

// RequestBuilder is the builder API for the samlprepareauthentication.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Samlprepareauthentication request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Acs(acs string) *RequestBuilder {
	rb.v.Acs = &acs
	return rb
}

func (rb *RequestBuilder) Realm(realm string) *RequestBuilder {
	rb.v.Realm = &realm
	return rb
}

func (rb *RequestBuilder) RelayState(relaystate string) *RequestBuilder {
	rb.v.RelayState = &relaystate
	return rb
}
