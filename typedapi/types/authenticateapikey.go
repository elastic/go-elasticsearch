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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/apikeymanagedby"
)

// AuthenticateApiKey type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/security/authenticate/SecurityAuthenticateResponse.ts#L45-L50
type AuthenticateApiKey struct {
	Id        string                          `json:"id"`
	Internal  *bool                           `json:"internal,omitempty"`
	ManagedBy apikeymanagedby.ApiKeyManagedBy `json:"managed_by"`
	Name      *string                         `json:"name,omitempty"`
}

func (s *AuthenticateApiKey) UnmarshalJSON(data []byte) error {

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

		case "id":
			if err := dec.Decode(&s.Id); err != nil {
				return fmt.Errorf("%s | %w", "Id", err)
			}

		case "internal":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Internal", err)
				}
				s.Internal = &value
			case bool:
				s.Internal = &v
			}

		case "managed_by":
			if err := dec.Decode(&s.ManagedBy); err != nil {
				return fmt.Errorf("%s | %w", "ManagedBy", err)
			}

		case "name":
			if err := dec.Decode(&s.Name); err != nil {
				return fmt.Errorf("%s | %w", "Name", err)
			}

		}
	}
	return nil
}

// NewAuthenticateApiKey returns a AuthenticateApiKey.
func NewAuthenticateApiKey() *AuthenticateApiKey {
	r := &AuthenticateApiKey{}

	return r
}
