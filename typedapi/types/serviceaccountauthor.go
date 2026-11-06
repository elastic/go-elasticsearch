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

// The user who wrote a user-managed service account. If the request ran as
// another user, that user is recorded. If the request was made with an API key,
// the owner of the key is recorded, from the owner's own realm, together with
// the key itself. For an older API key that does not record its owner's realm,
// the realm is the API key realm, `_es_api_key`.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/security/get_service_accounts/types.ts#L111-L135
type ServiceAccountAuthor struct {
	// ApiKey The API key that the request was made with. Absent when the request was not
	// made with an API key, or when the API key ran as another user.
	ApiKey *ServiceAccountAuthorApiKey `json:"api_key,omitempty"`
	// Realm The name of the realm that authenticated the user.
	Realm string `json:"realm"`
	// RealmType The type of the realm that authenticated the user.
	RealmType string `json:"realm_type"`
	// Username The username of the user.
	Username string `json:"username"`
}

func (s *ServiceAccountAuthor) UnmarshalJSON(data []byte) error {

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

		case "api_key":
			if err := dec.Decode(&s.ApiKey); err != nil {
				return fmt.Errorf("%s | %w", "ApiKey", err)
			}

		case "realm":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Realm", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Realm = o

		case "realm_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RealmType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RealmType = o

		case "username":
			if err := dec.Decode(&s.Username); err != nil {
				return fmt.Errorf("%s | %w", "Username", err)
			}

		}
	}
	return nil
}

// NewServiceAccountAuthor returns a ServiceAccountAuthor.
func NewServiceAccountAuthor() *ServiceAccountAuthor {
	r := &ServiceAccountAuthor{}

	return r
}
