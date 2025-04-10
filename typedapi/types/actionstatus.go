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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package types

// ActionStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/beeb1dc688bcc058488dcc45d9cbd2cd364e9943/specification/watcher/_types/Action.ts#L125-L130
type ActionStatus struct {
	Ack                     AcknowledgeState `json:"ack"`
	LastExecution           *ExecutionState  `json:"last_execution,omitempty"`
	LastSuccessfulExecution *ExecutionState  `json:"last_successful_execution,omitempty"`
	LastThrottle            *ThrottleState   `json:"last_throttle,omitempty"`
}

// NewActionStatus returns a ActionStatus.
func NewActionStatus() *ActionStatus {
	r := &ActionStatus{}

	return r
}

// true

type ActionStatusVariant interface {
	ActionStatusCaster() *ActionStatus
}

func (s *ActionStatus) ActionStatusCaster() *ActionStatus {
	return s
}
