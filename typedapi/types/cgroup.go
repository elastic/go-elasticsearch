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

// Cgroup type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L182-L186
type Cgroup struct {
	Cpu     *CgroupCpu    `json:"cpu,omitempty"`
	Cpuacct *CpuAcct      `json:"cpuacct,omitempty"`
	Memory  *CgroupMemory `json:"memory,omitempty"`
}

// CgroupBuilder holds Cgroup struct and provides a builder API.
type CgroupBuilder struct {
	v *Cgroup
}

// NewCgroup provides a builder for the Cgroup struct.
func NewCgroupBuilder() *CgroupBuilder {
	r := CgroupBuilder{
		&Cgroup{},
	}

	return &r
}

// Build finalize the chain and returns the Cgroup struct
func (rb *CgroupBuilder) Build() Cgroup {
	return *rb.v
}

func (rb *CgroupBuilder) Cpu(cpu *CgroupCpuBuilder) *CgroupBuilder {
	v := cpu.Build()
	rb.v.Cpu = &v
	return rb
}

func (rb *CgroupBuilder) Cpuacct(cpuacct *CpuAcctBuilder) *CgroupBuilder {
	v := cpuacct.Build()
	rb.v.Cpuacct = &v
	return rb
}

func (rb *CgroupBuilder) Memory(memory *CgroupMemoryBuilder) *CgroupBuilder {
	v := memory.Build()
	rb.v.Memory = &v
	return rb
}
