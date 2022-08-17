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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexprivilege"
)

// IndexPrivilegesCheck type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/has_privileges/types.ts#L33-L44
type IndexPrivilegesCheck struct {
	// AllowRestrictedIndices This needs to be set to true (default is false) if using wildcards or regexps
	// for patterns that cover restricted indices.
	// Implicitly, restricted indices do not match index patterns because restricted
	// indices usually have limited privileges and including them in pattern tests
	// would render most such tests false.
	// If restricted indices are explicitly included in the names list, privileges
	// will be checked against them regardless of the value of
	// allow_restricted_indices.
	AllowRestrictedIndices *bool `json:"allow_restricted_indices,omitempty"`
	// Names A list of indices.
	Names Indices `json:"names"`
	// Privileges A list of the privileges that you want to check for the specified indices.
	Privileges []indexprivilege.IndexPrivilege `json:"privileges"`
}

// IndexPrivilegesCheckBuilder holds IndexPrivilegesCheck struct and provides a builder API.
type IndexPrivilegesCheckBuilder struct {
	v *IndexPrivilegesCheck
}

// NewIndexPrivilegesCheck provides a builder for the IndexPrivilegesCheck struct.
func NewIndexPrivilegesCheckBuilder() *IndexPrivilegesCheckBuilder {
	r := IndexPrivilegesCheckBuilder{
		&IndexPrivilegesCheck{},
	}

	return &r
}

// Build finalize the chain and returns the IndexPrivilegesCheck struct
func (rb *IndexPrivilegesCheckBuilder) Build() IndexPrivilegesCheck {
	return *rb.v
}

// AllowRestrictedIndices This needs to be set to true (default is false) if using wildcards or regexps
// for patterns that cover restricted indices.
// Implicitly, restricted indices do not match index patterns because restricted
// indices usually have limited privileges and including them in pattern tests
// would render most such tests false.
// If restricted indices are explicitly included in the names list, privileges
// will be checked against them regardless of the value of
// allow_restricted_indices.

func (rb *IndexPrivilegesCheckBuilder) AllowRestrictedIndices(allowrestrictedindices bool) *IndexPrivilegesCheckBuilder {
	rb.v.AllowRestrictedIndices = &allowrestrictedindices
	return rb
}

// Names A list of indices.

func (rb *IndexPrivilegesCheckBuilder) Names(names *IndicesBuilder) *IndexPrivilegesCheckBuilder {
	v := names.Build()
	rb.v.Names = v
	return rb
}

// Privileges A list of the privileges that you want to check for the specified indices.

func (rb *IndexPrivilegesCheckBuilder) Privileges(privileges ...indexprivilege.IndexPrivilege) *IndexPrivilegesCheckBuilder {
	rb.v.Privileges = privileges
	return rb
}
