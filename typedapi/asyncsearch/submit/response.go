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

package submit

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package submit
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/async_search/submit/AsyncSearchSubmitResponse.ts#L22-L24

type Response struct {

	// CompletionTime Indicates when the async search completed. Only present
	// when the search has completed.
	CompletionTime         types.DateTime `json:"completion_time,omitempty"`
	CompletionTimeInMillis *int64         `json:"completion_time_in_millis,omitempty"`
	// ExpirationTime Indicates when the async search will expire.
	ExpirationTime         types.DateTime `json:"expiration_time,omitempty"`
	ExpirationTimeInMillis int64          `json:"expiration_time_in_millis"`
	Id                     *string        `json:"id,omitempty"`
	// IsPartial When the query is no longer running, this property indicates whether the
	// search failed or was successfully completed on all shards.
	// While the query is running, `is_partial` is always set to `true`.
	IsPartial bool `json:"is_partial"`
	// IsRunning Indicates whether the search is still running or has completed.
	// NOTE: If the search failed after some shards returned their results or the
	// node that is coordinating the async search dies, results may be partial even
	// though `is_running` is `false`.
	IsRunning         bool              `json:"is_running"`
	Response          types.AsyncSearch `json:"response"`
	StartTime         types.DateTime    `json:"start_time,omitempty"`
	StartTimeInMillis int64             `json:"start_time_in_millis"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
