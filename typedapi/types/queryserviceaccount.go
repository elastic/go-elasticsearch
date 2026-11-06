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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// QueryServiceAccount type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/security/query_service_accounts/types.ts#L103-L112
type QueryServiceAccount struct {
	// CreatedAt When the account was created, in milliseconds since the Unix epoch. Absent
	// for an account that was created before Elasticsearch recorded its creation
	// time.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// CreatedBy The user who created the account. Absent for an account that was created
	// before Elasticsearch recorded its creator.
	CreatedBy *ServiceAccountAuthor `json:"created_by,omitempty"`
	// CreatedByProfileUid The profile UID of the user who created the account. Returned only when
	// `with_profile_uid` is `true` and that user has an activated user profile.
	CreatedByProfileUid *string `json:"created_by_profile_uid,omitempty"`
	// Description A free-text description of the account, as sent on the last PUT of the
	// account. It has no meaning to Elasticsearch. Absent when the account has no
	// description.
	Description *string `json:"description,omitempty"`
	// Enabled Whether the account can authenticate.
	Enabled bool `json:"enabled"`
	// Roles The names of the roles granted to the account, as sent on the last PUT of the
	// account. They are resolved when the account authenticates.
	Roles []string `json:"roles"`
	// Sort_ The sort values of the service account, returned when the request has a
	// `sort`.
	Sort_ []FieldValue `json:"_sort,omitempty"`
	// Type The account was created with the put user-managed service account API.
	Type string `json:"type,omitempty"`
	// UpdatedAt When the account was last replaced, in milliseconds since the Unix epoch.
	// Absent until the account is first replaced.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// UpdatedBy The user who last replaced the account with the put user-managed service
	// account API. Absent until the account is first replaced.
	UpdatedBy *ServiceAccountAuthor `json:"updated_by,omitempty"`
	// UpdatedByProfileUid The profile UID of the user who last replaced the account. Returned only when
	// `with_profile_uid` is `true` and that user has an activated user profile.
	UpdatedByProfileUid *string `json:"updated_by_profile_uid,omitempty"`
	// Username The name of the service account, as `namespace/service`.
	Username string `json:"username"`
}

func (s *QueryServiceAccount) UnmarshalJSON(data []byte) error {

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

		case "created_at":
			if err := dec.Decode(&s.CreatedAt); err != nil {
				return fmt.Errorf("%s | %w", "CreatedAt", err)
			}

		case "created_by":
			if err := dec.Decode(&s.CreatedBy); err != nil {
				return fmt.Errorf("%s | %w", "CreatedBy", err)
			}

		case "created_by_profile_uid":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "CreatedByProfileUid", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CreatedByProfileUid = &o

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

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

		case "roles":
			if err := dec.Decode(&s.Roles); err != nil {
				return fmt.Errorf("%s | %w", "Roles", err)
			}

		case "_sort":
			if err := dec.Decode(&s.Sort_); err != nil {
				return fmt.Errorf("%s | %w", "Sort_", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "updated_at":
			if err := dec.Decode(&s.UpdatedAt); err != nil {
				return fmt.Errorf("%s | %w", "UpdatedAt", err)
			}

		case "updated_by":
			if err := dec.Decode(&s.UpdatedBy); err != nil {
				return fmt.Errorf("%s | %w", "UpdatedBy", err)
			}

		case "updated_by_profile_uid":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "UpdatedByProfileUid", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UpdatedByProfileUid = &o

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return fmt.Errorf("%s | %w", "Username", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s QueryServiceAccount) MarshalJSON() ([]byte, error) {
	type innerQueryServiceAccount QueryServiceAccount
	tmp := innerQueryServiceAccount{
		CreatedAt:           s.CreatedAt,
		CreatedBy:           s.CreatedBy,
		CreatedByProfileUid: s.CreatedByProfileUid,
		Description:         s.Description,
		Enabled:             s.Enabled,
		Roles:               s.Roles,
		Sort_:               s.Sort_,
		Type:                s.Type,
		UpdatedAt:           s.UpdatedAt,
		UpdatedBy:           s.UpdatedBy,
		UpdatedByProfileUid: s.UpdatedByProfileUid,
		Username:            s.Username,
	}

	tmp.Type = "user_managed"

	return json.Marshal(tmp)
}

// NewQueryServiceAccount returns a QueryServiceAccount.
func NewQueryServiceAccount() *QueryServiceAccount {
	r := &QueryServiceAccount{}

	return r
}
