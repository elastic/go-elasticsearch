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

// DataPathStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L223-L240
type DataPathStats struct {
	Available            *string `json:"available,omitempty"`
	AvailableInBytes     *int64  `json:"available_in_bytes,omitempty"`
	DiskQueue            *string `json:"disk_queue,omitempty"`
	DiskReadSize         *string `json:"disk_read_size,omitempty"`
	DiskReadSizeInBytes  *int64  `json:"disk_read_size_in_bytes,omitempty"`
	DiskReads            *int64  `json:"disk_reads,omitempty"`
	DiskWriteSize        *string `json:"disk_write_size,omitempty"`
	DiskWriteSizeInBytes *int64  `json:"disk_write_size_in_bytes,omitempty"`
	DiskWrites           *int64  `json:"disk_writes,omitempty"`
	Free                 *string `json:"free,omitempty"`
	FreeInBytes          *int64  `json:"free_in_bytes,omitempty"`
	Mount                *string `json:"mount,omitempty"`
	Path                 *string `json:"path,omitempty"`
	Total                *string `json:"total,omitempty"`
	TotalInBytes         *int64  `json:"total_in_bytes,omitempty"`
	Type                 *string `json:"type,omitempty"`
}

// DataPathStatsBuilder holds DataPathStats struct and provides a builder API.
type DataPathStatsBuilder struct {
	v *DataPathStats
}

// NewDataPathStats provides a builder for the DataPathStats struct.
func NewDataPathStatsBuilder() *DataPathStatsBuilder {
	r := DataPathStatsBuilder{
		&DataPathStats{},
	}

	return &r
}

// Build finalize the chain and returns the DataPathStats struct
func (rb *DataPathStatsBuilder) Build() DataPathStats {
	return *rb.v
}

func (rb *DataPathStatsBuilder) Available(available string) *DataPathStatsBuilder {
	rb.v.Available = &available
	return rb
}

func (rb *DataPathStatsBuilder) AvailableInBytes(availableinbytes int64) *DataPathStatsBuilder {
	rb.v.AvailableInBytes = &availableinbytes
	return rb
}

func (rb *DataPathStatsBuilder) DiskQueue(diskqueue string) *DataPathStatsBuilder {
	rb.v.DiskQueue = &diskqueue
	return rb
}

func (rb *DataPathStatsBuilder) DiskReadSize(diskreadsize string) *DataPathStatsBuilder {
	rb.v.DiskReadSize = &diskreadsize
	return rb
}

func (rb *DataPathStatsBuilder) DiskReadSizeInBytes(diskreadsizeinbytes int64) *DataPathStatsBuilder {
	rb.v.DiskReadSizeInBytes = &diskreadsizeinbytes
	return rb
}

func (rb *DataPathStatsBuilder) DiskReads(diskreads int64) *DataPathStatsBuilder {
	rb.v.DiskReads = &diskreads
	return rb
}

func (rb *DataPathStatsBuilder) DiskWriteSize(diskwritesize string) *DataPathStatsBuilder {
	rb.v.DiskWriteSize = &diskwritesize
	return rb
}

func (rb *DataPathStatsBuilder) DiskWriteSizeInBytes(diskwritesizeinbytes int64) *DataPathStatsBuilder {
	rb.v.DiskWriteSizeInBytes = &diskwritesizeinbytes
	return rb
}

func (rb *DataPathStatsBuilder) DiskWrites(diskwrites int64) *DataPathStatsBuilder {
	rb.v.DiskWrites = &diskwrites
	return rb
}

func (rb *DataPathStatsBuilder) Free(free string) *DataPathStatsBuilder {
	rb.v.Free = &free
	return rb
}

func (rb *DataPathStatsBuilder) FreeInBytes(freeinbytes int64) *DataPathStatsBuilder {
	rb.v.FreeInBytes = &freeinbytes
	return rb
}

func (rb *DataPathStatsBuilder) Mount(mount string) *DataPathStatsBuilder {
	rb.v.Mount = &mount
	return rb
}

func (rb *DataPathStatsBuilder) Path(path string) *DataPathStatsBuilder {
	rb.v.Path = &path
	return rb
}

func (rb *DataPathStatsBuilder) Total(total string) *DataPathStatsBuilder {
	rb.v.Total = &total
	return rb
}

func (rb *DataPathStatsBuilder) TotalInBytes(totalinbytes int64) *DataPathStatsBuilder {
	rb.v.TotalInBytes = &totalinbytes
	return rb
}

func (rb *DataPathStatsBuilder) Type_(type_ string) *DataPathStatsBuilder {
	rb.v.Type = &type_
	return rb
}
