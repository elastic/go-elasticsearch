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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// Profile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_global/search/_types/profile.ts#L93-L95
type Profile struct {
	Shards []ShardProfile `json:"shards"`
}

// ProfileBuilder holds Profile struct and provides a builder API.
type ProfileBuilder struct {
	v *Profile
}

// NewProfile provides a builder for the Profile struct.
func NewProfileBuilder() *ProfileBuilder {
	r := ProfileBuilder{
		&Profile{},
	}

	return &r
}

// Build finalize the chain and returns the Profile struct
func (rb *ProfileBuilder) Build() Profile {
	return *rb.v
}

func (rb *ProfileBuilder) Shards(shards []ShardProfileBuilder) *ProfileBuilder {
	tmp := make([]ShardProfile, len(shards))
	for _, value := range shards {
		tmp = append(tmp, value.Build())
	}
	rb.v.Shards = tmp
	return rb
}
