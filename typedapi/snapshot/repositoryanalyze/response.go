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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package repositoryanalyze

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package repositoryanalyze
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/snapshot/repository_analyze/SnapshotAnalyzeRepositoryResponse.ts#L24-L108
type Response struct {

	// BlobCount The number of blobs written to the repository during the test.
	BlobCount int `json:"blob_count"`
	// BlobPath The path in the repository under which all the blobs were written during the
	// test.
	BlobPath string `json:"blob_path"`
	// Concurrency The number of write operations performed concurrently during the test.
	Concurrency int `json:"concurrency"`
	// CoordinatingNode The node that coordinated the analysis and performed the final cleanup.
	CoordinatingNode types.SnapshotNodeInfo `json:"coordinating_node"`
	// DeleteElapsed The time it took to delete all the blobs in the container.
	DeleteElapsed types.Duration `json:"delete_elapsed"`
	// DeleteElapsedNanos The time it took to delete all the blobs in the container, in nanoseconds.
	DeleteElapsedNanos int64 `json:"delete_elapsed_nanos"`
	// Details A description of every read and write operation performed during the test.
	Details types.DetailsInfo `json:"details"`
	// EarlyReadNodeCount The limit on the number of nodes on which early read operations were
	// performed after writing each blob.
	EarlyReadNodeCount int `json:"early_read_node_count"`
	// IssuesDetected A list of correctness issues detected, which is empty if the API succeeded.
	// It is included to emphasize that a successful response does not guarantee
	// correct behaviour in future.
	IssuesDetected []string `json:"issues_detected"`
	// ListingElapsed The time it took to retrieve a list of all the blobs in the container.
	ListingElapsed types.Duration `json:"listing_elapsed"`
	// ListingElapsedNanos The time it took to retrieve a list of all the blobs in the container, in
	// nanoseconds.
	ListingElapsedNanos int64 `json:"listing_elapsed_nanos"`
	// MaxBlobSize The limit on the size of a blob written during the test.
	MaxBlobSize types.ByteSize `json:"max_blob_size"`
	// MaxBlobSizeBytes The limit, in bytes, on the size of a blob written during the test.
	MaxBlobSizeBytes int64 `json:"max_blob_size_bytes"`
	// MaxTotalDataSize The limit on the total size of all blob written during the test.
	MaxTotalDataSize types.ByteSize `json:"max_total_data_size"`
	// MaxTotalDataSizeBytes The limit, in bytes, on the total size of all blob written during the test.
	MaxTotalDataSizeBytes int64 `json:"max_total_data_size_bytes"`
	// RareActionProbability The probability of performing rare actions during the test.
	RareActionProbability types.Float64 `json:"rare_action_probability"`
	// ReadNodeCount The limit on the number of nodes on which read operations were performed
	// after writing each blob.
	ReadNodeCount int `json:"read_node_count"`
	// Repository The name of the repository that was the subject of the analysis.
	Repository string `json:"repository"`
	// Seed The seed for the pseudo-random number generator used to generate the
	// operations used during the test.
	Seed int64 `json:"seed"`
	// Summary A collection of statistics that summarize the results of the test.
	Summary types.SummaryInfo `json:"summary"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
