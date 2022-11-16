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
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// IndicesIndexingPressureMemory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/indices/_types/IndexSettings.ts#L541-L548
type IndicesIndexingPressureMemory struct {
	// Limit Number of outstanding bytes that may be consumed by indexing requests. When
	// this limit is reached or exceeded,
	// the node will reject new coordinating and primary operations. When replica
	// operations consume 1.5x this limit,
	// the node will reject new replica operations. Defaults to 10% of the heap.
	Limit *int `json:"limit,omitempty"`
}

// NewIndicesIndexingPressureMemory returns a IndicesIndexingPressureMemory.
func NewIndicesIndexingPressureMemory() *IndicesIndexingPressureMemory {
	r := &IndicesIndexingPressureMemory{}

	return r
}
