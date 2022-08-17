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

// ShardsStatsSummaryItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/_types/SnapshotShardsStatus.ts#L37-L40
type ShardsStatsSummaryItem struct {
	FileCount   int64 `json:"file_count"`
	SizeInBytes int64 `json:"size_in_bytes"`
}

// ShardsStatsSummaryItemBuilder holds ShardsStatsSummaryItem struct and provides a builder API.
type ShardsStatsSummaryItemBuilder struct {
	v *ShardsStatsSummaryItem
}

// NewShardsStatsSummaryItem provides a builder for the ShardsStatsSummaryItem struct.
func NewShardsStatsSummaryItemBuilder() *ShardsStatsSummaryItemBuilder {
	r := ShardsStatsSummaryItemBuilder{
		&ShardsStatsSummaryItem{},
	}

	return &r
}

// Build finalize the chain and returns the ShardsStatsSummaryItem struct
func (rb *ShardsStatsSummaryItemBuilder) Build() ShardsStatsSummaryItem {
	return *rb.v
}

func (rb *ShardsStatsSummaryItemBuilder) FileCount(filecount int64) *ShardsStatsSummaryItemBuilder {
	rb.v.FileCount = filecount
	return rb
}

func (rb *ShardsStatsSummaryItemBuilder) SizeInBytes(sizeinbytes int64) *ShardsStatsSummaryItemBuilder {
	rb.v.SizeInBytes = sizeinbytes
	return rb
}
