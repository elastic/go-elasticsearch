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

// GlobalPrivilege type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/_types/Privileges.ts#L160-L162
type GlobalPrivilege struct {
	Application ApplicationGlobalUserPrivileges `json:"application"`
}

// GlobalPrivilegeBuilder holds GlobalPrivilege struct and provides a builder API.
type GlobalPrivilegeBuilder struct {
	v *GlobalPrivilege
}

// NewGlobalPrivilege provides a builder for the GlobalPrivilege struct.
func NewGlobalPrivilegeBuilder() *GlobalPrivilegeBuilder {
	r := GlobalPrivilegeBuilder{
		&GlobalPrivilege{},
	}

	return &r
}

// Build finalize the chain and returns the GlobalPrivilege struct
func (rb *GlobalPrivilegeBuilder) Build() GlobalPrivilege {
	return *rb.v
}

func (rb *GlobalPrivilegeBuilder) Application(application *ApplicationGlobalUserPrivilegesBuilder) *GlobalPrivilegeBuilder {
	v := application.Build()
	rb.v.Application = v
	return rb
}
