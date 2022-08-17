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

// ExtendedMemoryStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L255-L258
type ExtendedMemoryStats struct {
	AdjustedTotalInBytes *int64  `json:"adjusted_total_in_bytes,omitempty"`
	FreeInBytes          *int64  `json:"free_in_bytes,omitempty"`
	FreePercent          *int    `json:"free_percent,omitempty"`
	Resident             *string `json:"resident,omitempty"`
	ResidentInBytes      *int64  `json:"resident_in_bytes,omitempty"`
	Share                *string `json:"share,omitempty"`
	ShareInBytes         *int64  `json:"share_in_bytes,omitempty"`
	TotalInBytes         *int64  `json:"total_in_bytes,omitempty"`
	TotalVirtual         *string `json:"total_virtual,omitempty"`
	TotalVirtualInBytes  *int64  `json:"total_virtual_in_bytes,omitempty"`
	UsedInBytes          *int64  `json:"used_in_bytes,omitempty"`
	UsedPercent          *int    `json:"used_percent,omitempty"`
}

// ExtendedMemoryStatsBuilder holds ExtendedMemoryStats struct and provides a builder API.
type ExtendedMemoryStatsBuilder struct {
	v *ExtendedMemoryStats
}

// NewExtendedMemoryStats provides a builder for the ExtendedMemoryStats struct.
func NewExtendedMemoryStatsBuilder() *ExtendedMemoryStatsBuilder {
	r := ExtendedMemoryStatsBuilder{
		&ExtendedMemoryStats{},
	}

	return &r
}

// Build finalize the chain and returns the ExtendedMemoryStats struct
func (rb *ExtendedMemoryStatsBuilder) Build() ExtendedMemoryStats {
	return *rb.v
}

func (rb *ExtendedMemoryStatsBuilder) AdjustedTotalInBytes(adjustedtotalinbytes int64) *ExtendedMemoryStatsBuilder {
	rb.v.AdjustedTotalInBytes = &adjustedtotalinbytes
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) FreeInBytes(freeinbytes int64) *ExtendedMemoryStatsBuilder {
	rb.v.FreeInBytes = &freeinbytes
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) FreePercent(freepercent int) *ExtendedMemoryStatsBuilder {
	rb.v.FreePercent = &freepercent
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) Resident(resident string) *ExtendedMemoryStatsBuilder {
	rb.v.Resident = &resident
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) ResidentInBytes(residentinbytes int64) *ExtendedMemoryStatsBuilder {
	rb.v.ResidentInBytes = &residentinbytes
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) Share(share string) *ExtendedMemoryStatsBuilder {
	rb.v.Share = &share
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) ShareInBytes(shareinbytes int64) *ExtendedMemoryStatsBuilder {
	rb.v.ShareInBytes = &shareinbytes
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) TotalInBytes(totalinbytes int64) *ExtendedMemoryStatsBuilder {
	rb.v.TotalInBytes = &totalinbytes
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) TotalVirtual(totalvirtual string) *ExtendedMemoryStatsBuilder {
	rb.v.TotalVirtual = &totalvirtual
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) TotalVirtualInBytes(totalvirtualinbytes int64) *ExtendedMemoryStatsBuilder {
	rb.v.TotalVirtualInBytes = &totalvirtualinbytes
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) UsedInBytes(usedinbytes int64) *ExtendedMemoryStatsBuilder {
	rb.v.UsedInBytes = &usedinbytes
	return rb
}

func (rb *ExtendedMemoryStatsBuilder) UsedPercent(usedpercent int) *ExtendedMemoryStatsBuilder {
	rb.v.UsedPercent = &usedpercent
	return rb
}
