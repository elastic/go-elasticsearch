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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package status

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package status
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/async_search/status/AsyncSearchStatusResponse.ts#L28-L30

type Response struct {
	CompletionStatus       *int                  `json:"completion_status,omitempty"`
	ExpirationTime         types.DateTime        `json:"expiration_time,omitempty"`
	ExpirationTimeInMillis int64                 `json:"expiration_time_in_millis"`
	Id                     *string               `json:"id,omitempty"`
	IsPartial              bool                  `json:"is_partial"`
	IsRunning              bool                  `json:"is_running"`
	Shards_                types.ShardStatistics `json:"_shards"`
	StartTime              types.DateTime        `json:"start_time,omitempty"`
	StartTimeInMillis      int64                 `json:"start_time_in_millis"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
