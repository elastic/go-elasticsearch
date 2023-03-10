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

// DataFrameAnalyticsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/cat/ml_data_frame_analytics/types.ts#L22-L102
type DataFrameAnalyticsRecord struct {
	// AssignmentExplanation why the job is or is not assigned to a node
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// CreateTime job creation time
	CreateTime *string `json:"create_time,omitempty"`
	// Description description
	Description *string `json:"description,omitempty"`
	// DestIndex destination index
	DestIndex *string `json:"dest_index,omitempty"`
	// FailureReason failure reason
	FailureReason *string `json:"failure_reason,omitempty"`
	// Id the id
	Id *string `json:"id,omitempty"`
	// ModelMemoryLimit model memory limit
	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`
	// NodeAddress network address of the assigned node
	NodeAddress *string `json:"node.address,omitempty"`
	// NodeEphemeralId ephemeral id of the assigned node
	NodeEphemeralId *string `json:"node.ephemeral_id,omitempty"`
	// NodeId id of the assigned node
	NodeId *string `json:"node.id,omitempty"`
	// NodeName name of the assigned node
	NodeName *string `json:"node.name,omitempty"`
	// Progress progress
	Progress *string `json:"progress,omitempty"`
	// SourceIndex source index
	SourceIndex *string `json:"source_index,omitempty"`
	// State job state
	State *string `json:"state,omitempty"`
	// Type analysis type
	Type *string `json:"type,omitempty"`
	// Version the version of Elasticsearch when the analytics was created
	Version *string `json:"version,omitempty"`
}

// NewDataFrameAnalyticsRecord returns a DataFrameAnalyticsRecord.
func NewDataFrameAnalyticsRecord() *DataFrameAnalyticsRecord {
	r := &DataFrameAnalyticsRecord{}

	return r
}
