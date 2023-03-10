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

// Security type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/xpack/usage/types.ts#L425-L438
type Security struct {
	Anonymous          FeatureToggle               `json:"anonymous"`
	ApiKeyService      FeatureToggle               `json:"api_key_service"`
	Audit              Audit                       `json:"audit"`
	Available          bool                        `json:"available"`
	Enabled            bool                        `json:"enabled"`
	Fips140            FeatureToggle               `json:"fips_140"`
	Ipfilter           IpFilter                    `json:"ipfilter"`
	OperatorPrivileges Base                        `json:"operator_privileges"`
	Realms             map[string]XpackRealm       `json:"realms"`
	RoleMapping        map[string]XpackRoleMapping `json:"role_mapping"`
	Roles              SecurityRoles               `json:"roles"`
	Ssl                Ssl                         `json:"ssl"`
	SystemKey          *FeatureToggle              `json:"system_key,omitempty"`
	TokenService       FeatureToggle               `json:"token_service"`
}

// NewSecurity returns a Security.
func NewSecurity() *Security {
	r := &Security{
		Realms:      make(map[string]XpackRealm, 0),
		RoleMapping: make(map[string]XpackRoleMapping, 0),
	}

	return r
}
