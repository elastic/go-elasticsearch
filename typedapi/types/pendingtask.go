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

// PendingTask type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/pending_tasks/types.ts#L23-L30
type PendingTask struct {
	Executing         bool                    `json:"executing"`
	InsertOrder       int                     `json:"insert_order"`
	Priority          string                  `json:"priority"`
	Source            string                  `json:"source"`
	TimeInQueue       *Duration               `json:"time_in_queue,omitempty"`
	TimeInQueueMillis DurationValueUnitMillis `json:"time_in_queue_millis"`
}

// PendingTaskBuilder holds PendingTask struct and provides a builder API.
type PendingTaskBuilder struct {
	v *PendingTask
}

// NewPendingTask provides a builder for the PendingTask struct.
func NewPendingTaskBuilder() *PendingTaskBuilder {
	r := PendingTaskBuilder{
		&PendingTask{},
	}

	return &r
}

// Build finalize the chain and returns the PendingTask struct
func (rb *PendingTaskBuilder) Build() PendingTask {
	return *rb.v
}

func (rb *PendingTaskBuilder) Executing(executing bool) *PendingTaskBuilder {
	rb.v.Executing = executing
	return rb
}

func (rb *PendingTaskBuilder) InsertOrder(insertorder int) *PendingTaskBuilder {
	rb.v.InsertOrder = insertorder
	return rb
}

func (rb *PendingTaskBuilder) Priority(priority string) *PendingTaskBuilder {
	rb.v.Priority = priority
	return rb
}

func (rb *PendingTaskBuilder) Source(source string) *PendingTaskBuilder {
	rb.v.Source = source
	return rb
}

func (rb *PendingTaskBuilder) TimeInQueue(timeinqueue *DurationBuilder) *PendingTaskBuilder {
	v := timeinqueue.Build()
	rb.v.TimeInQueue = &v
	return rb
}

func (rb *PendingTaskBuilder) TimeInQueueMillis(timeinqueuemillis *DurationValueUnitMillisBuilder) *PendingTaskBuilder {
	v := timeinqueuemillis.Build()
	rb.v.TimeInQueueMillis = v
	return rb
}
