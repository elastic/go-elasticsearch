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

// CpuAcct type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L188-L191
type CpuAcct struct {
	ControlGroup *string                 `json:"control_group,omitempty"`
	UsageNanos   *DurationValueUnitNanos `json:"usage_nanos,omitempty"`
}

// CpuAcctBuilder holds CpuAcct struct and provides a builder API.
type CpuAcctBuilder struct {
	v *CpuAcct
}

// NewCpuAcct provides a builder for the CpuAcct struct.
func NewCpuAcctBuilder() *CpuAcctBuilder {
	r := CpuAcctBuilder{
		&CpuAcct{},
	}

	return &r
}

// Build finalize the chain and returns the CpuAcct struct
func (rb *CpuAcctBuilder) Build() CpuAcct {
	return *rb.v
}

func (rb *CpuAcctBuilder) ControlGroup(controlgroup string) *CpuAcctBuilder {
	rb.v.ControlGroup = &controlgroup
	return rb
}

func (rb *CpuAcctBuilder) UsageNanos(usagenanos *DurationValueUnitNanosBuilder) *CpuAcctBuilder {
	v := usagenanos.Build()
	rb.v.UsageNanos = &v
	return rb
}
