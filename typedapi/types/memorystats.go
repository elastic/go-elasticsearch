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

// MemoryStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L242-L253
type MemoryStats struct {
	AdjustedTotalInBytes *int64  `json:"adjusted_total_in_bytes,omitempty"`
	FreeInBytes          *int64  `json:"free_in_bytes,omitempty"`
	Resident             *string `json:"resident,omitempty"`
	ResidentInBytes      *int64  `json:"resident_in_bytes,omitempty"`
	Share                *string `json:"share,omitempty"`
	ShareInBytes         *int64  `json:"share_in_bytes,omitempty"`
	TotalInBytes         *int64  `json:"total_in_bytes,omitempty"`
	TotalVirtual         *string `json:"total_virtual,omitempty"`
	TotalVirtualInBytes  *int64  `json:"total_virtual_in_bytes,omitempty"`
	UsedInBytes          *int64  `json:"used_in_bytes,omitempty"`
}

// MemoryStatsBuilder holds MemoryStats struct and provides a builder API.
type MemoryStatsBuilder struct {
	v *MemoryStats
}

// NewMemoryStats provides a builder for the MemoryStats struct.
func NewMemoryStatsBuilder() *MemoryStatsBuilder {
	r := MemoryStatsBuilder{
		&MemoryStats{},
	}

	return &r
}

// Build finalize the chain and returns the MemoryStats struct
func (rb *MemoryStatsBuilder) Build() MemoryStats {
	return *rb.v
}

func (rb *MemoryStatsBuilder) AdjustedTotalInBytes(adjustedtotalinbytes int64) *MemoryStatsBuilder {
	rb.v.AdjustedTotalInBytes = &adjustedtotalinbytes
	return rb
}

func (rb *MemoryStatsBuilder) FreeInBytes(freeinbytes int64) *MemoryStatsBuilder {
	rb.v.FreeInBytes = &freeinbytes
	return rb
}

func (rb *MemoryStatsBuilder) Resident(resident string) *MemoryStatsBuilder {
	rb.v.Resident = &resident
	return rb
}

func (rb *MemoryStatsBuilder) ResidentInBytes(residentinbytes int64) *MemoryStatsBuilder {
	rb.v.ResidentInBytes = &residentinbytes
	return rb
}

func (rb *MemoryStatsBuilder) Share(share string) *MemoryStatsBuilder {
	rb.v.Share = &share
	return rb
}

func (rb *MemoryStatsBuilder) ShareInBytes(shareinbytes int64) *MemoryStatsBuilder {
	rb.v.ShareInBytes = &shareinbytes
	return rb
}

func (rb *MemoryStatsBuilder) TotalInBytes(totalinbytes int64) *MemoryStatsBuilder {
	rb.v.TotalInBytes = &totalinbytes
	return rb
}

func (rb *MemoryStatsBuilder) TotalVirtual(totalvirtual string) *MemoryStatsBuilder {
	rb.v.TotalVirtual = &totalvirtual
	return rb
}

func (rb *MemoryStatsBuilder) TotalVirtualInBytes(totalvirtualinbytes int64) *MemoryStatsBuilder {
	rb.v.TotalVirtualInBytes = &totalvirtualinbytes
	return rb
}

func (rb *MemoryStatsBuilder) UsedInBytes(usedinbytes int64) *MemoryStatsBuilder {
	rb.v.UsedInBytes = &usedinbytes
	return rb
}
