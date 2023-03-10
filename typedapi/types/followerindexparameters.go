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

// FollowerIndexParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ccr/follow_info/types.ts#L38-L49
type FollowerIndexParameters struct {
	MaxOutstandingReadRequests    int      `json:"max_outstanding_read_requests"`
	MaxOutstandingWriteRequests   int      `json:"max_outstanding_write_requests"`
	MaxReadRequestOperationCount  int      `json:"max_read_request_operation_count"`
	MaxReadRequestSize            string   `json:"max_read_request_size"`
	MaxRetryDelay                 Duration `json:"max_retry_delay"`
	MaxWriteBufferCount           int      `json:"max_write_buffer_count"`
	MaxWriteBufferSize            string   `json:"max_write_buffer_size"`
	MaxWriteRequestOperationCount int      `json:"max_write_request_operation_count"`
	MaxWriteRequestSize           string   `json:"max_write_request_size"`
	ReadPollTimeout               Duration `json:"read_poll_timeout"`
}

// NewFollowerIndexParameters returns a FollowerIndexParameters.
func NewFollowerIndexParameters() *FollowerIndexParameters {
	r := &FollowerIndexParameters{}

	return r
}
