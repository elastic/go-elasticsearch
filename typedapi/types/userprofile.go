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

// UserProfile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/_types/UserProfile.ts#L43-L49
type UserProfile struct {
	Data    map[string]interface{} `json:"data"`
	Enabled *bool                  `json:"enabled,omitempty"`
	Labels  map[string]interface{} `json:"labels"`
	Uid     UserProfileId          `json:"uid"`
	User    UserProfileUser        `json:"user"`
}

// UserProfileBuilder holds UserProfile struct and provides a builder API.
type UserProfileBuilder struct {
	v *UserProfile
}

// NewUserProfile provides a builder for the UserProfile struct.
func NewUserProfileBuilder() *UserProfileBuilder {
	r := UserProfileBuilder{
		&UserProfile{
			Data:   make(map[string]interface{}, 0),
			Labels: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the UserProfile struct
func (rb *UserProfileBuilder) Build() UserProfile {
	return *rb.v
}

func (rb *UserProfileBuilder) Data(value map[string]interface{}) *UserProfileBuilder {
	rb.v.Data = value
	return rb
}

func (rb *UserProfileBuilder) Enabled(enabled bool) *UserProfileBuilder {
	rb.v.Enabled = &enabled
	return rb
}

func (rb *UserProfileBuilder) Labels(value map[string]interface{}) *UserProfileBuilder {
	rb.v.Labels = value
	return rb
}

func (rb *UserProfileBuilder) Uid(uid UserProfileId) *UserProfileBuilder {
	rb.v.Uid = uid
	return rb
}

func (rb *UserProfileBuilder) User(user *UserProfileUserBuilder) *UserProfileBuilder {
	v := user.Build()
	rb.v.User = v
	return rb
}
