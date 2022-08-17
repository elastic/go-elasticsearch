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

// Cpu type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L212-L221
type Cpu struct {
	LoadAverage   map[string]float64       `json:"load_average,omitempty"`
	Percent       *int                     `json:"percent,omitempty"`
	Sys           *Duration                `json:"sys,omitempty"`
	SysInMillis   *DurationValueUnitMillis `json:"sys_in_millis,omitempty"`
	Total         *Duration                `json:"total,omitempty"`
	TotalInMillis *DurationValueUnitMillis `json:"total_in_millis,omitempty"`
	User          *Duration                `json:"user,omitempty"`
	UserInMillis  *DurationValueUnitMillis `json:"user_in_millis,omitempty"`
}

// CpuBuilder holds Cpu struct and provides a builder API.
type CpuBuilder struct {
	v *Cpu
}

// NewCpu provides a builder for the Cpu struct.
func NewCpuBuilder() *CpuBuilder {
	r := CpuBuilder{
		&Cpu{
			LoadAverage: make(map[string]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Cpu struct
func (rb *CpuBuilder) Build() Cpu {
	return *rb.v
}

func (rb *CpuBuilder) LoadAverage(value map[string]float64) *CpuBuilder {
	rb.v.LoadAverage = value
	return rb
}

func (rb *CpuBuilder) Percent(percent int) *CpuBuilder {
	rb.v.Percent = &percent
	return rb
}

func (rb *CpuBuilder) Sys(sys *DurationBuilder) *CpuBuilder {
	v := sys.Build()
	rb.v.Sys = &v
	return rb
}

func (rb *CpuBuilder) SysInMillis(sysinmillis *DurationValueUnitMillisBuilder) *CpuBuilder {
	v := sysinmillis.Build()
	rb.v.SysInMillis = &v
	return rb
}

func (rb *CpuBuilder) Total(total *DurationBuilder) *CpuBuilder {
	v := total.Build()
	rb.v.Total = &v
	return rb
}

func (rb *CpuBuilder) TotalInMillis(totalinmillis *DurationValueUnitMillisBuilder) *CpuBuilder {
	v := totalinmillis.Build()
	rb.v.TotalInMillis = &v
	return rb
}

func (rb *CpuBuilder) User(user *DurationBuilder) *CpuBuilder {
	v := user.Build()
	rb.v.User = &v
	return rb
}

func (rb *CpuBuilder) UserInMillis(userinmillis *DurationValueUnitMillisBuilder) *CpuBuilder {
	v := userinmillis.Build()
	rb.v.UserInMillis = &v
	return rb
}
