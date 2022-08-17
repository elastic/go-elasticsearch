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

// RetentionLease type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/IndexSettings.ts#L65-L67
type RetentionLease struct {
	Period Duration `json:"period"`
}

// RetentionLeaseBuilder holds RetentionLease struct and provides a builder API.
type RetentionLeaseBuilder struct {
	v *RetentionLease
}

// NewRetentionLease provides a builder for the RetentionLease struct.
func NewRetentionLeaseBuilder() *RetentionLeaseBuilder {
	r := RetentionLeaseBuilder{
		&RetentionLease{},
	}

	return &r
}

// Build finalize the chain and returns the RetentionLease struct
func (rb *RetentionLeaseBuilder) Build() RetentionLease {
	return *rb.v
}

func (rb *RetentionLeaseBuilder) Period(period *DurationBuilder) *RetentionLeaseBuilder {
	v := period.Build()
	rb.v.Period = v
	return rb
}
