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
// https://github.com/elastic/elasticsearch-specification/tree/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67

package types

// Cgroup type.
//
// https://github.com/elastic/elasticsearch-specification/blob/b7d4fb5356784b8bcde8d3a2d62a1fd5621ffd67/specification/nodes/_types/Stats.ts#L461-L474
type Cgroup struct {
	// Cpu Contains statistics about `cpu` control group for the node.
	Cpu *CgroupCpu `json:"cpu,omitempty"`
	// Cpuacct Contains statistics about `cpuacct` control group for the node.
	Cpuacct *CpuAcct `json:"cpuacct,omitempty"`
	// Memory Contains statistics about the memory control group for the node.
	Memory *CgroupMemory `json:"memory,omitempty"`
}

// NewCgroup returns a Cgroup.
func NewCgroup() *Cgroup {
	r := &Cgroup{}

	return r
}
