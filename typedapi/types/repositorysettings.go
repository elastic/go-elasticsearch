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

// RepositorySettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotRepository.ts#L29-L38
type RepositorySettings struct {
	ChunkSize         *string `json:"chunk_size,omitempty"`
	Compress          string  `json:"compress,omitempty"`
	ConcurrentStreams string  `json:"concurrent_streams,omitempty"`
	Location          string  `json:"location"`
	ReadOnly          string  `json:"read_only,omitempty"`
}

// RepositorySettingsBuilder holds RepositorySettings struct and provides a builder API.
type RepositorySettingsBuilder struct {
	v *RepositorySettings
}

// NewRepositorySettings provides a builder for the RepositorySettings struct.
func NewRepositorySettingsBuilder() *RepositorySettingsBuilder {
	r := RepositorySettingsBuilder{
		&RepositorySettings{},
	}

	return &r
}

// Build finalize the chain and returns the RepositorySettings struct
func (rb *RepositorySettingsBuilder) Build() RepositorySettings {
	return *rb.v
}

func (rb *RepositorySettingsBuilder) ChunkSize(chunksize string) *RepositorySettingsBuilder {
	rb.v.ChunkSize = &chunksize
	return rb
}

func (rb *RepositorySettingsBuilder) Compress(arg string) *RepositorySettingsBuilder {
	rb.v.Compress = arg
	return rb
}

func (rb *RepositorySettingsBuilder) ConcurrentStreams(arg string) *RepositorySettingsBuilder {
	rb.v.ConcurrentStreams = arg
	return rb
}

func (rb *RepositorySettingsBuilder) Location(location string) *RepositorySettingsBuilder {
	rb.v.Location = location
	return rb
}

func (rb *RepositorySettingsBuilder) ReadOnly(arg string) *RepositorySettingsBuilder {
	rb.v.ReadOnly = arg
	return rb
}
