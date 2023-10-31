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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package executewatch

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actionexecutionmode"
)

// Request holds the request body struct for the package executewatch
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/watcher/execute_watch/WatcherExecuteWatchRequest.ts#L28-L79
type Request struct {

	// ActionModes Determines how to handle the watch actions as part of the watch execution.
	ActionModes map[string]actionexecutionmode.ActionExecutionMode `json:"action_modes,omitempty"`
	// AlternativeInput When present, the watch uses this object as a payload instead of executing
	// its own input.
	AlternativeInput map[string]json.RawMessage `json:"alternative_input,omitempty"`
	// IgnoreCondition When set to `true`, the watch execution uses the always condition. This can
	// also be specified as an HTTP parameter.
	IgnoreCondition *bool `json:"ignore_condition,omitempty"`
	// RecordExecution When set to `true`, the watch record representing the watch execution result
	// is persisted to the `.watcher-history` index for the current time. In
	// addition, the status of the watch is updated, possibly throttling subsequent
	// executions. This can also be specified as an HTTP parameter.
	RecordExecution  *bool                   `json:"record_execution,omitempty"`
	SimulatedActions *types.SimulatedActions `json:"simulated_actions,omitempty"`
	// TriggerData This structure is parsed as the data of the trigger event that will be used
	// during the watch execution
	TriggerData *types.ScheduleTriggerEvent `json:"trigger_data,omitempty"`
	// Watch When present, this watch is used instead of the one specified in the request.
	// This watch is not persisted to the index and record_execution cannot be set.
	Watch *types.Watch `json:"watch,omitempty"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{
		ActionModes:      make(map[string]actionexecutionmode.ActionExecutionMode, 0),
		AlternativeInput: make(map[string]json.RawMessage, 0),
	}
	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Executewatch request: %w", err)
	}

	return &req, nil
}
