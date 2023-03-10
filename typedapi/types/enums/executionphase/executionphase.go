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

// Package executionphase
package executionphase

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/watcher/_types/Execution.ts#L49-L58
type ExecutionPhase struct {
	Name string
}

var (
	Awaitsexecution = ExecutionPhase{"awaits_execution"}

	Started = ExecutionPhase{"started"}

	Input = ExecutionPhase{"input"}

	Condition = ExecutionPhase{"condition"}

	Actions = ExecutionPhase{"actions"}

	Watchtransform = ExecutionPhase{"watch_transform"}

	Aborted = ExecutionPhase{"aborted"}

	Finished = ExecutionPhase{"finished"}
)

func (e ExecutionPhase) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *ExecutionPhase) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {

	case "awaits_execution":
		*e = Awaitsexecution
	case "started":
		*e = Started
	case "input":
		*e = Input
	case "condition":
		*e = Condition
	case "actions":
		*e = Actions
	case "watch_transform":
		*e = Watchtransform
	case "aborted":
		*e = Aborted
	case "finished":
		*e = Finished
	default:
		*e = ExecutionPhase{string(text)}
	}

	return nil
}

func (e ExecutionPhase) String() string {
	return e.Name
}
