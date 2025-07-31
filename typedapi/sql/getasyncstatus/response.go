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

package getasyncstatus

// Response holds the response body struct for the package getasyncstatus
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/sql/get_async_status/SqlGetAsyncStatusResponse.ts#L23-L55
type Response struct {

	// CompletionStatus The HTTP status code for the search.
	// The API returns this property only for completed searches.
	CompletionStatus *uint `json:"completion_status,omitempty"`
	// ExpirationTimeInMillis The timestamp, in milliseconds since the Unix epoch, when Elasticsearch will
	// delete the search and its results, even if the search is still running.
	ExpirationTimeInMillis int64 `json:"expiration_time_in_millis"`
	// Id The identifier for the search.
	Id string `json:"id"`
	// IsPartial If `true`, the response does not contain complete search results.
	// If `is_partial` is `true` and `is_running` is `true`, the search is still
	// running.
	// If `is_partial` is `true` but `is_running` is `false`, the results are
	// partial due to a failure or timeout.
	IsPartial bool `json:"is_partial"`
	// IsRunning If `true`, the search is still running.
	// If `false`, the search has finished.
	IsRunning bool `json:"is_running"`
	// StartTimeInMillis The timestamp, in milliseconds since the Unix epoch, when the search started.
	// The API returns this property only for running searches.
	StartTimeInMillis int64 `json:"start_time_in_millis"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
