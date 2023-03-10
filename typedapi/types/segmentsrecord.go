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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// SegmentsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/cat/segments/types.ts#L22-L96
type SegmentsRecord struct {
	// Committed is segment committed
	Committed *string `json:"committed,omitempty"`
	// Compound is segment compound
	Compound *string `json:"compound,omitempty"`
	// DocsCount number of docs in segment
	DocsCount *string `json:"docs.count,omitempty"`
	// DocsDeleted number of deleted docs in segment
	DocsDeleted *string `json:"docs.deleted,omitempty"`
	// Generation segment generation
	Generation *string `json:"generation,omitempty"`
	// Id unique id of node where it lives
	Id *string `json:"id,omitempty"`
	// Index index name
	Index *string `json:"index,omitempty"`
	// Ip ip of node where it lives
	Ip *string `json:"ip,omitempty"`
	// Prirep primary or replica
	Prirep *string `json:"prirep,omitempty"`
	// Searchable is segment searched
	Searchable *string `json:"searchable,omitempty"`
	// Segment segment name
	Segment *string `json:"segment,omitempty"`
	// Shard shard name
	Shard *string `json:"shard,omitempty"`
	// Size segment size in bytes
	Size ByteSize `json:"size,omitempty"`
	// SizeMemory segment memory in bytes
	SizeMemory ByteSize `json:"size.memory,omitempty"`
	// Version version
	Version *string `json:"version,omitempty"`
}

// NewSegmentsRecord returns a SegmentsRecord.
func NewSegmentsRecord() *SegmentsRecord {
	r := &SegmentsRecord{}

	return r
}
