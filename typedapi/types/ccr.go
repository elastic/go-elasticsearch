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

// Ccr type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L317-L320
type Ccr struct {
	AutoFollowPatternsCount int  `json:"auto_follow_patterns_count"`
	Available               bool `json:"available"`
	Enabled                 bool `json:"enabled"`
	FollowerIndicesCount    int  `json:"follower_indices_count"`
}

// CcrBuilder holds Ccr struct and provides a builder API.
type CcrBuilder struct {
	v *Ccr
}

// NewCcr provides a builder for the Ccr struct.
func NewCcrBuilder() *CcrBuilder {
	r := CcrBuilder{
		&Ccr{},
	}

	return &r
}

// Build finalize the chain and returns the Ccr struct
func (rb *CcrBuilder) Build() Ccr {
	return *rb.v
}

func (rb *CcrBuilder) AutoFollowPatternsCount(autofollowpatternscount int) *CcrBuilder {
	rb.v.AutoFollowPatternsCount = autofollowpatternscount
	return rb
}

func (rb *CcrBuilder) Available(available bool) *CcrBuilder {
	rb.v.Available = available
	return rb
}

func (rb *CcrBuilder) Enabled(enabled bool) *CcrBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *CcrBuilder) FollowerIndicesCount(followerindicescount int) *CcrBuilder {
	rb.v.FollowerIndicesCount = followerindicescount
	return rb
}
