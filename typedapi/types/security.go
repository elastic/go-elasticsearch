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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// Security type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/xpack/usage/types.ts#L425-L438
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

func (s *Security) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "anonymous":
			if err := dec.Decode(&s.Anonymous); err != nil {
				return err
			}

		case "api_key_service":
			if err := dec.Decode(&s.ApiKeyService); err != nil {
				return err
			}

		case "audit":
			if err := dec.Decode(&s.Audit); err != nil {
				return err
			}

		case "available":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "enabled":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "fips_140":
			if err := dec.Decode(&s.Fips140); err != nil {
				return err
			}

		case "ipfilter":
			if err := dec.Decode(&s.Ipfilter); err != nil {
				return err
			}

		case "operator_privileges":
			if err := dec.Decode(&s.OperatorPrivileges); err != nil {
				return err
			}

		case "realms":
			if s.Realms == nil {
				s.Realms = make(map[string]XpackRealm, 0)
			}
			if err := dec.Decode(&s.Realms); err != nil {
				return err
			}

		case "role_mapping":
			if s.RoleMapping == nil {
				s.RoleMapping = make(map[string]XpackRoleMapping, 0)
			}
			if err := dec.Decode(&s.RoleMapping); err != nil {
				return err
			}

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return err
			}

		case "ssl":
			if err := dec.Decode(&s.Ssl); err != nil {
				return err
			}

		case "system_key":
			if err := dec.Decode(&s.SystemKey); err != nil {
				return err
			}

		case "token_service":
			if err := dec.Decode(&s.TokenService); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewSecurity returns a Security.
func NewSecurity() *Security {
	r := &Security{
		Realms:      make(map[string]XpackRealm, 0),
		RoleMapping: make(map[string]XpackRoleMapping, 0),
	}

	return r
}
