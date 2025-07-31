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

package deletebyquery

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package deletebyquery
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_global/delete_by_query/DeleteByQueryResponse.ts#L26-L88
type Response struct {

	// Batches The number of scroll responses pulled back by the delete by query.
	Batches *int64 `json:"batches,omitempty"`
	// Deleted The number of documents that were successfully deleted.
	Deleted *int64 `json:"deleted,omitempty"`
	// Failures An array of failures if there were any unrecoverable errors during the
	// process.
	// If this array is not empty, the request ended abnormally because of those
	// failures.
	// Delete by query is implemented using batches and any failures cause the
	// entire process to end but all failures in the current batch are collected
	// into the array.
	// You can use the `conflicts` option to prevent reindex from ending on version
	// conflicts.
	Failures []types.BulkIndexByScrollFailure `json:"failures,omitempty"`
	// Noops This field is always equal to zero for delete by query.
	// It exists only so that delete by query, update by query, and reindex APIs
	// return responses with the same structure.
	Noops *int64 `json:"noops,omitempty"`
	// RequestsPerSecond The number of requests per second effectively run during the delete by query.
	RequestsPerSecond *float32 `json:"requests_per_second,omitempty"`
	// Retries The number of retries attempted by delete by query.
	// `bulk` is the number of bulk actions retried.
	// `search` is the number of search actions retried.
	Retries   *types.Retries `json:"retries,omitempty"`
	SliceId   *int           `json:"slice_id,omitempty"`
	Task      types.TaskId   `json:"task,omitempty"`
	Throttled types.Duration `json:"throttled,omitempty"`
	// ThrottledMillis The number of milliseconds the request slept to conform to
	// `requests_per_second`.
	ThrottledMillis *int64         `json:"throttled_millis,omitempty"`
	ThrottledUntil  types.Duration `json:"throttled_until,omitempty"`
	// ThrottledUntilMillis This field should always be equal to zero in a `_delete_by_query` response.
	// It has meaning only when using the task API, where it indicates the next time
	// (in milliseconds since epoch) a throttled request will be run again in order
	// to conform to `requests_per_second`.
	ThrottledUntilMillis *int64 `json:"throttled_until_millis,omitempty"`
	// TimedOut If `true`, some requests run during the delete by query operation timed out.
	TimedOut *bool `json:"timed_out,omitempty"`
	// Took The number of milliseconds from start to end of the whole operation.
	Took *int64 `json:"took,omitempty"`
	// Total The number of documents that were successfully processed.
	Total *int64 `json:"total,omitempty"`
	// VersionConflicts The number of version conflicts that the delete by query hit.
	VersionConflicts *int64 `json:"version_conflicts,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
