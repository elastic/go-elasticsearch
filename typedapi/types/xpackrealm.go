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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// XpackRealm type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/xpack/usage/types.ts#L423-L432
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

func (s *XpackRealm) UnmarshalJSON(data []byte) error {

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

		case "available":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Available", err)
				}
				s.Available = value
			case bool:
				s.Available = v
			}

		case "cache":
			if err := dec.Decode(&s.Cache); err != nil {
				return fmt.Errorf("%s | %w", "Cache", err)
			}

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "has_authorization_realms":
			if err := dec.Decode(&s.HasAuthorizationRealms); err != nil {
				return fmt.Errorf("%s | %w", "HasAuthorizationRealms", err)
			}

		case "has_default_username_pattern":
			if err := dec.Decode(&s.HasDefaultUsernamePattern); err != nil {
				return fmt.Errorf("%s | %w", "HasDefaultUsernamePattern", err)
			}

		case "has_truststore":
			if err := dec.Decode(&s.HasTruststore); err != nil {
				return fmt.Errorf("%s | %w", "HasTruststore", err)
			}

		case "is_authentication_delegated":
			if err := dec.Decode(&s.IsAuthenticationDelegated); err != nil {
				return fmt.Errorf("%s | %w", "IsAuthenticationDelegated", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		case "order":
			if err := dec.Decode(&s.Order); err != nil {
				return fmt.Errorf("%s | %w", "Order", err)
			}

		case "size":
			if err := dec.Decode(&s.Size); err != nil {
				return fmt.Errorf("%s | %w", "Size", err)
			}

		}
	}
	return nil
}

// NewXpackRealm returns a XpackRealm.
func NewXpackRealm() *XpackRealm {
	r := &XpackRealm{}

	return r
}
