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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

// IndicesModifyAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/modify_data_stream/types.ts#L22-L37
type IndicesModifyAction struct {
	// AddBackingIndex Adds an existing index as a backing index for a data stream.
	// The index is hidden as part of this operation.
	// WARNING: Adding indices with the `add_backing_index` action can potentially
	// result in improper data stream behavior.
	// This should be considered an expert level API.
	AddBackingIndex *IndexAndDataStreamAction `json:"add_backing_index,omitempty"`
	// RemoveBackingIndex Removes a backing index from a data stream.
	// The index is unhidden as part of this operation.
	// A data stream’s write index cannot be removed.
	RemoveBackingIndex *IndexAndDataStreamAction `json:"remove_backing_index,omitempty"`
}

// NewIndicesModifyAction returns a IndicesModifyAction.
func NewIndicesModifyAction() *IndicesModifyAction {
	r := &IndicesModifyAction{}

	return r
}
