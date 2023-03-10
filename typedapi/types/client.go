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

package types

// Client type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/nodes/_types/Stats.ts#L272-L284
type Client struct {
	Agent                 *string `json:"agent,omitempty"`
	ClosedTimeMillis      *int64  `json:"closed_time_millis,omitempty"`
	Id                    *int64  `json:"id,omitempty"`
	LastRequestTimeMillis *int64  `json:"last_request_time_millis,omitempty"`
	LastUri               *string `json:"last_uri,omitempty"`
	LocalAddress          *string `json:"local_address,omitempty"`
	OpenedTimeMillis      *int64  `json:"opened_time_millis,omitempty"`
	RemoteAddress         *string `json:"remote_address,omitempty"`
	RequestCount          *int64  `json:"request_count,omitempty"`
	RequestSizeBytes      *int64  `json:"request_size_bytes,omitempty"`
	XOpaqueId             *string `json:"x_opaque_id,omitempty"`
}

// NewClient returns a Client.
func NewClient() *Client {
	r := &Client{}

	return r
}
