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

package types

// IlmActions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ilm/_types/Phase.ts#L39-L93
type IlmActions struct {
	// Allocate Phases allowed: warm, cold.
	Allocate *AllocateAction `json:"allocate,omitempty"`
	// Delete Phases allowed: delete.
	Delete *DeleteAction `json:"delete,omitempty"`
	// Downsample Phases allowed: hot, warm, cold.
	Downsample *DownsampleAction `json:"downsample,omitempty"`
	// Forcemerge Phases allowed: hot, warm.
	Forcemerge *ForceMergeAction `json:"forcemerge,omitempty"`
	// Freeze The freeze action is a noop in 8.x
	Freeze *EmptyObject `json:"freeze,omitempty"`
	// Migrate Phases allowed: warm, cold.
	Migrate *MigrateAction `json:"migrate,omitempty"`
	// Readonly Phases allowed: hot, warm, cold.
	Readonly *EmptyObject `json:"readonly,omitempty"`
	// Rollover Phases allowed: hot.
	Rollover *RolloverAction `json:"rollover,omitempty"`
	// SearchableSnapshot Phases allowed: hot, cold, frozen.
	SearchableSnapshot *SearchableSnapshotAction `json:"searchable_snapshot,omitempty"`
	// SetPriority Phases allowed: hot, warm, cold.
	SetPriority *SetPriorityAction `json:"set_priority,omitempty"`
	// Shrink Phases allowed: hot, warm.
	Shrink *ShrinkAction `json:"shrink,omitempty"`
	// Unfollow Phases allowed: hot, warm, cold, frozen.
	Unfollow *EmptyObject `json:"unfollow,omitempty"`
	// WaitForSnapshot Phases allowed: delete.
	WaitForSnapshot *WaitForSnapshotAction `json:"wait_for_snapshot,omitempty"`
}

// NewIlmActions returns a IlmActions.
func NewIlmActions() *IlmActions {
	r := &IlmActions{}

	return r
}
