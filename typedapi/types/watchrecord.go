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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/executionstatus"
)

// WatchRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/execute_watch/types.ts#L27-L39
type WatchRecord struct {
	Condition    ConditionContainer              `json:"condition"`
	Input        InputContainer                  `json:"input"`
	Messages     []string                        `json:"messages"`
	Metadata     *Metadata                       `json:"metadata,omitempty"`
	Node         string                          `json:"node"`
	Result       ExecutionResult                 `json:"result"`
	State        executionstatus.ExecutionStatus `json:"state"`
	Status       *WatchStatus                    `json:"status,omitempty"`
	TriggerEvent TriggerEventResult              `json:"trigger_event"`
	User         Username                        `json:"user"`
	WatchId      Id                              `json:"watch_id"`
}

// WatchRecordBuilder holds WatchRecord struct and provides a builder API.
type WatchRecordBuilder struct {
	v *WatchRecord
}

// NewWatchRecord provides a builder for the WatchRecord struct.
func NewWatchRecordBuilder() *WatchRecordBuilder {
	r := WatchRecordBuilder{
		&WatchRecord{},
	}

	return &r
}

// Build finalize the chain and returns the WatchRecord struct
func (rb *WatchRecordBuilder) Build() WatchRecord {
	return *rb.v
}

func (rb *WatchRecordBuilder) Condition(condition *ConditionContainerBuilder) *WatchRecordBuilder {
	v := condition.Build()
	rb.v.Condition = v
	return rb
}

func (rb *WatchRecordBuilder) Input(input *InputContainerBuilder) *WatchRecordBuilder {
	v := input.Build()
	rb.v.Input = v
	return rb
}

func (rb *WatchRecordBuilder) Messages(messages ...string) *WatchRecordBuilder {
	rb.v.Messages = messages
	return rb
}

func (rb *WatchRecordBuilder) Metadata(metadata *MetadataBuilder) *WatchRecordBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *WatchRecordBuilder) Node(node string) *WatchRecordBuilder {
	rb.v.Node = node
	return rb
}

func (rb *WatchRecordBuilder) Result(result *ExecutionResultBuilder) *WatchRecordBuilder {
	v := result.Build()
	rb.v.Result = v
	return rb
}

func (rb *WatchRecordBuilder) State(state executionstatus.ExecutionStatus) *WatchRecordBuilder {
	rb.v.State = state
	return rb
}

func (rb *WatchRecordBuilder) Status(status *WatchStatusBuilder) *WatchRecordBuilder {
	v := status.Build()
	rb.v.Status = &v
	return rb
}

func (rb *WatchRecordBuilder) TriggerEvent(triggerevent *TriggerEventResultBuilder) *WatchRecordBuilder {
	v := triggerevent.Build()
	rb.v.TriggerEvent = v
	return rb
}

func (rb *WatchRecordBuilder) User(user Username) *WatchRecordBuilder {
	rb.v.User = user
	return rb
}

func (rb *WatchRecordBuilder) WatchId(watchid Id) *WatchRecordBuilder {
	rb.v.WatchId = watchid
	return rb
}
