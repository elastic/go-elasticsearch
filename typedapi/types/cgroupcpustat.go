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

// CgroupCpuStat type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L200-L204
type CgroupCpuStat struct {
	NumberOfElapsedPeriods *int64                  `json:"number_of_elapsed_periods,omitempty"`
	NumberOfTimesThrottled *int64                  `json:"number_of_times_throttled,omitempty"`
	TimeThrottledNanos     *DurationValueUnitNanos `json:"time_throttled_nanos,omitempty"`
}

// CgroupCpuStatBuilder holds CgroupCpuStat struct and provides a builder API.
type CgroupCpuStatBuilder struct {
	v *CgroupCpuStat
}

// NewCgroupCpuStat provides a builder for the CgroupCpuStat struct.
func NewCgroupCpuStatBuilder() *CgroupCpuStatBuilder {
	r := CgroupCpuStatBuilder{
		&CgroupCpuStat{},
	}

	return &r
}

// Build finalize the chain and returns the CgroupCpuStat struct
func (rb *CgroupCpuStatBuilder) Build() CgroupCpuStat {
	return *rb.v
}

func (rb *CgroupCpuStatBuilder) NumberOfElapsedPeriods(numberofelapsedperiods int64) *CgroupCpuStatBuilder {
	rb.v.NumberOfElapsedPeriods = &numberofelapsedperiods
	return rb
}

func (rb *CgroupCpuStatBuilder) NumberOfTimesThrottled(numberoftimesthrottled int64) *CgroupCpuStatBuilder {
	rb.v.NumberOfTimesThrottled = &numberoftimesthrottled
	return rb
}

func (rb *CgroupCpuStatBuilder) TimeThrottledNanos(timethrottlednanos *DurationValueUnitNanosBuilder) *CgroupCpuStatBuilder {
	v := timethrottlednanos.Build()
	rb.v.TimeThrottledNanos = &v
	return rb
}
