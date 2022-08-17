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

// RetentionPolicy type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/_types/Transform.ts#L88-L96
type RetentionPolicy struct {
	// Field The date field that is used to calculate the age of the document.
	Field Field `json:"field"`
	// MaxAge Specifies the maximum age of a document in the destination index. Documents
	// that are older than the configured
	// value are removed from the destination index.
	MaxAge Duration `json:"max_age"`
}

// RetentionPolicyBuilder holds RetentionPolicy struct and provides a builder API.
type RetentionPolicyBuilder struct {
	v *RetentionPolicy
}

// NewRetentionPolicy provides a builder for the RetentionPolicy struct.
func NewRetentionPolicyBuilder() *RetentionPolicyBuilder {
	r := RetentionPolicyBuilder{
		&RetentionPolicy{},
	}

	return &r
}

// Build finalize the chain and returns the RetentionPolicy struct
func (rb *RetentionPolicyBuilder) Build() RetentionPolicy {
	return *rb.v
}

// Field The date field that is used to calculate the age of the document.

func (rb *RetentionPolicyBuilder) Field(field Field) *RetentionPolicyBuilder {
	rb.v.Field = field
	return rb
}

// MaxAge Specifies the maximum age of a document in the destination index. Documents
// that are older than the configured
// value are removed from the destination index.

func (rb *RetentionPolicyBuilder) MaxAge(maxage *DurationBuilder) *RetentionPolicyBuilder {
	v := maxage.Build()
	rb.v.MaxAge = v
	return rb
}
