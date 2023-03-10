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

package getstatus

// Response holds the response body struct for the package getstatus
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/eql/get_status/EqlGetStatusResponse.ts#L24-L51

type Response struct {

	// CompletionStatus For a completed search shows the http status code of the completed search.
	CompletionStatus *int `json:"completion_status,omitempty"`
	// ExpirationTimeInMillis Shows a timestamp when the eql search will be expired, in milliseconds since
	// the Unix epoch. When this time is reached, the search and its results are
	// deleted, even if the search is still ongoing.
	ExpirationTimeInMillis *int64 `json:"expiration_time_in_millis,omitempty"`
	// Id Identifier for the search.
	Id string `json:"id"`
	// IsPartial If true, the search request is still executing. If false, the search is
	// completed.
	IsPartial bool `json:"is_partial"`
	// IsRunning If true, the response does not contain complete search results. This could be
	// because either the search is still running (is_running status is false), or
	// because it is already completed (is_running status is true) and results are
	// partial due to failures or timeouts.
	IsRunning bool `json:"is_running"`
	// StartTimeInMillis For a running search shows a timestamp when the eql search started, in
	// milliseconds since the Unix epoch.
	StartTimeInMillis *int64 `json:"start_time_in_millis,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
