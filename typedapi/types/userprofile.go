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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// UserProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/security/_types/UserProfile.ts#L41-L47
type UserProfile struct {
	Data    map[string]json.RawMessage `json:"data"`
	Enabled *bool                      `json:"enabled,omitempty"`
	Labels  map[string]json.RawMessage `json:"labels"`
	Uid     string                     `json:"uid"`
	User    UserProfileUser            `json:"user"`
}

func (s *UserProfile) UnmarshalJSON(data []byte) error {

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

		case "data":
			if s.Data == nil {
				s.Data = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Data); err != nil {
				return fmt.Errorf("%s | %w", "Data", err)
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
				s.Enabled = &value
			case bool:
				s.Enabled = &v
			}

		case "labels":
			if s.Labels == nil {
				s.Labels = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Labels); err != nil {
				return fmt.Errorf("%s | %w", "Labels", err)
			}

		case "uid":
			if err := dec.Decode(&s.Uid); err != nil {
				return fmt.Errorf("%s | %w", "Uid", err)
			}

		case "user":
			if err := dec.Decode(&s.User); err != nil {
				return fmt.Errorf("%s | %w", "User", err)
			}

		}
	}
	return nil
}

// NewUserProfile returns a UserProfile.
func NewUserProfile() *UserProfile {
	r := &UserProfile{
		Data:   make(map[string]json.RawMessage),
		Labels: make(map[string]json.RawMessage),
	}

	return r
}
