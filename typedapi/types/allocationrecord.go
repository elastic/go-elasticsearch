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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// AllocationRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/cat/allocation/types.ts#L24-L69
type AllocationRecord struct {
	// DiskAvail disk available
	DiskAvail ByteSize `json:"disk.avail,omitempty"`
	// DiskIndices disk used by ES indices
	DiskIndices ByteSize `json:"disk.indices,omitempty"`
	// DiskPercent percent disk used
	DiskPercent Percentage `json:"disk.percent,omitempty"`
	// DiskTotal total capacity of all volumes
	DiskTotal ByteSize `json:"disk.total,omitempty"`
	// DiskUsed disk used (total, not just ES)
	DiskUsed ByteSize `json:"disk.used,omitempty"`
	// Host host of node
	Host string `json:"host,omitempty"`
	// Ip ip of node
	Ip string `json:"ip,omitempty"`
	// Node name of node
	Node *string `json:"node,omitempty"`
	// Shards number of shards on node
	Shards *string `json:"shards,omitempty"`
}

// NewAllocationRecord returns a AllocationRecord.
func NewAllocationRecord() *AllocationRecord {
	r := &AllocationRecord{}

	return r
}
