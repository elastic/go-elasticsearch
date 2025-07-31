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

package updatebyquery

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package updatebyquery
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/update_by_query/UpdateByQueryResponse.ts#L26-L67
type Response struct {

	// Batches The number of scroll responses pulled back by the update by query.
	Batches *int64 `json:"batches,omitempty"`
	// Deleted The number of documents that were successfully deleted.
	Deleted *int64 `json:"deleted,omitempty"`
	// Failures Array of failures if there were any unrecoverable errors during the process.
	// If this is non-empty then the request ended because of those failures.
	// Update by query is implemented using batches.
	// Any failure causes the entire process to end, but all failures in the current
	// batch are collected into the array.
	// You can use the `conflicts` option to prevent reindex from ending when
	// version conflicts occur.
	Failures []types.BulkIndexByScrollFailure `json:"failures,omitempty"`
	// Noops The number of documents that were ignored because the script used for the
	// update by query returned a noop value for `ctx.op`.
	Noops *int64 `json:"noops,omitempty"`
	// RequestsPerSecond The number of requests per second effectively run during the update by query.
	RequestsPerSecond *float32 `json:"requests_per_second,omitempty"`
	// Retries The number of retries attempted by update by query.
	// `bulk` is the number of bulk actions retried.
	// `search` is the number of search actions retried.
	Retries   *types.Retries `json:"retries,omitempty"`
	Task      types.TaskId   `json:"task,omitempty"`
	Throttled types.Duration `json:"throttled,omitempty"`
	// ThrottledMillis The number of milliseconds the request slept to conform to
	// `requests_per_second`.
	ThrottledMillis *int64         `json:"throttled_millis,omitempty"`
	ThrottledUntil  types.Duration `json:"throttled_until,omitempty"`
	// ThrottledUntilMillis This field should always be equal to zero in an _update_by_query response.
	// It only has meaning when using the task API, where it indicates the next time
	// (in milliseconds since epoch) a throttled request will be run again in order
	// to conform to `requests_per_second`.
	ThrottledUntilMillis *int64 `json:"throttled_until_millis,omitempty"`
	// TimedOut If true, some requests timed out during the update by query.
	TimedOut *bool `json:"timed_out,omitempty"`
	// Took The number of milliseconds from start to end of the whole operation.
	Took *int64 `json:"took,omitempty"`
	// Total The number of documents that were successfully processed.
	Total *int64 `json:"total,omitempty"`
	// Updated The number of documents that were successfully updated.
	Updated *int64 `json:"updated,omitempty"`
	// VersionConflicts The number of version conflicts that the update by query hit.
	VersionConflicts *int64 `json:"version_conflicts,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
