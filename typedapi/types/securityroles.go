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

// SecurityRoles type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L279-L283
type SecurityRoles struct {
	Dls    SecurityRolesDls    `json:"dls"`
	File   SecurityRolesFile   `json:"file"`
	Native SecurityRolesNative `json:"native"`
}

// SecurityRolesBuilder holds SecurityRoles struct and provides a builder API.
type SecurityRolesBuilder struct {
	v *SecurityRoles
}

// NewSecurityRoles provides a builder for the SecurityRoles struct.
func NewSecurityRolesBuilder() *SecurityRolesBuilder {
	r := SecurityRolesBuilder{
		&SecurityRoles{},
	}

	return &r
}

// Build finalize the chain and returns the SecurityRoles struct
func (rb *SecurityRolesBuilder) Build() SecurityRoles {
	return *rb.v
}

func (rb *SecurityRolesBuilder) Dls(dls *SecurityRolesDlsBuilder) *SecurityRolesBuilder {
	v := dls.Build()
	rb.v.Dls = v
	return rb
}

func (rb *SecurityRolesBuilder) File(file *SecurityRolesFileBuilder) *SecurityRolesBuilder {
	v := file.Build()
	rb.v.File = v
	return rb
}

func (rb *SecurityRolesBuilder) Native(native *SecurityRolesNativeBuilder) *SecurityRolesBuilder {
	v := native.Build()
	rb.v.Native = v
	return rb
}
