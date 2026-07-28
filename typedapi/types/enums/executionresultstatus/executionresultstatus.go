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
// https://github.com/elastic/elasticsearch-specification/tree/7fd0bd13eaf28bd179fc906f57da09e852eb818e

// Package executionresultstatus
package executionresultstatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/7fd0bd13eaf28bd179fc906f57da09e852eb818e/specification/watcher/_types/Execution.ts#L61-L71
type ExecutionResultStatus struct {
	Name string
}

var (
	Success = ExecutionResultStatus{"success"}

	Failure = ExecutionResultStatus{"failure"}
)

func (e ExecutionResultStatus) MarshalText() (text []byte, err error) {
	return []byte(e.String()), nil
}

func (e *ExecutionResultStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "success":
		*e = Success
	case "failure":
		*e = Failure
	default:
		*e = ExecutionResultStatus{string(text)}
	}

	return nil
}

func (e ExecutionResultStatus) String() string {
	return e.Name
}
