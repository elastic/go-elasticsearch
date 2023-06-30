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

// UserProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/security/_types/UserProfile.ts#L42-L48
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
				return err
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
				s.Enabled = &value
			case bool:
				s.Enabled = &v
			}

		case "labels":
			if s.Labels == nil {
				s.Labels = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Labels); err != nil {
				return err
			}

		case "uid":
			if err := dec.Decode(&s.Uid); err != nil {
				return err
			}

		case "user":
			if err := dec.Decode(&s.User); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewUserProfile returns a UserProfile.
func NewUserProfile() *UserProfile {
	r := &UserProfile{
		Data:   make(map[string]json.RawMessage, 0),
		Labels: make(map[string]json.RawMessage, 0),
	}

	return r
}
