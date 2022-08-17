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

// AllocationRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/allocation/types.ts#L24-L69
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
	Host Host `json:"host,omitempty"`
	// Ip ip of node
	Ip Ip `json:"ip,omitempty"`
	// Node name of node
	Node *string `json:"node,omitempty"`
	// Shards number of shards on node
	Shards *string `json:"shards,omitempty"`
}

// AllocationRecordBuilder holds AllocationRecord struct and provides a builder API.
type AllocationRecordBuilder struct {
	v *AllocationRecord
}

// NewAllocationRecord provides a builder for the AllocationRecord struct.
func NewAllocationRecordBuilder() *AllocationRecordBuilder {
	r := AllocationRecordBuilder{
		&AllocationRecord{},
	}

	return &r
}

// Build finalize the chain and returns the AllocationRecord struct
func (rb *AllocationRecordBuilder) Build() AllocationRecord {
	return *rb.v
}

// DiskAvail disk available

func (rb *AllocationRecordBuilder) DiskAvail(diskavail ByteSize) *AllocationRecordBuilder {
	rb.v.DiskAvail = diskavail
	return rb
}

// DiskIndices disk used by ES indices

func (rb *AllocationRecordBuilder) DiskIndices(diskindices ByteSize) *AllocationRecordBuilder {
	rb.v.DiskIndices = diskindices
	return rb
}

// DiskPercent percent disk used

func (rb *AllocationRecordBuilder) DiskPercent(diskpercent Percentage) *AllocationRecordBuilder {
	rb.v.DiskPercent = diskpercent
	return rb
}

// DiskTotal total capacity of all volumes

func (rb *AllocationRecordBuilder) DiskTotal(disktotal ByteSize) *AllocationRecordBuilder {
	rb.v.DiskTotal = disktotal
	return rb
}

// DiskUsed disk used (total, not just ES)

func (rb *AllocationRecordBuilder) DiskUsed(diskused ByteSize) *AllocationRecordBuilder {
	rb.v.DiskUsed = diskused
	return rb
}

// Host host of node

func (rb *AllocationRecordBuilder) Host(host Host) *AllocationRecordBuilder {
	rb.v.Host = host
	return rb
}

// Ip ip of node

func (rb *AllocationRecordBuilder) Ip(ip Ip) *AllocationRecordBuilder {
	rb.v.Ip = ip
	return rb
}

// Node name of node

func (rb *AllocationRecordBuilder) Node(node string) *AllocationRecordBuilder {
	rb.v.Node = &node
	return rb
}

// Shards number of shards on node

func (rb *AllocationRecordBuilder) Shards(shards string) *AllocationRecordBuilder {
	rb.v.Shards = &shards
	return rb
}
