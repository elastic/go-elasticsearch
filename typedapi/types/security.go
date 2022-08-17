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


package types

// Security type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L423-L436
type Security struct {
	Anonymous          FeatureToggle          `json:"anonymous"`
	ApiKeyService      FeatureToggle          `json:"api_key_service"`
	Audit              Audit                  `json:"audit"`
	Available          bool                   `json:"available"`
	Enabled            bool                   `json:"enabled"`
	Fips140            FeatureToggle          `json:"fips_140"`
	Ipfilter           IpFilter               `json:"ipfilter"`
	OperatorPrivileges Base                   `json:"operator_privileges"`
	Realms             map[string]Realm       `json:"realms"`
	RoleMapping        map[string]RoleMapping `json:"role_mapping"`
	Roles              SecurityRoles          `json:"roles"`
	Ssl                Ssl                    `json:"ssl"`
	SystemKey          *FeatureToggle         `json:"system_key,omitempty"`
	TokenService       FeatureToggle          `json:"token_service"`
}

// SecurityBuilder holds Security struct and provides a builder API.
type SecurityBuilder struct {
	v *Security
}

// NewSecurity provides a builder for the Security struct.
func NewSecurityBuilder() *SecurityBuilder {
	r := SecurityBuilder{
		&Security{
			Realms:      make(map[string]Realm, 0),
			RoleMapping: make(map[string]RoleMapping, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Security struct
func (rb *SecurityBuilder) Build() Security {
	return *rb.v
}

func (rb *SecurityBuilder) Anonymous(anonymous *FeatureToggleBuilder) *SecurityBuilder {
	v := anonymous.Build()
	rb.v.Anonymous = v
	return rb
}

func (rb *SecurityBuilder) ApiKeyService(apikeyservice *FeatureToggleBuilder) *SecurityBuilder {
	v := apikeyservice.Build()
	rb.v.ApiKeyService = v
	return rb
}

func (rb *SecurityBuilder) Audit(audit *AuditBuilder) *SecurityBuilder {
	v := audit.Build()
	rb.v.Audit = v
	return rb
}

func (rb *SecurityBuilder) Available(available bool) *SecurityBuilder {
	rb.v.Available = available
	return rb
}

func (rb *SecurityBuilder) Enabled(enabled bool) *SecurityBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *SecurityBuilder) Fips140(fips140 *FeatureToggleBuilder) *SecurityBuilder {
	v := fips140.Build()
	rb.v.Fips140 = v
	return rb
}

func (rb *SecurityBuilder) Ipfilter(ipfilter *IpFilterBuilder) *SecurityBuilder {
	v := ipfilter.Build()
	rb.v.Ipfilter = v
	return rb
}

func (rb *SecurityBuilder) OperatorPrivileges(operatorprivileges *BaseBuilder) *SecurityBuilder {
	v := operatorprivileges.Build()
	rb.v.OperatorPrivileges = v
	return rb
}

func (rb *SecurityBuilder) Realms(values map[string]*RealmBuilder) *SecurityBuilder {
	tmp := make(map[string]Realm, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Realms = tmp
	return rb
}

func (rb *SecurityBuilder) RoleMapping(values map[string]*RoleMappingBuilder) *SecurityBuilder {
	tmp := make(map[string]RoleMapping, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.RoleMapping = tmp
	return rb
}

func (rb *SecurityBuilder) Roles(roles *SecurityRolesBuilder) *SecurityBuilder {
	v := roles.Build()
	rb.v.Roles = v
	return rb
}

func (rb *SecurityBuilder) Ssl(ssl *SslBuilder) *SecurityBuilder {
	v := ssl.Build()
	rb.v.Ssl = v
	return rb
}

func (rb *SecurityBuilder) SystemKey(systemkey *FeatureToggleBuilder) *SecurityBuilder {
	v := systemkey.Build()
	rb.v.SystemKey = &v
	return rb
}

func (rb *SecurityBuilder) TokenService(tokenservice *FeatureToggleBuilder) *SecurityBuilder {
	v := tokenservice.Build()
	rb.v.TokenService = v
	return rb
}
