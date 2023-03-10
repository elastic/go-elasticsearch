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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// MemoryStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/nodes/_types/Stats.ts#L248-L259
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

// NewMemoryStats returns a MemoryStats.
func NewMemoryStats() *MemoryStats {
	r := &MemoryStats{}

	return r
}
