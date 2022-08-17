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
// https://github.com/elastic/elasticsearch-specification/tree/e0ea3dc890d394d682096cc862b3bd879d9422e9


package samllogout

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package samllogout
//
// https://github.com/elastic/elasticsearch-specification/blob/e0ea3dc890d394d682096cc862b3bd879d9422e9/specification/security/saml_logout/Request.ts#L22-L41
type Request struct {

	// RefreshToken The refresh token that was returned as a response to calling the SAML
	// authenticate API.
	// Alternatively, the most recent refresh token that was received after
	// refreshing the original access token.
	RefreshToken *string `json:"refresh_token,omitempty"`

	// Token The access token that was returned as a response to calling the SAML
	// authenticate API.
	// Alternatively, the most recent token that was received after refreshing the
	// original one by using a refresh_token.
	Token string `json:"token"`
}

// RequestBuilder is the builder API for the samllogout.Request
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
		return nil, fmt.Errorf("could not deserialise json into Samllogout request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) RefreshToken(refreshtoken string) *RequestBuilder {
	rb.v.RefreshToken = &refreshtoken
	return rb
}

func (rb *RequestBuilder) Token(token string) *RequestBuilder {
	rb.v.Token = token
	return rb
}
