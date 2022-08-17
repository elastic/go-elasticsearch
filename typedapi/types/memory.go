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

// Memory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/get_memory_stats/types.ts#L25-L48
type Memory struct {
	Attributes  map[string]string `json:"attributes"`
	EphemeralId Id                `json:"ephemeral_id"`
	// Jvm Contains Java Virtual Machine (JVM) statistics for the node.
	Jvm JvmStats `json:"jvm"`
	// Mem Contains statistics about memory usage for the node.
	Mem MemStats `json:"mem"`
	// Name Human-readable identifier for the node. Based on the Node name setting
	// setting.
	Name Name `json:"name"`
	// Roles Roles assigned to the node.
	Roles []string `json:"roles"`
	// TransportAddress The host and port where transport HTTP connections are accepted.
	TransportAddress TransportAddress `json:"transport_address"`
}

// MemoryBuilder holds Memory struct and provides a builder API.
type MemoryBuilder struct {
	v *Memory
}

// NewMemory provides a builder for the Memory struct.
func NewMemoryBuilder() *MemoryBuilder {
	r := MemoryBuilder{
		&Memory{
			Attributes: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Memory struct
func (rb *MemoryBuilder) Build() Memory {
	return *rb.v
}

func (rb *MemoryBuilder) Attributes(value map[string]string) *MemoryBuilder {
	rb.v.Attributes = value
	return rb
}

func (rb *MemoryBuilder) EphemeralId(ephemeralid Id) *MemoryBuilder {
	rb.v.EphemeralId = ephemeralid
	return rb
}

// Jvm Contains Java Virtual Machine (JVM) statistics for the node.

func (rb *MemoryBuilder) Jvm(jvm *JvmStatsBuilder) *MemoryBuilder {
	v := jvm.Build()
	rb.v.Jvm = v
	return rb
}

// Mem Contains statistics about memory usage for the node.

func (rb *MemoryBuilder) Mem(mem *MemStatsBuilder) *MemoryBuilder {
	v := mem.Build()
	rb.v.Mem = v
	return rb
}

// Name Human-readable identifier for the node. Based on the Node name setting
// setting.

func (rb *MemoryBuilder) Name(name Name) *MemoryBuilder {
	rb.v.Name = name
	return rb
}

// Roles Roles assigned to the node.

func (rb *MemoryBuilder) Roles(roles ...string) *MemoryBuilder {
	rb.v.Roles = roles
	return rb
}

// TransportAddress The host and port where transport HTTP connections are accepted.

func (rb *MemoryBuilder) TransportAddress(transportaddress TransportAddress) *MemoryBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}
