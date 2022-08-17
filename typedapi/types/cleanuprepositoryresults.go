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

// CleanupRepositoryResults type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/cleanup_repository/SnapshotCleanupRepositoryResponse.ts#L29-L34
type CleanupRepositoryResults struct {
	// DeletedBlobs Number of binary large objects (blobs) removed during cleanup.
	DeletedBlobs int64 `json:"deleted_blobs"`
	// DeletedBytes Number of bytes freed by cleanup operations.
	DeletedBytes int64 `json:"deleted_bytes"`
}

// CleanupRepositoryResultsBuilder holds CleanupRepositoryResults struct and provides a builder API.
type CleanupRepositoryResultsBuilder struct {
	v *CleanupRepositoryResults
}

// NewCleanupRepositoryResults provides a builder for the CleanupRepositoryResults struct.
func NewCleanupRepositoryResultsBuilder() *CleanupRepositoryResultsBuilder {
	r := CleanupRepositoryResultsBuilder{
		&CleanupRepositoryResults{},
	}

	return &r
}

// Build finalize the chain and returns the CleanupRepositoryResults struct
func (rb *CleanupRepositoryResultsBuilder) Build() CleanupRepositoryResults {
	return *rb.v
}

// DeletedBlobs Number of binary large objects (blobs) removed during cleanup.

func (rb *CleanupRepositoryResultsBuilder) DeletedBlobs(deletedblobs int64) *CleanupRepositoryResultsBuilder {
	rb.v.DeletedBlobs = deletedblobs
	return rb
}

// DeletedBytes Number of bytes freed by cleanup operations.

func (rb *CleanupRepositoryResultsBuilder) DeletedBytes(deletedbytes int64) *CleanupRepositoryResultsBuilder {
	rb.v.DeletedBytes = deletedbytes
	return rb
}
