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

// TranslogStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Stats.ts#L242-L250
type TranslogStats struct {
	EarliestLastModifiedAge int64   `json:"earliest_last_modified_age"`
	Operations              int64   `json:"operations"`
	Size                    *string `json:"size,omitempty"`
	SizeInBytes             int64   `json:"size_in_bytes"`
	UncommittedOperations   int     `json:"uncommitted_operations"`
	UncommittedSize         *string `json:"uncommitted_size,omitempty"`
	UncommittedSizeInBytes  int64   `json:"uncommitted_size_in_bytes"`
}

// TranslogStatsBuilder holds TranslogStats struct and provides a builder API.
type TranslogStatsBuilder struct {
	v *TranslogStats
}

// NewTranslogStats provides a builder for the TranslogStats struct.
func NewTranslogStatsBuilder() *TranslogStatsBuilder {
	r := TranslogStatsBuilder{
		&TranslogStats{},
	}

	return &r
}

// Build finalize the chain and returns the TranslogStats struct
func (rb *TranslogStatsBuilder) Build() TranslogStats {
	return *rb.v
}

func (rb *TranslogStatsBuilder) EarliestLastModifiedAge(earliestlastmodifiedage int64) *TranslogStatsBuilder {
	rb.v.EarliestLastModifiedAge = earliestlastmodifiedage
	return rb
}

func (rb *TranslogStatsBuilder) Operations(operations int64) *TranslogStatsBuilder {
	rb.v.Operations = operations
	return rb
}

func (rb *TranslogStatsBuilder) Size(size string) *TranslogStatsBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *TranslogStatsBuilder) SizeInBytes(sizeinbytes int64) *TranslogStatsBuilder {
	rb.v.SizeInBytes = sizeinbytes
	return rb
}

func (rb *TranslogStatsBuilder) UncommittedOperations(uncommittedoperations int) *TranslogStatsBuilder {
	rb.v.UncommittedOperations = uncommittedoperations
	return rb
}

func (rb *TranslogStatsBuilder) UncommittedSize(uncommittedsize string) *TranslogStatsBuilder {
	rb.v.UncommittedSize = &uncommittedsize
	return rb
}

func (rb *TranslogStatsBuilder) UncommittedSizeInBytes(uncommittedsizeinbytes int64) *TranslogStatsBuilder {
	rb.v.UncommittedSizeInBytes = uncommittedsizeinbytes
	return rb
}
