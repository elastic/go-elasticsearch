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

// SettingsSearch type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L231-L234
type SettingsSearch struct {
	Idle    *SearchIdle      `json:"idle,omitempty"`
	Slowlog *SlowlogSettings `json:"slowlog,omitempty"`
}

// SettingsSearchBuilder holds SettingsSearch struct and provides a builder API.
type SettingsSearchBuilder struct {
	v *SettingsSearch
}

// NewSettingsSearch provides a builder for the SettingsSearch struct.
func NewSettingsSearchBuilder() *SettingsSearchBuilder {
	r := SettingsSearchBuilder{
		&SettingsSearch{},
	}

	return &r
}

// Build finalize the chain and returns the SettingsSearch struct
func (rb *SettingsSearchBuilder) Build() SettingsSearch {
	return *rb.v
}

func (rb *SettingsSearchBuilder) Idle(idle *SearchIdleBuilder) *SettingsSearchBuilder {
	v := idle.Build()
	rb.v.Idle = &v
	return rb
}

func (rb *SettingsSearchBuilder) Slowlog(slowlog *SlowlogSettingsBuilder) *SettingsSearchBuilder {
	v := slowlog.Build()
	rb.v.Slowlog = &v
	return rb
}
