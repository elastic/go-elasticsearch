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

// OperatingSystemMemoryInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L282-L290
type OperatingSystemMemoryInfo struct {
	AdjustedTotalInBytes *int64 `json:"adjusted_total_in_bytes,omitempty"`
	FreeInBytes          int64  `json:"free_in_bytes"`
	FreePercent          int    `json:"free_percent"`
	TotalInBytes         int64  `json:"total_in_bytes"`
	UsedInBytes          int64  `json:"used_in_bytes"`
	UsedPercent          int    `json:"used_percent"`
}

// OperatingSystemMemoryInfoBuilder holds OperatingSystemMemoryInfo struct and provides a builder API.
type OperatingSystemMemoryInfoBuilder struct {
	v *OperatingSystemMemoryInfo
}

// NewOperatingSystemMemoryInfo provides a builder for the OperatingSystemMemoryInfo struct.
func NewOperatingSystemMemoryInfoBuilder() *OperatingSystemMemoryInfoBuilder {
	r := OperatingSystemMemoryInfoBuilder{
		&OperatingSystemMemoryInfo{},
	}

	return &r
}

// Build finalize the chain and returns the OperatingSystemMemoryInfo struct
func (rb *OperatingSystemMemoryInfoBuilder) Build() OperatingSystemMemoryInfo {
	return *rb.v
}

func (rb *OperatingSystemMemoryInfoBuilder) AdjustedTotalInBytes(adjustedtotalinbytes int64) *OperatingSystemMemoryInfoBuilder {
	rb.v.AdjustedTotalInBytes = &adjustedtotalinbytes
	return rb
}

func (rb *OperatingSystemMemoryInfoBuilder) FreeInBytes(freeinbytes int64) *OperatingSystemMemoryInfoBuilder {
	rb.v.FreeInBytes = freeinbytes
	return rb
}

func (rb *OperatingSystemMemoryInfoBuilder) FreePercent(freepercent int) *OperatingSystemMemoryInfoBuilder {
	rb.v.FreePercent = freepercent
	return rb
}

func (rb *OperatingSystemMemoryInfoBuilder) TotalInBytes(totalinbytes int64) *OperatingSystemMemoryInfoBuilder {
	rb.v.TotalInBytes = totalinbytes
	return rb
}

func (rb *OperatingSystemMemoryInfoBuilder) UsedInBytes(usedinbytes int64) *OperatingSystemMemoryInfoBuilder {
	rb.v.UsedInBytes = usedinbytes
	return rb
}

func (rb *OperatingSystemMemoryInfoBuilder) UsedPercent(usedpercent int) *OperatingSystemMemoryInfoBuilder {
	rb.v.UsedPercent = usedpercent
	return rb
}
