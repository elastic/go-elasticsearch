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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

// Package recoveryrunstatus
package recoveryrunstatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/data_recovery/_types/RecoveryRunStatus.ts#L20-L27
type RecoveryRunStatus struct {
	Name string
}

var (

	// Running The recovery has not finished.
	Running = RecoveryRunStatus{"running"}

	// Completed Every source was recovered.
	Completed = RecoveryRunStatus{"completed"}

	// Failed The recovery stopped without recovering every source. The `failures` of the
	// recovery describe why.
	Failed = RecoveryRunStatus{"failed"}
)

func (r RecoveryRunStatus) MarshalText() (text []byte, err error) {
	return []byte(r.String()), nil
}

func (r *RecoveryRunStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "running":
		*r = Running
	case "completed":
		*r = Completed
	case "failed":
		*r = Failed
	default:
		*r = RecoveryRunStatus{string(text)}
	}

	return nil
}

func (r RecoveryRunStatus) String() string {
	return r.Name
}
