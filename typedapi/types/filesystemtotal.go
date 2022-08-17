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

// FileSystemTotal type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L301-L308
type FileSystemTotal struct {
	Available        *string `json:"available,omitempty"`
	AvailableInBytes *int64  `json:"available_in_bytes,omitempty"`
	Free             *string `json:"free,omitempty"`
	FreeInBytes      *int64  `json:"free_in_bytes,omitempty"`
	Total            *string `json:"total,omitempty"`
	TotalInBytes     *int64  `json:"total_in_bytes,omitempty"`
}

// FileSystemTotalBuilder holds FileSystemTotal struct and provides a builder API.
type FileSystemTotalBuilder struct {
	v *FileSystemTotal
}

// NewFileSystemTotal provides a builder for the FileSystemTotal struct.
func NewFileSystemTotalBuilder() *FileSystemTotalBuilder {
	r := FileSystemTotalBuilder{
		&FileSystemTotal{},
	}

	return &r
}

// Build finalize the chain and returns the FileSystemTotal struct
func (rb *FileSystemTotalBuilder) Build() FileSystemTotal {
	return *rb.v
}

func (rb *FileSystemTotalBuilder) Available(available string) *FileSystemTotalBuilder {
	rb.v.Available = &available
	return rb
}

func (rb *FileSystemTotalBuilder) AvailableInBytes(availableinbytes int64) *FileSystemTotalBuilder {
	rb.v.AvailableInBytes = &availableinbytes
	return rb
}

func (rb *FileSystemTotalBuilder) Free(free string) *FileSystemTotalBuilder {
	rb.v.Free = &free
	return rb
}

func (rb *FileSystemTotalBuilder) FreeInBytes(freeinbytes int64) *FileSystemTotalBuilder {
	rb.v.FreeInBytes = &freeinbytes
	return rb
}

func (rb *FileSystemTotalBuilder) Total(total string) *FileSystemTotalBuilder {
	rb.v.Total = &total
	return rb
}

func (rb *FileSystemTotalBuilder) TotalInBytes(totalinbytes int64) *FileSystemTotalBuilder {
	rb.v.TotalInBytes = &totalinbytes
	return rb
}
