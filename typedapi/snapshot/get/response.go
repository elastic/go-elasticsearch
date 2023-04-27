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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package get

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Response holds the response body struct for the package get
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/snapshot/get/SnapshotGetResponse.ts#L25-L40

type Response struct {

	// Remaining The number of remaining snapshots that were not returned due to size limits
	// and that can be fetched by additional requests using the next field value.
	Remaining int                          `json:"remaining"`
	Responses []types.SnapshotResponseItem `json:"responses,omitempty"`
	Snapshots []types.SnapshotInfo         `json:"snapshots,omitempty"`
	// Total The total number of snapshots that match the request when ignoring size limit
	// or after query parameter.
	Total int `json:"total"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
