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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package samllogout

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package samllogout
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/security/saml_logout/Request.ts#L22-L41
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

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Samllogout request: %w", err)
	}

	return &req, nil
}
