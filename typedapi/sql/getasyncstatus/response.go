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

package getasyncstatus

// Response holds the response body struct for the package getasyncstatus
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/sql/get_async_status/SqlGetAsyncStatusResponse.ts#L23-L55

type Response struct {

	// CompletionStatus HTTP status code for the search. The API only returns this property for
	// completed searches.
	CompletionStatus *uint `json:"completion_status,omitempty"`
	// ExpirationTimeInMillis Timestamp, in milliseconds since the Unix epoch, when Elasticsearch will
	// delete
	// the search and its results, even if the search is still running.
	ExpirationTimeInMillis int64 `json:"expiration_time_in_millis"`
	// Id Identifier for the search.
	Id string `json:"id"`
	// IsPartial If `true`, the response does not contain complete search results. If
	// `is_partial`
	// is `true` and `is_running` is `true`, the search is still running. If
	// `is_partial`
	// is `true` but `is_running` is `false`, the results are partial due to a
	// failure or
	// timeout.
	IsPartial bool `json:"is_partial"`
	// IsRunning If `true`, the search is still running. If `false`, the search has finished.
	IsRunning bool `json:"is_running"`
	// StartTimeInMillis Timestamp, in milliseconds since the Unix epoch, when the search started.
	// The API only returns this property for running searches.
	StartTimeInMillis int64 `json:"start_time_in_millis"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
