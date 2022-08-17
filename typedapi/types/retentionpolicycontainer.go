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

// RetentionPolicyContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/_types/Transform.ts#L80-L86
type RetentionPolicyContainer struct {
	// Time Specifies that the transform uses a time field to set the retention policy.
	Time *RetentionPolicy `json:"time,omitempty"`
}

// RetentionPolicyContainerBuilder holds RetentionPolicyContainer struct and provides a builder API.
type RetentionPolicyContainerBuilder struct {
	v *RetentionPolicyContainer
}

// NewRetentionPolicyContainer provides a builder for the RetentionPolicyContainer struct.
func NewRetentionPolicyContainerBuilder() *RetentionPolicyContainerBuilder {
	r := RetentionPolicyContainerBuilder{
		&RetentionPolicyContainer{},
	}

	return &r
}

// Build finalize the chain and returns the RetentionPolicyContainer struct
func (rb *RetentionPolicyContainerBuilder) Build() RetentionPolicyContainer {
	return *rb.v
}

// Time Specifies that the transform uses a time field to set the retention policy.

func (rb *RetentionPolicyContainerBuilder) Time(time *RetentionPolicyBuilder) *RetentionPolicyContainerBuilder {
	v := time.Build()
	rb.v.Time = &v
	return rb
}
