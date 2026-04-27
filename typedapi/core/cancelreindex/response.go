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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package cancelreindex

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

// Response holds the response body struct for the package cancelreindex
//
// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/_global/cancel_reindex/CancelReindexResponse.ts#L31-L87
type Response struct {
	// Acknowledged Present and `true` when `wait_for_completion=false`.
	Acknowledged *bool `json:"acknowledged,omitempty"`
	// Cancelled Whether the reindex task has been cancelled.
	Cancelled *bool `json:"cancelled,omitempty"`
	// Completed Whether the reindex task has completed. Present when
	// `wait_for_completion=true`.
	Completed *bool `json:"completed,omitempty"`
	// Description A sanitized description of the reindex operation.
	Description *string `json:"description,omitempty"`
	// Error The error that caused the reindex task to fail, if any.
	Error *types.ErrorCause `json:"error,omitempty"`
	// Id The ID of the reindex task, in `nodeId:taskNum` format. Present when
	// `wait_for_completion=true`.
	Id *string `json:"id,omitempty"`
	// Response The final result of the reindex operation, if it completed before being
	// cancelled.
	Response *types.ReindexTaskResult `json:"response,omitempty"`
	// RunningTime The elapsed running time of the reindex task, in a human-readable format.
	// Only present when the request includes the `?human=true` query parameter.
	RunningTime types.Duration `json:"running_time,omitempty"`
	// RunningTimeInNanos The elapsed running time of the reindex task, in nanoseconds.
	RunningTimeInNanos *int64 `json:"running_time_in_nanos,omitempty"`
	// StartTime The time at which the reindex task started, as an ISO 8601 formatted string.
	StartTime *string `json:"start_time,omitempty"`
	// StartTimeInMillis The time at which the reindex task started, in milliseconds since the Unix
	// epoch.
	StartTimeInMillis *int64 `json:"start_time_in_millis,omitempty"`
	// Status The status of the reindex operation at the time of cancellation.
	Status *types.ReindexStatus `json:"status,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
