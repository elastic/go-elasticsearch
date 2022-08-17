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

// CgroupCpu type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L193-L198
type CgroupCpu struct {
	CfsPeriodMicros *int           `json:"cfs_period_micros,omitempty"`
	CfsQuotaMicros  *int           `json:"cfs_quota_micros,omitempty"`
	ControlGroup    *string        `json:"control_group,omitempty"`
	Stat            *CgroupCpuStat `json:"stat,omitempty"`
}

// CgroupCpuBuilder holds CgroupCpu struct and provides a builder API.
type CgroupCpuBuilder struct {
	v *CgroupCpu
}

// NewCgroupCpu provides a builder for the CgroupCpu struct.
func NewCgroupCpuBuilder() *CgroupCpuBuilder {
	r := CgroupCpuBuilder{
		&CgroupCpu{},
	}

	return &r
}

// Build finalize the chain and returns the CgroupCpu struct
func (rb *CgroupCpuBuilder) Build() CgroupCpu {
	return *rb.v
}

func (rb *CgroupCpuBuilder) CfsPeriodMicros(cfsperiodmicros int) *CgroupCpuBuilder {
	rb.v.CfsPeriodMicros = &cfsperiodmicros
	return rb
}

func (rb *CgroupCpuBuilder) CfsQuotaMicros(cfsquotamicros int) *CgroupCpuBuilder {
	rb.v.CfsQuotaMicros = &cfsquotamicros
	return rb
}

func (rb *CgroupCpuBuilder) ControlGroup(controlgroup string) *CgroupCpuBuilder {
	rb.v.ControlGroup = &controlgroup
	return rb
}

func (rb *CgroupCpuBuilder) Stat(stat *CgroupCpuStatBuilder) *CgroupCpuBuilder {
	v := stat.Build()
	rb.v.Stat = &v
	return rb
}
