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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package syncjobclaim

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package syncjobclaim
//
// https://github.com/elastic/elasticsearch-specification/blob/60a81659be928bfe6cec53708c7f7613555a5eaf/specification/connector/sync_job_claim/SyncJobClaimRequest.ts#L23-L61
type Request struct {

	// SyncCursor The cursor object from the last incremental sync job.
	// This should reference the `sync_cursor` field in the connector state for
	// which the job runs.
	SyncCursor json.RawMessage `json:"sync_cursor,omitempty"`
	// WorkerHostname The host name of the current system that will run the job.
	WorkerHostname string `json:"worker_hostname"`
}

// NewRequest returns a Request
func NewRequest() *Request {
	r := &Request{}

	return r
}

// FromJSON allows to load an arbitrary json into the request structure
func (r *Request) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Syncjobclaim request: %w", err)
	}

	return &req, nil
}
