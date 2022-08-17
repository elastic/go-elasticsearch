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

// Lifecycle type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ilm/get_lifecycle/types.ts#L24-L28
type Lifecycle struct {
	ModifiedDate DateTime      `json:"modified_date"`
	Policy       Policy        `json:"policy"`
	Version      VersionNumber `json:"version"`
}

// LifecycleBuilder holds Lifecycle struct and provides a builder API.
type LifecycleBuilder struct {
	v *Lifecycle
}

// NewLifecycle provides a builder for the Lifecycle struct.
func NewLifecycleBuilder() *LifecycleBuilder {
	r := LifecycleBuilder{
		&Lifecycle{},
	}

	return &r
}

// Build finalize the chain and returns the Lifecycle struct
func (rb *LifecycleBuilder) Build() Lifecycle {
	return *rb.v
}

func (rb *LifecycleBuilder) ModifiedDate(modifieddate *DateTimeBuilder) *LifecycleBuilder {
	v := modifieddate.Build()
	rb.v.ModifiedDate = v
	return rb
}

func (rb *LifecycleBuilder) Policy(policy *PolicyBuilder) *LifecycleBuilder {
	v := policy.Build()
	rb.v.Policy = v
	return rb
}

func (rb *LifecycleBuilder) Version(version VersionNumber) *LifecycleBuilder {
	rb.v.Version = version
	return rb
}
