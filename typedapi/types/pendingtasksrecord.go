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

// PendingTasksRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/pending_tasks/types.ts#L20-L41
type PendingTasksRecord struct {
	// InsertOrder task insertion order
	InsertOrder *string `json:"insertOrder,omitempty"`
	// Priority task priority
	Priority *string `json:"priority,omitempty"`
	// Source task source
	Source *string `json:"source,omitempty"`
	// TimeInQueue how long task has been in queue
	TimeInQueue *string `json:"timeInQueue,omitempty"`
}

// PendingTasksRecordBuilder holds PendingTasksRecord struct and provides a builder API.
type PendingTasksRecordBuilder struct {
	v *PendingTasksRecord
}

// NewPendingTasksRecord provides a builder for the PendingTasksRecord struct.
func NewPendingTasksRecordBuilder() *PendingTasksRecordBuilder {
	r := PendingTasksRecordBuilder{
		&PendingTasksRecord{},
	}

	return &r
}

// Build finalize the chain and returns the PendingTasksRecord struct
func (rb *PendingTasksRecordBuilder) Build() PendingTasksRecord {
	return *rb.v
}

// InsertOrder task insertion order

func (rb *PendingTasksRecordBuilder) InsertOrder(insertorder string) *PendingTasksRecordBuilder {
	rb.v.InsertOrder = &insertorder
	return rb
}

// Priority task priority

func (rb *PendingTasksRecordBuilder) Priority(priority string) *PendingTasksRecordBuilder {
	rb.v.Priority = &priority
	return rb
}

// Source task source

func (rb *PendingTasksRecordBuilder) Source(source string) *PendingTasksRecordBuilder {
	rb.v.Source = &source
	return rb
}

// TimeInQueue how long task has been in queue

func (rb *PendingTasksRecordBuilder) TimeInQueue(timeinqueue string) *PendingTasksRecordBuilder {
	rb.v.TimeInQueue = &timeinqueue
	return rb
}
