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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// NestedIdentity type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_global/search/_types/hits.ts#L84-L88
type NestedIdentity struct {
	Field   Field           `json:"field"`
	Nested_ *NestedIdentity `json:"_nested,omitempty"`
	Offset  int             `json:"offset"`
}

// NestedIdentityBuilder holds NestedIdentity struct and provides a builder API.
type NestedIdentityBuilder struct {
	v *NestedIdentity
}

// NewNestedIdentity provides a builder for the NestedIdentity struct.
func NewNestedIdentityBuilder() *NestedIdentityBuilder {
	r := NestedIdentityBuilder{
		&NestedIdentity{},
	}

	return &r
}

// Build finalize the chain and returns the NestedIdentity struct
func (rb *NestedIdentityBuilder) Build() NestedIdentity {
	return *rb.v
}

func (rb *NestedIdentityBuilder) Field(field Field) *NestedIdentityBuilder {
	rb.v.Field = field
	return rb
}

func (rb *NestedIdentityBuilder) Nested_(nested_ *NestedIdentityBuilder) *NestedIdentityBuilder {
	v := nested_.Build()
	rb.v.Nested_ = &v
	return rb
}

func (rb *NestedIdentityBuilder) Offset(offset int) *NestedIdentityBuilder {
	rb.v.Offset = offset
	return rb
}
