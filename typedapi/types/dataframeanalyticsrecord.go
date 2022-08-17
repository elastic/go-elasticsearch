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

// DataFrameAnalyticsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/ml_data_frame_analytics/types.ts#L22-L102
type DataFrameAnalyticsRecord struct {
	// AssignmentExplanation why the job is or is not assigned to a node
	AssignmentExplanation *string `json:"assignment_explanation,omitempty"`
	// CreateTime job creation time
	CreateTime *string `json:"create_time,omitempty"`
	// Description description
	Description *string `json:"description,omitempty"`
	// DestIndex destination index
	DestIndex *IndexName `json:"dest_index,omitempty"`
	// FailureReason failure reason
	FailureReason *string `json:"failure_reason,omitempty"`
	// Id the id
	Id *Id `json:"id,omitempty"`
	// ModelMemoryLimit model memory limit
	ModelMemoryLimit *string `json:"model_memory_limit,omitempty"`
	// NodeAddress network address of the assigned node
	NodeAddress *string `json:"node.address,omitempty"`
	// NodeEphemeralId ephemeral id of the assigned node
	NodeEphemeralId *Id `json:"node.ephemeral_id,omitempty"`
	// NodeId id of the assigned node
	NodeId *Id `json:"node.id,omitempty"`
	// NodeName name of the assigned node
	NodeName *Name `json:"node.name,omitempty"`
	// Progress progress
	Progress *string `json:"progress,omitempty"`
	// SourceIndex source index
	SourceIndex *IndexName `json:"source_index,omitempty"`
	// State job state
	State *string `json:"state,omitempty"`
	// Type analysis type
	Type *string `json:"type,omitempty"`
	// Version the version of Elasticsearch when the analytics was created
	Version *VersionString `json:"version,omitempty"`
}

// DataFrameAnalyticsRecordBuilder holds DataFrameAnalyticsRecord struct and provides a builder API.
type DataFrameAnalyticsRecordBuilder struct {
	v *DataFrameAnalyticsRecord
}

// NewDataFrameAnalyticsRecord provides a builder for the DataFrameAnalyticsRecord struct.
func NewDataFrameAnalyticsRecordBuilder() *DataFrameAnalyticsRecordBuilder {
	r := DataFrameAnalyticsRecordBuilder{
		&DataFrameAnalyticsRecord{},
	}

	return &r
}

// Build finalize the chain and returns the DataFrameAnalyticsRecord struct
func (rb *DataFrameAnalyticsRecordBuilder) Build() DataFrameAnalyticsRecord {
	return *rb.v
}

// AssignmentExplanation why the job is or is not assigned to a node

func (rb *DataFrameAnalyticsRecordBuilder) AssignmentExplanation(assignmentexplanation string) *DataFrameAnalyticsRecordBuilder {
	rb.v.AssignmentExplanation = &assignmentexplanation
	return rb
}

// CreateTime job creation time

func (rb *DataFrameAnalyticsRecordBuilder) CreateTime(createtime string) *DataFrameAnalyticsRecordBuilder {
	rb.v.CreateTime = &createtime
	return rb
}

// Description description

func (rb *DataFrameAnalyticsRecordBuilder) Description(description string) *DataFrameAnalyticsRecordBuilder {
	rb.v.Description = &description
	return rb
}

// DestIndex destination index

func (rb *DataFrameAnalyticsRecordBuilder) DestIndex(destindex IndexName) *DataFrameAnalyticsRecordBuilder {
	rb.v.DestIndex = &destindex
	return rb
}

// FailureReason failure reason

func (rb *DataFrameAnalyticsRecordBuilder) FailureReason(failurereason string) *DataFrameAnalyticsRecordBuilder {
	rb.v.FailureReason = &failurereason
	return rb
}

// Id the id

func (rb *DataFrameAnalyticsRecordBuilder) Id(id Id) *DataFrameAnalyticsRecordBuilder {
	rb.v.Id = &id
	return rb
}

// ModelMemoryLimit model memory limit

func (rb *DataFrameAnalyticsRecordBuilder) ModelMemoryLimit(modelmemorylimit string) *DataFrameAnalyticsRecordBuilder {
	rb.v.ModelMemoryLimit = &modelmemorylimit
	return rb
}

// NodeAddress network address of the assigned node

func (rb *DataFrameAnalyticsRecordBuilder) NodeAddress(nodeaddress string) *DataFrameAnalyticsRecordBuilder {
	rb.v.NodeAddress = &nodeaddress
	return rb
}

// NodeEphemeralId ephemeral id of the assigned node

func (rb *DataFrameAnalyticsRecordBuilder) NodeEphemeralId(nodeephemeralid Id) *DataFrameAnalyticsRecordBuilder {
	rb.v.NodeEphemeralId = &nodeephemeralid
	return rb
}

// NodeId id of the assigned node

func (rb *DataFrameAnalyticsRecordBuilder) NodeId(nodeid Id) *DataFrameAnalyticsRecordBuilder {
	rb.v.NodeId = &nodeid
	return rb
}

// NodeName name of the assigned node

func (rb *DataFrameAnalyticsRecordBuilder) NodeName(nodename Name) *DataFrameAnalyticsRecordBuilder {
	rb.v.NodeName = &nodename
	return rb
}

// Progress progress

func (rb *DataFrameAnalyticsRecordBuilder) Progress(progress string) *DataFrameAnalyticsRecordBuilder {
	rb.v.Progress = &progress
	return rb
}

// SourceIndex source index

func (rb *DataFrameAnalyticsRecordBuilder) SourceIndex(sourceindex IndexName) *DataFrameAnalyticsRecordBuilder {
	rb.v.SourceIndex = &sourceindex
	return rb
}

// State job state

func (rb *DataFrameAnalyticsRecordBuilder) State(state string) *DataFrameAnalyticsRecordBuilder {
	rb.v.State = &state
	return rb
}

// Type analysis type

func (rb *DataFrameAnalyticsRecordBuilder) Type_(type_ string) *DataFrameAnalyticsRecordBuilder {
	rb.v.Type = &type_
	return rb
}

// Version the version of Elasticsearch when the analytics was created

func (rb *DataFrameAnalyticsRecordBuilder) Version(version VersionString) *DataFrameAnalyticsRecordBuilder {
	rb.v.Version = &version
	return rb
}
