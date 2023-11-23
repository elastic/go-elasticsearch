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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// UserProfileWithMetadata type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/security/_types/UserProfile.ts#L50-L53
type UserProfileWithMetadata struct {
	Data             map[string]json.RawMessage `json:"data"`
	Doc_             UserProfileHitMetadata     `json:"_doc"`
	Enabled          *bool                      `json:"enabled,omitempty"`
	Labels           map[string]json.RawMessage `json:"labels"`
	LastSynchronized int64                      `json:"last_synchronized"`
	Uid              string                     `json:"uid"`
	User             UserProfileUser            `json:"user"`
}

func (s *UserProfileWithMetadata) UnmarshalJSON(data []byte) error {

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

		case "_doc":
			if err := dec.Decode(&s.Doc_); err != nil {
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

		case "last_synchronized":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return err
				}
				s.LastSynchronized = value
			case float64:
				f := int64(v)
				s.LastSynchronized = f
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

// NewUserProfileWithMetadata returns a UserProfileWithMetadata.
func NewUserProfileWithMetadata() *UserProfileWithMetadata {
	r := &UserProfileWithMetadata{
		Data:   make(map[string]json.RawMessage, 0),
		Labels: make(map[string]json.RawMessage, 0),
	}

	return r
}
