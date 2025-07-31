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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Package syncstatus
package syncstatus

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/connector/_types/Connector.ts#L138-L146
type SyncStatus struct {
	Name string
}

var (
	Canceling = SyncStatus{"canceling"}

	Canceled = SyncStatus{"canceled"}

	Completed = SyncStatus{"completed"}

	Error = SyncStatus{"error"}

	Inprogress = SyncStatus{"in_progress"}

	Pending = SyncStatus{"pending"}

	Suspended = SyncStatus{"suspended"}
)

func (s SyncStatus) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SyncStatus) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "canceling":
		*s = Canceling
	case "canceled":
		*s = Canceled
	case "completed":
		*s = Completed
	case "error":
		*s = Error
	case "in_progress":
		*s = Inprogress
	case "pending":
		*s = Pending
	case "suspended":
		*s = Suspended
	default:
		*s = SyncStatus{string(text)}
	}

	return nil
}

func (s SyncStatus) String() string {
	return s.Name
}
