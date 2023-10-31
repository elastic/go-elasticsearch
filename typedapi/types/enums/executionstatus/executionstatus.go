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

// Package executionstatus
package executionstatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/watcher/_types/Execution.ts#L38-L47
type ExecutionStatus struct {
	Name string
}

var (
	Awaitsexecution = ExecutionStatus{"awaits_execution"}

	Checking = ExecutionStatus{"checking"}

	Executionnotneeded = ExecutionStatus{"execution_not_needed"}

	Throttled = ExecutionStatus{"throttled"}

	Executed = ExecutionStatus{"executed"}

	Failed = ExecutionStatus{"failed"}

	Deletedwhilequeued = ExecutionStatus{"deleted_while_queued"}

	Notexecutedalreadyqueued = ExecutionStatus{"not_executed_already_queued"}
)

func (e ExecutionStatus) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *ExecutionStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "awaits_execution":
		*e = Awaitsexecution
	case "checking":
		*e = Checking
	case "execution_not_needed":
		*e = Executionnotneeded
	case "throttled":
		*e = Throttled
	case "executed":
		*e = Executed
	case "failed":
		*e = Failed
	case "deleted_while_queued":
		*e = Deletedwhilequeued
	case "not_executed_already_queued":
		*e = Notexecutedalreadyqueued
	default:
		*e = ExecutionStatus{string(text)}
	}

	return nil
}

func (e ExecutionStatus) String() string {
	return e.Name
}
