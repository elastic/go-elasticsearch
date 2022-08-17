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

// Hint type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/suggest_user_profiles/types.ts#L23-L34
type Hint struct {
	// Labels A single key-value pair to match against the labels section
	// of a profile. A profile is considered matching if it matches
	// at least one of the strings.
	Labels map[string][]string `json:"labels,omitempty"`
	// Uids A list of Profile UIDs to match against.
	Uids []UserProfileId `json:"uids,omitempty"`
}

// HintBuilder holds Hint struct and provides a builder API.
type HintBuilder struct {
	v *Hint
}

// NewHint provides a builder for the Hint struct.
func NewHintBuilder() *HintBuilder {
	r := HintBuilder{
		&Hint{
			Labels: make(map[string][]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Hint struct
func (rb *HintBuilder) Build() Hint {
	return *rb.v
}

// Labels A single key-value pair to match against the labels section
// of a profile. A profile is considered matching if it matches
// at least one of the strings.

func (rb *HintBuilder) Labels(value map[string][]string) *HintBuilder {
	rb.v.Labels = value
	return rb
}

// Uids A list of Profile UIDs to match against.

func (rb *HintBuilder) Uids(uids ...UserProfileId) *HintBuilder {
	rb.v.Uids = uids
	return rb
}
