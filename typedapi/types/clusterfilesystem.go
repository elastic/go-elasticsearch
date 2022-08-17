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

// ClusterFileSystem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L34-L38
type ClusterFileSystem struct {
	AvailableInBytes int64 `json:"available_in_bytes"`
	FreeInBytes      int64 `json:"free_in_bytes"`
	TotalInBytes     int64 `json:"total_in_bytes"`
}

// ClusterFileSystemBuilder holds ClusterFileSystem struct and provides a builder API.
type ClusterFileSystemBuilder struct {
	v *ClusterFileSystem
}

// NewClusterFileSystem provides a builder for the ClusterFileSystem struct.
func NewClusterFileSystemBuilder() *ClusterFileSystemBuilder {
	r := ClusterFileSystemBuilder{
		&ClusterFileSystem{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterFileSystem struct
func (rb *ClusterFileSystemBuilder) Build() ClusterFileSystem {
	return *rb.v
}

func (rb *ClusterFileSystemBuilder) AvailableInBytes(availableinbytes int64) *ClusterFileSystemBuilder {
	rb.v.AvailableInBytes = availableinbytes
	return rb
}

func (rb *ClusterFileSystemBuilder) FreeInBytes(freeinbytes int64) *ClusterFileSystemBuilder {
	rb.v.FreeInBytes = freeinbytes
	return rb
}

func (rb *ClusterFileSystemBuilder) TotalInBytes(totalinbytes int64) *ClusterFileSystemBuilder {
	rb.v.TotalInBytes = totalinbytes
	return rb
}
