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

// SnapshotResponseItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/snapshot/get/SnapshotGetResponse.ts#L42-L46
type SnapshotResponseItem struct {
	Error      *ErrorCause    `json:"error,omitempty"`
	Repository Name           `json:"repository"`
	Snapshots  []SnapshotInfo `json:"snapshots,omitempty"`
}

// SnapshotResponseItemBuilder holds SnapshotResponseItem struct and provides a builder API.
type SnapshotResponseItemBuilder struct {
	v *SnapshotResponseItem
}

// NewSnapshotResponseItem provides a builder for the SnapshotResponseItem struct.
func NewSnapshotResponseItemBuilder() *SnapshotResponseItemBuilder {
	r := SnapshotResponseItemBuilder{
		&SnapshotResponseItem{},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotResponseItem struct
func (rb *SnapshotResponseItemBuilder) Build() SnapshotResponseItem {
	return *rb.v
}

func (rb *SnapshotResponseItemBuilder) Error(error *ErrorCauseBuilder) *SnapshotResponseItemBuilder {
	v := error.Build()
	rb.v.Error = &v
	return rb
}

func (rb *SnapshotResponseItemBuilder) Repository(repository Name) *SnapshotResponseItemBuilder {
	rb.v.Repository = repository
	return rb
}

func (rb *SnapshotResponseItemBuilder) Snapshots(snapshots []SnapshotInfoBuilder) *SnapshotResponseItemBuilder {
	tmp := make([]SnapshotInfo, len(snapshots))
	for _, value := range snapshots {
		tmp = append(tmp, value.Build())
	}
	rb.v.Snapshots = tmp
	return rb
}
