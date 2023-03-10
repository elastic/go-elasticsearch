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

import (
	"encoding/json"
)

// UserProfileWithMetadata type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/security/_types/UserProfile.ts#L50-L53
type UserProfileWithMetadata struct {
	Data             map[string]json.RawMessage `json:"data"`
	Doc_             UserProfileHitMetadata     `json:"_doc"`
	Enabled          *bool                      `json:"enabled,omitempty"`
	Labels           map[string]json.RawMessage `json:"labels"`
	LastSynchronized int64                      `json:"last_synchronized"`
	Uid              string                     `json:"uid"`
	User             UserProfileUser            `json:"user"`
}

// NewUserProfileWithMetadata returns a UserProfileWithMetadata.
func NewUserProfileWithMetadata() *UserProfileWithMetadata {
	r := &UserProfileWithMetadata{
		Data:   make(map[string]json.RawMessage, 0),
		Labels: make(map[string]json.RawMessage, 0),
	}

	return r
}
