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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Package actionexecutionmode
package actionexecutionmode

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/bc885996c471cc7c2c7d51cba22aab19867672ac/specification/watcher/_types/Action.ts#L67-L88
type ActionExecutionMode struct {
	Name string
}

var (

	// Simulate The action execution is simulated. Each action type defines its own
	// simulation operation mode. For example, the email action creates the email
	// that would have been sent but does not actually send it. In this mode, the
	// action might be throttled if the current state of the watch indicates it
	// should be.
	Simulate = ActionExecutionMode{"simulate"}

	// Forcesimulate Similar to the `simulate` mode, except the action is not throttled even if
	// the current state of the watch indicates it should be.
	Forcesimulate = ActionExecutionMode{"force_simulate"}

	// Execute Executes the action as it would have been executed if the watch had been
	// triggered by its own trigger. The execution might be throttled if the current
	// state of the watch indicates it should be.
	Execute = ActionExecutionMode{"execute"}

	// Forceexecute Similar to the `execute` mode, except the action is not throttled even if the
	// current state of the watch indicates it should be.
	Forceexecute = ActionExecutionMode{"force_execute"}

	// Skip The action is skipped and is not executed or simulated. Effectively forces
	// the action to be throttled.
	Skip = ActionExecutionMode{"skip"}
)

func (a ActionExecutionMode) MarshalText() (text []byte, err error) {
	return []byte(a.String()), nil
}

func (a *ActionExecutionMode) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "simulate":
		*a = Simulate
	case "force_simulate":
		*a = Forcesimulate
	case "execute":
		*a = Execute
	case "force_execute":
		*a = Forceexecute
	case "skip":
		*a = Skip
	default:
		*a = ActionExecutionMode{string(text)}
	}

	return nil
}

func (a ActionExecutionMode) String() string {
	return a.Name
}
