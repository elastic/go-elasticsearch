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

// OperatingSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L367-L373
type OperatingSystem struct {
	Cgroup    *Cgroup              `json:"cgroup,omitempty"`
	Cpu       *Cpu                 `json:"cpu,omitempty"`
	Mem       *ExtendedMemoryStats `json:"mem,omitempty"`
	Swap      *MemoryStats         `json:"swap,omitempty"`
	Timestamp *int64               `json:"timestamp,omitempty"`
}

// OperatingSystemBuilder holds OperatingSystem struct and provides a builder API.
type OperatingSystemBuilder struct {
	v *OperatingSystem
}

// NewOperatingSystem provides a builder for the OperatingSystem struct.
func NewOperatingSystemBuilder() *OperatingSystemBuilder {
	r := OperatingSystemBuilder{
		&OperatingSystem{},
	}

	return &r
}

// Build finalize the chain and returns the OperatingSystem struct
func (rb *OperatingSystemBuilder) Build() OperatingSystem {
	return *rb.v
}

func (rb *OperatingSystemBuilder) Cgroup(cgroup *CgroupBuilder) *OperatingSystemBuilder {
	v := cgroup.Build()
	rb.v.Cgroup = &v
	return rb
}

func (rb *OperatingSystemBuilder) Cpu(cpu *CpuBuilder) *OperatingSystemBuilder {
	v := cpu.Build()
	rb.v.Cpu = &v
	return rb
}

func (rb *OperatingSystemBuilder) Mem(mem *ExtendedMemoryStatsBuilder) *OperatingSystemBuilder {
	v := mem.Build()
	rb.v.Mem = &v
	return rb
}

func (rb *OperatingSystemBuilder) Swap(swap *MemoryStatsBuilder) *OperatingSystemBuilder {
	v := swap.Build()
	rb.v.Swap = &v
	return rb
}

func (rb *OperatingSystemBuilder) Timestamp(timestamp int64) *OperatingSystemBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}
