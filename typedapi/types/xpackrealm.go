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
// https://github.com/elastic/elasticsearch-specification/tree/1ad7fe36297b3a8e187b2259dedaf68a47bc236e

package types

// XpackRealm type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1ad7fe36297b3a8e187b2259dedaf68a47bc236e/specification/xpack/usage/types.ts#L408-L417
type XpackRealm struct {
	Available                 bool         `json:"available"`
	Cache                     []RealmCache `json:"cache,omitempty"`
	Enabled                   bool         `json:"enabled"`
	HasAuthorizationRealms    []bool       `json:"has_authorization_realms,omitempty"`
	HasDefaultUsernamePattern []bool       `json:"has_default_username_pattern,omitempty"`
	HasTruststore             []bool       `json:"has_truststore,omitempty"`
	IsAuthenticationDelegated []bool       `json:"is_authentication_delegated,omitempty"`
	Name                      []string     `json:"name,omitempty"`
	Order                     []int64      `json:"order,omitempty"`
	Size                      []int64      `json:"size,omitempty"`
}

// NewXpackRealm returns a XpackRealm.
func NewXpackRealm() *XpackRealm {
	r := &XpackRealm{}

	return r
}
