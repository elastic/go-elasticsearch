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

// CgroupMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L206-L210
type CgroupMemory struct {
	ControlGroup *string `json:"control_group,omitempty"`
	LimitInBytes *string `json:"limit_in_bytes,omitempty"`
	UsageInBytes *string `json:"usage_in_bytes,omitempty"`
}

// CgroupMemoryBuilder holds CgroupMemory struct and provides a builder API.
type CgroupMemoryBuilder struct {
	v *CgroupMemory
}

// NewCgroupMemory provides a builder for the CgroupMemory struct.
func NewCgroupMemoryBuilder() *CgroupMemoryBuilder {
	r := CgroupMemoryBuilder{
		&CgroupMemory{},
	}

	return &r
}

// Build finalize the chain and returns the CgroupMemory struct
func (rb *CgroupMemoryBuilder) Build() CgroupMemory {
	return *rb.v
}

func (rb *CgroupMemoryBuilder) ControlGroup(controlgroup string) *CgroupMemoryBuilder {
	rb.v.ControlGroup = &controlgroup
	return rb
}

func (rb *CgroupMemoryBuilder) LimitInBytes(limitinbytes string) *CgroupMemoryBuilder {
	rb.v.LimitInBytes = &limitinbytes
	return rb
}

func (rb *CgroupMemoryBuilder) UsageInBytes(usageinbytes string) *CgroupMemoryBuilder {
	rb.v.UsageInBytes = &usageinbytes
	return rb
}
