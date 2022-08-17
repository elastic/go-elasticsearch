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

// FieldSecurity type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/_types/FieldSecurity.ts#L22-L25
type FieldSecurity struct {
	Except *Fields `json:"except,omitempty"`
	Grant  Fields  `json:"grant"`
}

// FieldSecurityBuilder holds FieldSecurity struct and provides a builder API.
type FieldSecurityBuilder struct {
	v *FieldSecurity
}

// NewFieldSecurity provides a builder for the FieldSecurity struct.
func NewFieldSecurityBuilder() *FieldSecurityBuilder {
	r := FieldSecurityBuilder{
		&FieldSecurity{},
	}

	return &r
}

// Build finalize the chain and returns the FieldSecurity struct
func (rb *FieldSecurityBuilder) Build() FieldSecurity {
	return *rb.v
}

func (rb *FieldSecurityBuilder) Except(except *FieldsBuilder) *FieldSecurityBuilder {
	v := except.Build()
	rb.v.Except = &v
	return rb
}

func (rb *FieldSecurityBuilder) Grant(grant *FieldsBuilder) *FieldSecurityBuilder {
	v := grant.Build()
	rb.v.Grant = v
	return rb
}
