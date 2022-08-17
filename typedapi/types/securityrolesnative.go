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

// SecurityRolesNative type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L285-L289
type SecurityRolesNative struct {
	Dls  bool  `json:"dls"`
	Fls  bool  `json:"fls"`
	Size int64 `json:"size"`
}

// SecurityRolesNativeBuilder holds SecurityRolesNative struct and provides a builder API.
type SecurityRolesNativeBuilder struct {
	v *SecurityRolesNative
}

// NewSecurityRolesNative provides a builder for the SecurityRolesNative struct.
func NewSecurityRolesNativeBuilder() *SecurityRolesNativeBuilder {
	r := SecurityRolesNativeBuilder{
		&SecurityRolesNative{},
	}

	return &r
}

// Build finalize the chain and returns the SecurityRolesNative struct
func (rb *SecurityRolesNativeBuilder) Build() SecurityRolesNative {
	return *rb.v
}

func (rb *SecurityRolesNativeBuilder) Dls(dls bool) *SecurityRolesNativeBuilder {
	rb.v.Dls = dls
	return rb
}

func (rb *SecurityRolesNativeBuilder) Fls(fls bool) *SecurityRolesNativeBuilder {
	rb.v.Fls = fls
	return rb
}

func (rb *SecurityRolesNativeBuilder) Size(size int64) *SecurityRolesNativeBuilder {
	rb.v.Size = size
	return rb
}
