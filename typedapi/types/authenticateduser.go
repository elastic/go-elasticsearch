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

package types

import (
	"encoding/json"
)

// AuthenticatedUser type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/security/get_token/types.ts#L40-L45
type AuthenticatedUser struct {
	AuthenticationProvider *AuthenticationProvider    `json:"authentication_provider,omitempty"`
	AuthenticationRealm    UserRealm                  `json:"authentication_realm"`
	AuthenticationType     string                     `json:"authentication_type"`
	Email                  string                     `json:"email,omitempty"`
	Enabled                bool                       `json:"enabled"`
	FullName               string                     `json:"full_name,omitempty"`
	LookupRealm            UserRealm                  `json:"lookup_realm"`
	Metadata               map[string]json.RawMessage `json:"metadata"`
	ProfileUid             *string                    `json:"profile_uid,omitempty"`
	Roles                  []string                   `json:"roles"`
	Username               string                     `json:"username"`
}

// NewAuthenticatedUser returns a AuthenticatedUser.
func NewAuthenticatedUser() *AuthenticatedUser {
	r := &AuthenticatedUser{}

	return r
}
