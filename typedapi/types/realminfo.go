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

// RealmInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/_types/RealmInfo.ts#L22-L25
type RealmInfo struct {
	Name Name   `json:"name"`
	Type string `json:"type"`
}

// RealmInfoBuilder holds RealmInfo struct and provides a builder API.
type RealmInfoBuilder struct {
	v *RealmInfo
}

// NewRealmInfo provides a builder for the RealmInfo struct.
func NewRealmInfoBuilder() *RealmInfoBuilder {
	r := RealmInfoBuilder{
		&RealmInfo{},
	}

	return &r
}

// Build finalize the chain and returns the RealmInfo struct
func (rb *RealmInfoBuilder) Build() RealmInfo {
	return *rb.v
}

func (rb *RealmInfoBuilder) Name(name Name) *RealmInfoBuilder {
	rb.v.Name = name
	return rb
}

func (rb *RealmInfoBuilder) Type_(type_ string) *RealmInfoBuilder {
	rb.v.Type = type_
	return rb
}
