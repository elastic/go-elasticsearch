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
// https://github.com/elastic/elasticsearch-specification/tree/5c8fed5fe577b0d5e9fde34fb13795c5a66fe9fe

package gettoken

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package gettoken
//
// https://github.com/elastic/elasticsearch-specification/blob/5c8fed5fe577b0d5e9fde34fb13795c5a66fe9fe/specification/security/get_token/GetUserAccessTokenResponse.ts#L23-L33
type Response struct {
	AccessToken                         string                  `json:"access_token"`
	Authentication                      types.AuthenticatedUser `json:"authentication"`
	ExpiresIn                           int64                   `json:"expires_in"`
	KerberosAuthenticationResponseToken *string                 `json:"kerberos_authentication_response_token,omitempty"`
	RefreshToken                        *string                 `json:"refresh_token,omitempty"`
	Scope                               *string                 `json:"scope,omitempty"`
	Type                                string                  `json:"type"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
